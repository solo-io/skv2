package render

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/invopop/jsonschema"
	"github.com/solo-io/skv2/codegen/model/values"
	"github.com/solo-io/skv2/codegen/util/stringutils"
	"google.golang.org/protobuf/types/known/structpb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/strings/slices"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig/v3"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/iancoleman/orderedmap"
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

		"toToml":       toTOML,
		"toYaml":       toYAML,
		"fromYaml":     fromYAML,
		"toJson":       toJSON,
		"fromJson":     fromJSON,
		"toNode":       toNode,
		"fromNode":     fromNode,
		"mergeNodes":   mergeNodes,
		"toJsonSchema": toJSONSchema,

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
				o.FormattedName(): o.Values,
			}
			if o.ValuePath != "" {
				splitPath := strings.Split(o.ValuePath, ".")
				reverse(splitPath)
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
					o.FormattedName(): o.CustomValues,
				}
				if o.ValuePath != "" {
					splitPath := strings.Split(o.ValuePath, ".")
					reverse(splitPath)
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
		"toListItem":       toListItem,
		"opVar":            opVar,
	}

	for k, v := range skv2Funcs {
		f[k] = v
	}

	for k, v := range customFuncs {
		f[k] = v
	}

	return f
}

func toListItem(item interface{}) []interface{} {
	return []interface{}{item}
}

type containerConfig struct {
	model.Container
	Name      string
	ValuesVar string
}

func containerConfigs(op model.Operator) []containerConfig {
	valuesVar := opVar(op)
	configs := []containerConfig{{
		Container: op.Deployment.Container,
		Name:      op.Name,
		ValuesVar: valuesVar,
	}}

	for _, sidecar := range op.Deployment.Sidecars {
		configs = append(configs, containerConfig{
			Container: sidecar.Container,
			Name:      sidecar.Name,
			ValuesVar: valuesVar + ".sidecars." + strcase.ToLowerCamel(sidecar.Name),
		})
	}

	return configs
}

