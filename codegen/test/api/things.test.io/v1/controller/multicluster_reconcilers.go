// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Paint Resource across clusters.
// implemented by the user
type MulticlusterPaintReconciler interface {
	ReconcilePaint(clusterName string, obj *things_test_io_v1.Paint) (reconcile.Result, error)
}

// Reconcile deletion events for the Paint Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterPaintDeletionReconciler interface {
	ReconcilePaintDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterPaintReconcilerFuncs struct {
	OnReconcilePaint         func(clusterName string, obj *things_test_io_v1.Paint) (reconcile.Result, error)
	OnReconcilePaintDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterPaintReconcilerFuncs) ReconcilePaint(clusterName string, obj *things_test_io_v1.Paint) (reconcile.Result, error) {
	if f.OnReconcilePaint == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcilePaint(clusterName, obj)
}

func (f *MulticlusterPaintReconcilerFuncs) ReconcilePaintDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcilePaintDeletion == nil {
		return
	}
	f.OnReconcilePaintDeletion(clusterName, req)
}

type MulticlusterPaintReconcileLoop interface {
	// AddMulticlusterPaintReconciler adds a MulticlusterPaintReconciler to the MulticlusterPaintReconcileLoop.
	AddMulticlusterPaintReconciler(ctx context.Context, rec MulticlusterPaintReconciler, predicates ...predicate.Predicate) error
}

type multiclusterPaintReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterPaintReconcileLoop) AddMulticlusterPaintReconciler(ctx context.Context, rec MulticlusterPaintReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericPaintMulticlusterReconciler{reconciler: rec}

	return m.loop.RunReconciler(ctx, genericReconciler, predicates...)
}

func NewPaintMulticlusterReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterPaintReconcileLoop {
	return &multiclusterPaintReconcileLoop{loop: multicluster.NewLoop(name, cw, &things_test_io_v1.Paint{})}
}

type genericPaintMulticlusterReconciler struct {
	reconciler MulticlusterPaintReconciler
}

func (g genericPaintMulticlusterReconciler) ReconcilePaintDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(PaintDeletionReconciler); ok {
		deletionReconciler.ReconcilePaintDeletion(req)
	}
}

func (g genericPaintMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return g.reconciler.ReconcilePaint(cluster, obj)
}

// Reconcile Upsert events for the ClusterResource Resource across clusters.
// implemented by the user
type MulticlusterClusterResourceReconciler interface {
	ReconcileClusterResource(clusterName string, obj *things_test_io_v1.ClusterResource) (reconcile.Result, error)
}

// Reconcile deletion events for the ClusterResource Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterClusterResourceDeletionReconciler interface {
	ReconcileClusterResourceDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterClusterResourceReconcilerFuncs struct {
	OnReconcileClusterResource         func(clusterName string, obj *things_test_io_v1.ClusterResource) (reconcile.Result, error)
	OnReconcileClusterResourceDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterClusterResourceReconcilerFuncs) ReconcileClusterResource(clusterName string, obj *things_test_io_v1.ClusterResource) (reconcile.Result, error) {
	if f.OnReconcileClusterResource == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileClusterResource(clusterName, obj)
}

func (f *MulticlusterClusterResourceReconcilerFuncs) ReconcileClusterResourceDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileClusterResourceDeletion == nil {
		return
	}
	f.OnReconcileClusterResourceDeletion(clusterName, req)
}

type MulticlusterClusterResourceReconcileLoop interface {
	// AddMulticlusterClusterResourceReconciler adds a MulticlusterClusterResourceReconciler to the MulticlusterClusterResourceReconcileLoop.
	AddMulticlusterClusterResourceReconciler(ctx context.Context, rec MulticlusterClusterResourceReconciler, predicates ...predicate.Predicate) error
}

type multiclusterClusterResourceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterClusterResourceReconcileLoop) AddMulticlusterClusterResourceReconciler(ctx context.Context, rec MulticlusterClusterResourceReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericClusterResourceMulticlusterReconciler{reconciler: rec}

	return m.loop.RunReconciler(ctx, genericReconciler, predicates...)
}

func NewClusterResourceMulticlusterReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterClusterResourceReconcileLoop {
	return &multiclusterClusterResourceReconcileLoop{loop: multicluster.NewLoop(name, cw, &things_test_io_v1.ClusterResource{})}
}

type genericClusterResourceMulticlusterReconciler struct {
	reconciler MulticlusterClusterResourceReconciler
}

func (g genericClusterResourceMulticlusterReconciler) ReconcileClusterResourceDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(ClusterResourceDeletionReconciler); ok {
		deletionReconciler.ReconcileClusterResourceDeletion(req)
	}
}

func (g genericClusterResourceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*things_test_io_v1.ClusterResource)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ClusterResource handler received event for %T", object)
	}
	return g.reconciler.ReconcileClusterResource(cluster, obj)
}
