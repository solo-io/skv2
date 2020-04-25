module github.com/solo-io/skv2

go 1.13

require (
	cloud.google.com/go v0.50.0 // indirect
	github.com/Azure/go-autorest/autorest v0.9.4 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/sprig/v3 v3.0.0
	github.com/avast/retry-go v2.4.3+incompatible
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gertd/go-pluralize v0.1.1
	github.com/go-logr/logr v0.1.0
	github.com/go-openapi/spec v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.6 // indirect
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4-0.20200406172829-6d816de489c1
	github.com/golang/protobuf v1.3.2
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/gophercloud/gophercloud v0.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/huandu/xstrings v1.2.1 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.8.1
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/rogpeppe/go-internal v1.5.2
	github.com/rotisserie/eris v0.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/solo-io/anyvendor v0.0.1
	github.com/solo-io/go-utils v0.15.2
	github.com/solo-io/protoc-gen-ext v0.0.7
	github.com/solo-io/solo-kit v0.13.3
	github.com/spf13/cast v1.3.1 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/sys v0.0.0-20200120151820-655fe14d7479 // indirect
	golang.org/x/tools v0.0.0-20200422205258-72e4a01eba43
	k8s.io/api v0.17.4
	k8s.io/apiextensions-apiserver v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.17.4
	k8s.io/gengo v0.0.0-20200114144118-36b2048a9120 // indirect
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c // indirect
	sigs.k8s.io/controller-runtime v0.5.1
	sigs.k8s.io/yaml v1.2.0
)

// Pinned to kubernetes-1.14.1
replace k8s.io/kubernetes => k8s.io/kubernetes v1.14.1

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2

	k8s.io/client-go => k8s.io/client-go v0.17.1
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
