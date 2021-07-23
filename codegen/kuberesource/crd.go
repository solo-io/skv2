package kuberesource

import (
	"fmt"
	"strings"

	errors "github.com/rotisserie/eris"

	"github.com/Masterminds/semver"
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/skv2/codegen/util/stringutils"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/codegen/model"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

type Options struct {
	WithoutSpecHash bool
}

type Option func(*Options)

func WithoutSpecHash() Option {
	return func(o *Options) {
		o.WithoutSpecHash = true
	}
}

// Create CRDs for a group
func CustomResourceDefinitions(
	group model.Group,
	opts ...Option,
) (objects []metav1.Object, err error) {
	for _, resource := range group.Resources {

		var validationSchema *apiextv1beta1.CustomResourceValidation
		if group.RenderValidationSchemas {
			validationSchema, err = constructValidationSchema(resource, group.OpenApiSchemas)
			if err != nil {
				return nil, err
			}
		}

		objects = append(objects, CustomResourceDefinition(resource, validationSchema, opts...))
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

func getJsonSchema(schemaName string, schemas map[string]*apiextv1beta1.JSONSchemaProps) (*apiextv1beta1.JSONSchemaProps, error) {

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
	opts ...Option,
) *apiextv1beta1.CustomResourceDefinition {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	group := resource.Group.Group
	version := resource.Group.Version
	kind := resource.Kind
	kindLowerPlural := strings.ToLower(stringutils.Pluralize(kind))
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
				Plural:     kindLowerPlural,
				Singular:   kindLower,
				Kind:       kind,
				ShortNames: resource.ShortNames,
				ListKind:   kind + "List",
			},
		},
	}
	if !options.WithoutSpecHash {
		// hashstructure of the crd spec:
		specHash, err := hashstructure.Hash(crd.Spec, nil)
		if err == nil {
			crd.Annotations = map[string]string{
				model.CRDSpecHashKey: fmt.Sprintf("%x", specHash),
			}
		}
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

type ErrMap map[string]error
type CrdNeedsUpgrade struct {
	CRDName string
}

func (e *CrdNeedsUpgrade) Error() string {
	return fmt.Sprintf("CRD %s needs to be upgraded", e.CRDName)
}

type CrdNotFound struct {
	CRDName string
}

func (e *CrdNotFound) Error() string {
	return fmt.Sprintf("CRD %s not found. Is it unused?", e.CRDName)
}

func DoCrdsNeedUpgrade(newProdCrdInfo model.CRDMetadata, ourCrds []apiextv1beta1.CustomResourceDefinition) ErrMap {
	newProductVersion := newProdCrdInfo.Version
	crdmap := make(map[string]string)
	for _, crd := range newProdCrdInfo.CRDS {
		crdmap[crd.Name] = crd.Hash
	}
	ret := ErrMap{}
	for _, ourCrd := range ourCrds {
		if hash, ok := crdmap[ourCrd.Name]; !ok {
			ret[ourCrd.Name] = &CrdNotFound{CRDName: ourCrd.Name}
			continue
		} else {
			needUpgrade, err := DoesCrdNeedUpgrade(newProducutVersion, hash, ourCrd.Annotations)

			if err != nil {
				ret[ourCrd.Name] = err
			} else if needUpgrade {
				ret[ourCrd.Name] = &CrdNeedsUpgrade{CRDName: ourCrd.Name}
			}
		}
	}
	if len(ret) == 0 {
		return nil
	}
	return ret
}

func DoesCrdNeedUpgrade(newProductVersion, newCrdHash string, deployedCrdAnnotations map[string]string) (bool, error) {

	if newProductVersion == "" || newCrdHash == "" {
		return false, errors.New(fmt.Sprintf("Cannot determine if CRDs need an upgrade, missing internal data: %s %s", newProductVersion, newCrdHash))
	}

	crdVersion, ok := deployedCrdAnnotations[model.CRDVersionKey]
	if !ok {
		return false, errors.New(fmt.Sprintf("Cannot determine crd product version from CRD annotations: %v", deployedCrdAnnotations))
	}
	crdSpecHash, ok := deployedCrdAnnotations[model.CRDSpecHashKey]
	if !ok {
		return false, errors.New(fmt.Sprintf("Cannot determine crd spec hash from CRD annotations: %v", deployedCrdAnnotations))
	}

	if newCrdHash == crdSpecHash {
		return false, nil
	}

	// parse semver of the current product version
	newProductVersionSemver, err := semver.NewVersion(newProductVersion)
	if err != nil {
		return false, errors.Wrapf(err, "Cannot parse current product version: %s", newProductVersion)
	}

	currentCrdVersionSemver, err := semver.NewVersion(crdVersion)
	if err != nil {
		return false, errors.Wrapf(err, "Cannot parse current crd version: %s", crdVersion)
	}

	// If the current product version is greater than the crd version, the CRD needs to be upgraded.
	if currentCrdVersionSemvar.Compare(newProductVersionSemver) <= 0 {
		return true, nil
	}
	return false, nil
}
