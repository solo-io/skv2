package multicluster_v2

import (
	"context"

	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type Reconciler[T client.Object] interface {
	// reconcile an object
	// requeue the object if returning an error, or a non-zero "requeue-after" duration
	Reconcile(ctx context.Context, cluster string, object T) (reconcile.Result, error)
}

type DeletionReconciler[T client.Object] interface {
	// we received a reconcile request for an object that was removed from the cache
	// requeue the object if returning an error,
	ReconcileDeletion(ctx context.Context, cluster string, request reconcile.Request) error
}

type ReconcilerFuncs[T client.Object] struct {
	ReconcileFunc func(ctx context.Context, cluster string, object T) (reconcile.Result, error)
	ReconcileDeletionFunc func(ctx context.Context, cluster string, request reconcile.Request) error
}

func (r *ReconcilerFuncs[T]) Reconcile(ctx context.Context, cluster string, object T) (reconcile.Result, error) {
	if r.ReconcileFunc != nil {
		return r.ReconcileFunc(ctx, cluster, object)
	}
	return reconcile.Result{}, nil
}

func (r *ReconcilerFuncs[T]) ReconcileDeletion(ctx context.Context, cluster string, request reconcile.Request) error {
	if r.ReconcileDeletionFunc != nil {
		return r.ReconcileDeletionFunc(ctx, cluster, request)
	}
	return nil
}

// Loop runs resource reconcilers until the context gets cancelled
type Loop[T client.Object] interface {
	// AddReconciler adds a reconciler to a slice of reconcilers that will be run against
	AddReconciler(ctx context.Context, reconciler Reconciler[T], predicates ...predicate.Predicate)
}
