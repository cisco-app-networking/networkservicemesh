package metrics

import "cisco-app-networking.github.io/networkservicemesh/controlplane/api/crossconnect"

type MetricsMonitor interface {
	HandleMetrics(statistics map[string]*crossconnect.Metrics)
}
