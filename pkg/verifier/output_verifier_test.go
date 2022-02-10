package verifier_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	. "github.com/solo-io/skv2/pkg/verifier"
)

var _ = Describe("Output Verifier", func() {
	It("verifies the server resources", func() {
		cfg, err := config.GetConfig()
		if err != nil {
			Skip("skipping verifier test, requires active kubernetes cluster, failed to get rest config: " + err.Error())
		}
		disc := discovery.NewDiscoveryClientForConfigOrDie(cfg)

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

		v := NewOutputVerifier(context.TODO(), disc, map[schema.GroupVersionKind]ServerVerifyOption{
			gvkDoesntExist: ServerVerifyOption_ErrorIfNotPresent,
		})

		resourceExists, err := v.VerifyServerResource("", gvkDoesntExist)
		Expect(err).To(HaveOccurred())

		resourceExists, err = v.VerifyServerResource("", gvkDoesExist)
		Expect(err).NotTo(HaveOccurred())
		Expect(resourceExists).To(BeTrue())

		// ignore errors on doesn't exist
		v = NewOutputVerifier(context.TODO(), disc, map[schema.GroupVersionKind]ServerVerifyOption{
			gvkDoesntExist: ServerVerifyOption_WarnIfNotPresent,
		})

		resourceExists, err = v.VerifyServerResource("", gvkDoesntExist)
		Expect(err).NotTo(HaveOccurred())
		Expect(resourceExists).To(BeFalse())
	})
})
