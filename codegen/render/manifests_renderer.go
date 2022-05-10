package render

import (
	"encoding/json"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/skv2/pkg/crdutils"

	protoutil "github.com/solo-io/skv2/codegen/proto"

	"github.com/rotisserie/eris"
	"github.com/solo-io/anyvendor/anyvendor"
	"github.com/solo-io/cue/cue"
	"github.com/solo-io/cue/encoding/openapi"
	"github.com/solo-io/cue/encoding/protobuf"
	"github.com/solo-io/go-utils/stringutils"
	"github.com/solo-io/skv2/codegen/kuberesource"
	"github.com/solo-io/skv2/codegen/model"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// creates a k8s resource for a group
// this gets turned into a k8s manifest file
type MakeResourceFunc func(group Group) ([]metav1.Object, error)

// renders kubernetes from templates
type ManifestsRenderer struct {
	AppName       string // used for labeling
	ResourceFuncs map[OutFile]MakeResourceFunc
	ManifestDir   string
	ProtoDir      string
}

func RenderManifests(
	appName, manifestDir, protoDir string,
	protoOpts protoutil.Options,
	grp Group,
) ([]OutFile, error) {
	defaultManifestsRenderer := ManifestsRenderer{
		AppName:     appName,
		ManifestDir: manifestDir,
		ProtoDir:    protoDir,
		ResourceFuncs: map[OutFile]MakeResourceFunc{
			{
				Path: manifestDir + "/crds/" + grp.Group + "_" + grp.Version + "_" + "crds.yaml",
			}: func(group Group) ([]metav1.Object, error) {
				return kuberesource.CustomResourceDefinitions(group)
			},
		},
	}
	return defaultManifestsRenderer.RenderManifests(grp, protoOpts)
}

func (r ManifestsRenderer) RenderManifests(grp Group, protoOpts protoutil.Options) ([]OutFile, error) {
	if !grp.RenderManifests {
		return nil, nil
	}

	if grp.RenderValidationSchemas {
		var err error
		oapiSchemas, err := generateOpenApi(grp, r.ProtoDir, protoOpts)
		if err != nil {
			return nil, err
		}
		grp.OpenApiSchemas = oapiSchemas
	}

	var renderedFiles []OutFile
	for out, mkFunc := range r.ResourceFuncs {
		content, err := r.renderManifest(r.AppName, mkFunc, grp)
		if err != nil {
			return nil, err
		}
		out.Content = content
		renderedFiles = append(renderedFiles, out)
	}
	return renderedFiles, nil
}

// Use cuelang as an intermediate language for transpiling protobuf schemas to openapi v3 with k8s structural schema constraints.
func generateOpenApi(grp model.Group, protoDir string, protoOpts protoutil.Options) (model.OpenApiSchemas, error) {
	if protoDir == "" {
		protoDir = anyvendor.DefaultDepDir
	}

	// Collect all protobuf definitions including transitive dependencies.
	var imports []string
	for _, fileDescriptor := range grp.Descriptors {
		imports = append(imports, fileDescriptor.Imports...)
	}
	imports = stringutils.Unique(imports)

	// Parse protobuf into cuelang
	cfg := &protobuf.Config{
		Root:   protoDir,
		Module: grp.Module,
		Paths:  imports,
	}

	ext := protobuf.NewExtractor(cfg)
	// collect the set of messsages for which validation is disabled
	for _, fileDescriptor := range grp.Descriptors {
		if err := ext.AddFile(fileDescriptor.ProtoFilePath, nil); err != nil {
			return nil, err
		}
	}

	instances, err := ext.Instances()
	if err != nil {
		return nil, err
	}

	// Convert cuelang to openapi
	unstructuredFieldsMap, err := getUnstructuredFieldsMap(grp, protoOpts)
	if err != nil {
		return nil, err
	}
	generator := &openapi.Generator{
		// k8s structural schemas do not allow $refs, i.e. all references must be expanded
		ExpandReferences:   true,
		UnstructuredFields: unstructuredFieldsMap,
	}
	if grp.SkipSchemaDescriptions {
		// returning empty from this func results in no description field being added
		generator.DescriptionFunc = func(_ cue.Value) string { return "" }
	}
	built := cue.Build(instances)
	for _, builtInstance := range built {
		// Avoid generating openapi for irrelevant proto imports.
		if !strings.HasSuffix(builtInstance.ImportPath, grp.Group+"/"+grp.Version) {
			continue
		}

		if err = builtInstance.Err; err != nil {
			return nil, eris.Errorf("Cue instance failed to build for %s: %+v", grp.Group, err)
		}
		if err = builtInstance.Value().Validate(); err != nil {
			return nil, eris.Errorf("Cue instance validation failed for %s: %+v", grp.Group, err)
		}
		schemas, err := generator.Schemas(builtInstance)
		if err != nil {
			return nil, eris.Errorf("Cue openapi generation failed for %s: %+v", grp.Group, err)
		}

		// Iterate openapi objects to construct mapping from proto message name to openapi schema
		oapiSchemas := model.OpenApiSchemas{}
		for _, kv := range schemas.Pairs() {
			orderedMap := kv.Value.(*openapi.OrderedMap)
			jsonSchema, err := postProcessValidationSchema(orderedMap)
			if err != nil {
				return nil, err
			}
			oapiSchemas[kv.Key] = jsonSchema
		}

		return oapiSchemas, err
	}
	return nil, nil
}

// returns the map of proto fields marked as unstructured, used by CUE to generate openapi schemas
func getUnstructuredFieldsMap(grp model.Group, opts protoutil.Options) (map[string][][]string, error) {

	unstructuredFields := map[string][][]string{}
	defaultProtoPkg := grp.Group
	defaultGoPkg := util.GoPackage(grp)
	for _, res := range grp.Resources {
		unstructuredSpecFields, err := opts.GetUnstructuredFields(
			ifDefined(res.Spec.Type.ProtoPackage, defaultProtoPkg),
			res.Spec.Type.Name,
		)
		if err != nil {
			return nil, err
		}
		goPkg := ifDefined(res.Spec.Type.GoPackage, defaultGoPkg)
		unstructuredFields[goPkg] = append(unstructuredFields[goPkg], unstructuredSpecFields...)
		if res.Status != nil {
			unstructuredStatusFields, err := opts.GetUnstructuredFields(
				ifDefined(res.Status.Type.ProtoPackage, defaultProtoPkg),
				res.Status.Type.Name,
			)
			if err != nil {
				return nil, err
			}
			goPkg := ifDefined(res.Status.Type.GoPackage, defaultGoPkg)
			if len(unstructuredStatusFields) > 0 {
				unstructuredFields[goPkg] = append(unstructuredFields[goPkg], unstructuredStatusFields...)
			}
		}
	}

	return unstructuredFields, nil
}

func SetVersionForObject(obj metav1.Object, version string) {
	if version == "" {
		return
	}
	a := obj.GetAnnotations()
	if a == nil {
		a = make(map[string]string)
	}
	// we only care about major minor and patch versions.
	if parsedSemVer, err := semver.NewVersion(version); err == nil {
		strippedVersion, err := parsedSemVer.SetPrerelease("")
		if err != nil {
			return
		}
		strippedVersion, err = strippedVersion.SetMetadata("")
		if err != nil {
			return
		}

		a[crdutils.CRDVersionKey] = strippedVersion.String()
		obj.SetAnnotations(a)
	}
}

func (r ManifestsRenderer) renderManifest(appName string, mk MakeResourceFunc, group Group) (string, error) {
	objs, err := mk(group)
	if err != nil {
		return "", err
	}

	var objManifests []string
	for _, obj := range objs {
		// find the annotation of the manifest, and add to them
		SetVersionForObject(obj, group.AddChartVersion)
		manifest, err := marshalObjToYaml(appName, obj)
		if err != nil {
			return "", err
		}
		objManifests = append(objManifests, manifest)
	}

	return strings.Join(objManifests, "\n---\n"), nil
}

func marshalObjToYaml(appName string, obj metav1.Object) (string, error) {
	labels := obj.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}

	labels["app"] = appName
	labels["app.kubernetes.io/name"] = appName

	obj.SetLabels(labels)

	yam, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}
	var v map[string]interface{}

	if err := yaml.Unmarshal(yam, &v); err != nil {
		return "", err
	}

	delete(v, "status")
	// why do we have to do this? Go problem???
	meta := v["metadata"].(map[string]interface{})

	delete(meta, "creationTimestamp")
	v["metadata"] = meta

	if spec, ok := v["spec"].(map[string]interface{}); ok {
		if template, ok := spec["template"].(map[string]interface{}); ok {
			if meta, ok := template["metadata"].(map[string]interface{}); ok {
				delete(meta, "creationTimestamp")
				template["metadata"] = meta
				spec["template"] = template
				v["spec"] = spec
			}
		}
	}

	yam, err = yaml.Marshal(v)

	return string(yam), err
}

