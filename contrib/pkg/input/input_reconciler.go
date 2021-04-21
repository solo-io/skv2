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
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	if r.singleClusterReconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no single-cluster reconcile func provided; cannot reconcile %v", sets.Key(id))
	}
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "id", sets.Key(id))
	// never queue more than one event
	if r.queue.Len() < 2 {
		contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "id", sets.Key(id))
		r.queue.AddRateLimited(id)
	}

	return reconcile.Result{}, nil
}

func (r *inputReconciler) ReconcileRemoteGeneric(id ezkube.ClusterResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	if r.multiClusterReconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no multi-cluster reconcile func provided; cannot reconcile %v", sets.Key(id))
	}
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", id.GetClusterName(), "id", sets.Key(id))
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

	var isRemoteCluster bool

	// determine whether the resource has been read from a remote cluster
	// based on whether its ClusterName field is set
	if clusterResource, ok := key.(ezkube.ClusterResourceId); ok {
		if clusterResource.GetClusterName() != "" {
			isRemoteCluster = true
		}
	}

	var (
		requeue bool
		err     error
	)
	// TODO (ilackarms): this is a workaround for an issue where Relay sets the clusterName field of an object, but the underlying reconciler is a single-cluster reconciler.
	if isRemoteCluster && r.multiClusterReconcileFunc != nil {
		requeue, err = r.multiClusterReconcileFunc(key.(ezkube.ClusterResourceId))
	} else {
		requeue, err = r.singleClusterReconcileFunc(key.(ezkube.ResourceId))
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
