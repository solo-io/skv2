package kuberesource

import (
	"fmt"
	"strings"

	"github.com/gertd/go-pluralize"

	"github.com/solo-io/skv2/codegen/model"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create CRDs for a group
func CustomResourceDefinitions(group model.Group) []metav1.Object {
	var objects []metav1.Object
	for _, resource := range group.Resources {
		objects = append(objects, CustomResourceDefinition(resource))
	}
	return objects
}

func CustomResourceDefinition(resource model.Resource) *apiextv1beta1.CustomResourceDefinition {
	group := resource.Group.Group
	version := resource.Group.Version
	kind := resource.Kind
	kindLowerPlural := strings.ToLower(pluralize.NewClient().Plural(kind))
	kindLower := strings.ToLower(kind)

	var status *apiextv1beta1.CustomResourceSubresourceStatus
	if resource.Status != nil {
		status = &apiextv1beta1.CustomResourceSubresourceStatus{}
	}

	scope := apiextv1beta1.NamespaceScoped
	if resource.ClusterScoped {
		scope = apiextv1beta1.ClusterScoped
	}

	crd := &apiextv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s.%s", kindLowerPlural, group),
		},
		Spec: apiextv1beta1.CustomResourceDefinitionSpec{
			Group: group,
			Scope: scope,
			Versions: []apiextv1beta1.CustomResourceDefinitionVersion{{
				Name:                     version,
				Served:                   true,
				Storage:                  true,
				AdditionalPrinterColumns: resource.AdditionalPrinterColumns,
			}},
			Subresources: &apiextv1beta1.CustomResourceSubresources{
				Status: status,
			},
			Names: apiextv1beta1.CustomResourceDefinitionNames{
				Plural:   kindLowerPlural,
				Singular: kindLower,
				Kind:     kind,
				ListKind: kind + "List",
			},
		},
	}
	return crd
}
