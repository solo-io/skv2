package kuberesource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/skv2/codegen/kuberesource"
	"github.com/solo-io/skv2/codegen/model"
)

var _ = Describe("Crd", func() {

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

})
