module github.com/solo-io/skv2

go 1.13

require (
	cloud.google.com/go v0.76.0
	cuelang.org/go v0.2.2
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/sprig/v3 v3.1.0
	github.com/avast/retry-go v2.2.0+incompatible
	github.com/aws/aws-sdk-go v1.30.15
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/gertd/go-pluralize v0.1.1
	github.com/go-logr/logr v0.4.0
	github.com/go-logr/zapr v0.2.0
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/packr v1.30.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0
	github.com/iancoleman/strcase v0.1.3
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.5
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/pseudomuto/protoc-gen-doc v1.4.1
	github.com/rogpeppe/go-internal v1.6.0
	github.com/rotisserie/eris v0.1.1
	github.com/sirupsen/logrus v1.6.0
	github.com/solo-io/anyvendor v0.0.3
	github.com/solo-io/go-list-licenses v0.0.4
	github.com/solo-io/go-utils v0.21.4
	github.com/solo-io/k8s-utils v0.0.1
	github.com/solo-io/protoc-gen-ext v0.0.13
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0 // indirect
	go.opencensus.io v0.22.6 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20210201163806-010130855d6c
	golang.org/x/tools v0.1.0
	google.golang.org/api v0.38.0
	google.golang.org/genproto v0.0.0-20210218151259-fe80b386bf06
	google.golang.org/protobuf v1.25.0
	k8s.io/api v0.19.6
	k8s.io/apiextensions-apiserver v0.19.6
	k8s.io/apimachinery v0.19.6
	k8s.io/client-go v0.19.6
	k8s.io/code-generator v0.19.6
	k8s.io/klog/v2 v2.5.0
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	sigs.k8s.io/aws-iam-authenticator v0.5.0
	sigs.k8s.io/controller-runtime v0.7.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	// Indirect operator-sdk dependencies use git.apache.org, which is frequently
	// down. The github mirror should be used instead.
	// Locking to a specific version (from 'go mod graph'):
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/operator-framework/operator-lifecycle-manager => github.com/operator-framework/operator-lifecycle-manager v0.0.0-20190605231540-b8a4faf68e36
)
