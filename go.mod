module github.com/bigbasket/k8s-cloudwatch-adapter

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.5
	github.com/awslabs/k8s-cloudwatch-adapter v0.9.0
	github.com/bigbasket/k8s-custom-hpa v0.7.2
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c // indirect
	github.com/juju/ratelimit v1.0.1 // indirect
	github.com/kubernetes-incubator/custom-metrics-apiserver v0.0.0-20200323093244-5046ce1afe6b
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/apimachinery v0.17.7
	k8s.io/apiserver v0.17.7 // indirect
	k8s.io/client-go v0.17.7
	k8s.io/code-generator v0.17.7
	k8s.io/component-base v0.17.7
	k8s.io/klog v1.0.0
	k8s.io/metrics v0.17.7
)
