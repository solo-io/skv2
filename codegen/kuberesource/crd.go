package kuberesource

import (
	"fmt"
	"strings"

	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/skv2/codegen/util/stringutils"
	"github.com/solo-io/skv2/pkg/crdutils"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/codegen/model"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create CRDs for a group
func CustomResourceDefinitions(
	group model.Group,
) (objects []metav1.Object, err error) {
	for _, resource := range group.Resources {

		var validationSchema *apiextv1.CustomResourceValidation
		validationSchema, err = constructValidationSchema(
			group.RenderValidationSchemas,
			resource,
			group.OpenApiSchemas,
		)
		if err != nil {
			return nil, err
		}

		objects = append(objects, CustomResourceDefinition(resource, validationSchema, group.SkipSpecHash))
	}
	return objects, nil
}

func constructValidationSchema(
	renderValidationSchema bool,
	resource model.Resource,
	oapiSchemas model.OpenApiSchemas,
) (*apiextv1.CustomResourceValidation, error) {
	// Even if we do not want to render validation schemas, we should include
	// the top level schema definition and preserve unknown fields since Helm
	// requires that some sort of schema is defined
	if !renderValidationSchema {
		preserveUnknownFields := true
		return &apiextv1.CustomResourceValidation{
			OpenAPIV3Schema: &apiextv1.JSONSchemaProps{
				Type:                   "object",
				XPreserveUnknownFields: &preserveUnknownFields,
			},
		}, nil
	}
	validationSchema := &apiextv1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextv1.JSONSchemaProps{
			Type:       "object",
			Properties: map[string]apiextv1.JSONSchemaProps{},
		},
	}

	// Spec validation schema
	specSchema, err := getJsonSchema(resource.Spec.Type.Name, oapiSchemas)
	if err != nil {
		return nil, eris.Wrapf(err, "constructing spec validation schema for Kind %s", resource.Kind)
	}
	validationSchema.OpenAPIV3Schema.Properties["spec"] = *specSchema

	// Status validation schema
	if resource.Status != nil {
		statusSchema, err := getJsonSchema(resource.Status.Type.Name, oapiSchemas)
		if err != nil {
			return nil, eris.Wrapf(err, "constructing status validation schema for Kind %s", resource.Kind)
		}
		validationSchema.OpenAPIV3Schema.Properties["status"] = *statusSchema
	}

	return validationSchema, nil
}

func getJsonSchema(
	schemaName string,
	schemas map[string]*apiextv1.JSONSchemaProps,
) (*apiextv1.JSONSchemaProps, error) {

	schema, ok := schemas[schemaName]
	if !ok {
		return nil, eris.Errorf("Could not find open api schema for %s", schemaName)
	}

	if err := validateStructural(schema); err != nil {
		return nil, err
	}

	return schema, nil
}

// Lifted from https://github.com/istio/tools/blob/477454adf7995dd3070129998495cdc8aaec5aff/cmd/cue-gen/crd.go#L108
func validateStructural(s *apiextv1.JSONSchemaProps) error {
	out := &apiext.JSONSchemaProps{}
	if err := apiextv1.Convert_v1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(s, out, nil); err != nil {
		return fmt.Errorf("cannot convert v1beta1 JSONSchemaProps to JSONSchemaProps: %v", err)
	}

	r, err := structuralschema.NewStructural(out)
	if err != nil {
		return fmt.Errorf("cannot convert to a structural schema: %v", err)
	}

	if errs := structuralschema.ValidateStructural(nil, r); len(errs) != 0 {
		return fmt.Errorf("schema is not structural: %v", errs.ToAggregate().Error())
	}

	return nil
}

func CustomResourceDefinition(
	resource model.Resource,
	validationSchema *apiextv1.CustomResourceValidation,
	withoutSpecHash bool,
) *apiextv1.CustomResourceDefinition {

	group := resource.Group.Group
	version := resource.Group.Version
	kind := resource.Kind
	kindLowerPlural := strings.ToLower(stringutils.Pluralize(kind))
	kindLower := strings.ToLower(kind)

	var status *apiextv1.CustomResourceSubresourceStatus
	if resource.Status != nil {
		status = &apiextv1.CustomResourceSubresourceStatus{}
	}

	scope := apiextv1.NamespaceScoped
	if resource.ClusterScoped {
		scope = apiextv1.ClusterScoped
	}

	crd := &apiextv1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextv1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s.%s", kindLowerPlural, group),
		},
		Spec: apiextv1.CustomResourceDefinitionSpec{
			Group: group,
			Scope: scope,
			Versions: []apiextv1.CustomResourceDefinitionVersion{
				{
					Name:                     version,
					Served:                   true,
					Storage:                  true,
					AdditionalPrinterColumns: resource.AdditionalPrinterColumns,
					Subresources: &apiextv1.CustomResourceSubresources{
						Status: status,
					},
					Schema: validationSchema,
				},
			},
			Names: apiextv1.CustomResourceDefinitionNames{
				Plural:     kindLowerPlural,
				Singular:   kindLower,
				Kind:       kind,
				ShortNames: resource.ShortNames,
				ListKind:   kind + "List",
				Categories: resource.Categories,
			},
		},
	}
	if !withoutSpecHash {
		// hashstructure of the crd spec:
		specHash, err := hashstructure.Hash(crd.Spec, nil)
		if err != nil {
			panic(err)
			return nil
		}
		crd.Annotations = map[string]string{
			crdutils.CRDSpecHashKey: fmt.Sprintf("%x", specHash),
		}
	}

	if validationSchema != nil {
		// Setting PreserveUnknownFields to false ensures that objects with unknown fields are rejected.
		crd.Spec.PreserveUnknownFields = false
	}
	return crd
}