func opVar(op model.Operator) string {
	name := op.FormattedName()
	opVar := fmt.Sprintf("$.Values.%s", name)
	if op.ValuePath != "" {
		opVar = fmt.Sprintf("$.Values.%s.%s", op.ValuePath, name)
	}
	return opVar
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
		// still try to add comments structs that are initialized to all zero values
		if value.Kind() != reflect.Struct {
			return
		}
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
					if field.Anonymous {
						field, _ = field.Type.FieldByName(fieldName)
					}

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
	buf := new(bytes.Buffer)
	encoder := goyaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err := encoder.Encode(sortYAML(&n))
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// Implement sorting for prettier yaml
type nodes []*goyaml.Node

func (i nodes) Len() int { return len(i) / 2 }

func (i nodes) Swap(x, y int) {
	x *= 2
	y *= 2
	i[x], i[y] = i[y], i[x]         // keys
	i[x+1], i[y+1] = i[y+1], i[x+1] // values
}

func (i nodes) Less(x, y int) bool {
	x *= 2
	y *= 2
	return i[x].Value < i[y].Value
}

func sortYAML(node *goyaml.Node) *goyaml.Node {
	if node.Kind == goyaml.DocumentNode {
		for i, n := range node.Content {
			node.Content[i] = sortYAML(n)
		}
	}
	if node.Kind == goyaml.SequenceNode {
		for i, n := range node.Content {
			node.Content[i] = sortYAML(n)
		}
	}
	if node.Kind == goyaml.MappingNode {
		for i, n := range node.Content {
			node.Content[i] = sortYAML(n)
		}
		sort.Sort(nodes(node.Content))
	}
	return node
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

// toJSONSchema generates the json schema for the values of the helm chart
func toJSONSchema(values values.UserHelmValues) string {
	reflector := createJsonSchemaReflector(values)
	schema := new(jsonschema.Schema)
	schema.Version = jsonschema.Version

	// add json schema properties from the chart's custom values
	if values.CustomValues != nil {
		mergeJsonSchema(schema, reflector.Reflect(values.CustomValues))
	}

	// add json schema for the operators
	if values.Operators != nil {
		for _, o := range values.Operators {
			opSchema := jsonSchemaForOperator(reflector, &o)
			mergeJsonSchema(schema, opSchema)
		}
	}

	jsonSchema, err := schema.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return indentJson(string(jsonSchema))
}

type customTypeMapper func(reflect.Type, *jsonschema.Schema) *jsonschema.Schema

// createTypeMappings initializes the type mappings based on what the user has
// passed and also some reasonable defaults
func createCustomTypeMapper(values values.UserHelmValues) customTypeMapper {
	// add the default type mappings ... these types have customized json deser
	// behaviour so we need to have custom mappings describing allowed values
	typeMappings := map[reflect.Type]customTypeMapper{
		reflect.TypeOf(structpb.Value{}): func(t reflect.Type, defaultSchema *jsonschema.Schema) *jsonschema.Schema {
			schema := new(jsonschema.Schema)
			schema.AnyOf = []*jsonschema.Schema{
				{Type: "null"},
				{Type: "number"},
				{Type: "string"},
				{Type: "boolean"},
				{Type: "object"},
				{Type: "array"},
			}
			return schema
		},

		reflect.TypeOf(metav1.Time{}): func(t reflect.Type, defaultSchema *jsonschema.Schema) *jsonschema.Schema {
			schema := new(jsonschema.Schema)
			schema.AnyOf = []*jsonschema.Schema{
				{
					Type:   "string",
					Format: "date-time",
				},
				{
					Type: "null",
				},
			}
			return schema
		},

		reflect.TypeOf(resource.Quantity{}): func(t reflect.Type, defaultSchema *jsonschema.Schema) *jsonschema.Schema {
			return &jsonschema.Schema{
				AnyOf: []*jsonschema.Schema{
					{Type: "string"},
					{Type: "number"},
				},
			}
		},

		reflect.TypeOf(intstr.IntOrString{}): func(t reflect.Type, defaultSchema *jsonschema.Schema) *jsonschema.Schema {
			return &jsonschema.Schema{
				AnyOf: []*jsonschema.Schema{
					{Type: "string"},
					{Type: "number"},
				},
			}
		},

		reflect.TypeOf(v1.SecurityContext{}): func(t reflect.Type, defaultSchema *jsonschema.Schema) *jsonschema.Schema {
			return &jsonschema.Schema{
				AnyOf: []*jsonschema.Schema{
					defaultSchema,
					{Type: "boolean", Const: false},
				},
			}
		},
	}

	return func(t reflect.Type, schema *jsonschema.Schema) *jsonschema.Schema {
		if values.JsonSchema.CustomTypeMapper != nil {
			// the custom mappings are accept the json schema as a map and are
			// expected to return an interface that can be serialized as a json schema
			// or null if it doesn't handle the type

			// serialize default schema into a map
			buf, err := schema.MarshalJSON()
			if err != nil {
				panic(err)
			}

			var defaultAsMap map[string]interface{}

			// if it's a boolean schema (per section 4.3.2 of the spec) convert it
			// to a map representation
			var section432Schema bool
			err = json.Unmarshal(buf, &section432Schema)
			if err == nil {
				if section432Schema {
					// "true" == Always passes validation, as if the empty schema {}
					defaultAsMap = map[string]interface{}{}
				} else {
					// "false" == Always fails validation, as if the schema { "not": {} }
					defaultAsMap = map[string]interface{}{
						"not": map[string]interface{}{},
					}
				}
			} else {
				// if here - schema is not a bool so unmarshal the schema as map
				err = json.Unmarshal(buf, &defaultAsMap)
				if err != nil {
					panic(err)
				}
			}

			// if the type is handled, deserialize it into json schema and return that
			modifiedSchema := values.JsonSchema.CustomTypeMapper(t, defaultAsMap)
			if modifiedSchema != nil {
				buf, err = json.Marshal(modifiedSchema)
				if err != nil {
					panic(err)
				}
				result := new(jsonschema.Schema)
				err = result.UnmarshalJSON(buf)
				if err != nil {
					panic(err)
				}
				return result
			}
		}

		defaultTypeMapper, ok := typeMappings[t]
		if ok {
			return defaultTypeMapper(t, schema)
		}

		return nil
	}
}

// initialize a new instance of the jsonschema Reflector with configured to
// create json schema for chart.
func createJsonSchemaReflector(values values.UserHelmValues) *jsonschema.Reflector {
	// we might want to pass configuration of this from values.UserHelmValues
	// in the future, to control how the schema is generated, but for now just
	// hard-coding some sensible defaults that won't break existing charts..
	reflector := jsonschema.Reflector{
		Anonymous:                  true,
		AllowAdditionalProperties:  true,
		DoNotReference:             true,
		RequiredFromJSONSchemaTags: true,
	}

	customTypeMapper := createCustomTypeMapper(values)
	extractBaseMapping := makeBaseMappingExtractor(&reflector)

	// specify custom overrides for some specific types
	reflector.Mapper = func(t reflect.Type) *jsonschema.Schema {
		baseMapping := extractBaseMapping(t)
		if baseMapping == nil {
			return baseMapping
		}

		customMapping := customTypeMapper(t, baseMapping)
		if customMapping != nil {
			return customMapping
		}

		// allow maps & arrays to be nil by default, which makes this backwards
		// compatible with older charts without schemas. In the future, might like
		// to control this behaviour w/ tags such as `jsonschema:"nullable"`
		// to fields of the values structs.
		if t.Kind() == reflect.Map || t.Kind() == reflect.Slice {
			return wrapAsNullable(baseMapping)
		}

		// make the pointer fields on structs nullable
		if t.Kind() == reflect.Struct {
			return pointerFieldsNullableTransformer(t, customTypeMapper)(baseMapping)
		}

		// by returning nil, it defaults to the default mapping behaviour reflector
		return nil
	}

	return &reflector
}

type jsonSchemaTransform func(*jsonschema.Schema) *jsonschema.Schema

// makeBaseMappingExtractor allows the base mapping function of the reflector
// to be called to generate the default json schema for the type (as if the
// relfector had no custom Mapper)
func makeBaseMappingExtractor(
	reflector *jsonschema.Reflector,
) func(t reflect.Type) *jsonschema.Schema {
	// we try to map the type recursively by re-calling ReflectFromType & letting
	// the calling funciton follow the same code path into this function which
	// will check this flag for the recursion exit condition
	insideBaseMappingExtractor := false

	return func(t reflect.Type) *jsonschema.Schema {
		if insideBaseMappingExtractor {
			insideBaseMappingExtractor = false
			return nil // deletage to the default mapper & return
		}

		insideBaseMappingExtractor = true

		// map the type to it's default non-nullable type
		schema := reflector.ReflectFromType(t)

		// remove default fields
		schema.Version = ""

		return schema
	}
}

// wrapAsNullable wraps the json schema in a way that allows the json value
// to be null
func wrapAsNullable(schema *jsonschema.Schema) *jsonschema.Schema {
	return &jsonschema.Schema{
		AnyOf: []*jsonschema.Schema{
			schema,
			{Type: "null"},
		},
	}
}

// pointerFieldsNullableTransformer creates function that when called will
// return a json schema with the fields which are pointers on the struct type
// as nullable
func pointerFieldsNullableTransformer(
	t reflect.Type,
	customTypeMapper customTypeMapper,
) jsonSchemaTransform {
	return func(schema *jsonschema.Schema) *jsonschema.Schema {
		makePointerFieldsNullable(t, schema, customTypeMapper)
		return schema
	}
}

// makePointerFieldsNullableInternal makes any fields that are of type pointer
// on the passed type (which must be a struct) nullable in the json schema
func makePointerFieldsNullable(
	t reflect.Type,
	schema *jsonschema.Schema,
	customTypeMapper customTypeMapper,
) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonPropertyName := field.Name

		// check if the type has been renmamed via the json tag
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			jsonTagSegments := strings.Split(jsonTag, ",")
			jsonPropertyName = jsonTagSegments[0]
		}

		if field.Type.Kind() == reflect.Pointer {
			f, ok := schema.Properties.Get(jsonPropertyName)
			if !ok {
				continue
			}
			fieldSchema := f.(*jsonschema.Schema)

			// if there's a custom type mapping for the field, delegate to that
			// whether it wants the field to be nullable
			if customTypeMapper(field.Type.Elem(), fieldSchema) != nil {
				continue
			}

			schema.Properties.Set(jsonPropertyName, wrapAsNullable(fieldSchema))
		}
	}
}

