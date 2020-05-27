package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/rotisserie/eris"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	core_v1 "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	EmptyRolesListError        = eris.New("Empty Roles list found, must specify at least one role to bind to.")
	EmptyClusterRolesListError = eris.New("Empty ClusterRoles list found, must specify at least one role to bind to.")
)

const (
	// visible for testing
	SecretTokenKey = "token"
)

var (
	// exponential backoff retry with an initial period of 0.1s for 7 iterations, which will mean a cumulative retry period of ~6s
	// visible for testing
	SecretLookupOpts = []retry.Option{
		retry.Delay(time.Millisecond * 100),
		retry.Attempts(7),
		retry.DelayType(retry.BackOffDelay),
	}
)

type clusterAuthorization struct {
	configCreator            RemoteAuthorityConfigCreator
	clusterRoleBindingClient rbac_v1.ClusterRoleBindingClient
	roleBindingClient        rbac_v1.RoleBindingClient
	secretClient             k8s_core_v1.SecretClient
	serviceAccountClient     k8s_core_v1.ServiceAccountClient
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
		), nil
	}
}

func NewClusterAuthorization(
	configCreator RemoteAuthorityConfigCreator,
	clusterRoleBindingClient rbac_v1.ClusterRoleBindingClient,
	roleBindingClient rbac_v1.RoleBindingClient,
) ClusterAuthorization {
	return &clusterAuthorization{
		configCreator:            configCreator,
		clusterRoleBindingClient: clusterRoleBindingClient,
		roleBindingClient:        roleBindingClient,
	}
}

func (c *clusterAuthorization) BuildClusterScopedRemoteBearerToken(
	ctx context.Context,
	serviceAccount client.ObjectKey,
	clusterRoles []client.ObjectKey,
) (bearerToken string, err error) {

	if len(clusterRoles) == 0 {
		return "", EmptyClusterRolesListError
	}

	if err = c.bindClusterRolesToServiceAccount(ctx, serviceAccount, clusterRoles); err != nil {
		return "", err
	}

	return c.getTokenForSa(ctx, serviceAccount)
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

func (c *clusterAuthorization) BuildRemoteBearerToken(
	ctx context.Context,
	serviceAccount client.ObjectKey,
	roles []client.ObjectKey,
) (bearerToken string, err error) {

	if len(roles) == 0 {
		return "", EmptyRolesListError
	}

	if err = c.bindRolesToServiceAccount(ctx, serviceAccount, roles); err != nil {
		return "", err
	}

	return c.getTokenForSa(ctx, serviceAccount)
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

func (c *clusterAuthorization) getTokenForSa(
	ctx context.Context,
	saRef client.ObjectKey,
) (string, error) {
	sa, err := c.serviceAccountClient.GetServiceAccount(ctx, saRef)
	if err != nil {
		return "", err
	}
	if len(sa.Secrets) == 0 {
		return "", eris.Errorf(
			"service account %s.%s does not have a token secret associated with it",
			saRef.Name,
			saRef.Namespace,
		)
	}
	var foundSecret *core_v1.Secret
	if err = retry.Do(func() error {
		secretName := sa.Secrets[0].Name
		secret, err := c.secretClient.GetSecret(ctx, client.ObjectKey{Name: secretName, Namespace: saRef.Namespace})
		if err != nil {
			return err
		}

		foundSecret = secret
		return nil
	}, SecretLookupOpts...); err != nil {
		return "", SecretNotReady(err)
	}

	serviceAccountToken, ok := foundSecret.Data[SecretTokenKey]
	if !ok {
		return "", MalformedSecret
	}

	return string(serviceAccountToken), nil
}
