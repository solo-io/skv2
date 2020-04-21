package multicluster

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// ClusterHandler is passed to RunClusterWatcher to handle select cluster events.
// It is implemented internally by skv2 components but can be implemented by the user for specialized use cases.
type ClusterHandler interface {
	// ClusterHandler is called when a new cluster is identified by a cluster watch.
	// The provided context is cancelled when a cluster is removed, so any teardown behavior for removed clusters
	// should take place when ctx is cancelled.
	HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) error
}
