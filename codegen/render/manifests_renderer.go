package render

import (
	"encoding/json"
	"regexp"
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
type MakeResourceFunc func(groups []*Group) ([]metav1.Object, error)

// renders kubernetes from templates
type ManifestsRenderer struct {
	templateRenderer
	AppName       string // used for labeling
	ResourceFuncs map[OutFile]MakeResourceFunc
	ManifestDir   string
	ProtoDir      string
}

type templateArgs struct {
	Crds       []apiextv1.CustomResourceDefinition
	ShouldSkip map[string]bool
}

func RenderManifests(
	appName, manifestDir, protoDir string,
	protoOpts protoutil.Options,
	groupOptions model.GroupOptions,
	grps []*Group,
) ([]OutFile, error) {
	defaultManifestsRenderer := ManifestsRenderer{
		AppName:     appName,
		ManifestDir: manifestDir,
		ProtoDir:    protoDir,
	}
	return defaultManifestsRenderer.RenderManifests(grps, protoOpts, groupOptions)
}

func (r ManifestsRenderer) RenderManifests(grps []*Group, protoOpts protoutil.Options, groupOptions model.GroupOptions) ([]OutFile, error) {
	grpsByGroupName := make(map[string][]*Group)
	shouldRenderGroups := make(map[string]bool)
	shouldSkipCRDManifest := make(map[string]bool)
	shouldSkipTemplatedCRDManifest := make(map[string]bool)
	grandfatheredGroups := make(map[string]bool)
	for _, grp := range grps {
		grpsByGroupName[grp.Group] = append(grpsByGroupName[grp.Group], grp)
		shouldRenderGroups[grp.Group] = shouldRenderGroups[grp.Group] || grp.RenderManifests
		grandfatheredGroups[grp.GroupVersion.String()] = grandfatheredGroups[grp.GroupVersion.String()] || grp.Grandfathered
		shouldSkipCRDManifest[grp.Group] =
			shouldSkipCRDManifest[grp.Group] || grp.SkipCRDManifest
		shouldSkipTemplatedCRDManifest[grp.Group] =
			shouldSkipTemplatedCRDManifest[grp.Group] || grp.SkipTemplatedCRDManifest
	}

	for _, grp := range grps {
		if grp.RenderValidationSchemas && shouldRenderGroups[grp.Group] {
			var err error
			oapiSchemas, err := generateOpenApi(*grp, r.ProtoDir, protoOpts, groupOptions)
			if err != nil {
				return nil, err
			}
			grp.OpenApiSchemas = oapiSchemas
		}
	}

	var renderedFiles []OutFile

	for groupName, selectedGrps := range grpsByGroupName {
		if !shouldRenderGroups[groupName] {
			continue
		}

		crds, err := r.createCrds(r.AppName, selectedGrps)
		if err != nil {
			return nil, err
		}
		out, err := r.renderCRDManifest(r.AppName, groupName, crds)
		if err != nil {
			return nil, err
		}
		renderedFiles = append(renderedFiles, out)

		out, err = r.renderTemplatedCRDManifest(r.AppName, groupName, crds, grandfatheredGroups)
		if err != nil {
			return nil, err
		}
		renderedFiles = append(renderedFiles, out)
	}

	return renderedFiles, nil
}

// Use cuelang as an intermediate language for transpiling protobuf schemas to openapi v3 with k8s structural schema constraints.
func generateOpenApi(grp model.Group, protoDir string, protoOpts protoutil.Options, groupOptions model.GroupOptions) (model.OpenApiSchemas, error) {
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
			jsonSchema, err := postProcessValidationSchema(orderedMap, groupOptions)
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

// TODO (dmitri-d): this can be removed once we migrate to use platform charts exclusively
func (r ManifestsRenderer) renderCRDManifest(appName, groupName string, objs []apiextv1.CustomResourceDefinition) (OutFile, error) {
	outFile := OutFile{
		Path: r.ManifestDir + "/crds/" + groupName + "_" + "crds.yaml",
	}

	var objManifests []string
	for _, obj := range objs {
		manifest, err := marshalObjToYaml(appName, &obj)
		if err != nil {
			return OutFile{}, err
		}
		objManifests = append(objManifests, manifest)
	}

	outFile.Content = strings.Join(objManifests, "\n---\n")
	return outFile, nil
}

func (r ManifestsRenderer) renderTemplatedCRDManifest(appName, groupName string,
	objs []apiextv1.CustomResourceDefinition,
	grandfatheredGroups map[string]bool) (OutFile, error) {

	renderer := DefaultTemplateRenderer

	// when rendering helm charts, we need
	// to use a custom delimiter
	renderer.left = "[["
	renderer.right = "]]"

	defaultManifestRenderer := ChartRenderer{
		templateRenderer: renderer,
	}

	outFile := OutFile{Path: r.ManifestDir + "/templates/" + groupName + "_" + "crds.yaml"}
	templatesToRender := inputTemplates{
		"manifests/crd.yamltmpl": outFile,
	}
	files, err := defaultManifestRenderer.renderCoreTemplates(
		templatesToRender,
		templateArgs{Crds: objs, ShouldSkip: grandfatheredGroups})
	if err != nil {
		return OutFile{}, err
	}
	return files[0], nil
}

func (r ManifestsRenderer) createCrds(appName string, groups []*Group) ([]apiextv1.CustomResourceDefinition, error) {
	objs, err := kuberesource.CustomResourceDefinitions(groups)
	if err != nil {
		return nil, err
	}

	for i, obj := range objs {
		// find the annotation of the manifest, and add to them
		SetVersionForObject(objs[i].GetObjectMeta(), groups[0].AddChartVersion)

		labels := obj.GetLabels()
		if labels == nil {
			labels = map[string]string{}
		}

		labels["app"] = appName
		labels["app.kubernetes.io/name"] = appName

		objs[i].SetLabels(labels)
	}
	return objs, nil
}

func marshalObjToYaml(appName string, obj metav1.Object) (string, error) {
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

func postProcessValidationSchema(oapi *openapi.OrderedMap, groupOptions model.GroupOptions) (*apiextv1.JSONSchemaProps, error) {
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

	if groupOptions.EscapeGoTemplateOperators {
		// escape {{ or }} in descriptions as they will cause helm to error
		escapeGoTemplateOperators(obj)
	}

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

// escape go template operators in descriptions, {{ or }} in description will cause helm to error
// We do not always control descriptions since we import some protos
func escapeGoTemplateOperators(d map[string]interface{}) {
	for k, v := range d {

		if k == "description" {
			d[k] = sanitizeDescription(v)
		}

		values, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		escapeGoTemplateOperators(values)
	}
}

// Regex to escape double brackets in go templates
func sanitizeDescription(desc interface{}) interface{} {
	if description, isString := desc.(string); isString {
		exp := regexp.MustCompile("{{([^}]+)}}")
		description = exp.ReplaceAllString(description, `{{"{{"}}$1{{"}}"}}`)
		return description
	}
	return desc
}

func ifDefined(val, defaultValue string) string {
	if val != "" {
		return val
	}

	return defaultValue
}
