module cisco-app-networking.github.io/networkservicemesh/k8s

go 1.15

require (
	cisco-app-networking.github.io/networkservicemesh/controlplane v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/controlplane/api v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/k8s/pkg/apis v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/pkg v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/sdk v1.0.13-vanity
	cisco-app-networking.github.io/networkservicemesh/utils v1.0.13-vanity
	github.com/golang/protobuf v1.4.2
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20201006153459-a7d1128ccaa0
	google.golang.org/grpc v1.29.1
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.2-beta.0
	k8s.io/client-go v0.18.1
	k8s.io/kubelet v0.18.1
)

replace (
	github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8
	gonum.org/v1/gonum => github.com/gonum/gonum v0.0.0-20190331200053-3d26580ed485
	gonum.org/v1/netlib => github.com/gonum/netlib v0.0.0-20190331212654-76723241ea4e
	k8s.io/api => k8s.io/api v0.18.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.2-beta.0
	k8s.io/apiserver => k8s.io/apiserver v0.18.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.1
	k8s.io/client-go => k8s.io/client-go v0.18.1
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.1
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.1
	k8s.io/code-generator => k8s.io/code-generator v0.18.2-beta.0
	k8s.io/component-base => k8s.io/component-base v0.18.1
	k8s.io/cri-api => k8s.io/cri-api v0.18.2-beta.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.1
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.1
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.1
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.1
	k8s.io/kubectl => k8s.io/kubectl v0.18.1
	k8s.io/kubelet => k8s.io/kubelet v0.18.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.1
	k8s.io/metrics => k8s.io/metrics v0.18.1
	k8s.io/node-api => k8s.io/node-api v0.17.1
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.1
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.18.1
	k8s.io/sample-controller => k8s.io/sample-controller v0.18.1
)
