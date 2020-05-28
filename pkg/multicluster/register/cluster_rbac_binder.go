package register

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	EmptyRolesListError        = eris.New("Empty Roles list found, must specify at least one role to bind to.")
	EmptyClusterRolesListError = eris.New("Empty ClusterRoles list found, must specify at least one role to bind to.")
)

type clusterAuthorization struct {
	clusterRoleBindingClient rbac_v1.ClusterRoleBindingClient
	roleBindingClient        rbac_v1.RoleBindingClient
	secretClient             k8s_core_v1.SecretClient
	serviceAccountClient     k8s_core_v1.ServiceAccountClient
}

func NewClusterRBACBinderFactory() ClusterRBACBinderFactory {
	return func(cfg clientcmd.ClientConfig) (ClusterRBACBinder, error) {
		restCfg, err := cfg.ClientConfig()
		if err != nil {
			return nil, err
		}
		rbacClientset, err := rbac_v1.NewClientsetFromConfig(restCfg)
		if err != nil {
			return nil, err
		}
		return NewClusterRBACBinder(
			rbacClientset.ClusterRoleBindings(),
			rbacClientset.RoleBindings(),
		), nil
	}
}

func NewClusterRBACBinder(
	clusterRoleBindingClient rbac_v1.ClusterRoleBindingClient,
	roleBindingClient rbac_v1.RoleBindingClient,
) ClusterRBACBinder {
	return &clusterAuthorization{
		clusterRoleBindingClient: clusterRoleBindingClient,
		roleBindingClient:        roleBindingClient,
	}
}

func (c *clusterAuthorization) BindClusterRoles(
	ctx context.Context,
	serviceAccount client.ObjectKey,
	clusterRoles []client.ObjectKey,
) error {

	if len(clusterRoles) == 0 {
		return EmptyClusterRolesListError
	}

	return c.bindClusterRolesToServiceAccount(ctx, serviceAccount, clusterRoles)
}

func (c *clusterAuthorization) bindClusterRolesToServiceAccount(
	ctx context.Context,
	targetServiceAccount client.ObjectKey,
	roles []client.ObjectKey,
) error {

	for _, role := range roles {
		crbToCreate := &k8s_rbac_types.ClusterRoleBinding{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", targetServiceAccount.Name, role.Name),
			},
			Subjects: []k8s_rbac_types.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.Name,
				Namespace: targetServiceAccount.Namespace,
			}},
			RoleRef: k8s_rbac_types.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     role.Name,
			},
		}

		if err := c.clusterRoleBindingClient.UpsertClusterRoleBinding(ctx, crbToCreate); err != nil {
			return err
		}
	}

	return nil
}

func (c *clusterAuthorization) BindRoles(
	ctx context.Context,
	serviceAccount client.ObjectKey,
	roles []client.ObjectKey,
)  error {

	if len(roles) == 0 {
		return EmptyRolesListError
	}

	return c.bindRolesToServiceAccount(ctx, serviceAccount, roles)
}

func (c *clusterAuthorization) bindRolesToServiceAccount(
	ctx context.Context,
	targetServiceAccount client.ObjectKey,
	roles []client.ObjectKey,
) error {

	for _, role := range roles {
		rbToCreate := &k8s_rbac_types.RoleBinding{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-role-binding", targetServiceAccount.Name, role.Name),
			},
			Subjects: []k8s_rbac_types.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.Name,
				Namespace: targetServiceAccount.Namespace,
			}},
			RoleRef: k8s_rbac_types.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
				Name:     role.Name,
			},
		}

		if err := c.roleBindingClient.UpsertRoleBinding(ctx, rbToCreate); err != nil {
			return err
		}
	}

	return nil
}
