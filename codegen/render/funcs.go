package render

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/solo-io/skv2/codegen/model/values"
	"github.com/solo-io/skv2/codegen/util/stringutils"
	"k8s.io/utils/strings/slices"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig/v3"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/util"
	goyaml "gopkg.in/yaml.v3"
	"sigs.k8s.io/yaml"
)

func makeTemplateFuncs(customFuncs template.FuncMap) template.FuncMap {
	f := sprig.TxtFuncMap()

	// Add some functionality for skv2 templates
	skv2Funcs := template.FuncMap{
		// string utils

		"toToml":     toTOML,
		"toYaml":     toYAML,
		"fromYaml":   fromYAML,
		"toJson":     toJSON,
		"fromJson":   fromJSON,
		"toNode":     toNode,
		"fromNode":   fromNode,
		"mergeNodes": mergeNodes,

		"join":            strings.Join,
		"lower":           strings.ToLower,
		"lower_camel":     strcase.ToLowerCamel,
		"upper_camel":     strcase.ToCamel,
		"pluralize":       stringutils.Pluralize,
		"snake":           strcase.ToSnake,
		"split":           splitTrimEmpty,
		"string_contains": strings.Contains,
		"alias_for":       util.AliasFor,

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

		"type_name": func(t model.Type, g model.Group) string {
			// If the type's go package isn't the current package, give its package name an alias (don't dot import!)
			if t.GoPackage != "" && t.GoPackage != util.GoPackage(g) {
				return fmt.Sprintf("%s.%s", util.AliasFor(t.GoPackage), t.Name)
			}
			return t.Name
		},

		"get_operator_values": func(o values.UserOperatorValues) map[string]interface{} {
			opValues := map[string]interface{}{
				strcase.ToLowerCamel(o.Name): o.Values,
			}
			if o.ValuePath != "" {
				splitPath := strings.Split(o.ValuePath, ".")
				for _, p := range splitPath {
					opValues = map[string]interface{}{
						strcase.ToLowerCamel(p): opValues,
					}
				}
			}
			return opValues
		},

		"get_operator_custom_values": func(o values.UserOperatorValues) map[string]interface{} {
			opValues := map[string]interface{}{}
			if o.CustomValues != nil {
				opValues = map[string]interface{}{
					strcase.ToLowerCamel(o.Name): o.CustomValues,
				}
				if o.ValuePath != "" {
					splitPath := strings.Split(o.ValuePath, ".")
					for _, p := range splitPath {
						opValues = map[string]interface{}{
							strcase.ToLowerCamel(p): opValues,
						}
					}
				}
			}
			return opValues
		},

		"containerConfigs": containerConfigs,
	}

	for k, v := range skv2Funcs {
		f[k] = v
	}

	for k, v := range customFuncs {
		f[k] = v
	}

	return f
}

type containerConfig struct {
	model.Container
	Name      string
	ValuesVar string
}