func postProcessValidationSchema(oapi *openapi.OrderedMap) (*apiextv1.JSONSchemaProps, error) {
	oapiJson, err := oapi.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	if err = json.Unmarshal(oapiJson, &obj); err != nil {
		return nil, err
	}

	// remove 'properties' and 'required' fields to prevent validating proto.Any fields
	removeProtoAnyValidation(obj)

	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	jsonSchema := &apiextv1.JSONSchemaProps{}
	if err = json.Unmarshal(bytes, jsonSchema); err != nil {
		return nil, err
	}
	return jsonSchema, nil
}

// prevent k8s from validating proto.Any fields (since it's unstructured)
func removeProtoAnyValidation(d map[string]interface{}) {
	for _, v := range d {
		values, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		desc, ok := values["properties"]
		properties, isObj := desc.(map[string]interface{})
		// detect proto.Any field from presence of "@type" as field under "properties"
		if !ok || !isObj || properties["@type"] == nil {
			removeProtoAnyValidation(values)
			continue
		}
		// remove "properties" value
		delete(values, "properties")
		// remove "required" value
		delete(values, "required")
		// x-kubernetes-preserve-unknown-fields allows for unknown fields from a particular node
		// see https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#specifying-a-structural-schema
		values["x-kubernetes-preserve-unknown-fields"] = true
	}
}

func ifDefined(val, defaultVal string) string {
	if val != "" {
		return val
	}

	return defaultVal
}
