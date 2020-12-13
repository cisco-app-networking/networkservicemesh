package main

import (
	"context"
	"net"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"cisco-app-networking.github.io/networkservicemesh/k8s/pkg/registryserver"
	k8s_utils "cisco-app-networking.github.io/networkservicemesh/k8s/pkg/utils"
	"cisco-app-networking.github.io/networkservicemesh/pkg/tools"
	"cisco-app-networking.github.io/networkservicemesh/pkg/tools/jaeger"
	"cisco-app-networking.github.io/networkservicemesh/pkg/tools/spanhelper"
	"cisco-app-networking.github.io/networkservicemesh/utils"
)

var version string

func main() {
	logrus.Info("Starting nsmd-k8s...")
	logrus.Infof("Version: %v", version)
	utils.PrintAllEnv(logrus.StandardLogger())
	// Capture signals to cleanup before exiting
	c := tools.NewOSSignalChannel()

	closer := jaeger.InitJaeger("nsmd-k8s")
	defer func() { _ = closer.Close() }()

	address := os.Getenv("NSMD_K8S_ADDRESS")
	if strings.TrimSpace(address) == "" {
		address = "0.0.0.0:5000"
	}

	span := spanhelper.FromContext(context.Background(), "Start-NSMD-k8s")
	defer span.Finish()

	nsmName, ok := os.LookupEnv("NODE_NAME")

	span.LogObject("address", address)
	span.LogObject("nsmName", nsmName)

	if !ok {
		span.Logger().Fatalf("You must set env variable NODE_NAME to match the name of your Node.  See https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/")
	}
	span.LogValue("NODE_NAME", nsmName)
	span.Logger().Println("Starting NSMD Kubernetes on " + address + " with NsmName " + nsmName)

	nsmClientSet, _, err := k8s_utils.NewClientSet()
	if err != nil {
		span.LogError(err)
		span.Logger().Fatalln("Fail to start NSMD Kubernetes service", err)
	}

	server := registryserver.New(span.Context(), nsmClientSet, nsmName)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		span.Logger().Fatalln(err)
	}

	span.Logger().Print("nsmd-k8s initialized and waiting for connection")
	go func() {
		err = server.Serve(listener)
		if err != nil {
			span.LogError(err)
			span.Logger().Fatalln(err)
		}
	}()

	span.Finish()
	<-c
}
