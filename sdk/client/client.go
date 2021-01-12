// Copyright 2018, 2019 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection/mechanisms/cls"

	"cisco-app-networking.github.io/networkservicemesh/pkg/tools/spanhelper"

	"cisco-app-networking.github.io/networkservicemesh/pkg/tools/jaeger"

	"github.com/sirupsen/logrus"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connectioncontext"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/networkservice"
	"cisco-app-networking.github.io/networkservicemesh/pkg/tools"
	"cisco-app-networking.github.io/networkservicemesh/sdk/common"
	ctrl_common "cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/common"
)

const (
	// ConnectTimeout - a default connection timeout
	ConnectTimeout = 15 * time.Second
	// ConnectionRetry - A number of retries for establish a network service, default == 10
	ConnectionRetry = 10
	// RequestDelay - A delay between attempts, default = 5sec
	RequestDelay = time.Second * 5
)

// NsmClient is the NSM client struct
type NsmClient struct {
	*common.NsmConnection
	ClientNetworkService string
	ClientLabels         map[string]string
	OutgoingConnections  []*connection.Connection
	NscInterfaceName     string
	tracerCloser         io.Closer
}

// Connect with no retry and delay
func (nsmc *NsmClient) Connect(ctx context.Context, name, mechanism, description string) (*connection.Connection, error) {
	return nsmc.ConnectRetry(ctx, name, mechanism, description,1, 0)
}

func (nsmc *NsmClient) ConnectRetry(ctx context.Context, name, mechanism, description string, retryCount int, retryDelay time.Duration) (*connection.Connection, error) {
	return nsmc.ConnectToEndpointRetry(ctx, "", "", "", name, mechanism, description, nsmc.Configuration.Routes, retryCount, retryDelay)
}

func (nsmc *NsmClient) ConnectToEndpoint(ctx context.Context, remoteIp, destEndpointName, destEndpointManager, name, mechanism, description string, routes []string) (*connection.Connection, error) {
	return nsmc.ConnectToEndpointRetry(ctx, remoteIp, destEndpointName, destEndpointManager, name, mechanism, description, routes, 1, 0)
}

