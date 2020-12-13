module cisco-app-networking.github.io/networkservicemesh/controlplane

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/onsi/gomega v1.10.3
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20201006153459-a7d1128ccaa0
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20200114203027-fcfc50b29cbb
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	github.com/networkservicemesh/networkservicemesh => ../
	github.com/networkservicemesh/networkservicemesh/controlplane => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api => cisco-app-networking.github.io/networkservicemesh/controlplane/api latest
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/api/nsm => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/common => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/local => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/metrics => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/monitor/remote => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nseregistry => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nsm => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nsmd => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/properties => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/remote => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/remote/proxy_network_service_server => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/selector => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/serviceregistry => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/services => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/sid => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/tests => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/tests/utils => ./pkg
	github.com/networkservicemesh/networkservicemesh/controlplane/pkg/vni => ./pkg
	github.com/networkservicemesh/networkservicemesh/forwarder => ../forwarder
	github.com/networkservicemesh/networkservicemesh/forwarder/api => ../forwarder/api
	github.com/networkservicemesh/networkservicemesh/k8s/pkg/apis => ../k8s/pkg/apis
	github.com/networkservicemesh/networkservicemesh/pkg => ../pkg
	github.com/networkservicemesh/networkservicemesh/sdk => ../sdk
	github.com/networkservicemesh/networkservicemesh/side-cars => ../side-cars
	github.com/networkservicemesh/networkservicemesh/utils => ../utils
)
