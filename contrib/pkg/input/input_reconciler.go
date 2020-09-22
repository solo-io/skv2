package input

// This file provides the base interface which is exposed via
// generated reconcilers in input_reconciler.gotmpl

import (
	"context"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/client-go/util/workqueue"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/reconcile"
)

// reconcile a resource in a single cluster.
// the passed resource can either be a ref to a resource (caused by a deletion), or an actual resource itself.
type SingleClusterReconcileFunc func(id ezkube.ResourceId) (bool, error)

// the SingleClusterReconciler reconciles events for input resources in a single cluster
type SingleClusterReconciler interface {
	// reconcile the generic resource type.
	// this function is called from generated code.
	ReconcileGeneric(id ezkube.ResourceId) (reconcile.Result, error)
}

// reconcile a resource across multiple clusters.
// the passed resource can either be a ref to a resource (caused by a deletion), or an actual resource itself. ClusterName will always be set on the object.
type MultiClusterReconcileFunc func(id ezkube.ClusterResourceId) (bool, error)

// the MultiClusterReconciler reconciles events for input resources across clusters
type MultiClusterReconciler interface {
	// reconcile the generic resource type.
	// this function is called from generated code.
	ReconcileClusterGeneric(id ezkube.ClusterResourceId) (reconcile.Result, error)
}

// input reconciler implements both the single and multi cluster reconcilers, for convenience.
type inputReconciler struct {
	ctx                        context.Context
	queue                      workqueue.RateLimitingInterface
	multiClusterReconcileFunc  MultiClusterReconcileFunc
	singleClusterReconcileFunc SingleClusterReconcileFunc
}

// Note(ilackarms): in the current implementation, the constructor
// also starts the reconciler's event processor in a goroutine.
// Make sure to cancel the parent context in order to ensure the goroutine started here is gc'ed.
// only one event will be processed per reconcileInterval.
func NewMultiClusterReconcilerImpl(
	ctx context.Context,
	reconcileFunc MultiClusterReconcileFunc,
	reconcileInterval time.Duration,
) MultiClusterReconciler {
	r := &inputReconciler{
		ctx:                       ctx,
		queue:                     workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		multiClusterReconcileFunc: reconcileFunc,
	}
	go r.reconcileEventsForever(reconcileInterval)
	return r
}

// Note(ilackarms): in the current implementation, the constructor
// also starts the reconciler's event processor in a goroutine.
// Make sure to cancel the parent context in order to ensure the goroutine started here is gc'ed.
// only one event will be processed per reconcileInterval.
func NewSingleClusterReconciler(
	ctx context.Context,
	reconcileFunc SingleClusterReconcileFunc,
	reconcileInterval time.Duration,
) SingleClusterReconciler {
	r := &inputReconciler{
		ctx:                        ctx,
		queue:                      workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		singleClusterReconcileFunc: reconcileFunc,
	}
	go r.reconcileEventsForever(reconcileInterval)
	return r
}

func (r *inputReconciler) ReconcileGeneric(id ezkube.ResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	// never queue more than one event
	if r.queue.Len() < 2 {
		contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "id", sets.Key(id))
		r.queue.AddRateLimited(id)
	}

	return reconcile.Result{}, nil
}

func (r *inputReconciler) ReconcileClusterGeneric(id ezkube.ClusterResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	// never queue more than one event
	if r.queue.Len() < 2 {
		contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", id.GetClusterName(), "id", sets.Key(id))
		r.queue.AddRateLimited(id)
	}

	return reconcile.Result{}, nil
}

// reconcile queued events until context is cancelled.
// blocking (runs from a goroutine)
func (r *inputReconciler) reconcileEventsForever(reconcileInterval time.Duration) {
	for r.processNextWorkItem() {
		select {
		case <-r.ctx.Done():
			return
		default:
			time.Sleep(reconcileInterval)
		}
	}
}

// processNextWorkItem deals with one key off the queue.  It returns false when it's time to quit.
func (r *inputReconciler) processNextWorkItem() bool {
	key, quit := r.queue.Get()
	if quit {
		return false
	}
	defer r.queue.Done(key)

	var (
		requeue bool
		err     error
	)
	switch {
	case r.singleClusterReconcileFunc != nil:
		requeue, err = r.singleClusterReconcileFunc(key.(ezkube.ResourceId))
	case r.multiClusterReconcileFunc != nil:
		requeue, err = r.multiClusterReconcileFunc(key.(ezkube.ClusterResourceId))
	}

	switch {
	case err != nil:
		contextutils.LoggerFrom(r.ctx).Errorw("encountered error reconciling state; retrying", "error", err)
		fallthrough
	case requeue:
		r.queue.AddRateLimited(key)
	default:
		r.queue.Forget(key)
	}

	return true
}