// mergeJsonSchema Adds the properties from the source json schema
// to the target
func mergeJsonSchema(
	target *jsonschema.Schema,
	src *jsonschema.Schema,
) {

	if target.Required == nil {
		target.Required = []string{}
	}

	target.Required = append(target.Required, src.Required...)

	if target.Properties == nil {
		target.Properties = orderedmap.New()
	}

	for _, prop := range src.Properties.Keys() {
		sourceVal, _ := src.Properties.Get(prop)
		sourceValSchema, sourceIsSchema := sourceVal.(*jsonschema.Schema)
		merged := false
		if sourceIsSchema {
			targetValue, targetHasProp := target.Properties.Get(prop)
			if sourceValSchema.Properties != nil && len(sourceValSchema.Properties.Keys()) > 0 && targetHasProp {
				targetValSchema, targetIsSchema := targetValue.(*jsonschema.Schema)
				if targetIsSchema {
					mergeJsonSchema(targetValSchema, sourceValSchema)
					target.Properties.Set(prop, targetValSchema)
					merged = true
				}
			}
		}

		if !merged {
			target.Properties.Set(prop, sourceVal)
		}
	}
}

// jsonSchemaForOperator creates the json schema for the operator's values.
//
// The resulting schema will have the values nested at the appropriate path,
// taking into account the operator's name and value path. This means it can be
// added directly to the base object's schema
func jsonSchemaForOperator(
	reflector *jsonschema.Reflector,
	o *values.UserOperatorValues,
) *jsonschema.Schema {
	schema := reflector.Reflect(o.Values)
	schema.Version = "" // will be set on parent object

	// if the opeartor defines custom values, add the properties to the schema
	if o.CustomValues != nil {
		mergeJsonSchema(schema, reflector.Reflect(o.CustomValues))
	}

	// nest the operator values schema in the correct place
	path := []string{o.FormattedName()}
	if o.ValuePath != "" {
		splitPath := strings.Split(o.ValuePath, ".")
		reverse(splitPath)
		path = append(path, splitPath...)
	}
	schema = nestSchemaAtPath(schema, path...)

	return schema
}

// returns a new schema, nested within an object at the given path
func nestSchemaAtPath(schema *jsonschema.Schema, path ...string) *jsonschema.Schema {
	for _, p := range path {
		parentSchema := new(jsonschema.Schema)
		parentSchema.Type = "object"
		parentSchema.Properties = orderedmap.New()
		parentSchema.Properties.Set(p, schema)
		schema = parentSchema
	}

	return schema
}

func indentJson(src string) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(src), "", "  ")
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// reverse mutates a slice of strings to reverse the order.
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitTrimEmpty(s, sep string) []string {
	return strings.Split(strings.TrimSpace(s), sep)
}
