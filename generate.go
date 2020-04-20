package main

import (
	"log"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//go:generate go run generate.go

func main() {
	cmd := &codegen.Command{
		AppName: "skv2",
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
