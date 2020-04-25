package auth_test

import (
	"context"
	"errors"
	"time"

	"github.com/avast/retry-go"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/go-utils/testutils"
	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/core/v1"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	k8s_core_types "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Config creator", func() {
	var (
		ctx         context.Context
		ctrl        *gomock.Controller
		saName      = "test-sa"
		saNamespace = "test-ns"

		tokenSecretRef = k8s_core_types.ObjectReference{
			Name: "test-secret",
		}
		secret = &k8s_core_types.Secret{
			Data: map[string][]byte{
				auth.SecretTokenKey: []byte("my-test-token"),
			},
		}
		testKubeConfig = &rest.Config{
			Host: "www.new-test-who-dis.edu",
			TLSClientConfig: rest.TLSClientConfig{
				CertData: []byte("super secure cert data"),
			},
		}
	)

	BeforeEach(func() {
		ctx = context.TODO()
		ctrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("works when the service account is immediately ready", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		secretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)

		remoteAuthConfigCreator := auth.NewRemoteAuthorityConfigCreator(secretClient, saClient)

		saClient.
			EXPECT().
			GetServiceAccount(ctx, client.ObjectKey{Name: saName, Namespace: saNamespace}).
			Return(&k8s_core_types.ServiceAccount{
				Secrets: []k8s_core_types.ObjectReference{tokenSecretRef},
			}, nil)

		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: tokenSecretRef.Name, Namespace: saNamespace}).
			Return(secret, nil)

		newCfg, err := remoteAuthConfigCreator.ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace)

		Expect(err).NotTo(HaveOccurred())
		Expect(newCfg.TLSClientConfig.CertData).To(BeEmpty())
		Expect([]byte(newCfg.BearerToken)).To(Equal(secret.Data[auth.SecretTokenKey]))
	})

	It("works when the service account eventually has a secret attached to it", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		secretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)

		remoteAuthConfigCreator := auth.NewRemoteAuthorityConfigCreator(secretClient, saClient)

		attemptsRemaining := 3
		saClient.
			EXPECT().
			GetServiceAccount(ctx, client.ObjectKey{Name: saName, Namespace: saNamespace}).
			DoAndReturn(func(ctx context.Context, key client.ObjectKey) (*k8s_core_types.ServiceAccount, error) {
				attemptsRemaining -= 1
				if attemptsRemaining > 0 {
					return nil, errors.New("whoops not ready yet")
				}

				return &k8s_core_types.ServiceAccount{
					Secrets: []k8s_core_types.ObjectReference{tokenSecretRef},
				}, nil
			}).
			AnyTimes()

		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: tokenSecretRef.Name, Namespace: saNamespace}).
			Return(secret, nil)

		newCfg, err := remoteAuthConfigCreator.ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace)

		Expect(err).NotTo(HaveOccurred())
		Expect(newCfg.TLSClientConfig.CertData).To(BeEmpty())
		Expect([]byte(newCfg.BearerToken)).To(Equal(secret.Data[auth.SecretTokenKey]))
	})

	It("works when the service account is immediately ready, and the CA data is in a file", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		secretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)

		fileTestKubeConfig := &rest.Config{
			Host: "www.grahams-a-great-programmer.edu",
			TLSClientConfig: rest.TLSClientConfig{
				CAFile:   "path-to-ca-file",
				CertData: []byte("super secure cert data"),
			},
		}

		remoteAuthConfigCreator := auth.NewRemoteAuthorityConfigCreator(secretClient, saClient)

		saClient.
			EXPECT().
			GetServiceAccount(ctx, client.ObjectKey{Name: saName, Namespace: saNamespace}).
			Return(&k8s_core_types.ServiceAccount{
				Secrets: []k8s_core_types.ObjectReference{tokenSecretRef},
			}, nil)

		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: tokenSecretRef.Name, Namespace: saNamespace}).
			Return(secret, nil)

		newCfg, err := remoteAuthConfigCreator.ConfigFromRemoteServiceAccount(ctx, fileTestKubeConfig, saName, saNamespace)

		Expect(err).NotTo(HaveOccurred())
		Expect(newCfg.TLSClientConfig.CertData).To(BeEmpty())
		Expect(newCfg.TLSClientConfig.CAData).To(BeEmpty())
		Expect(newCfg.TLSClientConfig.CAFile).To(Equal(fileTestKubeConfig.TLSClientConfig.CAFile))
		Expect([]byte(newCfg.BearerToken)).To(Equal(secret.Data[auth.SecretTokenKey]))
	})

	It("returns an error when the secret is malformed", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		secretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)

		remoteAuthConfigCreator := auth.NewRemoteAuthorityConfigCreator(secretClient, saClient)

		saClient.
			EXPECT().
			GetServiceAccount(ctx, client.ObjectKey{Name: saName, Namespace: saNamespace}).
			Return(&k8s_core_types.ServiceAccount{
				Secrets: []k8s_core_types.ObjectReference{tokenSecretRef},
			}, nil)

		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: tokenSecretRef.Name, Namespace: saNamespace}).
			Return(&k8s_core_types.Secret{Data: map[string][]byte{"whoops wrong key": []byte("yikes")}}, nil)

		newCfg, err := remoteAuthConfigCreator.ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace)

		Expect(err).To(Equal(auth.MalformedSecret))
		Expect(err).To(HaveInErrorChain(auth.MalformedSecret))
		Expect(newCfg).To(BeNil())
	})

	It("returns an error if the secret never appears", func() {
		auth.SecretLookupOpts = []retry.Option{
			retry.Delay(time.Millisecond * 1),
			retry.Attempts(7),
			retry.DelayType(retry.FixedDelay),
		}
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		secretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)

		remoteAuthConfigCreator := auth.NewRemoteAuthorityConfigCreator(secretClient, saClient)

		saClient.
			EXPECT().
			GetServiceAccount(ctx, client.ObjectKey{Name: saName, Namespace: saNamespace}).
			Return(&k8s_core_types.ServiceAccount{
				Secrets: []k8s_core_types.ObjectReference{tokenSecretRef},
			}, nil).
			AnyTimes()

		testErr := errors.New("not ready yet")

		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: tokenSecretRef.Name, Namespace: saNamespace}).
			Return(nil, testErr).
			AnyTimes()

		newCfg, err := remoteAuthConfigCreator.ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace)

		Expect(err).To(HaveInErrorChain(auth.SecretNotReady(errors.New("test-err"))))
		Expect(newCfg).To(BeNil())
	})
})
