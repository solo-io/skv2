package multicluster

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// ClusterWatcher watches for KubeConfig secrets on the master cluster.
// It is responsible for starting cluster managers and calling ClusterHandler functions.
type ClusterWatcher interface {
	// Run starts a watch for KubeConfig secrets on the cluster managed by the given manager.Manager.
	// Note that Run will call Start on the given manager and run all registered ClusterHandlers.
	Run(master manager.Manager) error
	// RegisterClusterHandler adds a ClusterHandler to the ClusterWatcher.
	RegisterClusterHandler(handler ClusterHandler)
}

// ManagerSet maintains a manager for every cluster in the system.
type ManagerSet interface {
	// Cluster returns a manager for the given cluster, or an error if one does not exist.
	Cluster(cluster string) (manager.Manager, error)
}

// ManagerSet maintains a manager for every cluster in the system.
type ClusterSet interface {
	// List the clusters (sorted) currently known to the set
	ListClusters() []string
}
