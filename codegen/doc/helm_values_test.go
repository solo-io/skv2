package doc_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/codegen/doc"
)

var _ = Describe("GenerateHelmValuesDoc", func() {
	It("handles the case where the field has 'omitChildren' and json tag", func() {
		type ChildType2 struct {
		}
		type ChildType struct {
			Field1 *ChildType2 `json:"myCoolField" desc:"my field" omitChildren:"true"`
		}
		result := doc.GenerateHelmValuesDoc(
			ChildType{},
			"test",
			"my test",
		)
		expected := doc.HelmValues{
			{
				Key:          "test",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "my test",
			},
			{
				Key:          "test.myCoolField",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "my field",
			},
		}
		Expect(result).To(Equal(expected))
	})

	It("handles case where map value struct has a nil pointer field that omits children", func() {
		type ChildType2 struct {
		}
		type ChildType struct {
			Field1 *ChildType2 `desc:"my field" omitChildren:"true"`
		}

		type TestType struct {
			Sidecars map[string]ChildType `json:"sidecars"`
		}

		result := doc.GenerateHelmValuesDoc(
			TestType{},
			"test",
			"my test",
		)
		expected := doc.HelmValues{
			{
				Key:          "test",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "my test",
			},
			{
				Key:          "test.sidecars",
				Type:         "map[string, struct]",
				DefaultValue: "null",
				Description:  "",
			},
			{
				Key:          "test.sidecars.<MAP_KEY>",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "",
			},
			{
				Key:          "test.sidecars.<MAP_KEY>.Field1",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "my field",
			},
		}
		Expect(result).To(Equal(expected))
	})

	It("handles slices", func() {
		type ChildType struct {
			Field1 []string `json:"myCoolField" desc:"my field"`
		}

		c := ChildType{
			Field1: []string{"default"},
		}

		result := doc.GenerateHelmValuesDoc(
			c,
			"test",
			"my test",
		)
		expected := doc.HelmValues{
			{
				Key:          "test",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "my test",
			},
			{
				Key:          "test.myCoolField[]",
				Type:         "[]string",
				DefaultValue: "[\"default\"]",
				Description:  "my field",
			},
		}
		Expect(result).To(Equal(expected))
	})

	It("handles slices of structs", func() {
		type ChildType2 struct {
			Field2 string `json:"myCoolField2" desc:"my field 2"`
		}
		type ChildType struct {
			Field1 []string `json:"myCoolField" desc:"my field"`
		}
		type Parent struct {
			ChildType  []ChildType `json:"childType" desc:"child type"`
			ChildType2 ChildType2  `json:"childType2" desc:"child type 2"`
		}
		p := Parent{
			ChildType: []ChildType{
				{
					Field1: []string{"default"},
				},
			},
			ChildType2: ChildType2{
				Field2: "default",
			},
		}
		result := doc.GenerateHelmValuesDoc(
			p,
			"test",
			"my test",
		)
		expected := doc.HelmValues{
			{
				Key:          "test",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "my test",
			},
			{
				Key:          "test.childType[]",
				Type:         "[]struct",
				DefaultValue: "[{\"myCoolField\":[\"default\"]}]",
				Description:  "child type",
			},
			{
				Key:          "test.childType[]",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "child type",
			},
			{
				Key:          "test.childType[].myCoolField[]",
				Type:         "[]string",
				DefaultValue: " ",
				Description:  "my field",
			},
			{
				Key:          "test.childType2",
				Type:         "struct",
				DefaultValue: " ",
				Description:  "child type 2",
			},
			{
				Key:          "test.childType2.myCoolField2",
				Type:         "string",
				DefaultValue: "default",
				Description:  "my field 2",
			},
		}
		Expect(result).To(Equal(expected))
	})
})
