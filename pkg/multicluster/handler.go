package multicluster

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// ClusterHandler is passed to RunClusterWatcher to handle select cluster events.
// It is implemented internally by skv2 components but can be implemented by the user for specialized use cases.
type ClusterHandler interface {
	// AddCluster is called when a new cluster is identified by a cluster watch.
	// The provided context is cancelled when a cluster is removed, so any teardown behavior for removed clusters
	// should take place when ctx is cancelled.
	AddCluster(ctx context.Context, cluster string, mgr manager.Manager)
}

// ClusterRemovedHandler can be implemented by ClusterHandlers to perform cleanup when a cluster is deleted.
// NOTE: in most cases, cleanup should be handled when the cluster manager's context is cancelled without
// the need for implementing ClusterRemovedHandler.
type ClusterRemovedHandler interface {
	// RemoveCluster is called when a cluster watch identifies a cluster as deleted.
	RemoveCluster(cluster string)
}
