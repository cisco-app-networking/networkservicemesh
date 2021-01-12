package proxyregistryserver

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"cisco-app-networking.github.io/networkservicemesh/pkg/tools"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/clusterinfo"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/registry"
	nsmClientset "cisco-app-networking.github.io/networkservicemesh/k8s/pkg/networkservice/clientset/versioned"
	"cisco-app-networking.github.io/networkservicemesh/k8s/pkg/registryserver"
)

// New starts proxy Network Service Discovery Server and Cluster Info Server
func New(clientset *nsmClientset.Clientset, clusterInfoService clusterinfo.ClusterInfoServer) *grpc.Server {
	server := tools.NewServer(context.Background())
	cache := registryserver.NewRegistryCache(clientset, &registryserver.ResourceFilterConfig{})
	discovery := newDiscoveryService(cache, clusterInfoService)
	nseRegistry := newNseRegistryService(clusterInfoService)

	registry.RegisterNetworkServiceDiscoveryServer(server, discovery)
	registry.RegisterNetworkServiceRegistryServer(server, nseRegistry)
	clusterinfo.RegisterClusterInfoServer(server, clusterInfoService)

	if err := cache.Start(); err != nil {
		logrus.Error(err)
	}
	logrus.Info("RegistryCache started")

	return server
}
