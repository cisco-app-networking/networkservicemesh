// Copyright (c) 2019 Cisco and/or its affiliates.
// Copyright (c) 2019 Red Hat Inc. and/or its affiliates.
// Copyright (c) 2019 VMware, Inc.
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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v1alpha1 "cisco-app-networking.github.io/networkservicemesh/k8s/pkg/networkservice/clientset/versioned/typed/networkservice/v1alpha1"
)

type FakeNetworkserviceV1alpha1 struct {
	*testing.Fake
}

func (c *FakeNetworkserviceV1alpha1) NetworkServices(namespace string) v1alpha1.NetworkServiceInterface {
	return &FakeNetworkServices{c, namespace}
}

func (c *FakeNetworkserviceV1alpha1) NetworkServiceEndpoints(namespace string) v1alpha1.NetworkServiceEndpointInterface {
	return &FakeNetworkServiceEndpoints{c, namespace}
}

func (c *FakeNetworkserviceV1alpha1) NetworkServiceManagers(namespace string) v1alpha1.NetworkServiceManagerInterface {
	return &FakeNetworkServiceManagers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkserviceV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
