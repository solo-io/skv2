package kuberesource

import (
	"fmt"
	"sort"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/skv2/codegen/util/stringutils"
	"github.com/solo-io/skv2/pkg/crdutils"
	"golang.org/x/exp/maps"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/codegen/model"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create CRDs for a group
func CustomResourceDefinitions(
	groups []*model.Group,
) (objects []apiextv1.CustomResourceDefinition, err error) {
	resourcesByKind := make(map[string][]model.Resource)
	skipHashByKind := make(map[string]bool)
	for _, group := range groups {
		for i, resource := range group.Resources {
			if resource.CodegenOnly {
				continue
			}
			resourcesByKind[resource.Kind] = append(resourcesByKind[resource.Kind], group.Resources[i])
			skipHashByKind[resource.Kind] = skipHashByKind[resource.Kind] || resource.Group.SkipSpecHash
		}
	}

	// Make ordering of crds in a group deterministic
	kinds := maps.Keys(resourcesByKind)
	sort.Strings(kinds)
	for _, kind := range kinds {
		resources := resourcesByKind[kind]
		// make version ordering deterministic
		sort.Slice(resources, func(i, j int) bool { return resources[i].Version < resources[j].Version })
		validationSchemas, err := constructValidationSchemas(resources)
		if err != nil {
			return nil, err
		}

		crd, err := CustomResourceDefinition(resources, validationSchemas, skipHashByKind[kind])
		if err != nil {
			return nil, err
		}
		objects = append(objects, *crd)
	}
	return objects, nil
}

func constructValidationSchemas(resources []model.Resource) (map[string]*apiextv1.CustomResourceValidation, error) {
	validationSchemas := make(map[string]*apiextv1.CustomResourceValidation)
	for _, resource := range resources {
		var validationSchema *apiextv1.CustomResourceValidation
		validationSchema, err := constructValidationSchema(
			resource.Group.RenderValidationSchemas,
			resource,
			resource.Group.OpenApiSchemas,
		)
		if err != nil {
			return nil, err
		}

		validationSchemas[resource.Group.String()] = validationSchema
	}
	return validationSchemas, nil
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

func validateCRDResources(resources []model.Resource) error {
	if len(resources) < 2 {
		return nil
	}

	scope := resources[0].ClusterScoped
	shortNames := resources[0].ShortNames
	categories := resources[0].Categories

	for i := 1; i < len(resources); i++ {
		if resources[i].ClusterScoped != scope {
			return fmt.Errorf("mismatched 'currentScope' in versions of CRD for resource kind %s", resources[i].Kind)
		}
		if !cmp.Equal(resources[i].ShortNames, shortNames) {
			return fmt.Errorf("mismatched 'ShortNames' in versions of CRD for resource kind %s", resources[i].Kind)
		}
		if !cmp.Equal(resources[i].Categories, categories) {
			return fmt.Errorf("mismatched 'Categories' in versions of CRD for resource kind %s", resources[i].Kind)
		}
	}

	return nil
}

func CustomResourceDefinition(
	resources []model.Resource,
	validationSchemas map[string]*apiextv1.CustomResourceValidation,
	withoutSpecHash bool,
) (*apiextv1.CustomResourceDefinition, error) {

	err := validateCRDResources(resources)
	if err != nil {
		return nil, err
	}

	group := resources[0].Group.Group
	kind := resources[0].Kind
	kindLowerPlural := strings.ToLower(stringutils.Pluralize(kind))
	kindLower := strings.ToLower(kind)

	scope := apiextv1.NamespaceScoped
	if resources[0].ClusterScoped {
		scope = apiextv1.ClusterScoped
	}

	versions := make([]apiextv1.CustomResourceDefinitionVersion, 0, len(resources))
	for _, resource := range resources {
		var status *apiextv1.CustomResourceSubresourceStatus
		if resource.Status != nil {
			status = &apiextv1.CustomResourceSubresourceStatus{}
		}

		v := apiextv1.CustomResourceDefinitionVersion{
			Name:                     resource.Group.Version,
			Served:                   true,
			Storage:                  resource.Stored,
			Deprecated:               resource.Deprecated,
			AdditionalPrinterColumns: resource.AdditionalPrinterColumns,
			Subresources: &apiextv1.CustomResourceSubresources{
				Status: status,
			},
			Schema: validationSchemas[resource.Group.String()],
		}
		versions = append(versions, v)
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
			Group:    group,
			Scope:    scope,
			Versions: versions,
			Names: apiextv1.CustomResourceDefinitionNames{
				Plural:     kindLowerPlural,
				Singular:   kindLower,
				Kind:       kind,
				ShortNames: resources[0].ShortNames,
				ListKind:   kind + "List",
				Categories: resources[0].Categories,
			},
		},
	}
	if !withoutSpecHash {
		// hashstructure of the crd spec:
		specHash, err := hashstructure.Hash(crd.Spec, nil)
		if err != nil {
			panic(err)
		}
		crd.Annotations = map[string]string{
			crdutils.CRDSpecHashKey: fmt.Sprintf("%x", specHash),
		}
	}

	if len(validationSchemas) > 0 {
		// Setting PreserveUnknownFields to false ensures that objects with unknown fields are rejected.
		crd.Spec.PreserveUnknownFields = false
	}
	return crd, nil
}
