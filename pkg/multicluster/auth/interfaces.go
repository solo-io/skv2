package auth

import (
	"context"

	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_auth.go

// create a kube config that can authorize to the target cluster as the service account from that target cluster
type RemoteAuthorityConfigCreator interface {

	// Returns a `*rest.Config` that points to the same cluster as `targetClusterCfg`, but authorizes itself using the
	// bearer token belonging to the service account at `serviceAccountRef` in the target cluster
	//
	// NB: This function blocks the current go routine for up to 6 seconds while waiting for the service account's secret
	// to appear, by performing an exponential backoff retry loop
	ConfigFromRemoteServiceAccount(
		ctx context.Context,
		targetClusterCfg *rest.Config,
		name, namespace string,
	) (*rest.Config, error)
}

// Given a way to authorize to a cluster, produce a bearer token that can authorize to that same cluster
// using a newly-created service account token in that cluster.
// Creates a service account in the target cluster with the name/namespace of `serviceAccountRef`
// If any clusterRoles are passed in it will attempt bind to them, otherwise it will default to cluster-admin
type ClusterAuthorization interface {
	BuildRemoteBearerToken(
		ctx context.Context,
		targetClusterCfg *rest.Config,
		name, namespace string,
		clusterRoles ...*k8s_rbac_types.ClusterRole,
	) (bearerToken string, err error)
}

type ClusterAuthorizationFactory func(cfg clientcmd.ClientConfig) (ClusterAuthorization, error)
