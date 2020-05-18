package register

import (
	"context"

	k8s_rbac_types "k8s.io/api/rbac/v1"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

type ClusterInfo struct {

	// Name by which the cluster will be identified
	ClusterName string

	// Namespace to write namespaced resources to in the "master" and "remote" clusters
	Namespace string

	/*
		This option should be used mostly for testing.
		When passed in, it will overwrite the Api Server endpoint in the the kubeconfig before it is written.
		This is primarily useful when running multi cluster KinD environments on a mac as  the local IP needs
		to be re-written to `host.docker.internal` so that the local instance knows to hit localhost.
	*/
	LocalClusterDomainOverride string

	// A list of cluster roles to bind the New kubeconfig token to, if empty will be default to `cluster-admin`
	ClusterRoles []*k8s_rbac_types.ClusterRole

	// If true attempt to upsert the specified ClusterRoles
	UpsertRoles bool
}

/*
	Standard Cluster Registrant (one who registers) interface.

	This component is responsible for registering a "remote" kubernetes cluster to a "management" cluster.
	As the "management" cluster is not present in the interface itself, it is defined by the config used to build
	the registrant instance.
*/
type ClusterRegistrant interface {
	/*
		RegisterClusterFromConfig takes an instance of the remote config, and the registration info, and registers
		the cluster.
	*/
	RegisterClusterFromConfig(
		ctx context.Context,
		remoteCfg clientcmd.ClientConfig,
		info ClusterInfo,
	) error
	/*
		The standard cluster register function.

		The behavior is the same as RegisterClusterFromConfig, but the config is first read in from the path, and
		the given context, and then transformed into a `ClientConfig`
	*/
	RegisterCluster(
		ctx context.Context,
		info ClusterInfo,
		remoteCfgPath, remoteCtx string,
	) error
}
