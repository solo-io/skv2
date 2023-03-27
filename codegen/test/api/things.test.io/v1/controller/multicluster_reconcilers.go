// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller



import (
	"context"

	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CueBug Resource across clusters.
// implemented by the user
type MulticlusterCueBugReconciler interface {
	ReconcileCueBug(clusterName string, obj *things_test_io_v1.CueBug) (reconcile.Result, error)
}

// Reconcile deletion events for the CueBug Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterCueBugDeletionReconciler interface {
	ReconcileCueBugDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterCueBugReconcilerFuncs struct {
	OnReconcileCueBug         func(clusterName string, obj *things_test_io_v1.CueBug) (reconcile.Result, error)
	OnReconcileCueBugDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterCueBugReconcilerFuncs) ReconcileCueBug(clusterName string, obj *things_test_io_v1.CueBug) (reconcile.Result, error) {
	if f.OnReconcileCueBug == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCueBug(clusterName, obj)
}

func (f *MulticlusterCueBugReconcilerFuncs) ReconcileCueBugDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileCueBugDeletion == nil {
		return nil
	}
	return f.OnReconcileCueBugDeletion(clusterName, req)
}

type MulticlusterCueBugReconcileLoop interface {
	// AddMulticlusterCueBugReconciler adds a MulticlusterCueBugReconciler to the MulticlusterCueBugReconcileLoop.
	AddMulticlusterCueBugReconciler(ctx context.Context, rec MulticlusterCueBugReconciler, predicates ...predicate.Predicate)
}

type multiclusterCueBugReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterCueBugReconcileLoop) AddMulticlusterCueBugReconciler(ctx context.Context, rec MulticlusterCueBugReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericCueBugMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterCueBugReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterCueBugReconcileLoop {
	return &multiclusterCueBugReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &things_test_io_v1.CueBug{}, options)}
}

type genericCueBugMulticlusterReconciler struct {
	reconciler MulticlusterCueBugReconciler
}

func (g genericCueBugMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterCueBugDeletionReconciler); ok {
		return deletionReconciler.ReconcileCueBugDeletion(cluster, req)
	}
	return nil
}

func (g genericCueBugMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*things_test_io_v1.CueBug)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CueBug handler received event for %T", object)
	}
	return g.reconciler.ReconcileCueBug(cluster, obj)
}
