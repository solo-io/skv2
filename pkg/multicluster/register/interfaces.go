package register

import (
	"context"

	"github.com/rotisserie/eris"
	corev1 "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

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
}

type RbacOptions struct {
	Options

	// A list of roles to bind the New kubeconfig token to
	Roles []*k8s_rbac_types.Role

	// A list of cluster roles to bind the New kubeconfig token to
	ClusterRoles []*k8s_rbac_types.ClusterRole

	// List of roles which will be bound to by the created role bindings
	// The Roles created from the above list will automatically appended
	RoleBindings []client.ObjectKey

	// List of cluster roles which will be bound to by the created cluster role bindings
	// The Roles created from the above list will automatically appended to the list
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
	EnsureRemoteServiceAccount(
		ctx context.Context,
		remoteClientCfg clientcmd.ClientConfig,
		opts Options,
	) (*corev1.ServiceAccount, error)

	CreateRemoteAccessToken(
		ctx context.Context,
		remoteClientCfg clientcmd.ClientConfig,
		sa client.ObjectKey,
		opts RbacOptions,
	) (token string, err error)

	/*
		RegisterClusterFromConfig takes an instance of the remote config, and the registration info, and registers
		the cluster.
	*/
	RegisterClusterWithToken(
		ctx context.Context,
		remoteClientCfg clientcmd.ClientConfig,
		token string,
		opts Options,
	) error
}

// Given a way to authorize to a cluster, produce a bearer token that can authorize to that same cluster
// using a newly-created service account token in that cluster.
// Creates a service account in the target cluster with the name/namespace of `serviceAccountRef`
type ClusterRBACBinder interface {
	// If any clusterRoles are passed in it will attempt bind to them, otherwise it will default to cluster-admin
	BindClusterRoles(
		ctx context.Context,
		serviceAccount client.ObjectKey,
		clusterRoles []client.ObjectKey,
	) error
	// At least one Role is required to bind to, an empty list will be considered invalid
	BindRoles(
		ctx context.Context,
		serviceAccount client.ObjectKey,
		roles []client.ObjectKey,
	) error
}

type ClusterRBACBinderFactory func(cfg clientcmd.ClientConfig) (ClusterRBACBinder, error)
