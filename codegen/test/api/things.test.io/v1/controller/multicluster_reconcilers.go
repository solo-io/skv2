// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	"github.com/pkg/errors"
	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type MulticlusterPaintReconciler interface {
	ReconcilePaint(clusterName string, obj *things_test_io_v1.Paint) (reconcile.Result, error)
}

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
	RunPaintReconciler(ctx context.Context, rec MulticlusterPaintReconciler, predicates ...predicate.Predicate) error
}

type multiclusterPaintReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterPaintReconcileLoop) RunPaintReconciler(ctx context.Context, rec MulticlusterPaintReconciler, predicates ...predicate.Predicate) error {
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
