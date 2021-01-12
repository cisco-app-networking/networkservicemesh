package endpoint

import (
	"context"

	"github.com/sirupsen/logrus"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connectioncontext"
	"cisco-app-networking.github.io/networkservicemesh/pkg/tools"
	"cisco-app-networking.github.io/networkservicemesh/sdk/common"
)

//CreateRouteMutator - Creates an instance of ConnectionMutator with routes mutating
func CreateRouteMutator(routes []string) ConnectionMutator {
	return func(ctc context.Context, c *connection.Connection) error {
		for _, r := range routes {
			c.GetContext().GetIpContext().DstRoutes = append(c.GetContext().GetIpContext().GetDstRoutes(), &connectioncontext.Route{
				Prefix: r,
			})
		}
		return nil
	}
}

func CreatePodNameMutator() ConnectionMutator {
	return func(ctc context.Context, c *connection.Connection) error {
		podName, err := tools.GetCurrentPodNameFromHostname()
		if err != nil {
			logrus.Infof("failed to get current pod name from hostname: %v", err)
		} else {
			c.Labels[connection.PodNameKey] = podName
			c.Labels[connection.NamespaceKey] = common.GetNamespace()
		}
		return nil
	}

}
