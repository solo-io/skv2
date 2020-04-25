package auth_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	. "github.com/solo-io/go-utils/testutils"
	mock_k8s_rbac_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/rbac.authorization.k8s.io/v1"
	auth "github.com/solo-io/skv2/pkg/multicluster/auth"
	v1 "k8s.io/api/core/v1"
	rbacapiv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ = Describe("default rbac binder", func() {
	var (
		ctx           context.Context
		ctrl          *gomock.Controller
		rbacBinder    auth.RbacBinder
		rbacClientset *mock_k8s_rbac_clients.MockClientset
		crbClient     *mock_k8s_rbac_clients.MockClusterRoleBindingClient

		testErr = eris.New("hello")
		name    = "hello"
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())
		rbacClientset = mock_k8s_rbac_clients.NewMockClientset(ctrl)
		crbClient = mock_k8s_rbac_clients.NewMockClusterRoleBindingClient(ctrl)
		rbacClientset.EXPECT().ClusterRoleBindings().Return(crbClient)

		rbacBinder = auth.NewRbacBinder(rbacClientset)
	})

	It("will attempt to update if create fails with already exist", func() {
		sa := &v1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: name,
			},
		}
		roles := []*rbacapiv1.ClusterRole{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
			},
		}

		crb := &rbacapiv1.ClusterRoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ClusterRoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), name),
			},
			Subjects: []rbacapiv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      sa.GetName(),
					Namespace: sa.GetNamespace(),
				},
			},
			RoleRef: rbacapiv1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     roles[0].GetName(),
			},
		}

		crbClient.EXPECT().
			CreateClusterRoleBinding(ctx, crb).
			Return(errors.NewAlreadyExists(schema.GroupResource{}, ""))

		crbClient.EXPECT().
			UpdateClusterRoleBinding(ctx, crb).
			Return(testErr)

		err := rbacBinder.BindClusterRolesToServiceAccount(ctx, sa, roles)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))

	})

	It("will return result if create fails with any other error", func() {
		sa := &v1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: name,
			},
		}
		roles := []*rbacapiv1.ClusterRole{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
			},
		}

		crb := &rbacapiv1.ClusterRoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ClusterRoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), name),
			},
			Subjects: []rbacapiv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      sa.GetName(),
					Namespace: sa.GetNamespace(),
				},
			},
			RoleRef: rbacapiv1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     roles[0].GetName(),
			},
		}

		crbClient.EXPECT().
			CreateClusterRoleBinding(ctx, crb).
			Return(testErr)

		err := rbacBinder.BindClusterRolesToServiceAccount(ctx, sa, roles)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will succeed if no errors occur", func() {
		sa := &v1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: name,
			},
		}
		roles := []*rbacapiv1.ClusterRole{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
			},
		}

		crb := &rbacapiv1.ClusterRoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ClusterRoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", sa.GetName(), name),
			},
			Subjects: []rbacapiv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      sa.GetName(),
					Namespace: sa.GetNamespace(),
				},
			},
			RoleRef: rbacapiv1.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     roles[0].GetName(),
			},
		}

		crbClient.EXPECT().
			CreateClusterRoleBinding(ctx, crb).
			Return(nil)

		err := rbacBinder.BindClusterRolesToServiceAccount(ctx, sa, roles)
		Expect(err).NotTo(HaveOccurred())
	})
})
