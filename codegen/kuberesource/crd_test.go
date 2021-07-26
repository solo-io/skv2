package kuberesource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/solo-io/skv2/codegen/kuberesource"
	"github.com/solo-io/skv2/codegen/model"
)

var _ = Describe("Crd", func() {

	Describe("CRD upgrade checks", func() {
		DescribeTable("DoesCrdNeedUpgrade",
			func(currentProductVersion, currentCrdHash, crdProductVersion, crdCrdHash string, expected bool) {
				annotations := make(map[string]string)
				annotations[model.CRDVersionKey] = crdProductVersion
				annotations[model.CRDSpecHashKey] = crdCrdHash
				answer, err := DoesCrdNeedUpgrade(currentProductVersion, currentCrdHash, annotations)
				Expect(err).NotTo(HaveOccurred())
				Expect(answer).To(Equal(expected))
			},
			Entry("everything same", "1.0", "123", "1.0", "123", false),
			// this might happen during dev cycle
			Entry("same version, different hash", "1.0", "123", "1.0", "13", true),
			Entry("higher beta version, different hash", "1.1.0-beta2", "123", "1.1.0-beta1", "13", true),
			Entry("lower beta version, different hash", "1.1.0-beta1", "123", "1.1.0-beta2", "13", false),
			Entry("different hash, smaller version", "1.0", "123", "0.9", "13", true),
			Entry("different hash, higher version", "1.0", "123", "1.9", "13", false),
			Entry("same hash, smaller version", "1.0", "123", "0.9", "123", false),
			Entry("same hash, higher version", "1.0", "123", "1.9", "123", false),
		)

		It("should return CrdNeedsUpgrade", func() {
			crd := apiextv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
					Annotations: map[string]string{
						model.CRDSpecHashKey: "123",
						model.CRDVersionKey:  "1.0",
					},
				},
			}
			newProdCrdInfo := model.CRDMetadata{
				Version: "1.1",
				CRDS: []model.CRDAnnotations{
					{
						Name: "foo",
						Hash: "456",
					},
				},
			}

			errmap := DoCrdsNeedUpgrade(newProdCrdInfo, []apiextv1beta1.CustomResourceDefinition{crd})
			Expect(errmap).To(HaveKeyWithValue(crd.Name, BeAssignableToTypeOf(&CrdNeedsUpgrade{})))
		})
		It("should return CrdNeedsUpgrade", func() {
			crd1 := apiextv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
					Annotations: map[string]string{
						model.CRDSpecHashKey: "123",
						model.CRDVersionKey:  "1.2",
					},
				},
			}
			crd2 := apiextv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "bar",
					Annotations: map[string]string{
						model.CRDSpecHashKey: "123",
						model.CRDVersionKey:  "1.1",
					},
				},
			}
			newProdCrdInfo := model.CRDMetadata{
				Version: "1.1",
				CRDS: []model.CRDAnnotations{
					{
						Name: "foo",
						Hash: "456",
					},
				},
			}

			errmap := DoCrdsNeedUpgrade(newProdCrdInfo, []apiextv1beta1.CustomResourceDefinition{crd1, crd2})
			Expect(errmap).NotTo(HaveKey(crd1.Name))
			Expect(errmap).To(HaveKeyWithValue(crd2.Name, BeAssignableToTypeOf(&CrdNotFound{})))
		})
	})

	Describe("CRD gen", func() {
		var (
			grp model.Group
		)

		BeforeEach(func() {
			grp = model.Group{
				Resources: []model.Resource{
					{
						Kind: "kind",
						Spec: model.Field{
							Type: model.Type{
								Name:    "test",
								Message: &v1.AcrylicType{},
							},
						},
					},
				},
			}
			grp.Init()
		})

		It("should generate spec hash", func() {
			grp.SkipSpecHash = false
			o, err := CustomResourceDefinitions(grp)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(1))
			// note: we intentionally provide the "d18828e563010e32" hash in the test, as it shouldn't change
			// between runs.
			Expect(o[0].GetAnnotations()).To(HaveKeyWithValue(model.CRDSpecHashKey, "d18828e563010e32"))

		})
		It("should not generate spec hash", func() {
			grp.SkipSpecHash = true
			o, err := CustomResourceDefinitions(grp)
			Expect(err).NotTo(HaveOccurred())
			Expect(o).To(HaveLen(1))
			// note: we intentionally provide the "d18828e563010e32" hash in the test, as it shouldn't change
			// between runs.
			Expect(o[0].GetAnnotations()).NotTo(HaveKey(model.CRDSpecHashKey))
		})
	})

})
