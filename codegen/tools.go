//+build tools

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/onsi/ginkgo/ginkgo"
	_ "github.com/solo-io/protoc-gen-ext"
	_ "golang.org/x/tools/cmd/goimports"
)
