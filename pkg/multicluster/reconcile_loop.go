package multicluster

import (
	"context"
	"sync"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type Reconciler interface {
	// reconcile an object
	// requeue the object if returning an error, or a non-zero "requeue-after" duration
	Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error)
}

type DeletionReconciler interface {
	// we received a reconcile request for an object that was removed from the cache
	ReconcileDeletion(cluster string, request reconcile.Request)
}

// a Reconcile Loop runs resource reconcilers until the context gets cancelled
type Loop interface {
	// AddReconciler adds a reconciler to a slice of reconcilers that will be run against
	AddReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate)
}

var _ Loop = &clusterLoopRunner{}

type clusterLoopRunner struct {
	name         string
	resource     ezkube.Object
	clusterLoops *clusterLoopSet
	reconcilers  *reconcilerList
}

func NewLoop(name string, cw ClusterWatcher, resource ezkube.Object) *clusterLoopRunner {
	runner := &clusterLoopRunner{
		name:         name,
		resource:     resource,
		clusterLoops: newClusterLoopSet(),
		reconcilers:  newReconcilerList(),
	}
	cw.RegisterClusterHandler(runner)

	return runner
}

// AddCluster creates a reconcile loop for the cluster.
func (r *clusterLoopRunner) AddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	loopForCluster := reconcile.NewLoop(r.name+"-"+cluster, mgr, r.resource)
	r.clusterLoops.add(cluster, loopForCluster)
	r.clusterLoops.ensureReconcilers(r.reconcilers)

	go func() {
		<-ctx.Done()
		r.clusterLoops.remove(cluster)
		r.reconcilers.clusterRemoved(cluster)
	}()
}

// AddReconciler registers a cluster handler for the reconciler.
func (r *clusterLoopRunner) AddReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate) {
	r.reconcilers.add(ctx, reconciler, predicates...)
	r.clusterLoops.ensureReconcilers(r.reconcilers)
}

type multiclusterReconciler struct {
	cluster        string
	userReconciler Reconciler
}

func (m multiclusterReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	return m.userReconciler.Reconcile(m.cluster, object)
}

func (m multiclusterReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := m.userReconciler.(DeletionReconciler); ok {
		deletionReconciler.ReconcileDeletion(m.cluster, request)
	}
}

type clusterLoopSet struct {
	mutex        sync.RWMutex
	clusterLoops map[string]reconcile.Loop
}

func newClusterLoopSet() *clusterLoopSet {
	return &clusterLoopSet{
		mutex:        sync.RWMutex{},
		clusterLoops: make(map[string]reconcile.Loop),
	}
}

func (s *clusterLoopSet) add(cluster string, loop reconcile.Loop) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clusterLoops[cluster] = loop
}

func (s *clusterLoopSet) remove(cluster string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clusterLoops, cluster)
}

// ensureReconcilers ensures the given reconcilers are running on every cluster loop.
func (s *clusterLoopSet) ensureReconcilers(list *reconcilerList) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for cluster, loop := range s.clusterLoops {
		list.runAll(cluster, loop)
	}
}

type runnableReconciler struct {
	ctx            context.Context
	reconciler     Reconciler
	predicates     []predicate.Predicate
	activeClusters sets.String
}

type reconcilerList struct {
	mutex       sync.RWMutex
	reconcilers []runnableReconciler
}

func newReconcilerList() *reconcilerList {
	return &reconcilerList{
		mutex:       sync.RWMutex{},
		reconcilers: make([]runnableReconciler, 0, 1),
	}
}

func (r *reconcilerList) add(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.reconcilers = append(r.reconcilers, runnableReconciler{
		ctx:            ctx,
		reconciler:     reconciler,
		predicates:     predicates,
		activeClusters: sets.String{},
	})
}

// runAll runs all reconcilers in the list on the given loop.
// If a reconciler is already active on a cluster, is is skipped.
func (r *reconcilerList) runAll(cluster string, loop reconcile.Loop) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, rr := range r.reconcilers {
		if rr.activeClusters.Has(cluster) {
			continue
		}

		mcReconciler := &multiclusterReconciler{
			cluster:        cluster,
			userReconciler: rr.reconciler,
		}
		err := loop.RunReconciler(rr.ctx, mcReconciler, rr.predicates...)
		if err != nil {
			contextutils.LoggerFrom(rr.ctx).Debug("Error occurred when adding reconciler to cluster loop",
				zap.Error(err),
				zap.String("cluster", cluster))
		}
		rr.activeClusters.Insert(cluster)
	}
}

// clusterRemoved removes cluster from the set of active clusters.
func (r *reconcilerList) clusterRemoved(cluster string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, rr := range r.reconcilers {
		rr.activeClusters.Delete(cluster)
	}
}
