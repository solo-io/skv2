package kuberesource

import (
	"encoding/json"
	"fmt"
	"strings"

	"cuelang.org/go/encoding/openapi"
	"github.com/gertd/go-pluralize"
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/codegen/model"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

// Create CRDs for a group
func CustomResourceDefinitions(group model.Group) (objects []metav1.Object, err error) {
	for _, resource := range group.Resources {

		var validationSchema *apiextv1beta1.CustomResourceValidation
		if group.RenderValidationSchemas {
			validationSchema, err = constructValidationSchema(resource, group.OpenApiSchemas)
			if err != nil {
				return nil, err
			}
		}

		objects = append(objects, CustomResourceDefinition(resource, validationSchema))
	}
	return objects, nil
}

func constructValidationSchema(resource model.Resource, oapiSchemas model.OpenApiSchemas) (*apiextv1beta1.CustomResourceValidation, error) {
	validationSchema := &apiextv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextv1beta1.JSONSchemaProps{
			Type:       "object",
			Properties: map[string]apiextv1beta1.JSONSchemaProps{},
		},
	}

	// Spec validation schema
	specSchema, err := getJsonSchema(resource, resource.Spec.Type.Name, oapiSchemas)
	if err != nil {
		return nil, eris.Wrapf(err, "constructing spec validation schema for Kind %s", resource.Kind)
	}
	validationSchema.OpenAPIV3Schema.Properties["spec"] = *specSchema

	// Status validation schema
	if resource.Status != nil {
		statusSchema, err := getJsonSchema(resource, resource.Status.Type.Name, oapiSchemas)
		if err != nil {
			return nil, eris.Wrapf(err, "constructing status validation schema for Kind %s", resource.Kind)
		}
		validationSchema.OpenAPIV3Schema.Properties["status"] = *statusSchema
	}

	return validationSchema, nil
}

func getJsonSchema(resource model.Resource, schemaName string, schemas map[string]*openapi.OrderedMap) (*apiextv1beta1.JSONSchemaProps, error) {

	schema, ok := schemas[schemaName]
	if !ok {
		return nil, eris.Errorf("Could not find open api schema for %s", schemaName)
	}

	byt, err := schema.MarshalJSON()
	if err != nil {
		return nil, eris.Errorf("Cannot marshal OpenAPI schema for %v: %v", resource.Group.Group, err)
	}

	jsonSchema := &apiextv1beta1.JSONSchemaProps{}
	if err = json.Unmarshal(byt, jsonSchema); err != nil {
		return nil, eris.Errorf("Cannot unmarshal raw OpenAPI schema to JSONSchemaProps for %v: %v", resource.Group.Group, err)
	}

	if err = validateStructural(jsonSchema); err != nil {
		return nil, err
	}

	return jsonSchema, nil
}

// Lifted from https://github.com/istio/tools/blob/477454adf7995dd3070129998495cdc8aaec5aff/cmd/cue-gen/crd.go#L108
func validateStructural(s *apiextv1beta1.JSONSchemaProps) error {
	out := &apiext.JSONSchemaProps{}
	if err := apiextv1beta1.Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(s, out, nil); err != nil {
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
	validationSchema *apiextv1beta1.CustomResourceValidation,
) *apiextv1beta1.CustomResourceDefinition {

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
	if validationSchema != nil {
		// Setting PreserveUnknownFields to false ensures that objects with unknown fields are rejected.
		crd.Spec.PreserveUnknownFields = pointer.BoolPtr(false)

		// TODO move this block into versions once we support multiple versions of a CRD
		// Including a validation schema inside the Version field when there is only a single version present yields the error:
		// "per-version schemas may not all be set to identical values (top-level validation should be used instead"
		crd.Spec.Validation = validationSchema
	}
	return crd
}
