package main

import (
	"log"

	"github.com/solo-io/skv2/api/k8s"
	"github.com/solo-io/skv2/api/multicluster/v1alpha1"
	"github.com/solo-io/skv2/codegen/skv2_anyvendor"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
)

func main() {
	log.Println("starting kube client generation")

	skv2Imports := skv2_anyvendor.CreateDefaultMatchOptions([]string{
		"api/**/*.proto",
	})

	groups := []model.Group{
		v1alpha1.Group,
	}

	// add internal k8s groups we depend on
	groups = append(groups, k8s.Groups()...)

	skv2Cmd := codegen.Command{
		AppName:         "skv2",
		Groups:          groups,
		AnyVendorConfig: skv2Imports,
		RenderProtos:    true,
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating kube clients\n")
}
