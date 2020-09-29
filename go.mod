module github.com/solo-io/skv2

go 1.13

require (
	cloud.google.com/go v0.40.0
	cuelang.org/go v0.2.2
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/sprig/v3 v3.1.0
	github.com/avast/retry-go v2.2.0+incompatible
	github.com/aws/aws-sdk-go v1.30.15
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gertd/go-pluralize v0.1.1
	github.com/go-logr/logr v0.1.0
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.2
	github.com/gophercloud/gophercloud v0.2.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1
	github.com/rogpeppe/go-internal v1.6.0
	github.com/rotisserie/eris v0.1.1
	github.com/sirupsen/logrus v1.6.0
	github.com/solo-io/anyvendor v0.0.1
	github.com/solo-io/go-utils v0.17.0
	github.com/solo-io/protoc-gen-ext v0.0.9
	go.uber.org/zap v1.13.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/tools v0.0.0-20200811153730-74512f09e4b0
	google.golang.org/api v0.6.0
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc v1.28.1 // indirect
	google.golang.org/protobuf v1.23.0
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/code-generator v0.18.6
	sigs.k8s.io/aws-iam-authenticator v0.5.0
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
)

replace (
	// Indirect operator-sdk dependencies use git.apache.org, which is frequently
	// down. The github mirror should be used instead.
	// Locking to a specific version (from 'go mod graph'):
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/operator-framework/operator-lifecycle-manager => github.com/operator-framework/operator-lifecycle-manager v0.0.0-20190605231540-b8a4faf68e36
)

// Remove when controller-tools v0.2.2 is released
// Required for the bugfix https://github.com/kubernetes-sigs/controller-tools/pull/322
replace sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.2.2-0.20190919011008-6ed4ff330711

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
