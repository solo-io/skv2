package render

import (
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"

	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig/v3"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/util"
	"sigs.k8s.io/yaml"
)

func makeTemplateFuncs(customFuncs template.FuncMap) template.FuncMap {
	f := sprig.TxtFuncMap()

	// Add some functionality for skv2 templates
	skv2Funcs := template.FuncMap{
		// string utils

		"toToml":   toTOML,
		"toYaml":   toYAML,
		"fromYaml": fromYAML,
		"toJson":   toJSON,
		"fromJson": fromJSON,

		"join":            strings.Join,
		"lower":           strings.ToLower,
		"lower_camel":     strcase.ToLowerCamel,
		"upper_camel":     strcase.ToCamel,
		"pluralize":       pluralize.NewClient().Plural,
		"snake":           strcase.ToSnake,
		"split":           splitTrimEmpty,
		"string_contains": strings.Contains,

		// resource-related funcs
		"group_import_path": func(grp Group) string {
			return util.GoPackage(grp)
		},
		"group_import_name": func(grp Group) string {
			name := strings.ReplaceAll(grp.GroupVersion.String(), "/", "_")
			name = strings.ReplaceAll(name, ".", "_")
			name = strings.ReplaceAll(name, "-", "_")

			return name
		},
		"generated_code_import_path": func(grp Group) string {
			return util.GeneratedGoPackage(grp)
		},
		// Used by types.go to get all unique external imports for a groups resources
		"imports_for_group": func(grp Group) []string {
			allImports := uniqueGoImportPathsForGroup(grp)
			var excludingGroupImport []string
			for _, imp := range allImports {
				if imp == util.GoPackage(grp) {
					continue
				}
				excludingGroupImport = append(excludingGroupImport, imp)
			}
			return excludingGroupImport
		},
	}

	for k, v := range skv2Funcs {
		f[k] = v
	}

	for k, v := range customFuncs {
		f[k] = v
	}

	return f
}

/*
	Find the proto messages for a given set of descriptors which need proto_deepcopoy funcs and whose types are not in
	the API root package

	return true if the descriptor corresponds to the Spec or the Status field
*/
func shouldDeepCopyExternalMessage(resources []model.Resource, desc *descriptor.DescriptorProto) bool {
	for _, resource := range resources {
		if resource.Spec.Type.Name == desc.GetName() ||
			(resource.Status != nil && resource.Status.Type.Name == desc.GetName()) {
			return true
		}
	}
	return false
}

/*
	Find the proto messages for a given set of descriptors which need proto_deepcopoy funcs.
	The two cases are as follows:

	1. One of the subfields has an external type
	2. There is a oneof present
*/
func shouldDeepCopyInternalMessage(packageName string, desc *descriptor.DescriptorProto) bool {
	var shouldGenerate bool
	// case 1 above
	for _, v := range desc.GetField() {
		if v.TypeName != nil && !strings.Contains(v.GetTypeName(), packageName) {
			shouldGenerate = true
			break
		}
	}
	// case 2 above
	return len(desc.GetOneofDecl()) > 0 || shouldGenerate
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

// fromYAML converts a YAML document into a map[string]interface{}.
//
// This is not a general-purpose YAML parser, and will not parse all valid
// YAML documents. Additionally, because its intended use is within templates
// it tolerates errors. It will insert the returned error message string into
// m["Error"] in the returned map.
func fromYAML(str string) map[string]interface{} {
	m := map[string]interface{}{}

	if err := yaml.Unmarshal([]byte(str), &m); err != nil {
		m["Error"] = err.Error()
	}
	return m
}

// toTOML takes an interface, marshals it to toml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toTOML(v interface{}) string {
	b := bytes.NewBuffer(nil)
	e := toml.NewEncoder(b)
	err := e.Encode(v)
	if err != nil {
		return err.Error()
	}
	return b.String()
}

// toJSON takes an interface, marshals it to json, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return string(data)
}

// fromJSON converts a JSON document into a map[string]interface{}.
//
// This is not a general-purpose JSON parser, and will not parse all valid
// JSON documents. Additionally, because its intended use is within templates
// it tolerates errors. It will insert the returned error message string into
// m["Error"] in the returned map.
func fromJSON(str string) map[string]interface{} {
	m := make(map[string]interface{})

	if err := json.Unmarshal([]byte(str), &m); err != nil {
		m["Error"] = err.Error()
	}
	return m
}

func splitTrimEmpty(s, sep string) []string {
	return strings.Split(strings.TrimSpace(s), sep)
}
