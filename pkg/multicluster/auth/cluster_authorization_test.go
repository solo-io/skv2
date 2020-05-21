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
	"github.com/solo-io/skv2/test"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

var _ = Describe("Cluster authorization", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context

		saClient  *mock_k8s_core_clients.MockServiceAccountClient
		crbClient *mock_k8s_rbac_clients.MockClusterRoleBindingClient
		rbClient  *mock_k8s_rbac_clients.MockRoleBindingClient

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

		rbClient = mock_k8s_rbac_clients.NewMockRoleBindingClient(ctrl)
		crbClient = mock_k8s_rbac_clients.NewMockClusterRoleBindingClient(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("reports an error when the service account can't be created", func() {
		mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)
		clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, crbClient, rbClient, saClient)

		saClient.EXPECT().
			UpsertServiceAccount(ctx, &corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      saName,
					Namespace: saNamespace,
				},
			}).Return(testErr)

		outputBearerToken, err := clusterAuthClient.BuildClusterScopedRemoteBearerToken(
			ctx,
			testKubeConfig,
			saName,
			saNamespace,
			test.ServiceAccountClusterAdminRoles,
		)

		Expect(outputBearerToken).To(BeEmpty(), "Should not have created a new config")
		Expect(err).To(Equal(testErr), "Should have reported the expected error")
	})

	Context("ClusterRoleBindings", func() {

		It("will fail if ClusterRoleBinding fails to upsert", func() {
			mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)

			clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, crbClient, rbClient, saClient)

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
						Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), test.ServiceAccountClusterAdminRoles[0].GetName()),
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
						Name:     test.ServiceAccountClusterAdminRoles[0].GetName(),
					},
				}).
				Return(testErr)

			outputBearerToken, err := clusterAuthClient.BuildClusterScopedRemoteBearerToken(
				ctx,
				testKubeConfig,
				saName,
				saNamespace,
				test.ServiceAccountClusterAdminRoles,
			)

			Expect(outputBearerToken).To(BeEmpty(), "Should not have created a new config")
			Expect(err).To(Equal(testErr), "Should have reported the expected error")
		})

		It("works when its clients work", func() {
			mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)

			clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, crbClient, rbClient, saClient)

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
						Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), test.ServiceAccountClusterAdminRoles[0].GetName()),
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
						Name:     test.ServiceAccountClusterAdminRoles[0].GetName(),
					},
				}).
				Return(nil)

			mockConfigCreator.
				EXPECT().
				ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace).
				Return(serviceAccountKubeConfig, nil)

			outputBearerToken, err := clusterAuthClient.BuildClusterScopedRemoteBearerToken(
				ctx,
				testKubeConfig,
				saName,
				saNamespace,
				test.ServiceAccountClusterAdminRoles,
			)

			Expect(err).NotTo(HaveOccurred(), "An error should not have occurred")
			Expect(outputBearerToken).To(Equal(serviceAccountBearerToken), "Should have returned the expected kube config")
		})

	})

	Context("RoleBinding", func() {

		It("will fail if RoleBinding fails to upsert", func() {
			mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)

			clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, crbClient, rbClient, saClient)

			sa := &corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      saName,
					Namespace: saNamespace,
				},
			}
			saClient.EXPECT().
				UpsertServiceAccount(ctx, sa).
				Return(nil)

			role := &rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-role-name",
				},
			}

			rbClient.EXPECT().
				UpsertRoleBinding(ctx, &rbacv1.RoleBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s-role-binding", sa.GetName(), role.GetName()),
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
						Kind:     "Role",
						Name:     role.GetName(),
					},
				}).Return(testErr)

			outputBearerToken, err := clusterAuthClient.BuildRemoteBearerToken(
				ctx,
				testKubeConfig,
				saName,
				saNamespace,
				[]*rbacv1.Role{role},
			)

			Expect(outputBearerToken).To(BeEmpty(), "Should not have created a new config")
			Expect(err).To(Equal(testErr), "Should have reported the expected error")
		})

		It("works when its clients work", func() {
			mockConfigCreator := mock_auth.NewMockRemoteAuthorityConfigCreator(ctrl)

			clusterAuthClient := auth.NewClusterAuthorization(mockConfigCreator, crbClient, rbClient, saClient)

			sa := &corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      saName,
					Namespace: saNamespace,
				},
			}
			saClient.EXPECT().
				UpsertServiceAccount(ctx, sa).
				Return(nil)

			role := &rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-role-name",
				},
			}

			rbClient.EXPECT().
				UpsertRoleBinding(ctx, &rbacv1.RoleBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s-role-binding", sa.GetName(), role.GetName()),
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
						Kind:     "Role",
						Name:     role.GetName(),
					},
				}).Return(nil)

			mockConfigCreator.
				EXPECT().
				ConfigFromRemoteServiceAccount(ctx, testKubeConfig, saName, saNamespace).
				Return(serviceAccountKubeConfig, nil)

			outputBearerToken, err := clusterAuthClient.BuildRemoteBearerToken(
				ctx,
				testKubeConfig,
				saName,
				saNamespace,
				[]*rbacv1.Role{role},
			)

			Expect(err).NotTo(HaveOccurred(), "An error should not have occurred")
			Expect(outputBearerToken).To(Equal(serviceAccountBearerToken), "Should have returned the expected kube config")
		})

	})
})
