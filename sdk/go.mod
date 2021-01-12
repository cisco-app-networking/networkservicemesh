module cisco-app-networking.github.io/networkservicemesh/sdk

go 1.13

require (
	cisco-app-networking.github.io/networkservicemesh/controlplane v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/controlplane/api v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/pkg v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/utils v1.0.13-vanity
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/go-multierror v1.0.0
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.5.0
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	go.ligato.io/vpp-agent/v3 v3.2.0
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
)
