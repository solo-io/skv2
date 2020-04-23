package multicluster

import (
	"context"
	"sync"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
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
	cw           ClusterWatcher
	clusterLoops clusterLoopSet
}

func NewLoop(name string, cw ClusterWatcher, resource ezkube.Object) *clusterLoopRunner {
	runner := &clusterLoopRunner{
		name:         name,
		resource:     resource,
		clusterLoops: newClusterLoopSet(),
		cw:           cw,
	}
	cw.RegisterClusterHandler(runner)

	return runner
}

// AddCluster creates a reconcile loop for the cluster.
func (r *clusterLoopRunner) AddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	loopForCluster := reconcile.NewLoop(r.name+"-"+cluster, mgr, r.resource)
	r.clusterLoops.add(cluster, loopForCluster)
}

// AddReconciler registers a cluster handler for the reconciler.
func (r *clusterLoopRunner) AddReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate) {
	recRunner := reconcilerRunner{
		clusterLoopRunner: r,
		reconciler:        reconciler,
		predicates:        predicates,
	}
	r.cw.RegisterClusterHandler(recRunner)
}

type reconcilerRunner struct {
	reconcilerName    string
	clusterLoopRunner *clusterLoopRunner
	reconciler        Reconciler
	predicates        []predicate.Predicate
}

// AddCluster runs a reconciler on an existing cluster reconcile loop.
func (r reconcilerRunner) AddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	loop, ok := r.clusterLoopRunner.clusterLoops.get(cluster)
	if !ok {
		contextutils.LoggerFrom(ctx).Debug("Attempted to run reconciler for nonexistent cluster loop",
			zap.String("cluster", cluster),
			zap.String("loop", r.clusterLoopRunner.name))
		return
	}

	mcReconciler := &multiclusterReconciler{
		cluster:        cluster,
		userReconciler: r.reconciler,
	}
	err := loop.RunReconciler(ctx, mcReconciler, r.predicates...)
	if err != nil {
		contextutils.LoggerFrom(ctx).Debug("Error occurred when adding reconciler for cluster",
			zap.Error(err),
			zap.String("cluster", cluster),
			zap.String("loop", r.clusterLoopRunner.name))
	}
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

func newClusterLoopSet() clusterLoopSet {
	return clusterLoopSet{
		mutex:        sync.RWMutex{},
		clusterLoops: make(map[string]reconcile.Loop),
	}
}

func (s clusterLoopSet) add(cluster string, loop reconcile.Loop) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clusterLoops[cluster] = loop
}

func (s clusterLoopSet) get(cluster string) (reconcile.Loop, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	loop, ok := s.clusterLoops[cluster]
	return loop, ok
}