func containerConfigs(op model.Operator) []containerConfig {
	opVar := fmt.Sprintf("$.Values.%s", strcase.ToLowerCamel(op.Name))
	if op.ValuePath != "" {
		opVar = fmt.Sprintf("$.Values.%s.%s", op.ValuePath, strcase.ToLowerCamel(op.Name))

	}
	configs := []containerConfig{{
		Container: op.Deployment.Container,
		Name:      op.Name,
		ValuesVar: opVar,
	}}

	for _, sidecar := range op.Deployment.Sidecars {
		configs = append(configs, containerConfig{
			Container: sidecar.Container,
			Name:      sidecar.Name,
			ValuesVar: opVar + ".sidecars." + strcase.ToLowerCamel(sidecar.Name),
		})
	}

	return configs
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
	// NOTE(ilackarms): due to a bug in the underlying yaml library
	// inserting unnecessary newlines when marshalling string arrays,
	// we handle that special-case here
	if strSlice, ok := v.([]string); ok {
		return strSliceToYaml(strSlice)
	}

	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

func strSliceToYaml(strSlice []string) string {
	var yamlElements []string
	for _, s := range strSlice {
		yamlElements = append(yamlElements, "- "+s)
	}
	return strings.Join(yamlElements, "\n")
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

type yamlCommentsConfig interface {
	Enabled() bool
	LineLength() int
}

func toNode(v interface{}, commentsConfig yamlCommentsConfig) goyaml.Node {
	var node goyaml.Node
	if v != nil {
		err := goyaml.Unmarshal([]byte(toYAML(v)), &node)
		if err != nil {
			panic(err)
		}
		if commentsConfig.Enabled() {
			newYamlCommenter(commentsConfig).addYamlComments(reflect.ValueOf(v), &node)
		}
	}
	return node
}

type yamlCommenter struct {
	config yamlCommentsConfig
	fields map[reflect.Type]yamlFieldMapping
}

func newYamlCommenter(config yamlCommentsConfig) *yamlCommenter {
	return &yamlCommenter{
		config: config,
		fields: make(map[reflect.Type]yamlFieldMapping),
	}
}

// addYamlComments mutates the intermediate representation (the node) and adds
// header comments, (comments on the line above above the field), for structs
// whose members are tagged with field descriptions (i.e. the "desc" tag). It
// does this by traversing the IR together with the value being marshalled and
// inspecting the value's type via reflection to find the tags.
func (yc *yamlCommenter) addYamlComments(
	value reflect.Value,
	node *goyaml.Node,
) {
	if value.IsZero() {
		return
	}

	valueType := value.Type()

	// unwrap interface to concrete type
	if valueType.Kind() == reflect.Interface {
		value = value.Elem()
		valueType = value.Type()
	}

	// unwrap pointer to dereferenced type
	for valueType.Kind() == reflect.Pointer {
		value = value.Elem()
		valueType = valueType.Elem()
	}

	if node.Kind == goyaml.DocumentNode {
		for _, childNode := range node.Content {
			yc.addYamlComments(value, childNode)
		}
	}

	// sequence node is a list of values. continue by iterating each list item
	if node.Kind == goyaml.SequenceNode {
		for i, child := range node.Content {
			yc.addYamlComments(value.Index(i), child)
		}
	}

	// mapping kind represents a key-value map
	if node.Kind == goyaml.MappingNode {
		// lookup of keys in a map type, keyed by their string value.
		var mapKeysByName map[string]reflect.Value

		// we're iterating like this b/c the the Content list organizes the kv pairs
		// in a sequence like: [key1, value1, key2, value2, ...etc.]
		for i := 0; i < len(node.Content); i = i + 2 {
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]
			var nextValue reflect.Value

			switch valueType.Kind() {
			case reflect.Struct:
				{
					// get the field from the strcut
					fieldName := yc.getStructFieldName(valueType, keyNode.Value)
					field, _ := valueType.FieldByName(fieldName)

					// if the field is tagged w/ description, add the comment
					keyNode.HeadComment = field.Tag.Get("desc")
					if yc.config.LineLength() > 0 {
						keyNode.HeadComment = wrapWords(
							keyNode.HeadComment,
							yc.config.LineLength(),
						)
					}

					nextValue = value.FieldByName(fieldName)
				}
			case reflect.Map:
				{
					// lazily init the map key lookup
					if mapKeysByName == nil {
						mapKeysByName = make(map[string]reflect.Value)
						for _, k := range value.MapKeys() {
							mapKeysByName[k.String()] = k
						}
					}

					nextValue = value.MapIndex(mapKeysByName[keyNode.Value])
				}
			}

			// continue traversing the object
			yc.addYamlComments(nextValue, valueNode)
		}
	}
}

// getStructFieldName returns the name of the field on a struct type
// corresponding to the field on the serialized yaml object.
func (yc *yamlCommenter) getStructFieldName(
	valueType reflect.Type,
	yamlFieldName string,
) string {
	mapping, ok := yc.fields[valueType]
	if !ok {
		mapping = newYamlFieldMapping(valueType)
		yc.fields[valueType] = mapping
	}

	return mapping[yamlFieldName]
}

// yamlFieldMapping is a mapping between the fields struct and its yaml
// serialized form.
type yamlFieldMapping map[string]string

func newYamlFieldMapping(valueType reflect.Type) yamlFieldMapping {
	result := make(map[string]string)

	// iterate the fields to populate result
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)

		// set the field name based on the yaml tag if it's present
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			jsonTagSegments := strings.Split(jsonTag, ",")
			jsonFieldName := jsonTagSegments[0]

			// handle inlined struct
			if len(jsonTagSegments) > 1 && slices.Contains(jsonTagSegments[1:], "inline") {
				for k, v := range newYamlFieldMapping(field.Type) {
					result[k] = v
				}
			} else {
				result[jsonFieldName] = field.Name
			}

			// handle embedded type
		} else if field.Anonymous {
			for k, v := range newYamlFieldMapping(field.Type) {
				result[k] = v
			}

			// otherwise field name in serialized yaml is same as struct field
		} else {
			result[field.Name] = field.Name
		}
	}

	return result
}

