package auth

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	core_v1 "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	// visible for testing
	ServiceAccountRoles = []*k8s_rbac_types.ClusterRole{{
		ObjectMeta: k8s_meta.ObjectMeta{Name: "cluster-admin"},
	}}

	EmptyRolesListError = eris.New("Empty Roles list found, must specify at least one role to bind to.")
)

type clusterAuthorization struct {
	configCreator            RemoteAuthorityConfigCreator
	serviceAccountClient     k8s_core_v1.ServiceAccountClient
	clusterRoleBindingClient rbac_v1.ClusterRoleBindingClient
	roleBindingClient        rbac_v1.RoleBindingClient
}

func NewClusterAuthorizationFactory() ClusterAuthorizationFactory {
	return func(cfg clientcmd.ClientConfig) (ClusterAuthorization, error) {
		restCfg, err := cfg.ClientConfig()
		if err != nil {
			return nil, err
		}
		v1Clientset, err := k8s_core_v1.NewClientsetFromConfig(restCfg)
		if err != nil {
			return nil, err
		}
		rbacClientset, err := rbac_v1.NewClientsetFromConfig(restCfg)
		if err != nil {
			return nil, err
		}
		cfgCreator := NewRemoteAuthorityConfigCreator(v1Clientset.Secrets(), v1Clientset.ServiceAccounts())
		return NewClusterAuthorization(cfgCreator, rbacClientset, v1Clientset), nil
	}
}

func NewClusterAuthorization(
	configCreator RemoteAuthorityConfigCreator,
	rbacClientset rbac_v1.Clientset,
	coreClientset k8s_core_v1.Clientset,
) ClusterAuthorization {
	return &clusterAuthorization{
		configCreator:            configCreator,
		clusterRoleBindingClient: rbacClientset.ClusterRoleBindings(),
		roleBindingClient:        rbacClientset.RoleBindings(),
		serviceAccountClient:     coreClientset.ServiceAccounts(),
	}
}

func (c *clusterAuthorization) BuildClusterScopedRemoteBearerToken(
	ctx context.Context,
	targetClusterCfg *rest.Config,
	name, namespace string,
	clusterRoles ...*k8s_rbac_types.ClusterRole,
) (bearerToken string, err error) {

	saToCreate := &core_v1.ServiceAccount{
		ObjectMeta: k8s_meta.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	if err = c.serviceAccountClient.UpsertServiceAccount(ctx, saToCreate); err != nil {
		return "", err
	}

	roles := ServiceAccountRoles
	if len(clusterRoles) != 0 {
		roles = clusterRoles
	}

	if err = c.bindClusterRolesToServiceAccount(ctx, saToCreate, roles); err != nil {
		return "", err
	}

	saConfig, err := c.configCreator.ConfigFromRemoteServiceAccount(ctx, targetClusterCfg, name, namespace)
	if err != nil {
		return "", err
	}

	// we only want the bearer token for that service account
	return saConfig.BearerToken, nil
}

func (c *clusterAuthorization) bindClusterRolesToServiceAccount(
	ctx context.Context,
	targetServiceAccount *core_v1.ServiceAccount,
	roles []*k8s_rbac_types.ClusterRole,
) error {

	for _, role := range roles {
		crbToCreate := &k8s_rbac_types.ClusterRoleBinding{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", targetServiceAccount.GetName(), role.GetName()),
			},
			Subjects: []k8s_rbac_types.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.GetName(),
				Namespace: targetServiceAccount.GetNamespace(),
			}},
			RoleRef: k8s_rbac_types.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     role.GetName(),
			},
		}

		if err := c.clusterRoleBindingClient.UpsertClusterRoleBinding(ctx, crbToCreate); err != nil {
			return err
		}
	}

	return nil
}

func (c *clusterAuthorization) BuildRemoteBearerToken(
	ctx context.Context,
	targetClusterCfg *rest.Config,
	name, namespace string,
	roles []*k8s_rbac_types.Role,
) (bearerToken string, err error) {

	if len(roles) == 0 {
		return "", EmptyRolesListError
	}

	saToCreate := &core_v1.ServiceAccount{
		ObjectMeta: k8s_meta.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	if err = c.serviceAccountClient.UpsertServiceAccount(ctx, saToCreate); err != nil {
		return "", err
	}

	if err = c.bindRolesToServiceAccount(ctx, saToCreate, roles); err != nil {
		return "", err
	}

	saConfig, err := c.configCreator.ConfigFromRemoteServiceAccount(ctx, targetClusterCfg, name, namespace)
	if err != nil {
		return "", err
	}

	// we only want the bearer token for that service account
	return saConfig.BearerToken, nil
}

func (c *clusterAuthorization) bindRolesToServiceAccount(
	ctx context.Context,
	targetServiceAccount *core_v1.ServiceAccount,
	roles []*k8s_rbac_types.Role,
) error {

	for _, role := range roles {
		rbToCreate := &k8s_rbac_types.RoleBinding{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-role-binding", targetServiceAccount.GetName(), role.GetName()),
			},
			Subjects: []k8s_rbac_types.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.GetName(),
				Namespace: targetServiceAccount.GetNamespace(),
			}},
			RoleRef: k8s_rbac_types.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
				Name:     role.GetName(),
			},
		}

		if err := c.roleBindingClient.UpsertRoleBinding(ctx, rbToCreate); err != nil {
			return err
		}
	}

	return nil
}
