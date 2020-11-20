module github.com/cisco-app-networking/networkservicemesh/pkg

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/networkservicemesh/networkservicemesh/pkg v0.0.0-00010101000000-000000000000
	github.com/onsi/gomega v1.10.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spiffe/go-spiffe v0.0.0-20191104192205-d29ac0a1ba99
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	google.golang.org/grpc v1.27.0
)

replace (
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
	github.com/networkservicemesh/networkservicemesh/pkg => ./
	github.com/uber-go/atomic => go.uber.org/atomic v1.6.0
)
