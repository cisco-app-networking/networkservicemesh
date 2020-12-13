package remote

import (
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/services"
	"cisco-app-networking.github.io/networkservicemesh/sdk/monitor"
	"cisco-app-networking.github.io/networkservicemesh/sdk/monitor/connectionmonitor"
)

// MonitorServer is a monitor.Server for remote/connection GRPC API
type MonitorServer interface {
	monitor.Server
	connection.MonitorConnectionServer
}

type monitorServer struct {
	connectionmonitor.MonitorServer
	manager *services.ClientConnectionManager
}

// NewMonitorServer creates a new MonitorServer
func NewMonitorServer(manager *services.ClientConnectionManager) MonitorServer {
	rv := &monitorServer{
		MonitorServer: connectionmonitor.NewMonitorServer("RemoteConnection"),
		manager:       manager,
	}
	return rv
}

// MonitorConnections adds recipient for MonitorServer events
func (s *monitorServer) MonitorConnections(selector *connection.MonitorScopeSelector, recipient connection.MonitorConnection_MonitorConnectionsServer) error {
	err := s.MonitorServer.MonitorConnections(selector, recipient)
	if s.manager != nil {
		s.manager.UpdateRemoteMonitorDone(selector.GetPathSegments())
	}
	return err
}
