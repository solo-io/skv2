package verifier_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	. "github.com/solo-io/skv2/pkg/verifier"
)

var _ = Describe("Verifier", func() {
	It("verifies the server resources", func() {
		cfg, err := config.GetConfig()
		if err != nil {
			Skip("skipping verifier test, requires active kubernetes cluster, failed to get rest config: " + err.Error())
		}

		gvkDoesntExist := schema.GroupVersionKind{
			Group:   "doesnt",
			Version: "v1",
			Kind:    "Exist",
		}
		gvkDoesExist := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Secret",
		}

		v := NewVerifier(context.TODO(), map[schema.GroupVersionKind]ServerVerifyOption{
			gvkDoesntExist: ServerVerifyOption_ErrorIfNotPresent,
		})

		resourceExists, err := v.VerifyServerResource("", cfg, gvkDoesntExist)
		Expect(err).To(HaveOccurred())

		resourceExists, err = v.VerifyServerResource("", cfg, gvkDoesExist)
		Expect(err).NotTo(HaveOccurred())
		Expect(resourceExists).To(BeTrue())

		// ignore errors on doesn't exist
		v = NewVerifier(context.TODO(), map[schema.GroupVersionKind]ServerVerifyOption{
			gvkDoesntExist: ServerVerifyOption_WarnIfNotPresent,
		})

		resourceExists, err = v.VerifyServerResource("", cfg, gvkDoesntExist)
		Expect(err).NotTo(HaveOccurred())
		Expect(resourceExists).To(BeFalse())
	})
})
