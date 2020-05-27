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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	EmptyRolesListError        = eris.New("Empty Roles list found, must specify at least one role to bind to.")
	EmptyClusterRolesListError = eris.New("Empty ClusterRoles list found, must specify at least one role to bind to.")
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
		return NewClusterAuthorization(
			cfgCreator,
			rbacClientset.ClusterRoleBindings(),
			rbacClientset.RoleBindings(),
			v1Clientset.ServiceAccounts(),
		), nil
	}
}

func NewClusterAuthorization(
	configCreator RemoteAuthorityConfigCreator,
	clusterRoleBindingClient rbac_v1.ClusterRoleBindingClient,
	roleBindingClient rbac_v1.RoleBindingClient,
	serviceAccountClient k8s_core_v1.ServiceAccountClient,
) ClusterAuthorization {
	return &clusterAuthorization{
		configCreator:            configCreator,
		clusterRoleBindingClient: clusterRoleBindingClient,
		roleBindingClient:        roleBindingClient,
		serviceAccountClient:     serviceAccountClient,
	}
}

func (c *clusterAuthorization) BuildClusterScopedRemoteBearerToken(
	ctx context.Context,
	targetClusterCfg *rest.Config,
	name, namespace string,
	clusterRoles []client.ObjectKey,
) (bearerToken string, err error) {

	if len(clusterRoles) == 0 {
		return "", EmptyClusterRolesListError
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

	if err = c.bindClusterRolesToServiceAccount(ctx, saToCreate, clusterRoles); err != nil {
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
	roles []client.ObjectKey,
) error {

	for _, role := range roles {
		crbToCreate := &k8s_rbac_types.ClusterRoleBinding{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", targetServiceAccount.GetName(), role.Name),
			},
			Subjects: []k8s_rbac_types.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.GetName(),
				Namespace: targetServiceAccount.GetNamespace(),
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

func (c *clusterAuthorization) BuildRemoteBearerToken(
	ctx context.Context,
	targetClusterCfg *rest.Config,
	name, namespace string,
	roles []client.ObjectKey,
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
	roles []client.ObjectKey,
) error {

	for _, role := range roles {
		rbToCreate := &k8s_rbac_types.RoleBinding{
			ObjectMeta: k8s_meta.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-role-binding", targetServiceAccount.GetName(), role.Name),
			},
			Subjects: []k8s_rbac_types.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.GetName(),
				Namespace: targetServiceAccount.GetNamespace(),
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
