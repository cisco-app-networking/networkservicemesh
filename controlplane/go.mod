module github.com/networkservicemesh/networkservicemesh/controlplane

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/networkservicemesh/networkservicemesh/controlplane/api v1.0.0
	github.com/networkservicemesh/networkservicemesh/forwarder/api v1.0.0
	github.com/networkservicemesh/networkservicemesh/pkg v1.0.0
	github.com/networkservicemesh/networkservicemesh/sdk v1.0.3
	github.com/networkservicemesh/networkservicemesh/utils v1.0.0
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
	github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
)
