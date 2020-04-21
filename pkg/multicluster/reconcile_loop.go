package multicluster

import (
	"context"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type Request = reconcile.Request
type Result = reconcile.Result

type Reconciler interface {
	// reconcile an object
	// requeue the object if returning an error, or a non-zero "requeue-after" duration
	Reconcile(cluster string, object ezkube.Object) (Result, error)
}

type DeletionReconciler interface {
	// we received a reconcile request for an object that was removed from the cache
	ReconcileDeletion(cluster string, request Request)
}

// a Reconcile Loop runs resource reconcilers until the context gets cancelled
type Loop interface {
	RunReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate) error
}
type runner struct {
	name     string
	cw       clusterWatcher
	resource ezkube.Object
}

func NewLoop(name string, cw       ClusterWatcher, resource ezkube.Object) *runner {
	mcLoop := &mcLoop{}
	watcher := &clusterWatcher{
		ctx:      ctx,
		handlers: mcLoop,
		cancels:  newCancelSet(),
	}

	cw.RegisterClusterHandler(mcLoop)

	return &runner{name: name, mgr: mgr, resource: resource}
}

type userReconciler struct {
	ctx        context.Context
	reconciler Reconciler
	predicates []predicate.Predicate
}

type mcLoop struct {
	userReconcilers []userReconciler
	resource        ezkube.Object
}

func (m *mcLoop) RunReconciler(reconciler Reconciler, predicates ...predicate.Predicate) error {
	m.userReconcilers = append(m.userReconcilers, )
}

func (m *mcLoop) HandleAddCluster(ctx context.Context, cluster string, mgr manager.Manager) error {
	loopForCluster := reconcile.NewLoop("whatever", mgr, m.resource)
	for _, userRec := range m.userReconcilers {
		mcReconciler := &mcReconciler{
			cluster:        cluster,
			userReconciler: userRec.reconciler,
		}
		loopForCluster.RunReconciler(ctx, mcReconciler, userRec.predicates...)
	}
}

type mcReconciler struct {
	cluster        string
	userReconciler Reconciler
}

func (m mcReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	return m.userReconciler.Reconcile(m.cluster, object)
}
