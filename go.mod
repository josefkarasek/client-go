module github.com/openshift/client-go

go 1.16

require (
	github.com/openshift/api v0.0.0-20211209135129-c58d9f695577
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	k8s.io/code-generator v0.23.0
)

replace github.com/openshift/api => github.com/josefkarasek/api-1 v0.0.0-20220208141251-8deb1762a41c
