package register

import (
	"context"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1"
	k8s_core_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/providers"
	rbac_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/providers"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

/*
	RegisterClusterFromConfig is meant to be a helper function to easily "register" a remote cluster.
	Curently this entails:
		1. Creating a `ServiceAccount` on the remote cluster.
		2. Binding RBAC `Roles/ClusterRoles` to said `ServiceAccount`
		3. And finally creating a kubeconfig `Secret` with the BearerToken of the remote `ServiceAccount`
*/
func RegisterClusterFromConfig(
	ctx context.Context,
	remoteCfg clientcmd.ClientConfig,
	opts RbacOptions,
	registrant ClusterRegistrant,
) error {
	sa, err := registrant.EnsureRemoteServiceAccount(ctx, remoteCfg, opts.Options)
	if err != nil {
		return err
	}

	token, err := registrant.CreateRemoteAccessToken(ctx, remoteCfg, client.ObjectKey{
		Namespace: sa.GetNamespace(),
		Name:      sa.GetName(),
	}, opts)
	if err != nil {
		return err
	}

	return registrant.RegisterClusterWithToken(ctx, remoteCfg, token, opts.Options)
}

/*
	DefaultRegistrant provider function. This function will create a `ClusterRegistrant` using the
	current kubeconfig, and the specified context. It will build all of the dependencies from the
	availble `ClientConfig`.

	Meant to be used in tandem with RegisterClusterFromConfig above.
	They are exposed separately so the `Registrant` may be mocked for the function above.
*/
func DefaultRegistrant(context string) (ClusterRegistrant, error) {
	cfg, err := config.GetConfigWithContext(context)
	if err != nil {
		return nil, err
	}
	clientset, err := k8s_core_v1.NewClientsetFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	nsClientFactory := k8s_core_v1_providers.NamespaceClientFromConfigFactoryProvider()
	secretClientFactory := k8s_core_v1_providers.SecretClientFromConfigFactoryProvider()
	saClientFactory := k8s_core_v1_providers.ServiceAccountClientFromConfigFactoryProvider()
	roleClientFactory := rbac_v1_providers.RoleClientFromConfigFactoryProvider()
	clusterRoleClientFactory := rbac_v1_providers.ClusterRoleClientFromConfigFactoryProvider()
	clusterRBACBinderFactory := NewClusterRBACBinderFactory()
	registrant := NewClusterRegistrant(
		clusterRBACBinderFactory,
		clientset.Secrets(),
		secretClientFactory,
		nsClientFactory,
		saClientFactory,
		clusterRoleClientFactory,
		roleClientFactory,
	)
	return registrant, nil
}
