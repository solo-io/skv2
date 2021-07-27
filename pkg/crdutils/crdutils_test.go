package crdutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/solo-io/skv2/pkg/crdutils"
)

var _ = Describe("CrdUtils", func() {

	Describe("CRD upgrade checks", func() {
		DescribeTable("DoesCrdNeedUpgrade",
			func(currentProductVersion, currentCrdHash, crdProductVersion, crdCrdHash string, expected bool) {
				annotations := make(map[string]string)
				annotations[CRDVersionKey] = crdProductVersion
				annotations[CRDSpecHashKey] = crdCrdHash
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
						CRDSpecHashKey: "123",
						CRDVersionKey:  "1.0",
					},
				},
			}
			newProdCrdInfo := CRDMetadata{
				Version: "1.1",
				CRDS: []CRDAnnotations{
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
						CRDSpecHashKey: "123",
						CRDVersionKey:  "1.2",
					},
				},
			}
			crd2 := apiextv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "bar",
					Annotations: map[string]string{
						CRDSpecHashKey: "123",
						CRDVersionKey:  "1.1",
					},
				},
			}
			newProdCrdInfo := CRDMetadata{
				Version: "1.1",
				CRDS: []CRDAnnotations{
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
	Describe("Parse annotation", func() {
		It("Parse CRDMetadataKeyannotation", func() {
			crdMeta, err := ParseCRDMetadataFromAnnotations(map[string]string{
				CRDMetadataKey: `{"version":"1.2.3","crds":[{"name":"test","hash":"123"}]}`,
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(crdMeta.Version).To(Equal("1.2.3"))
			Expect(crdMeta.CRDS).To(HaveLen(1))
			Expect(crdMeta.CRDS[0].Name).To(Equal("test"))
			Expect(crdMeta.CRDS[0].Hash).To(Equal("123"))
		})
		It("errors on invalid json", func() {
			crdMeta, err := ParseCRDMetadataFromAnnotations(map[string]string{
				CRDMetadataKey: `not json`,
			})
			Expect(crdMeta).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
		It("doesnt error when annotation missing", func() {
			crdMeta, err := ParseCRDMetadataFromAnnotations(map[string]string{})
			Expect(crdMeta).To(BeNil())
			Expect(err).NotTo(HaveOccurred())
		})
		It("doesnt error when annotation nil", func() {
			crdMeta, err := ParseCRDMetadataFromAnnotations(nil)
			Expect(crdMeta).To(BeNil())
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
