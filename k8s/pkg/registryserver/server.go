package registryserver

import (
	"context"

	"cisco-app-networking.github.io/networkservicemesh/pkg/tools/spanhelper"

	"google.golang.org/grpc"

	"cisco-app-networking.github.io/networkservicemesh/k8s/pkg/apis/networkservice/v1alpha1"
	"cisco-app-networking.github.io/networkservicemesh/k8s/pkg/networkservice/namespace"
	"cisco-app-networking.github.io/networkservicemesh/k8s/pkg/registryserver/resourcecache"

	"cisco-app-networking.github.io/networkservicemesh/pkg/tools"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/registry"
	nsmClientset "cisco-app-networking.github.io/networkservicemesh/k8s/pkg/networkservice/clientset/versioned"
)

// New - construct a registration server
func New(ctx context.Context, clientset *nsmClientset.Clientset, nsmName string) *grpc.Server {
	span := spanhelper.FromContext(ctx, "K8SServer.New")
	defer span.Finish()
	server := tools.NewServer(span.Context())

	cache := NewRegistryCache(clientset, &ResourceFilterConfig{
		NetworkServiceManagerPolicy: resourcecache.FilterByNamespacePolicy(namespace.GetNamespace(), func(resource interface{}) string {
			nsm := resource.(*v1alpha1.NetworkServiceManager)
			return nsm.Namespace
		}),
	})

	nseRegistry := newNseRegistryService(nsmName, cache)
	nsmRegistry := newNsmRegistryService(nsmName, cache)
	discovery := newDiscoveryService(cache)

	registry.RegisterNetworkServiceRegistryServer(server, nseRegistry)
	registry.RegisterNetworkServiceDiscoveryServer(server, discovery)
	registry.RegisterNsmRegistryServer(server, nsmRegistry)

	err := cache.Start()
	span.LogError(err)
	span.Logger().Info("RegistryCache started")

	return server
}
