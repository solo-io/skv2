package input

// This file provides the base interface which is exposed via
// generated reconcilers in input_reconciler.gotmpl

import (
	"context"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"k8s.io/client-go/util/workqueue"
)

// reconcile a resource in a single cluster.
// the passed resource can either be a ref to a resource (caused by a deletion), or an actual resource itself.
type SingleClusterReconcileFunc func(id ezkube.ResourceId) (bool, error)

// reconcile a resource across multiple clusters.
// the passed resource can either be a ref to a resource (caused by a deletion), or an actual resource itself. ClusterName will always be set on the object.
type MultiClusterReconcileFunc func(id ezkube.ClusterResourceId) (bool, error)

// the InputReconciler reconciles events for input resources in a single cluster
type InputReconciler interface {
	// reconcile the generic resource type in the local cluster.
	// this function is called from generated code.
	ReconcileLocalGeneric(id ezkube.ResourceId) (reconcile.Result, error)

	// reconcile the generic resource type in a remote cluster.
	// this function is called from generated code.
	ReconcileRemoteGeneric(id ezkube.ClusterResourceId) (reconcile.Result, error)
}

// input reconciler implements both the single and multi cluster reconcilers, for convenience.
type inputReconciler struct {
	ctx                        context.Context
	queue                      workqueue.RateLimitingInterface
	multiClusterReconcileFunc  MultiClusterReconcileFunc
	singleClusterReconcileFunc SingleClusterReconcileFunc
}

const keySeparator = "/"

// Note(ilackarms): in the current implementation, the constructor
// also starts the reconciler's event processor in a goroutine.
// Make sure to cancel the parent context in order to ensure the goroutine started here is gc'ed.
// only one event will be processed per reconcileInterval.
func NewInputReconciler(
	ctx context.Context,
	multiClusterReconcileFunc MultiClusterReconcileFunc,
	singleClusterReconcileFunc SingleClusterReconcileFunc,
	reconcileInterval time.Duration,
) InputReconciler {
	r := &inputReconciler{
		ctx:                        ctx,
		queue:                      workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		multiClusterReconcileFunc:  multiClusterReconcileFunc,
		singleClusterReconcileFunc: singleClusterReconcileFunc,
	}
	go r.reconcileEventsForever(reconcileInterval)
	return r
}

func (r *inputReconciler) ReconcileLocalGeneric(id ezkube.ResourceId) (reconcile.Result, error) {
	if r.singleClusterReconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no single-cluster reconcile func provided; cannot reconcile %v", sets.TypedKey(id))
	}

	return r.reconcileGeneric(id)
}

func (r *inputReconciler) ReconcileRemoteGeneric(id ezkube.ClusterResourceId) (reconcile.Result, error) {
	if r.multiClusterReconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no multi-cluster reconcile func provided; cannot reconcile %v", sets.TypedKey(id))
	}

	return r.reconcileGeneric(id)
}

func (r *inputReconciler) reconcileGeneric(id ezkube.ResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}

	// no need to queue more than one event in reconcile-the-world approach
	if r.queue.Len() == 0 {
		contextutils.LoggerFrom(r.ctx).Debugw("adding event to reconciler queue", "id", sets.TypedKey(id))
		r.queue.AddRateLimited(ezkube.KeyWithSeparator(id, keySeparator))
	} else {
		contextutils.LoggerFrom(r.ctx).Debugw("dropping event as there are objects in the reconciler's queue", "id", sets.TypedKey(id))
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
	queueItem, quit := r.queue.Get()
	if quit {
		return false
	}
	defer r.queue.Done(queueItem)

	// get the string representation of the queue item
	key, ok := queueItem.(string)
	if !ok {
		contextutils.LoggerFrom(r.ctx).Errorw("got a work queue item of non-string type", "item", queueItem)
		r.queue.Forget(queueItem)
		return true
	}

	// convert the key to a ResourceId/ClusterResourceId
	var err error
	resource, err := ezkube.ResourceIdFromKeyWithSeparator(key, keySeparator)
	if err != nil {
		contextutils.LoggerFrom(r.ctx).Errorw("could not convert work queue item to resource", "error", err)
		r.queue.Forget(key)
		return true
	}

	// determine whether the resource has been read from a remote cluster
	// based on whether its ClusterName field is set
	var isRemoteCluster bool
	if clusterResource, ok := resource.(ezkube.ClusterResourceId); ok {
		if ezkube.GetClusterName(clusterResource) != "" {
			isRemoteCluster = true
		}
	}

	var requeue bool
	// TODO (ilackarms): this is a workaround for an issue where Relay sets the clusterName field of an object, but the underlying reconciler is a single-cluster reconciler.
	if isRemoteCluster && r.multiClusterReconcileFunc != nil {
		requeue, err = r.multiClusterReconcileFunc(resource.(ezkube.ClusterResourceId))
	} else {
		requeue, err = r.singleClusterReconcileFunc(resource)
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
