module github.com/bb-rajdeep/k8s-cloudwatch-adapter

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.5
	github.com/kubernetes-incubator/custom-metrics-apiserver v0.0.0-20200323093244-5046ce1afe6b
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/apimachinery v0.17.7
	k8s.io/apiserver v0.17.7 // indirect
	k8s.io/client-go v0.17.7
	k8s.io/code-generator v0.17.7
	k8s.io/component-base v0.17.7
	k8s.io/klog v1.0.0
	k8s.io/metrics v0.17.7
	github.com/bigbasket/k8s-custom-hpa v0.7.2
)
