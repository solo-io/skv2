package main

import (
	"log"
	"path/filepath"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//go:generate go run generated.go

const (
	generatedPackageName = "pkg/generated/contrib"
)

var (
	kubeGeneratedPackage = filepath.Join(generatedPackageName, "kubernetes")
)

func main() {
	log.Println("starting contrib resource generation")
	skv2Cmd := codegen.Command{
		Groups: []model.Group{
			{
				GroupVersion: schema.GroupVersion{
					Group:   "core",
					Version: "v1",
				},
				Module: "k8s.io/api",
				Resources: []model.Resource{
					{
						Kind: "Service",
					},
					{
						Kind: "Pod",
					},
				},
				RenderContrib: model.RenderContrib{
					Sets:        true,
					SetMatchers: true,
				},
				CustomTypesImportPath: "k8s.io/api/core/v1",
				ApiRoot:               kubeGeneratedPackage,
			},
		},
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating contrib resources\n")
}
