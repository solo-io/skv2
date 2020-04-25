package auth_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	mock_auth "github.com/solo-io/skv2/pkg/multicluster/auth/mocks"
	"k8s.io/client-go/rest"
)

var _ = Describe("Cluster authorization", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context

		saName      = "test-service-account"
		saNamespace = "test-ns"

		testKubeConfig = &rest.Config{
			Host: "www.new-test-who-dis.edu",
			TLSClientConfig: rest.TLSClientConfig{
				CertData: []byte("super secure cert data"),
			},
		}
		serviceAccountBearerToken = "test-token"
		serviceAccountKubeConfig  = &rest.Config{
			Host:        "www.new-test-who-dis.edu",
			BearerToken: serviceAccountBearerToken,
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("works when its clients work", func() {
		mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)
		mockRemoteAuthorityManager := mock_auth.NewMockRemoteAuthorityManager(ctrl)

		clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, mockRemoteAuthorityManager)

		mockRemoteAuthorityManager.
			EXPECT().
			ApplyRemoteServiceAccount(ctx, saName, saNamespace, auth.ServiceAccountRoles).
			Return(nil, nil)

		mockConfigCreator.
			EXPECT().
			ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace).
			Return(serviceAccountKubeConfig, nil)

		outputBearerToken, err := clusterAuthClient.BuildRemoteBearerToken(
			ctx,
			testKubeConfig,
			saName,
			saNamespace,
		)

		Expect(err).NotTo(HaveOccurred(), "An error should not have occurred")
		Expect(outputBearerToken).To(Equal(serviceAccountBearerToken), "Should have returned the expected kube config")
	})

	It("reports an error when the service account can't be created", func() {
		mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)
		mockRemoteAuthorityManager := mock_auth.NewMockRemoteAuthorityManager(ctrl)

		clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, mockRemoteAuthorityManager)

		testErr := errors.New("test-err")

		mockRemoteAuthorityManager.
			EXPECT().
			ApplyRemoteServiceAccount(ctx, saName, saNamespace, auth.ServiceAccountRoles).
			Return(nil, testErr)

		outputBearerToken, err := clusterAuthClient.BuildRemoteBearerToken(ctx, testKubeConfig, saName, saNamespace)

		Expect(outputBearerToken).To(BeEmpty(), "Should not have created a new config")
		Expect(err).To(Equal(testErr), "Should have reported the expected error")
	})
})
