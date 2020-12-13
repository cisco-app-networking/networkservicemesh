package nsmd

import (
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/networkservice"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/api/nsm"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/common"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/local"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/model"
)

// NewNetworkServiceServer - construct a local network service chain
func NewNetworkServiceServer(model model.Model, ws *Workspace,
	nsmManager nsm.NetworkServiceManager) networkservice.NetworkServiceServer {
	return common.NewCompositeService("Local",
		common.NewRequestValidator(),
		common.NewMonitorService(ws.MonitorConnectionServer()),
		local.NewWorkspaceService(ws.Name()),
		local.NewConnectionService(model),
		local.NewForwarderService(model, nsmManager.ServiceRegistry()),
		local.NewEndpointSelectorService(nsmManager.NseManager()),
		common.NewExcludedPrefixesService(),
		local.NewEndpointService(nsmManager.NseManager(), nsmManager.GetHealProperties(), nsmManager.Model()),
		common.NewCrossConnectService(),
	)
}
