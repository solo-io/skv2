package verifier_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/skv2/pkg/verifier"
	mock_discovery "github.com/solo-io/skv2/pkg/verifier/mocks"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//go:generate mockgen -destination mocks/discovery.go k8s.io/client-go/discovery DiscoveryInterface

var _ = Describe(
	"Output Verifier", func() {

		var (
			ctrl *gomock.Controller

			mockDiscovery *mock_discovery.MockDiscoveryInterface
		)

		BeforeEach(
			func() {
				ctrl = gomock.NewController(GinkgoT())

				mockDiscovery = mock_discovery.NewMockDiscoveryInterface(ctrl)
			},
		)

		AfterEach(
			func() {
				ctrl.Finish()
			},
		)

		It(
			"verifies the server resources", func() {

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

				v := NewVerifier(
					context.TODO(), mockDiscovery, map[schema.GroupVersionKind]ServerVerifyOption{
						gvkDoesntExist: ServerVerifyOption_ErrorIfNotPresent,
						gvkDoesExist:   ServerVerifyOption_WarnIfNotPresent,
					},
				)

				mockDiscovery.EXPECT().
					ServerResourcesForGroupVersion(gvkDoesntExist.GroupVersion().String()).
					Return(nil, errors.NewNotFound(schema.GroupResource{}, "")).
					Times(2)

				resourceExists, err := v.VerifyServerResource("", gvkDoesntExist)
				Expect(err).To(HaveOccurred())

				mockDiscovery.EXPECT().
					ServerResourcesForGroupVersion(gvkDoesExist.GroupVersion().String()).
					Return(
						&metav1.APIResourceList{
							TypeMeta:     metav1.TypeMeta{},
							GroupVersion: gvkDoesExist.GroupVersion().String(),
							APIResources: []metav1.APIResource{
								{
									Kind: "Secret",
								},
							},
						},
						nil,
					)

				resourceExists, err = v.VerifyServerResource("", gvkDoesExist)
				Expect(err).NotTo(HaveOccurred())
				Expect(resourceExists).To(BeTrue())

				// Test mock doesn't get called twice
				resourceExists, err = v.VerifyServerResource("", gvkDoesExist)
				Expect(err).NotTo(HaveOccurred())
				Expect(resourceExists).To(BeTrue())

				// ignore errors on doesn't exist
				v = NewVerifier(
					context.TODO(), mockDiscovery, map[schema.GroupVersionKind]ServerVerifyOption{
						gvkDoesntExist: ServerVerifyOption_WarnIfNotPresent,
					},
				)

				resourceExists, err = v.VerifyServerResource("", gvkDoesntExist)
				Expect(err).NotTo(HaveOccurred())
				Expect(resourceExists).To(BeFalse())
			},
		)
	},
)
