package register

import (
	"context"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	v1alpha1_providers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/providers"

	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/rest"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1"
	k8s_core_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/providers"
	rbac_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/providers"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// Options for registering a cluster
type RegistrationOptions struct {
	// Management kubeconfig
	KubeCfg clientcmd.ClientConfig

	// Remote kubeconfig
	RemoteKubeCfg clientcmd.ClientConfig

	// Remote context name
	// We need to explicitly pass this because of this open issue: https://github.com/kubernetes/client-go/issues/735
	RemoteCtx string

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

	// Set of labels to include on the registration output resources, currently consisting of KubernetesCluster and Secret.
	ResourceLabels map[string]string
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
	return opts.RegisterProviderCluster(ctx, nil)
}

/*
	RegisterProviderCluster augments RegisterCluster functionality
	with additional metadata to persist to the resulting KubernetesCluster object.
	ProviderInfo contains cloud provider metadata.
*/
func (opts RegistrationOptions) RegisterProviderCluster(
	ctx context.Context,
	providerInfo *v1alpha1.KubernetesClusterSpec_ProviderInfo,
) error {
	masterRestCfg, remoteCfg, rbacOpts, registrant, err := opts.initialize()
	if err != nil {
		return err
	}
	// Parse ClusterRole policy rules and pass into RegisterProviderClusterFromConfig in order to persist them to the resulting
	// KubernetesCluster status.
	var clusterRolePolicyRules []*v1alpha1.PolicyRule
	for _, clusterRole := range opts.ClusterRoles {
		for _, policyRules := range clusterRole.Rules {
			clusterRolePolicyRules = append(clusterRolePolicyRules, &v1alpha1.PolicyRule{
				Verbs:           policyRules.Verbs,
				ApiGroups:       policyRules.APIGroups,
				Resources:       policyRules.Resources,
				ResourceNames:   policyRules.ResourceNames,
				NonResourceUrls: policyRules.NonResourceURLs,
			})
		}
	}
	return RegisterProviderClusterFromConfig(
		ctx,
		masterRestCfg,
		remoteCfg,
		rbacOpts,
		registrant,
		providerInfo,
		opts.ResourceLabels,
		clusterRolePolicyRules,
	)
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
	masterRestCfg, err = opts.KubeCfg.ClientConfig()
	if err != nil {
		return masterRestCfg, remoteCfg, rbacOpts, registrant, err
	}

	remoteCfg = opts.RemoteKubeCfg

	registrant, err = defaultRegistrant(masterRestCfg, opts.APIServerAddress)
	if err != nil {
		return masterRestCfg, remoteCfg, rbacOpts, registrant, err
	}

	rbacOpts = RbacOptions{
		Options: Options{
			ClusterName:     opts.ClusterName,
			Namespace:       opts.Namespace,
			RemoteCtx:       opts.RemoteCtx,
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
	return RegisterProviderClusterFromConfig(ctx, masterClusterCfg, remoteCfg, opts, registrant, nil, nil, nil)
}

func RegisterProviderClusterFromConfig(
	ctx context.Context,
	masterClusterCfg *rest.Config,
	remoteCfg clientcmd.ClientConfig,
	opts RbacOptions,
	registrant ClusterRegistrant,
	providerInfo *v1alpha1.KubernetesClusterSpec_ProviderInfo,
	resourceLabels map[string]string,
	policyRules []*v1alpha1.PolicyRule,
) error {
	err := registrant.EnsureRemoteNamespace(ctx, remoteCfg, opts.RemoteNamespace)
	if err != nil {
		return err
	}

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

	return registrant.RegisterProviderClusterWithToken(
		ctx,
		masterClusterCfg,
		remoteCfg,
		token,
		opts.Options,
		providerInfo,
		resourceLabels,
		policyRules,
	)
}

func DeregisterClusterFromConfig(
	ctx context.Context,
	masterClusterCfg *rest.Config,
	remoteCfg clientcmd.ClientConfig,
	opts RbacOptions,
	registrant ClusterRegistrant,
) error {
	var multierr *multierror.Error

	if err := registrant.DeregisterCluster(ctx, masterClusterCfg, opts.Options); err != nil {
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

// expects `path` to be nonempty
func assertKubeConfigExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	return nil
}
