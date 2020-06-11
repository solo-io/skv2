package internal_test

import (
	"context"
	"fmt"

	mock_k8s_rbac_clients "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/multicluster/register/internal"
	"github.com/solo-io/skv2/test"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Cluster authorization", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context

		crbClient *mock_k8s_rbac_clients.MockClusterRoleBindingClient
		rbClient  *mock_k8s_rbac_clients.MockRoleBindingClient

		saName      = "sa-name"
		saNamespace = "sa-namespace"

		saObjectKey = func() client.ObjectKey {
			return client.ObjectKey{
				Namespace: saNamespace,
				Name:      saName,
			}
		}

		testErr = eris.New("test-err")
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		rbClient = mock_k8s_rbac_clients.NewMockRoleBindingClient(ctrl)
		crbClient = mock_k8s_rbac_clients.NewMockClusterRoleBindingClient(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("ClusterRoleBindings", func() {

		It("will fail if ClusterRoleBinding fails to upsert", func() {
			clusterRbacBinder := internal.NewClusterRBACBinder(crbClient, rbClient)

			sa := saObjectKey()

			crbClient.EXPECT().
				UpsertClusterRoleBinding(ctx, &rbacv1.ClusterRoleBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.Name, test.ServiceAccountClusterAdminRoles[0].Name),
					},
					Subjects: []rbacv1.Subject{
						{
							Kind:      "ServiceAccount",
							Name:      sa.Name,
							Namespace: sa.Namespace,
						},
					},
					RoleRef: rbacv1.RoleRef{
						APIGroup: "rbac.authorization.k8s.io",
						Kind:     "ClusterRole",
						Name:     test.ServiceAccountClusterAdminRoles[0].Name,
					},
				}).
				Return(testErr)

			err := clusterRbacBinder.BindClusterRoles(
				ctx,
				sa,
				[]client.ObjectKey{{
					Name: test.ServiceAccountClusterAdminRoles[0].Name,
				}},
			)

			Expect(err).To(Equal(testErr), "Should have reported the expected error")
		})

		It("works when its clients work", func() {
			clusterRbacBinder := internal.NewClusterRBACBinder(crbClient, rbClient)

			sa := saObjectKey()

			crbClient.EXPECT().
				UpsertClusterRoleBinding(ctx, &rbacv1.ClusterRoleBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.Name, test.ServiceAccountClusterAdminRoles[0].Name),
					},
					Subjects: []rbacv1.Subject{
						{
							Kind:      "ServiceAccount",
							Name:      sa.Name,
							Namespace: sa.Namespace,
						},
					},
					RoleRef: rbacv1.RoleRef{
						APIGroup: "rbac.authorization.k8s.io",
						Kind:     "ClusterRole",
						Name:     test.ServiceAccountClusterAdminRoles[0].Name,
					},
				}).
				Return(nil)

			err := clusterRbacBinder.BindClusterRoles(
				ctx,
				sa,
				[]client.ObjectKey{{
					Name: test.ServiceAccountClusterAdminRoles[0].Name,
				}},
			)

			Expect(err).NotTo(HaveOccurred(), "An error should not have occurred")
		})

	})

	Context("RoleBinding", func() {

		It("will fail if RoleBinding fails to upsert", func() {
			clusterRbacBinder := internal.NewClusterRBACBinder(crbClient, rbClient)

			sa := saObjectKey()

			role := &rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-role-name",
				},
			}

			rbClient.EXPECT().
				UpsertRoleBinding(ctx, &rbacv1.RoleBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s-role-binding", sa.Name, role.GetName()),
					},
					Subjects: []rbacv1.Subject{
						{
							Kind:      "ServiceAccount",
							Name:      sa.Name,
							Namespace: sa.Namespace,
						},
					},
					RoleRef: rbacv1.RoleRef{
						APIGroup: "rbac.authorization.k8s.io",
						Kind:     "Role",
						Name:     role.GetName(),
					},
				}).Return(testErr)

			err := clusterRbacBinder.BindRoles(
				ctx,
				sa,
				[]client.ObjectKey{{
					Namespace: role.GetNamespace(),
					Name:      role.GetName(),
				}},
			)

			Expect(err).To(Equal(testErr), "Should have reported the expected error")
		})

		It("works when its clients work", func() {
			clusterRbacBinder := internal.NewClusterRBACBinder(crbClient, rbClient)

			sa := saObjectKey()

			role := &rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-role-name",
				},
			}

			rbClient.EXPECT().
				UpsertRoleBinding(ctx, &rbacv1.RoleBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s-role-binding", sa.Name, role.GetName()),
					},
					Subjects: []rbacv1.Subject{
						{
							Kind:      "ServiceAccount",
							Name:      sa.Name,
							Namespace: sa.Namespace,
						},
					},
					RoleRef: rbacv1.RoleRef{
						APIGroup: "rbac.authorization.k8s.io",
						Kind:     "Role",
						Name:     role.GetName(),
					},
				}).Return(nil)

			err := clusterRbacBinder.BindRoles(
				ctx,
				sa,
				[]client.ObjectKey{{
					Namespace: role.GetNamespace(),
					Name:      role.GetName(),
				}},
			)

			Expect(err).NotTo(HaveOccurred(), "An error should not have occurred")
		})

	})
})
