package v1alpha1

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// export group
var Group = model.Group{
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
			Stored: true,
		},
	},
	RenderManifests:           true,
	RenderValidationSchemas:   true,
	RenderController:          true,
	RenderClients:             true,
	RenderTypes:               true,
	MockgenDirective:          true,
	ApiRoot:                   "pkg/api",
	CustomTemplates:           contrib.AllGroupCustomTemplates,
	SkipConditionalCRDLoading: true,
}
