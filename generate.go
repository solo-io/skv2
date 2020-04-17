package main

import (
	"log"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//go:generate go run generate.go

func main() {
	protoImports := sk_anyvendor.CreateDefaultMatchOptions([]string{
		"api/**/*.proto",
		"pkg/**/*.proto",
	})

	// custom proto imports
	protoImports.External["k8s.io/api"] = []string{
		"core/v1/*.proto",
	}
	protoImports.External["k8s.io/apimachinery"] = []string{
		"pkg/**/*.proto",
	}

	cmd := &codegen.Command{
		AppName:         "dev-portal",
		AnyVendorConfig: protoImports,
		Groups: []model.Group{
			{
				GroupVersion: schema.GroupVersion{
					Group:   "core",
					Version: "v1",
				},
				Module: "k8s.io/api",
				Resources: []model.Resource{
					{
						Kind: "Secret",
					},
					{
						Kind: "ConfigMap",
					},
				},
				RenderClients:         true,
				RenderController:      true,
				CustomTypesImportPath: "k8s.io/api/core/v1",
				ApiRoot:               "pkg/api/kube",
			},
		},
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
