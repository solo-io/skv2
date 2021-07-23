package render_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/render"
)

var _ = Describe("ManifestsRenderer", func() {

	It("should set version without patch", func() {
		m := new(metav1.ObjectMeta)
		render.SetVersionForObject(m, "1.0.0-patch1")
		Expect(m.Annotations[model.CRDVersionKey]).To(Equal("1.0.0"))
	})

	It("should set version without metadata", func() {
		m := new(metav1.ObjectMeta)
		render.SetVersionForObject(m, "1.0.0+metadata")
		Expect(m.Annotations[model.CRDVersionKey]).To(Equal("1.0.0"))
	})
	It("should set version with v", func() {
		m := new(metav1.ObjectMeta)
		render.SetVersionForObject(m, "v1.2.3+metadata")
		Expect(m.Annotations[model.CRDVersionKey]).To(Equal("1.2.3"))
	})
})
