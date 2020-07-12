package register

import (
	"context"

	"k8s.io/client-go/rest"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/multicluster/register/internal"
	corev1 "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

// used for testing multicluster components
//go:generate  mockgen -package mock_clientcmd -destination ./mock_clientcmd/config.go k8s.io/client-go/tools/clientcmd ClientConfig

// Expose internal providers for Dependency Injection
var (
	NewClusterRBACBinderFactory = internal.NewClusterRBACBinderFactory
)

type Options struct {

	// Name by which the cluster will be identified
	// If left empty will return error
	ClusterName string

	// Name of the remote cluster Kubeconfig Context
	RemoteCtx string

	// Namespace to write namespaced resources to in the "master" and "remote" clusters
	// If left empty will return error
	Namespace string

	// Namespace to write namespaced resources to in the "master" and "remote" clusters
	// If left empty will return error
	RemoteNamespace string

	// The Cluster Domain used by the Kubernetes DNS Service.
	// Defaults to 'cluster.local'
	// Read more: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/
	ClusterDomain string
}

type RbacOptions struct {
	Options

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

func (o *Options) validate() error {
	if o.Namespace == "" {
		return eris.Errorf("Must specify namespace")
	}
	if o.RemoteNamespace == "" {
		o.RemoteNamespace = o.Namespace
	}
	if o.ClusterName == "" {
		return eris.Errorf("Must specify cluster name")
	}
	return nil
}

/*
	Standard Cluster Registrant (one who registers) interface.

	This component is responsible for registering a "remote" kubernetes cluster to a "management" cluster.
	As the "management" cluster is not present in the interface itself, it is defined by the config used to build
	the registrant instance.
*/
type ClusterRegistrant interface {

	/*
		EnsureRemoteServiceAccount takes an instance of a remote config, and ensure a ServiceAccount exists on the
		remote cluster, in the namespace specified.

		This `ServiceAccount` can then be used and/or referenced by `CreateRemoteAccessToken` below for the remainder
		of the registration workflow
	*/
	EnsureRemoteServiceAccount(
		ctx context.Context,
		remoteClientCfg clientcmd.ClientConfig,
		opts Options,
	) (*corev1.ServiceAccount, error)

	/*
		CreateRemoteAccessToken takes an instance of a remote config, and a reference to an existing `ServiceAccount`,
		and attempts to bind the given `RBAC` objects to said `ServiceAccount`, in the specified cluster.

		The function will then return the `ServiceAccount` token,
	*/
	CreateRemoteAccessToken(
		ctx context.Context,
		remoteClientCfg clientcmd.ClientConfig,
		sa client.ObjectKey,
		opts RbacOptions,
	) (token string, err error)

	/*
		RegisterClusterWithToken takes an instance of the remote config, as well as a `BearerToken` and creates a
		kubeconfig secret on the local cluster, in the specified namespace.
	*/
	RegisterClusterWithToken(
		ctx context.Context,
		masterClusterCfg *rest.Config,
		remoteClientCfg clientcmd.ClientConfig,
		token string,
		opts Options,
	) error
}
