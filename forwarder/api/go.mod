module cisco-app-networking.github.io/networkservicemesh/forwarder/api

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/networkservicemesh/networkservicemesh/controlplane/api v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.27.0
)

replace (
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	github.com/networkservicemesh/networkservicemesh/controlplane/api => ../../controlplane/api
	github.com/networkservicemesh/networkservicemesh/controlplane/api v0.0.0-00010101000000-000000000000 => ../../controlplane/api
	github.com/networkservicemesh/networkservicemesh/forwarder/api => ./
	github.com/networkservicemesh/networkservicemesh/pkg => ../../pkg
	github.com/networkservicemesh/networkservicemesh/utils => ../../utils
)