// Connect implements the business logic
func (nsmc *NsmClient) ConnectToEndpointRetry(ctx context.Context, remoteIp, destEndpointName, destEndpointManager, name, mechanism, description string, routes []string, retryCount int, retryDelay time.Duration) (*connection.Connection, error) {
	span := spanhelper.FromContext(ctx, "nsmClient.Connect")
	defer span.Finish()
	span.Logger().WithFields(logrus.Fields{
		"destEndpointName": destEndpointName,
		"destEndpointManager": destEndpointManager,
		"remoteIp": remoteIp,
		"mechanismName": name,
		"mechanism": mechanism,
		"description": description,
	}).Infof("Initiating an outgoing connection.")
	span.Logger().Infof("Initiating an outgoing connection.")
	nsmc.Lock()
	defer nsmc.Unlock()

	if nsmc.NscInterfaceName != "" {
		// The environment variable will override local call parameters
		name = nsmc.NscInterfaceName
	}

	outgoingMechanism, err := common.NewMechanism(cls.LOCAL, mechanism, name, description)

	span.LogObject("Selected mechanism", outgoingMechanism)

	if err != nil {
		err = errors.Wrap(err, "failure to prepare the outgoing mechanism preference with error")
		span.LogError(err)
		return nil, err
	}

	srcRoutes := []*connectioncontext.Route{}
	for _, r := range routes {
		srcRoutes = append(srcRoutes, &connectioncontext.Route{
			Prefix: r,
		})
	}
	nsName := nsmc.Configuration.ClientNetworkService
	if remoteIp != "" {
		nsName = nsName + "@" + remoteIp
	}
	outgoingRequest := &networkservice.NetworkServiceRequest{
		Connection: &connection.Connection{
			NetworkService: nsName,
			Context: &connectioncontext.ConnectionContext{
				IpContext: &connectioncontext.IPContext{
					SrcIpRequired: true,
					DstIpRequired: true,
					SrcRoutes:     srcRoutes,
				},
			},
			Labels: nsmc.ClientLabels,
		},
		MechanismPreferences: []*connection.Mechanism{
			outgoingMechanism,
		},
	}
	if destEndpointName != "" {
		outgoingRequest.Connection.NetworkServiceEndpointName = destEndpointName
	}
	if destEndpointManager != "" {
		outgoingRequest.Connection.Path = ctrl_common.Strings2Path(destEndpointManager)
	}
	var outgoingConnection *connection.Connection
	maxRetry := retryCount
	for retryCount >= 0 {
		var attemptSpan = spanhelper.FromContext(span.Context(), fmt.Sprintf("nsmClient.Connect.attempt:%v", maxRetry-retryCount))
		defer attemptSpan.Finish()

		attempCtx, cancelProc := context.WithTimeout(attemptSpan.Context(), ConnectTimeout)
		defer cancelProc()

		attemptLogger := attemptSpan.Logger()
		attemptLogger.Infof("Requesting %v", outgoingRequest)
		outgoingConnection, err = nsmc.NsClient.Request(attempCtx, outgoingRequest)

		if err != nil {
			attemptSpan.LogError(err)

			cancelProc()
			if retryCount == 0 {
				return nil, errors.Wrap(err, "nsm client: Failed to connect")
			} else {
				attemptLogger.Errorf("nsm client: Failed to connect %v. Retry attempts: %v Delaying: %v", err, retryCount, retryDelay)
			}
			retryCount--
			attemptSpan.Finish()
			<-time.After(retryDelay)
			continue
		}
		break
	}
	span.Logger().WithFields(logrus.Fields{
		"destEndpointName": destEndpointName,
		"destEndpointManager": destEndpointManager,
		"remoteIp": remoteIp,
		"mechanismName": name,
		"mechanism": mechanism,
		"description": description,
	}).Infof("Successfully requested connection")
	span.LogObject("connection", outgoingConnection)

	nsmc.OutgoingConnections = append(nsmc.OutgoingConnections, outgoingConnection)
	return outgoingConnection, nil
}

// Close will terminate a particular connection
func (nsmc *NsmClient) Close(ctx context.Context, outgoingConnection *connection.Connection) error {
	nsmc.Lock()
	defer nsmc.Unlock()

	span := spanhelper.FromContext(ctx, "Client.Close")
	defer span.Finish()
	span.LogObject("connection", outgoingConnection)

	_, err := nsmc.NsClient.Close(span.Context(), outgoingConnection)

	span.LogError(err)

	arr := nsmc.OutgoingConnections
	for i, c := range arr {
		if c == outgoingConnection {
			copy(arr[i:], arr[i+1:])
			arr[len(arr)-1] = nil
			arr = arr[:len(arr)-1]
		}
	}
	return err
}

// Destroy - Destroy stops the whole module
func (nsmc *NsmClient) Destroy(ctx context.Context) error {
	nsmc.Lock()
	defer nsmc.Unlock()

	span := spanhelper.FromContext(ctx, "Client.Destroy")
	defer span.Finish()

	err := nsmc.NsmConnection.Close()
	span.LogError(errors.Wrap(err, "failed to close opentracing context"))
	if nsmc.tracerCloser != nil {
		trErr := nsmc.tracerCloser.Close()
		if trErr != nil {
			logrus.Error(trErr)
		}
	}
	return err
}

// NewNSMClient creates the NsmClient
func NewNSMClient(ctx context.Context, configuration *common.NSConfiguration) (*NsmClient, error) {
	if configuration == nil {
		configuration = &common.NSConfiguration{}
	}

	client := &NsmClient{
		ClientNetworkService: configuration.ClientNetworkService,
		ClientLabels:         tools.ParseKVStringToMap(configuration.ClientLabels, ",", "="),
		NscInterfaceName:     configuration.NscInterfaceName,
	}

	client.tracerCloser = jaeger.InitJaeger("nsm-client")

	nsmConnection, err := common.NewNSMConnection(ctx, configuration)
	if err != nil {
		logrus.Errorf("Error: %v", err)
		return nil, err
	}

	client.NsmConnection = nsmConnection

	return client, nil
}
