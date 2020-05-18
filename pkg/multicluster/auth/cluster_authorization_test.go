package auth_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/core/v1"
	mock_k8s_rbac_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/rbac.authorization.k8s.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	mock_auth "github.com/solo-io/skv2/pkg/multicluster/auth/mocks"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

var _ = Describe("Cluster authorization", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context

		coreClientset *mock_k8s_core_clients.MockClientset
		saClient      *mock_k8s_core_clients.MockServiceAccountClient
		rbacClientset *mock_k8s_rbac_clients.MockClientset
		crbClient     *mock_k8s_rbac_clients.MockClusterRoleBindingClient

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

		testErr = eris.New("test-err")
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		saClient = mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		coreClientset = mock_k8s_core_clients.NewMockClientset(ctrl)
		coreClientset.EXPECT().ServiceAccounts().Return(saClient)

		crbClient = mock_k8s_rbac_clients.NewMockClusterRoleBindingClient(ctrl)
		rbacClientset = mock_k8s_rbac_clients.NewMockClientset(ctrl)
		rbacClientset.EXPECT().ClusterRoleBindings().Return(crbClient)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("reports an error when the service account can't be created", func() {
		mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)
		clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, rbacClientset, coreClientset)

		saClient.EXPECT().
			UpsertServiceAccount(ctx, &corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      saName,
					Namespace: saNamespace,
				},
			}).Return(testErr)

		outputBearerToken, err := clusterAuthClient.BuildRemoteBearerToken(ctx, testKubeConfig, saName, saNamespace)

		Expect(outputBearerToken).To(BeEmpty(), "Should not have created a new config")
		Expect(err).To(Equal(testErr), "Should have reported the expected error")
	})

	It("will fail if ClusterRoleBinding fails to upsert", func() {
		mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)

		clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, rbacClientset, coreClientset)

		sa := &corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      saName,
				Namespace: saNamespace,
			},
		}
		saClient.EXPECT().
			UpsertServiceAccount(ctx, sa).
			Return(nil)

		crbClient.EXPECT().
			UpsertClusterRoleBinding(ctx, &rbacv1.ClusterRoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), auth.ServiceAccountRoles[0].GetName()),
				},
				Subjects: []rbacv1.Subject{
					{
						Kind:      "ServiceAccount",
						Name:      sa.GetName(),
						Namespace: sa.GetNamespace(),
					},
				},
				RoleRef: rbacv1.RoleRef{
					APIGroup: "rbac.authorization.k8s.io",
					Kind:     "ClusterRole",
					Name:     auth.ServiceAccountRoles[0].GetName(),
				},
			}).
			Return(testErr)

		outputBearerToken, err := clusterAuthClient.BuildRemoteBearerToken(
			ctx,
			testKubeConfig,
			saName,
			saNamespace,
		)


		Expect(outputBearerToken).To(BeEmpty(), "Should not have created a new config")
		Expect(err).To(Equal(testErr), "Should have reported the expected error")
	})

	It("works when its clients work", func() {
		mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)

		clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, rbacClientset, coreClientset)

		sa := &corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      saName,
				Namespace: saNamespace,
			},
		}
		saClient.EXPECT().
			UpsertServiceAccount(ctx, sa).
			Return(nil)

		crbClient.EXPECT().
			UpsertClusterRoleBinding(ctx, &rbacv1.ClusterRoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), auth.ServiceAccountRoles[0].GetName()),
				},
				Subjects: []rbacv1.Subject{
					{
						Kind:      "ServiceAccount",
						Name:      sa.GetName(),
						Namespace: sa.GetNamespace(),
					},
				},
				RoleRef: rbacv1.RoleRef{
					APIGroup: "rbac.authorization.k8s.io",
					Kind:     "ClusterRole",
					Name:     auth.ServiceAccountRoles[0].GetName(),
				},
			}).
			Return(nil)

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
})
