package selector

import (
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/registry"
)

type Selector interface {
	SelectEndpoint(requestConnection *connection.Connection, ns *registry.NetworkService, networkServiceEndpoints []*registry.NetworkServiceEndpoint) *registry.NetworkServiceEndpoint
}
