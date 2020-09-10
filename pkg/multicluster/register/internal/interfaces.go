package internal

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

/*
	ClusterRBACBinder is a helper interface for the registrant, meant to create, and bind RBAC objects to
	ServiceAccounts in a remote cluster.

	This interface supports both `Roles` and `ClusterRoles`.
	All resources being referenced must already exist, or the operations will fail.
*/
type ClusterRBACBinder interface {
	/*
		Given a set of references to `ClusterRoles`, bind the given `ServiceAccount` to said `ClusterRoles` using
		Newly created `ClusterRoleBindings`.
	*/
	BindClusterRoles(
		ctx context.Context,
		serviceAccount client.ObjectKey,
		clusterRoles []client.ObjectKey,
	) error
	/*
		Given a set of references to `Roles`, bind the given `ServiceAccount` to said `Roles` using
		Newly created `RoleBindings`.
	*/
	BindRoles(
		ctx context.Context,
		serviceAccount client.ObjectKey,
		roles []client.ObjectKey,
	) error

	DeleteClusterRoleBindings(
		ctx context.Context,
		serviceAccount client.ObjectKey,
		clusterRoles []client.ObjectKey,
	) error

	DeleteRoleBindings(
		ctx context.Context,
		serviceAccount client.ObjectKey,
		roles []client.ObjectKey,
	) error
}

/*
	Factory function to build a ClusterRBACBinder from a `ClientConfig`
	This is useful because an the operations performed by the `RbacBinder` require access to a cluster
	which will be determined by the caller.
*/
type ClusterRBACBinderFactory func(cfg clientcmd.ClientConfig) (ClusterRBACBinder, error)