// wrapWords inserts newline characters in the string between words to try to
// keep line-lengths below the limit. If some word is longer than the limit, it
// is not broken.
func wrapWords(s string, limit int) string {
	if strings.TrimSpace(s) == "" {
		return s
	}

	var lines []string
	var lineWords []string
	lineWordChars := 0

	for _, word := range strings.Fields(s) {
		if len(lineWords) == 0 {
			lineWords = append(lineWords, word)
			lineWordChars += len(word)
			continue
		}

		if len(lineWords)-1+lineWordChars+len(word) > limit {
			lines = append(lines, strings.Join(lineWords, " "))
			lineWordChars = 0
			lineWords = []string{}
		}

		lineWords = append(lineWords, word)
		lineWordChars += len(word)
	}

	lines = append(lines, strings.Join(lineWords, " "))
	return strings.Join(lines, "\n")
}

func fromNode(n goyaml.Node) string {
	b, err := goyaml.Marshal(&n)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func mergeNodes(nodes ...goyaml.Node) goyaml.Node {
	if len(nodes) <= 1 {
		panic("at least two nodes required for merge")
	}
	var mergedNode goyaml.Node
	for _, n := range nodes {

		if mergedNode.IsZero() {
			if n.IsZero() {
				continue
			}
			mergedNode = n
			continue
		}

		if err := recursiveNodeMerge(&n, &mergedNode); err != nil {
			panic(err)
		}
	}
	if mergedNode.IsZero() {
		panic("all nodes were found to be IsZero, cannot continue")

	}
	return mergedNode
}

func nodesEqual(l, r *goyaml.Node) bool {
	if l.Kind == goyaml.ScalarNode && r.Kind == goyaml.ScalarNode {
		return l.Value == r.Value
	}
	panic("equals on non-scalars not implemented!")
}

func recursiveNodeMerge(from, into *goyaml.Node) error {
	if from.Kind != into.Kind {
		return errors.New("cannot merge nodes of different kinds")
	}
	switch from.Kind {
	case goyaml.MappingNode:
		for i := 0; i < len(from.Content); i += 2 {
			found := false
			for j := 0; j < len(into.Content); j += 2 {
				if nodesEqual(from.Content[i], into.Content[j]) {
					found = true
					if into.Content[j+1].Value == "null" {
						// No matter the type if the value in the existing map is null use the new node
						into.Content[j+1] = from.Content[i+1]
					} else if from.Content[i+1].Value == "null" {
						// value on new map is null, skip and leave existing
						break
					} else if into.Content[j+1].Kind == goyaml.ScalarNode && from.Content[i+1].Kind == goyaml.ScalarNode {
						// if both Scalars simply overwrite
						into.Content[j+1] = from.Content[i+1]
					} else {
						// Sequences and maps will be merged here
						if err := recursiveNodeMerge(from.Content[i+1], into.Content[j+1]); err != nil {
							return errors.New("at key " + from.Content[i].Value + ": " + err.Error())
						}
					}
					break
				}
			}
			if !found {
				into.Content = append(into.Content, from.Content[i:i+2]...)
			}
		}
	case goyaml.SequenceNode:
		into.Content = from.Content
	case goyaml.DocumentNode:
		return recursiveNodeMerge(from.Content[0], into.Content[0])
	default:
		return errors.New("can only merge mapping and sequence nodes")
	}
	return nil
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
