package input

// This file provides the base interface which is exposed via
// generated reconcilers in input_reconciler.gotmpl

import (
	"context"
	"sync"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/reconcile"
)

// reconcile resources from local and remote clusters.
// the passed resources can either be refs to a resource (caused by a deletion), or actual resources themselves.
type EventBasedReconcileFunc func(
	localEventObjs []ezkube.ResourceId,
	remoteEventObjs []ezkube.ClusterResourceId,
) (bool, error)

// input reconciler implements both the single and multi cluster reconcilers, for convenience.
type eventBasedInputReconciler struct {
	ctx           context.Context
	queue         workqueue.RateLimitingInterface
	reconcileFunc EventBasedReconcileFunc

	localEventCache map[string]ezkube.ResourceId
	localLock       sync.RWMutex

	remoteEventCache map[string]ezkube.ClusterResourceId
	remoteLock       sync.RWMutex
}

// Note(ilackarms): in the current implementation, the constructor
// also starts the reconciler's event processor in a goroutine.
// Make sure to cancel the parent context in order to ensure the goroutine started here is gc'ed.
// only one event will be processed per reconcileInterval.
func NewEventBasedInputReconciler(
	ctx context.Context,
	reconcileFunc EventBasedReconcileFunc,
	reconcileInterval time.Duration,
) InputReconciler {
	r := &eventBasedInputReconciler{
		ctx:              ctx,
		queue:            workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		reconcileFunc:    reconcileFunc,
		localEventCache:  map[string]ezkube.ResourceId{},
		remoteEventCache: map[string]ezkube.ClusterResourceId{},
	}
	go r.reconcileEventsForever(reconcileInterval)
	return r
}

// handle an event from the k8s event stream by adding to queue and localEventCache
func (r *eventBasedInputReconciler) ReconcileLocalGeneric(obj ezkube.ResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	if r.reconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no event based reconcile func provided; cannot reconcile %v", sets.Key(obj))
	}
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", sets.Key(obj))

	// add local event to cache
	r.addLocalEventObj(obj)

	// never queue more than one event
	if r.queue.Len() < 2 {
		contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", sets.Key(obj))
		r.queue.AddRateLimited(obj)
	}

	return reconcile.Result{}, nil
}

// handle an event from the k8s event stream by adding to queue and localEventCache
func (r *eventBasedInputReconciler) ReconcileRemoteGeneric(obj ezkube.ClusterResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	if r.reconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no event based reconcile func provided; cannot reconcile %v", sets.Key(obj))
	}
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", obj.GetClusterName(), "obj", sets.Key(obj))

	// add remote event to cache
	r.addRemoteEventObj(obj)

	// never queue more than one event
	if r.queue.Len() < 2 {
		contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", obj.GetClusterName(), "obj", sets.Key(obj))
		r.queue.AddRateLimited(obj)
	}

	return reconcile.Result{}, nil
}

// reconcile queued events until context is cancelled.
// blocking (runs from a goroutine)
func (r *eventBasedInputReconciler) reconcileEventsForever(reconcileInterval time.Duration) {
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
func (r *eventBasedInputReconciler) processNextWorkItem() bool {
	key, quit := r.queue.Get()
	if quit {
		return false
	}
	defer r.queue.Done(key)

	var (
		requeue bool
		err     error
	)

	localEventObjs, localEventKeys := r.getAllLocalEventObjs()
	remoteEventObjs, remoteEventKeys := r.getAllRemoteEventObjs()

	requeue, err = r.reconcileFunc(localEventObjs, remoteEventObjs)

	switch {
	case err != nil:
		contextutils.LoggerFrom(r.ctx).Errorw("encountered error reconciling state; retrying", "error", err)
		fallthrough
	case requeue:
		r.queue.AddRateLimited(key)
	default:
		r.deleteLocalEventObjs(localEventKeys)
		r.deleteRemoteEventObjs(remoteEventKeys)
		r.queue.Forget(key)
	}

	return true
}

// add a new object into the local event cache
func (r *eventBasedInputReconciler) addLocalEventObj(id ezkube.ResourceId) {
	// only add k8s objects
	if obj, ok := id.(client.Object); ok {
		r.localLock.Lock()
		defer r.localLock.Unlock()
		r.localEventCache[uniqueId(obj)] = obj
	}
}

// add a new object into the remote event cache
func (r *eventBasedInputReconciler) addRemoteEventObj(id ezkube.ClusterResourceId) {
	if obj, ok := id.(client.Object); ok {
		r.remoteLock.Lock()
		defer r.remoteLock.Unlock()
		r.remoteEventCache[uniqueId(obj)] = obj
	}
}

// retrieve all objects from the local event cache, as well as their associated keys
func (r *eventBasedInputReconciler) getAllLocalEventObjs() ([]ezkube.ResourceId, []string) {
	r.localLock.RLock()
	defer r.localLock.RUnlock()

	var eventObjs []ezkube.ResourceId
	var keys []string
	for key, obj := range r.localEventCache {
		eventObjs = append(eventObjs, obj)
		keys = append(keys, key)
	}

	return eventObjs, keys
}

// retrieve all objects from the event cache, as well as their associated keys
func (r *eventBasedInputReconciler) getAllRemoteEventObjs() ([]ezkube.ClusterResourceId, []string) {
	r.remoteLock.RLock()
	defer r.remoteLock.RUnlock()

	var remoteEventObjs []ezkube.ClusterResourceId
	var keys []string
	for key, obj := range r.remoteEventCache {
		remoteEventObjs = append(remoteEventObjs, obj)
		keys = append(keys, key)
	}

	return remoteEventObjs, keys
}

// deletes the local objects specified by its key from the event cache, signifying successful processing of those objects
func (r *eventBasedInputReconciler) deleteLocalEventObjs(keys []string) {
	r.localLock.Lock()
	defer r.localLock.Unlock()
	for _, key := range keys {
		delete(r.localEventCache, key)
	}
}

// deletes the remote objects specified by its key from the event cache, signifying successful processing of those objects
func (r *eventBasedInputReconciler) deleteRemoteEventObjs(keys []string) {
	r.remoteLock.Lock()
	defer r.remoteLock.Unlock()
	for _, key := range keys {
		delete(r.remoteEventCache, key)
	}
}

// build unique id for object with GVK + name/namespace/cluster
func uniqueId(obj client.Object) string {
	gvk := obj.GetObjectKind().GroupVersionKind()
	return gvk.Kind + "." + gvk.Group + "." + gvk.Version + "/" + sets.Key(obj)
}
