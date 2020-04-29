package multicluster

import (
	"context"

	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
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

// Loop runs resource reconcilers until the context gets cancelled
type Loop interface {
	// AddReconciler adds a reconciler to a slice of reconcilers that will be run against
	AddReconciler(ctx context.Context, reconciler Reconciler, predicates ...predicate.Predicate)
}
