package internal

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

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
