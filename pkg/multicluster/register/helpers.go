package register

import (
	"context"
	"os"

	"github.com/hashicorp/go-multierror"
	v1alpha1_providers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/providers"

	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/rest"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1"
	k8s_core_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/providers"
	rbac_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/providers"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// Options for registering a cluster
type RegistrationOptions struct {
	// override the url of the k8s server
	MasterURL string

	// override the path of the local kubeconfig
	KubeCfgPath string

	// override the context to use from the local kubeconfig.
	// if unset, use current context
	KubeContext string

	// override the path of the remote kubeconfig
	RemoteKubeCfgPath string

	// override the context to use from the remote kubeconfig
	// if unset, use current context
	RemoteKubeContext string

	// localAPIServerAddress is optional. When passed in, it will overwrite the Api Server endpoint in
	//	the kubeconfig before it is written. This is primarily useful when running multi cluster KinD environments
	//	on a mac as  the local IP needs to be re-written to `host.docker.internal` so that the local instance
	//	knows to hit localhost.
	APIServerAddress string

	// Name by which the cluster will be identified
	// If left empty will return error
	ClusterName string

	// Namespace to write namespaced resources to in the "master" and "remote" clusters
	// If left empty will return error
	Namespace string

	// Namespace to write namespaced resources to in the "master" and "remote" clusters
	// If left empty will return error
	RemoteNamespace string

	// The Cluster Domain used by the Kubernetes DNS Service in the registered cluster.
	// Defaults to 'cluster.local'
	// Read more: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/
	ClusterDomain string

	// A list of roles to bind the New kubeconfig token to
	// Any Roles in this list will be Upserted by the registrant, prior to binding
	Roles []*k8s_rbac_types.Role

	// A list of cluster roles to bind the New kubeconfig token to
	// Any ClusterRoles in this list will be Upserted by the registrant, prior to binding
	ClusterRoles []*k8s_rbac_types.ClusterRole

	// List of roles which will be bound to by the created role bindings
	// The Roles upserted from the above list will automatically appended
	RoleBindings []client.ObjectKey

	// List of cluster roles which will be bound to by the created cluster role bindings
	// The ClusterRoles upserted from the above list will automatically appended to the list
	ClusterRoleBindings []client.ObjectKey
}

/*
	RegisterCluster is meant to be a helper function to easily "register" a remote cluster.
	Currently this entails:
		1. Creating a `ServiceAccount` on the remote cluster.
		2. Binding RBAC `Roles/ClusterRoles` to said `ServiceAccount`
		3. And finally creating a kubeconfig `Secret` with the BearerToken of the remote `ServiceAccount`
*/
func (opts RegistrationOptions) RegisterCluster(
	ctx context.Context,
) error {
	masterRestCfg, remoteCfg, rbacOpts, registrant, err := opts.initialize()
	if err != nil {
		return err
	}
	return RegisterClusterFromConfig(ctx, masterRestCfg, remoteCfg, rbacOpts, registrant)
}

/*
	DeregisterCluster deregisters a cluster by cleaning up the resources created when RegisterCluster is invoked.
	This entails:
		1. Deleting the ServiceAccount on the remote cluster.
		2. Deleting the remote Roles, RoleBindings, ClusterRoles, and ClusterRoleBindings associated with the ServiceAccount.
		3. Deletes the secret containing the kubeconfig for the remote cluster.
*/
func (opts RegistrationOptions) DeregisterCluster(
	ctx context.Context,
) error {
	masterRestCfg, remoteCfg, rbacOpts, registrant, err := opts.initialize()
	if err != nil {
		return err
	}
	return DeregisterClusterFromConfig(ctx, masterRestCfg, remoteCfg, rbacOpts, registrant)
}

// Initialize registration dependencies
func (opts RegistrationOptions) initialize() (masterRestCfg *rest.Config, remoteCfg clientcmd.ClientConfig, rbacOpts RbacOptions, registrant ClusterRegistrant, err error) {
	masterCfg, err := getClientConfigWithContext(opts.MasterURL, opts.KubeCfgPath, opts.KubeContext)
	if err != nil {
		return masterRestCfg, remoteCfg, rbacOpts, registrant, err
	}

	masterRestCfg, err = masterCfg.ClientConfig()
	if err != nil {
		return masterRestCfg, remoteCfg, rbacOpts, registrant, err
	}

	registrant, err = defaultRegistrant(masterRestCfg, opts.APIServerAddress)
	if err != nil {
		return masterRestCfg, remoteCfg, rbacOpts, registrant, err
	}

	remoteCfg, err = getClientConfigWithContext(opts.MasterURL, opts.RemoteKubeCfgPath, opts.RemoteKubeContext)
	if err != nil {
		return masterRestCfg, remoteCfg, rbacOpts, registrant, err
	}

	rbacOpts = RbacOptions{
		Options: Options{
			ClusterName:     opts.ClusterName,
			RemoteCtx:       opts.RemoteKubeContext,
			Namespace:       opts.Namespace,
			RemoteNamespace: opts.RemoteNamespace,
			ClusterDomain:   opts.ClusterDomain,
		},
		Roles:               opts.Roles,
		ClusterRoles:        opts.ClusterRoles,
		RoleBindings:        opts.RoleBindings,
		ClusterRoleBindings: opts.ClusterRoleBindings,
	}

	return masterRestCfg, remoteCfg, rbacOpts, registrant, nil
}

func RegisterClusterFromConfig(
	ctx context.Context,
	masterClusterCfg *rest.Config,
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

	return registrant.RegisterClusterWithToken(ctx, masterClusterCfg, remoteCfg, token, opts.Options)
}

func DeregisterClusterFromConfig(
	ctx context.Context,
	masterClusterCfg *rest.Config,
	remoteCfg clientcmd.ClientConfig,
	opts RbacOptions,
	registrant ClusterRegistrant,
) error {
	var multierr *multierror.Error

	if err := registrant.DeregisterCluster(ctx, masterClusterCfg, remoteCfg, opts.Options); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	if err := registrant.DeleteRemoteServiceAccount(ctx, remoteCfg, opts.Options); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	if err := registrant.DeleteRemoteAccessResources(ctx, remoteCfg, opts); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	return multierr.ErrorOrNil()
}

/*
	DefaultRegistrant provider function. This function will create a `ClusterRegistrant` using the
	current kubeconfig, and the specified context. It will build all of the dependencies from the
	available `ClientConfig`.

	The apiServerAddress parameter is optional. When passed in, it will overwrite the Api Server
	endpoint in the kubeconfig before it is written. This is primarily useful when running multi cluster
	KinD environments on a mac as  the local IP needs to be re-written to `host.docker.internal` so
	that the local instance knows to hit localhost.

	Meant to be used in tandem with RegisterClusterFromConfig above.
	They are exposed separately so the `Registrant` may be mocked for the function above.
*/
func DefaultRegistrant(context, apiServerAddress string) (ClusterRegistrant, error) {
	cfg, err := config.GetConfigWithContext(context)
	if err != nil {
		return nil, err
	}

	return defaultRegistrant(cfg, apiServerAddress)
}

func defaultRegistrant(cfg *rest.Config, apiServerAddress string) (ClusterRegistrant, error) {
	clientset, err := k8s_core_v1.NewClientsetFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	nsClientFactory := k8s_core_v1_providers.NamespaceClientFromConfigFactoryProvider()
	secretClientFactory := k8s_core_v1_providers.SecretClientFromConfigFactoryProvider()
	saClientFactory := k8s_core_v1_providers.ServiceAccountClientFromConfigFactoryProvider()
	roleClientFactory := rbac_v1_providers.RoleClientFromConfigFactoryProvider()
	clusterRoleClientFactory := rbac_v1_providers.ClusterRoleClientFromConfigFactoryProvider()
	kubeClusterClientFactory := v1alpha1_providers.KubernetesClusterClientFromConfigFactoryProvider()
	clusterRBACBinderFactory := NewClusterRBACBinderFactory()
	registrant := NewClusterRegistrant(
		apiServerAddress,
		clusterRBACBinderFactory,
		clientset.Secrets(),
		secretClientFactory,
		nsClientFactory,
		saClientFactory,
		clusterRoleClientFactory,
		roleClientFactory,
		kubeClusterClientFactory,
	)
	return registrant, nil
}

// Attempts to load a Client KubeConfig from a default list of sources.
func getClientConfigWithContext(masterURL, kubeCfgPath, context string) (clientcmd.ClientConfig, error) {
	verifiedKubeConfigPath := clientcmd.RecommendedHomeFile
	if kubeCfgPath != "" {
		verifiedKubeConfigPath = kubeCfgPath
	}

	if err := assertKubeConfigExists(verifiedKubeConfigPath); err != nil {
		return nil, err
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = verifiedKubeConfigPath
	configOverrides := &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterURL}}

	if context != "" {
		configOverrides.CurrentContext = context
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides), nil
}

// expects `path` to be nonempty
func assertKubeConfigExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	return nil
}
