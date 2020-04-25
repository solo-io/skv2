package auth_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	. "github.com/solo-io/go-utils/testutils"
	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/core/v1"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	mock_auth "github.com/solo-io/skv2/pkg/multicluster/auth/mocks"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Remote service account client", func() {
	var (
		ctx         context.Context
		ctrl        *gomock.Controller
		saName      = "test-sa"
		saNamespace = "test-ns"
		roles       = append([]*k8s_rbac_types.ClusterRole{}, auth.ServiceAccountRoles...)
		testErr     = eris.New("hello")
		notFoundErr = &errors.StatusError{
			ErrStatus: k8s_meta_types.Status{
				Reason: k8s_meta_types.StatusReasonAlreadyExists,
			},
		}
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("works", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		rbacBinder := mock_auth.NewMockRbacBinder(ctrl)

		remoteAuthManager := auth.NewRemoteAuthorityManager(saClient, rbacBinder)

		serviceAccount := &k8s_core_types.ServiceAccount{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      saName,
				Namespace: saNamespace,
			},
		}

		saClient.
			EXPECT().
			CreateServiceAccount(ctx, serviceAccount).
			Return(nil)

		rbacBinder.
			EXPECT().
			BindClusterRolesToServiceAccount(ctx, serviceAccount, roles).
			Return(nil)

		sa, err := remoteAuthManager.ApplyRemoteServiceAccount(ctx, saName, saNamespace, roles)
		Expect(err).NotTo(HaveOccurred())
		Expect(sa).To(Equal(serviceAccount))
	})

	It("will try and update if create fails", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		rbacBinder := mock_auth.NewMockRbacBinder(ctrl)

		remoteAuthManager := auth.NewRemoteAuthorityManager(saClient, rbacBinder)

		serviceAccount := &k8s_core_types.ServiceAccount{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      saName,
				Namespace: saNamespace,
			},
		}

		saClient.
			EXPECT().
			CreateServiceAccount(ctx, serviceAccount).
			Return(notFoundErr)

		saClient.
			EXPECT().
			UpdateServiceAccount(ctx, serviceAccount).
			Return(testErr)

		sa, err := remoteAuthManager.ApplyRemoteServiceAccount(ctx, saName, saNamespace, roles)
		Expect(err).To(HaveInErrorChain(testErr))
		Expect(sa).To(BeNil())
	})

	It("reports an error if service account creation fails, on not IsAlreadyExists error", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		rbacBinder := mock_auth.NewMockRbacBinder(ctrl)

		remoteAuthManager := auth.NewRemoteAuthorityManager(saClient, rbacBinder)

		serviceAccount := &k8s_core_types.ServiceAccount{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      saName,
				Namespace: saNamespace,
			},
		}

		saClient.
			EXPECT().
			CreateServiceAccount(ctx, serviceAccount).
			Return(testErr)

		sa, err := remoteAuthManager.ApplyRemoteServiceAccount(ctx, saName, saNamespace, roles)
		Expect(err).To(HaveInErrorChain(testErr))
		Expect(sa).To(BeNil())
	})

	It("reports an error if role binding fails", func() {
		saClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		rbacBinder := mock_auth.NewMockRbacBinder(ctrl)

		remoteAuthManager := auth.NewRemoteAuthorityManager(saClient, rbacBinder)

		serviceAccount := &k8s_core_types.ServiceAccount{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      saName,
				Namespace: saNamespace,
			},
		}

		saClient.
			EXPECT().
			CreateServiceAccount(ctx, serviceAccount).
			Return(nil)

		rbacBinder.
			EXPECT().
			BindClusterRolesToServiceAccount(ctx, serviceAccount, roles).
			Return(testErr)

		sa, err := remoteAuthManager.ApplyRemoteServiceAccount(ctx, saName, saNamespace, roles)
		Expect(err).To(HaveInErrorChain(testErr))
		Expect(sa).To(BeNil())
	})
})
