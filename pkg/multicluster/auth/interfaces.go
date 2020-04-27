package auth

import (
	"context"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_auth.go

type RbacBinder interface {
	// bind the given roles to the target service account at cluster scope
	BindClusterRolesToServiceAccount(
		ctx context.Context,
		targetServiceAccount *k8s_core_types.ServiceAccount,
		roles []*k8s_rbac_types.ClusterRole,
	) error
}

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
// Creates a service account in the target cluster with the name/namespace of `serviceAccountRef` and cluster-admin permissions
type ClusterAuthorization interface {
	BuildRemoteBearerToken(
		ctx context.Context,
		targetClusterCfg *rest.Config,
		name, namespace string,
	) (bearerToken string, err error)
}

type ClusterAuthorizationFactory func(cfg clientcmd.ClientConfig) (ClusterAuthorization, error)

// Create a service account on a cluster that `targetClusterCfg` can reach
// Set up that service account with the indicated cluster roles
type RemoteAuthorityManager interface {
	// creates a new service account in the cluster pointed to by the cfg at the name/namespace indicated by the ResourceRef,
	// and assigns it the given ClusterRoles
	// NB: if role assignment fails, the service account is left in the cluster; this is not an atomic operation
	ApplyRemoteServiceAccount(
		ctx context.Context,
		name, namespace string,
		roles []*k8s_rbac_types.ClusterRole,
	) (*k8s_core_types.ServiceAccount, error)
}

type Clients struct {
	ServiceAccountClient k8s_core_v1.ServiceAccountClient
	RbacBinder           RbacBinder
	SecretClient         k8s_core_v1.ServiceAccountClient
}
