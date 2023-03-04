package doc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/codegen/doc"
)

var _ = Describe("GenerateHelmValuesDoc", func() {
	It("handles case where map value struct has a nil pointer field that omits children", func() {
		type ChildType2 struct {
		}
		type ChildType struct {
			Field1 *ChildType2 `json:"field1" desc:"my field" omitChildren:"true"`
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
})
