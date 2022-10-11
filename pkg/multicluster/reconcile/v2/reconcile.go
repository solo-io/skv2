package reconcile_v2

import (
	"context"
	"sync"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	multicluster_v2 "github.com/solo-io/skv2/pkg/multicluster/v2"
	reconcile_v2 "github.com/solo-io/skv2/pkg/reconcile/v2"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type clusterLoopRunner[T client.Object] struct {
	name         string
	clusterLoops *clusterLoopSet[T]
	reconcilers  *reconcilerList[T]
	options      reconcile_v2.Options
	t            T
}

var _ multicluster_v2.Loop[client.Object] = &clusterLoopRunner[client.Object]{}
var _ multicluster.ClusterHandler = &clusterLoopRunner[client.Object]{}
var _ multicluster.ClusterRemovedHandler = &clusterLoopRunner[client.Object]{}

func NewLoop[T client.Object](
	name string,
	cw multicluster.ClusterWatcher,
	t T,
	options reconcile_v2.Options,
) *clusterLoopRunner[T] {
	runner := &clusterLoopRunner[T]{
		name:         name,
		clusterLoops: newClusterLoopSet[T](),
		reconcilers:  newReconcilerList[T](),
		options:      options,
		t:            t,
	}
	cw.RegisterClusterHandler(runner)

	return runner
}

// AddCluster creates a reconcile_v2 loop for the cluster.
func (r *clusterLoopRunner[T]) AddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	loopForCluster := reconcile_v2.NewLoop(r.name+"-"+cluster, cluster, mgr, r.t, r.options)

	// Add the cluster loop to the set of active loops and start reconcilers.
	r.clusterLoops.add(cluster, loopForCluster)
	r.clusterLoops.ensureReconcilers(r.reconcilers)
}

// RemoveCluster clears any state related to the removed cluster.
func (r *clusterLoopRunner[T]) RemoveCluster(cluster string) {
	r.clusterLoops.remove(cluster)
	r.reconcilers.unset(cluster)
}

// AddReconciler registers a cluster handler for the reconciler.
func (r *clusterLoopRunner[T]) AddReconciler(
	ctx context.Context,
	reconciler multicluster_v2.Reconciler[T],
	predicates ...predicate.Predicate,
) {
	r.reconcilers.add(ctx, reconciler, predicates...)
	r.clusterLoops.ensureReconcilers(r.reconcilers)
}

type multiclusterReconciler[T client.Object] struct {
	cluster        string
	userReconciler multicluster_v2.Reconciler[T]
}

func (m multiclusterReconciler[T]) Reconcile(ctx context.Context, object T) (reconcile_v2.Result, error) {
	return m.userReconciler.Reconcile(ctx, m.cluster, object)
}

func (m multiclusterReconciler[T]) ReconcileDeletion(ctx context.Context, request reconcile_v2.Request) error {
	if deletionReconciler, ok := m.userReconciler.(multicluster_v2.DeletionReconciler[T]); ok {
		return deletionReconciler.ReconcileDeletion(ctx, m.cluster, request)
	}
	return nil
}

type clusterLoopSet[T client.Object] struct {
	mutex        sync.RWMutex
	clusterLoops map[string]reconcile_v2.Loop[T]
}

func newClusterLoopSet[T client.Object]() *clusterLoopSet[T] {
	return &clusterLoopSet[T]{
		mutex:        sync.RWMutex{},
		clusterLoops: make(map[string]reconcile_v2.Loop[T]),
	}
}

func (s *clusterLoopSet[T]) add(cluster string, loop reconcile_v2.Loop[T]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clusterLoops[cluster] = loop
}

func (s *clusterLoopSet[T]) remove(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clusterLoops, cluster)
}

// ensureReconcilers ensures the given reconcilers are running on every cluster loop.
func (s *clusterLoopSet[T]) ensureReconcilers(list *reconcilerList[T]) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for cluster, loop := range s.clusterLoops {
		list.runAll(cluster, loop)
	}
}

type runnableReconciler[T client.Object] struct {
	ctx            context.Context
	reconciler     multicluster_v2.Reconciler[T]
	predicates     []predicate.Predicate
	activeClusters sets.String
}

type reconcilerList[T client.Object] struct {
	mutex       sync.RWMutex
	reconcilers []runnableReconciler[T]
}

func newReconcilerList[T client.Object]() *reconcilerList[T] {
	return &reconcilerList[T]{
		mutex:       sync.RWMutex{},
		reconcilers: make([]runnableReconciler[T], 0, 1),
	}
}

func (r *reconcilerList[T]) add(
	ctx context.Context,
	reconciler multicluster_v2.Reconciler[T],
	predicates ...predicate.Predicate,
) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.reconcilers = append(
		r.reconcilers, runnableReconciler[T]{
			ctx:            ctx,
			reconciler:     reconciler,
			predicates:     predicates,
			activeClusters: sets.String{},
		},
	)
}

// runAll runs all reconcilers in the list on the given loop.
// If a reconciler is already active on a cluster, is is skipped.
func (r *reconcilerList[T]) runAll(cluster string, loop reconcile_v2.Loop[T]) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, rr := range r.reconcilers {
		if rr.activeClusters.Has(cluster) {
			continue
		}

		mcReconciler := &multiclusterReconciler[T]{
			cluster:        cluster,
			userReconciler: rr.reconciler,
		}
		err := loop.RunReconciler(rr.ctx, mcReconciler, rr.predicates...)
		if err != nil {
			contextutils.LoggerFrom(rr.ctx).Debug(
				"Error occurred when adding reconciler to cluster loop",
				zap.Error(err),
				zap.String("cluster", cluster),
			)
		}
		rr.activeClusters.Insert(cluster)
	}
}

// unset removes cluster from the set of active clusters on each reconciler.
func (r *reconcilerList[T]) unset(cluster string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, rr := range r.reconcilers {
		rr.activeClusters.Delete(cluster)
	}
}
