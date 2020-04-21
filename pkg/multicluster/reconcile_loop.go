package multicluster

import (
	"context"

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
	// RunReconciler adds a reconciler to a slice of reconcilers that will be run against
	RunReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate) error
}

var _ Loop = &runner{}

type runner struct {
	name            string
	resource        ezkube.Object
	userReconcilers []userReconciler
}

type userReconciler struct {
	reconciler Reconciler
	predicates []predicate.Predicate
}

func NewLoop(name string, cw ClusterWatcher, resource ezkube.Object) *runner {
	runner := &runner{
		name:            name,
		resource:        resource,
		userReconcilers: make([]userReconciler, 0, 1),
	}
	cw.RegisterClusterHandler(runner)

	return runner
}

func (r *runner) RunReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate) error {
	r.userReconcilers = append(r.userReconcilers, userReconciler{
		reconciler: reconciler,
		predicates: predicates,
	})
	return nil
}

func (r *runner) HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) {
	loopForCluster := reconcile.NewLoop(r.name+"-"+cluster, mgr, r.resource)
	for _, userRec := range r.userReconcilers {
		mcReconciler := &multiclusterReconciler{
			cluster:        cluster,
			userReconciler: userRec.reconciler,
		}
		err := loopForCluster.RunReconciler(ctx, mcReconciler, userRec.predicates...)
		if err != nil {
			contextutils.LoggerFrom(ctx).Debug("Error occurred when adding reconciler for cluster",
				zap.Error(err),
				zap.String("cluster", cluster),
				zap.String("reconciler", r.name),
			)
		}
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
