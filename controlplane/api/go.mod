module cisco-app-networking.github.io/networkservicemesh/controlplane/api

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/pkg/errors v0.8.1
	google.golang.org/grpc v1.27.0
)

replace (
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	github.com/networkservicemesh/networkservicemesh/controlplane/api => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/clusterinfo => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/cls => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/common => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/kernel => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/memif => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/srv6 => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/vxlan => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/wireguard => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/connectioncontext => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/crossconnect => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/nsmdapi => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api/registry => ./

)
