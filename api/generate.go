package main

import (
	"log"

	"github.com/solo-io/skv2/api/k8s"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//go:generate go run generate.go

func main() {
	log.Println("starting kube client generation")

	skv2Imports := sk_anyvendor.CreateDefaultMatchOptions([]string{
		"api/**/*.proto",
	})

	groups := []model.Group{
		{
			GroupVersion: schema.GroupVersion{
				Group:   "multicluster.solo.io",
				Version: "v1alpha1",
			},
			Module: "github.com/solo-io/skv2",
			Resources: []model.Resource{
				{
					Kind: "KubernetesCluster",
					Spec: model.Field{
						Type: model.Type{
							Name: "KubernetesClusterSpec",
						},
					},
					Status: &model.Field{
						Type: model.Type{
							Name: "KubernetesClusterStatus",
						},
					},
				},
			},
			RenderManifests:  true,
			RenderController: true,
			RenderClients:    true,
			RenderTypes:      true,
			ApiRoot:          "pkg/api",
			CustomTemplates:  contrib.AllCustomTemplates,
		},
	}

	// add internal k8s groups we depend on
	groups = append(groups, k8s.Groups()...)

	skv2Cmd := codegen.Command{
		Groups:          groups,
		AnyVendorConfig: skv2Imports,
		RenderProtos:    true,
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating kube clients\n")
}
