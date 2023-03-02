package doc_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/codegen/doc"
)

var _ = Describe("GenerateHelmValuesDoc", func() {
	It("handles case where map value struct has a nil pointer field that omits children", func() {
		type ChildType2 struct {
		}
		type UserContainerValues struct {
			Field1 *ChildType2 `json:"field1" desc:"my field" omitChildren:"true"`
		}

		type TestType struct {
			// Child    *ChildType `json:"child" desc:"my child" omitChildren:"true"`
			Sidecars map[string]UserContainerValues
		}

		result := doc.GenerateHelmValuesDoc(
			TestType{
				// Child: &childType
			},
			"albert",
			"mr cool guy",
		)
		log.Printf("Result %v", result)
		Expect(true).To(BeFalse())
	})
})
