package render_test

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/lithammer/dedent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/skv2/codegen/model/values"
	"github.com/solo-io/skv2/codegen/render"
)

func prepareExpected(expected string) string {
	expected = strings.Trim(expected, "\n")
	expected = dedent.Dedent(expected)
	expected = strings.ReplaceAll(expected, "\t", "  ")
	return expected
}

var _ = Describe("toYAMLWithComments", func() {

	It("decodes yaml with the comments", func() {
		type NestedType struct {
			FieldN1 string `desc:"nested field 1"`
		}
		type TestType struct {
			Field1  string                  `desc:"field descripting comment"`
			Field2  []NestedType            `desc:"list of field 2"`
			Field3  *NestedType             `desc:"a field that is a pointer to another type"`
			Field4  []*NestedType           `desc:"a field that is a pointer to a list of types"`
			Field5  string                  `json:"fieldfive" desc:"non standard field name"`
			Field6  []string                `desc:"a list of scalars"`
			Field7  map[string]string       `desc:"a map of scalars to scalars"`
			Field8  map[string]NestedType   `desc:"a map to a struct"`
			Field9  map[string]*NestedType  `desc:"a map to a pointer to a struct"`
			Field10 map[string][]NestedType `desc:"a map to a list of structs"`
		}
		node := render.ToNode(&TestType{
			Field1: "Hello, comments",
			Field2: []NestedType{
				{FieldN1: "Hello!"},
				{FieldN1: "Hello!"},
			},
			Field3: &NestedType{FieldN1: "Hello field 3"},
			Field4: []*NestedType{
				{FieldN1: "Field4"},
				nil,
			},
			Field5: "hello field 5",
			Field6: []string{"hello", "field", "six"},
			Field7: map[string]string{
				"hello":   "field seven",
				"field 7": "hello",
			},
			Field8: map[string]NestedType{
				"hello": {FieldN1: "field 8"},
			},
			Field9: map[string]*NestedType{
				"hello": {FieldN1: "field 9"},
			},
			Field10: map[string][]NestedType{
				"hello": {
					{FieldN1: "field"},
					{FieldN1: "ten"},
				},
			},
		}, &values.UserValuesInlineDocs{})
		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# field descripting comment
			Field1: Hello, comments
			# list of field 2
			Field2:
					- # nested field 1
						FieldN1: Hello!
					- # nested field 1
						FieldN1: Hello!
			# a field that is a pointer to another type
			Field3:
					# nested field 1
					FieldN1: Hello field 3
			# a field that is a pointer to a list of types
			Field4:
					- # nested field 1
						FieldN1: Field4
					- null
			# a list of scalars
			Field6:
					- hello
					- field
					- six
			# a map of scalars to scalars
			Field7:
					field 7: hello
					hello: field seven
			# a map to a struct
			Field8:
					hello:
							# nested field 1
							FieldN1: field 8
			# a map to a pointer to a struct
			Field9:
					hello:
							# nested field 1
							FieldN1: field 9
			# a map to a list of structs
			Field10:
					hello:
							- # nested field 1
								FieldN1: field
							- # nested field 1
								FieldN1: ten
			# non standard field name
			fieldfive: hello field 5
		`)))
	})

	It("handles nested types with fields implicitly inlined", func() {
		type NestedType struct {
			FieldN1 string `desc:"nested field 1"`
		}
		type TestType struct {
			NestedType
		}
		node := render.ToNode(&TestType{
			NestedType: NestedType{
				FieldN1: "Hello",
			},
		}, &values.UserValuesInlineDocs{})
		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# nested field 1
			FieldN1: Hello
		`)))
	})

	It("handles nested type where field name is same as type name", func() {
		type NestedType struct {
			NestedType string `desc:"nested field 1"`
		}
		type TestType struct {
			NestedType
		}
		node := render.ToNode(&TestType{
			NestedType: NestedType{
				NestedType: "Hello",
			},
		}, &values.UserValuesInlineDocs{})
		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# nested field 1
			NestedType: Hello
		`)))
	})

	It("handles structs thate are completely initalized to zero values", func() {
		type ChildType struct {
			FieldC1 string `desc:"field c1"`
		}
		type TestType struct {
			ChildType ChildType `json:"childType"`
		}
		node := render.ToNode(&TestType{}, &values.UserValuesInlineDocs{})
		actual := render.FromNode(node)
		log.Printf("%s", actual)
		Expect(actual).To(Equal(prepareExpected(`
			childType:
				  # field c1
				  FieldC1: ""
		`)))
	})

	It("handles nested types with fields explicitly inlined", func() {
		type NestedType struct {
			FieldN1 string `desc:"nested field 1"`
		}
		type TestType struct {
			NestedType `json:",inline"`
		}
		node := render.ToNode(&TestType{
			NestedType: NestedType{
				FieldN1: "Hello",
			},
		}, &values.UserValuesInlineDocs{})
		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# nested field 1
			FieldN1: Hello
		`)))
	})

	It("handles case where the type is an interface", func() {
		type NestedType struct {
			FieldN1 string `desc:"nested field 1"`
		}
		type TestType struct {
			Field1 interface{}            `desc:"my interface field"`
			Field2 map[string]interface{} `desc:"my map of interface field"`
			Field3 interface{}            `desc:"yet another interface field"`
			Field4 interface{}            `desc:"yet another interface field 2"`
		}

		node := render.ToNode(&TestType{
			Field1: NestedType{
				FieldN1: "field1 n1",
			},
			Field2: map[string]interface{}{
				"field2": NestedType{
					FieldN1: "Hello!",
				},
				"field3": &NestedType{
					FieldN1: "map value is pointer to nested type!",
				},
			},
			Field3: &NestedType{
				FieldN1: "interface is pointer to nested type!",
			},
			Field4: "hello, world!",
		}, &values.UserValuesInlineDocs{})

		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# my interface field
			Field1:
					# nested field 1
					FieldN1: field1 n1
			# my map of interface field
			Field2:
					field2:
							# nested field 1
							FieldN1: Hello!
					field3:
							# nested field 1
							FieldN1: map value is pointer to nested type!
			# yet another interface field
			Field3:
					# nested field 1
					FieldN1: interface is pointer to nested type!
			# yet another interface field 2
			Field4: hello, world!
	`)))
	})

	It("can handle case where type is a pointer to a pointer", func() {
		type NestedType struct {
			FieldN1 string `desc:"nested field 1"`
		}
		type TestType struct {
			Field1 **NestedType  `desc:"nested pointer type"`
			Field2 ***NestedType `desc:"doubly nested pointer type"`
		}

		n1 := &NestedType{
			FieldN1: "Hello, world!",
		}
		n2 := &n1
		node := render.ToNode(&TestType{
			Field1: &n1,
			Field2: &n2,
		}, &values.UserValuesInlineDocs{})

		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# nested pointer type
			Field1:
					# nested field 1
					FieldN1: Hello, world!
			# doubly nested pointer type
			Field2:
					# nested field 1
					FieldN1: Hello, world!
		`)))
	})

	It("will wrap the lines of the field comments", func() {
		type TestType struct {
			Field1 string `desc:"this comment should be wraped because it is long long long long long long long long lnog long long long"`
			Field2 string `desc:"this short comment shouldnt be wrapped"`
			Field3 string `desc:"this comment has a_word_that_shouldnt_be_broken_even_though_it_is_very_long_long_long_long_long words after long"`
		}

		node := render.ToNode(&TestType{
			Field1: "one",
			Field2: "two",
			Field3: "three",
		}, &values.UserValuesInlineDocs{
			LineLengthLimit: 40,
		})

		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			# this comment should be wraped because it
			# is long long long long long long long
			# long lnog long long long
			Field1: one
			# this short comment shouldnt be wrapped
			Field2: two
			# this comment has
			# a_word_that_shouldnt_be_broken_even_though_it_is_very_long_long_long_long_long
			# words after long
			Field3: three
	`)))
	})

	It("can be disabled via config", func() {
		type TestType struct {
			Field1 string `desc:"comment field 1"`
		}

		var config *values.UserValuesInlineDocs = nil // nil means disabled
		node := render.ToNode(&TestType{
			Field1: "one",
		}, config)

		Expect(render.FromNode(node)).To(Equal(prepareExpected(`
			Field1: one
		`)))
	})

	It("can merge operator values", func() {

		baseValues := map[string]interface{}{
			"operatorName": values.UserValues{
				Enabled: true,
				UserContainerValues: values.UserContainerValues{
					Env: []v1.EnvVar{{
						Name: "POD_NAMESPACE",
						ValueFrom: &v1.EnvVarSource{
							FieldRef: &v1.ObjectFieldSelector{
								FieldPath: "metadata.namespace",
							},
						},
					}},
					Image: values.Image{
						PullPolicy: "IfNotPresent",
						Registry:   "gcr.io/gloo-mesh",
						Repository: "foo",
						Tag:        "bar",
					},
				},
			},
		}

		customValues := map[string]interface{}{
			"operatorName": values.UserValues{
				Enabled: false,
				UserContainerValues: values.UserContainerValues{
					Env: []v1.EnvVar{{
						Name: "OTHER_VAR",
						ValueFrom: &v1.EnvVarSource{
							FieldRef: &v1.ObjectFieldSelector{
								FieldPath: "metadata.namespace",
							},
						},
					}},
				},
				ServicePorts: map[string]uint32{
					"https": 80,
				},
				FloatingUserID: true,
			},
		}

		var config *values.UserValuesInlineDocs = nil // nil means disabled
		mergedNode := render.MergeNodes(render.ToNode(baseValues, config), render.ToNode(customValues, config))

		fmt.Println(render.FromNode(mergedNode))
		Expect(render.FromNode(mergedNode)).To(ContainSubstring("enabled: false"))
		Expect(render.FromNode(mergedNode)).To(ContainSubstring("floatingUserId: true"))
		Expect(render.FromNode(mergedNode)).To(ContainSubstring("https: 80"))
		Expect(render.FromNode(mergedNode)).To(ContainSubstring("OTHER_VAR"))
		Expect(render.FromNode(mergedNode)).ToNot(ContainSubstring("POD_NAMESPACE"))

	})
})

var _ = Describe("toJSONSchema", func() {
	It("creates a schema for an object with properties from custom values and operators", func() {
		type Type1 struct {
			Field1a string `json:"field1a"`
		}

		type Type2 struct {
			Field2a string `json:"field2a"`
		}

		result := render.ToJSONSchema(values.UserHelmValues{
			CustomValues: &Type1{
				Field1a: "Hello",
			},
			Operators: []values.UserOperatorValues{
				{
					Name:         "operator1",
					Values:       values.UserValues{},
					CustomValues: &Type2{Field2a: "hello2"},
				},
				{
					Name:         "My_Operator_Two",
					ValuePath:    "my.values",
					Values:       values.UserValues{},
					CustomValues: &Type2{Field2a: "hello2"},
				},
			},
		})

		resultContainer := struct {
			Properties *struct {
				Field1a   map[string]interface{} `json:"field1a"`
				Operator1 *struct {
					Properties map[string]map[string]interface{}
				} `json:"operator1"`
				Operator2 *struct {
					Properties *struct {
						My *struct {
							Properties *struct {
								Values *struct {
									Properties map[string]map[string]interface{} `json:"properties"`
								} `json:"values"`
							} `json:"properties"`
						} `json:"my"`
					} `json:"properties"`
				} `json:"myOperatorTwo"`
			} `json:"properties"`
		}{}

		checkHasStandardValuesFields := func(valueProperties map[string]map[string]interface{}) {
			Expect(valueProperties["image"]["type"]).To(Equal("object"))
			Expect(valueProperties["enabled"]["type"]).To(Equal("boolean"))
			Expect(valueProperties["runAsUser"]["type"]).To(Equal("integer"))
		}

		Expect(json.Unmarshal([]byte(result), &resultContainer)).NotTo(HaveOccurred())

		Expect(resultContainer.Properties).NotTo(BeNil())
		Expect(resultContainer.Properties.Field1a).NotTo(BeNil())
		Expect(resultContainer.Properties.Field1a["type"]).To(Equal("string"))

		Expect(resultContainer.Properties.Operator1).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator1.Properties).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator1.Properties["field2a"]["type"]).To(Equal("string"))
		checkHasStandardValuesFields(resultContainer.Properties.Operator1.Properties)

		Expect(resultContainer.Properties.Operator2).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator2.Properties).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator2.Properties.My).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator2.Properties.My.Properties).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator2.Properties.My.Properties.Values).NotTo(BeNil())
		Expect(resultContainer.Properties.Operator2.Properties.My.Properties.Values.Properties["field2a"]["type"]).To(Equal("string"))
		checkHasStandardValuesFields(resultContainer.Properties.Operator2.Properties.My.Properties.Values.Properties)
	})

	It("adds some json schema behaviour based on the jsonschema tags", func() {
		type Type1 struct {
			Field1a string `json:"field1a" jsonschema:"title=field a,description=the field called a,example=aaa,example=bbb,default=a"`
		}
		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type1{}})
		expected := prepareExpected(`
			{
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"field1a": {
						"type": "string",
						"title": "field a",
						"description": "the field called a",
						"default": "a",
						"examples": [
							"aaa",
							"bbb"
						]
					}
				}
			}`)
		Expect(result).To(Equal(expected))
	})

	It("will correctly generate schema for types that should be nullable", func() {
		type Type1 struct {
			Field1 map[string]interface{}
			Field2 []*struct {
				Field3 string
			}
		}

		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type1{}})
		expected := prepareExpected(`
			{
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"Field1": {
						"anyOf": [
							{
								"type": "object"
							},
							{
								"type": "null"
							}
						]
					},
					"Field2": {
						"anyOf": [
							{
								"items": {
									"properties": {
										"Field3": {
											"type": "string"
										}
									},
									"type": "object"
								},
								"type": "array"
							},
							{
								"type": "null"
							}
						]
					}
				}
			}`)
		Expect(result).To(Equal(expected))
	})

	It("handles nullable struct fields correctly", func() {
		type Type1 struct {
		}
		type Type2 struct {
			Field2A *Type1 // should be nullable
			Field2B *Type1 `json:"field2_b"` // check it handles renamed fields as well
		}

		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type2{}})
		expected := prepareExpected(`
			{
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"Field2A": {
						"anyOf": [
							{
								"properties": {},
								"type": "object"
							},
							{
								"type": "null"
							}
						]
					},
					"field2_b": {
						"anyOf": [
							{
								"properties": {},
								"type": "object"
							},
							{
								"type": "null"
							}
						]
					}
				}
			}`)
		Expect(expected).To(Equal(result))
	})

	It("will allow pbstructs to be deserialized as anything", func() {
		type Type1 struct {
			Field1 *structpb.Value
		}
		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type1{}})
		expected := prepareExpected(`
			{
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"Field1": {
						"anyOf": [
							{
								"type": "null"
							},
							{
								"type": "number"
							},
							{
								"type": "string"
							},
							{
								"type": "boolean"
							},
							{
								"type": "object"
							},
							{
								"type": "array"
							}
						]
					}
				}
			}`)
		Expect(result).To(Equal(expected))
	})

	It("maps time as a nullable string", func() {
		type Type1 struct {
			Field1 metav1.Time `json:"metatime"`
		}
		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type1{}})
		expected := prepareExpected(`
			{
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"metatime": {
						"anyOf": [
							{
								"type": "string",
								"format": "date-time"
							},
							{
								"type": "null"
							}
						]
					}
				}
			}`)
		Expect(result).To(Equal(expected))
	})

	It("maps resource quantity as either a string or number", func() {
		type Type1 struct {
			Field1 resource.Quantity
		}
		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type1{}})
		expected := prepareExpected(`
			{
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"Field1": {
						"anyOf": [
							{
								"type": "string"
							},
							{
								"type": "number"
							}
						]
					}
				}
			}`)
		Expect(result).To(Equal(expected))
	})

	It("maps security context as something that can be mapped to a boolean", func() {
		type Type1 struct {
			Field1 *v1.SecurityContext
		}
		result := render.ToJSONSchema(values.UserHelmValues{CustomValues: &Type1{}})
		resultContainer := struct {
			Properties struct {
				Field1 *struct {
					AnyOf []*struct {
						Const interface{} `json:"const"`
						Type  string      `json:"type"`
					} `json:"anyOf"`
				}
			} `json:"properties"`
		}{}
		Expect(json.Unmarshal([]byte(result), &resultContainer)).NotTo(HaveOccurred())
		Expect(resultContainer.Properties.Field1).NotTo(BeNil())
		Expect(resultContainer.Properties.Field1).NotTo(BeNil())
		Expect(resultContainer.Properties.Field1.AnyOf).NotTo(BeNil())
		Expect(resultContainer.Properties.Field1.AnyOf).To(HaveLen(2))
		Expect(resultContainer.Properties.Field1.AnyOf[0].Type).To(Equal("object"))
		Expect(resultContainer.Properties.Field1.AnyOf[1].Type).To(Equal("boolean"))
		Expect(resultContainer.Properties.Field1.AnyOf[1].Const).To(BeFalse())
	})

	Context("custom type mappings", func() {
		type MyJsonSchemaType struct {
			Type  string `json:"type"`
			AnyOf []interface{}
		}

		It("will invoke custom json schema type mappings", func() {
			type Type1 struct{}
			type Type2 struct {
				FieldA Type1
			}

			mapAsNumber := func(map[string]interface{}) interface{} {
				return &MyJsonSchemaType{
					Type: "number",
				}
			}

			typeMapper := func(t reflect.Type, s map[string]interface{}) interface{} {
				if t == reflect.TypeOf(Type1{}) {
					return mapAsNumber(s)
				}
				return nil
			}

			result := render.ToJSONSchema(values.UserHelmValues{
				CustomValues: &Type2{},
				JsonSchema: values.UserJsonSchema{
					CustomTypeMapper: typeMapper,
				},
			})
			expected := prepareExpected(`
				{
					"$schema": "https://json-schema.org/draft/2020-12/schema",
					"properties": {
						"FieldA": {
							"type": "number"
						}
					}
				}`)
			Expect(result).To(Equal(expected))
		})

		It("allows performing modification of the original type", func() {
			type Type1 struct{}
			type Type2 struct {
				FieldA Type1
			}

			mapAsPossiblyBoolean := func(original map[string]interface{}) interface{} {
				return MyJsonSchemaType{
					AnyOf: []interface{}{
						original,
						&MyJsonSchemaType{Type: "boolean"},
					},
				}
			}

			typeMapper := func(t reflect.Type, s map[string]interface{}) interface{} {
				if t == reflect.TypeOf(Type1{}) {
					return mapAsPossiblyBoolean(s)
				}
				return nil
			}

			result := render.ToJSONSchema(values.UserHelmValues{
				CustomValues: &Type2{},
				JsonSchema: values.UserJsonSchema{
					CustomTypeMapper: typeMapper,
				},
			})
			expected := prepareExpected(`
				{
					"$schema": "https://json-schema.org/draft/2020-12/schema",
					"properties": {
						"FieldA": {
							"anyOf": [
								{
									"properties": {},
									"type": "object"
								},
								{
									"type": "boolean"
								}
							]
						}
					}
				}`)
			Expect(result).To(Equal(expected))
		})
	})
})
