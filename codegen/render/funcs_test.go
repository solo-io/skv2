package render_test

import (
	"fmt"
	"strings"

	"github.com/lithammer/dedent"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	"github.com/solo-io/skv2/codegen/model/values"
	"github.com/solo-io/skv2/codegen/render"
)

var _ = Describe("toYAMLWithComments", func() {
	prepareExpected := func(expected string) string {
		expected = strings.Trim(expected, "\n")
		expected = dedent.Dedent(expected)
		expected = strings.ReplaceAll(expected, "\t", "  ")
		return expected
	}

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
