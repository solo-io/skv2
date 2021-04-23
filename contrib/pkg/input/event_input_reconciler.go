package input

// This file provides the base interface which is exposed via
// generated reconcilers in input_reconciler.gotmpl

import (
	"context"
	"sync"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/workqueue"
)

// reconcile resources from local and remote clusters.
// the passed resources can either be refs to a resource (caused by a deletion), or actual resources themselves.
type EventBasedReconcileFunc func(
	localEventObjs map[schema.GroupVersionKind][]ezkube.ResourceId,
	remoteEventObjs map[schema.GroupVersionKind][]ezkube.ClusterResourceId,
) (bool, error)

// the InputReconciler reconciles events for input resources in a single cluster
type EventBasedInputReconciler interface {
	// reconcile the generic resource type in the local cluster.
	// this function is called from generated code.
	ReconcileLocalGeneric(gvk schema.GroupVersionKind, id ezkube.ResourceId) (reconcile.Result, error)

	// reconcile the generic resource type in a remote cluster.
	// this function is called from generated code.
	ReconcileRemoteGeneric(gvk schema.GroupVersionKind, id ezkube.ClusterResourceId) (reconcile.Result, error)
}

// input reconciler implements both the single and multi cluster reconcilers, for convenience.
type eventBasedInputReconciler struct {
	ctx           context.Context
	queue         workqueue.RateLimitingInterface
	reconcileFunc EventBasedReconcileFunc

	localEventCache map[schema.GroupVersionKind]sets.ResourceSet
	localLock       sync.RWMutex

	remoteEventCache map[schema.GroupVersionKind]sets.ResourceSet
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
) EventBasedInputReconciler {
	r := &eventBasedInputReconciler{
		ctx:              ctx,
		queue:            workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		reconcileFunc:    reconcileFunc,
		localEventCache:  map[schema.GroupVersionKind]sets.ResourceSet{},
		remoteEventCache: map[schema.GroupVersionKind]sets.ResourceSet{},
	}
	go r.reconcileEventsForever(reconcileInterval)
	return r
}

// handle an event from the k8s event stream by adding to queue and localEventCache
func (r *eventBasedInputReconciler) ReconcileLocalGeneric(gvk schema.GroupVersionKind, obj ezkube.ResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	if r.reconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no event based reconcile func provided; cannot reconcile %v", sets.Key(obj))
	}
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", sets.Key(obj))

	// add local event to cache
	r.addLocalEventObj(gvk, obj)

	// never queue more than one event
	if r.queue.Len() < 1 {
		contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "obj", sets.Key(obj))
		r.queue.AddRateLimited(obj)
	}

	return reconcile.Result{}, nil
}

// handle an event from the k8s event stream by adding to queue and localEventCache
func (r *eventBasedInputReconciler) ReconcileRemoteGeneric(gvk schema.GroupVersionKind, obj ezkube.ClusterResourceId) (reconcile.Result, error) {
	if r.ctx == nil {
		return reconcile.Result{}, eris.Errorf("internal error: reconciler not started")
	}
	if r.reconcileFunc == nil {
		return reconcile.Result{}, eris.Errorf("internal error: no event based reconcile func provided; cannot reconcile %v", sets.Key(obj))
	}
	contextutils.LoggerFrom(r.ctx).Debugw("reconciling event", "cluster", obj.GetClusterName(), "obj", sets.Key(obj))

	// add remote event to cache
	r.addRemoteEventObj(gvk, obj)

	// never queue more than one event
	if r.queue.Len() < 1 {
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

	localEventObjs := r.getAllLocalEventObjs()
	remoteEventObjs := r.getAllRemoteEventObjs()

	requeue, err = r.reconcileFunc(localEventObjs, remoteEventObjs)

	switch {
	case err != nil:
		contextutils.LoggerFrom(r.ctx).Errorw("encountered error reconciling state; retrying", "error", err)
		fallthrough
	case requeue:
		r.queue.AddRateLimited(key)
	default:
		r.deleteLocalEventObjs(localEventObjs)
		r.deleteRemoteEventObjs(remoteEventObjs)
		r.queue.Forget(key)
	}

	return true
}

// add a new event into the local event cache
// id will be a client.Object for create or update events, and a ref delete events
func (r *eventBasedInputReconciler) addLocalEventObj(gvk schema.GroupVersionKind, id ezkube.ResourceId) {
	r.localLock.Lock()
	defer r.localLock.Unlock()

	cachedObjs, ok := r.localEventCache[gvk]
	if !ok {
		r.localEventCache[gvk] = sets.NewResourceSet()
		cachedObjs = sets.NewResourceSet()
	}
	cachedObjs.Insert(id)
}

// add a new object into the remote event cache
// id will be a client.Object for create or update events, and a ref delete events
func (r *eventBasedInputReconciler) addRemoteEventObj(gvk schema.GroupVersionKind, id ezkube.ClusterResourceId) {
	r.remoteLock.Lock()
	defer r.remoteLock.Unlock()

	cachedObjs, ok := r.remoteEventCache[gvk]
	if !ok {
		r.remoteEventCache[gvk] = sets.NewResourceSet()
		cachedObjs = sets.NewResourceSet()
	}
	cachedObjs.Insert(id)
}

// retrieve all objects from the local event cache, as well as their associated keys
func (r *eventBasedInputReconciler) getAllLocalEventObjs() map[schema.GroupVersionKind][]ezkube.ResourceId {
	r.localLock.RLock()
	defer r.localLock.RUnlock()

	eventObjs := map[schema.GroupVersionKind][]ezkube.ResourceId{}

	for gvk, objs := range r.localEventCache {
		for _, obj := range objs.List() {
			eventObjs[gvk] = append(eventObjs[gvk], obj)
		}
	}

	return eventObjs
}

// retrieve all objects from the event cache, as well as their associated keys
func (r *eventBasedInputReconciler) getAllRemoteEventObjs() map[schema.GroupVersionKind][]ezkube.ClusterResourceId {
	r.remoteLock.RLock()
	defer r.remoteLock.RUnlock()

	eventObjs := map[schema.GroupVersionKind][]ezkube.ClusterResourceId{}

	for gvk, objs := range r.remoteEventCache {
		for _, obj := range objs.List() {
			// all objects in remote event cache are guaranteed to be of type ezkube.ClusterResourceId
			eventObjs[gvk] = append(eventObjs[gvk], obj.(ezkube.ClusterResourceId))
		}
	}

	return eventObjs
}

// deletes the local objects specified by its key from the event cache, signifying successful processing of those objects
func (r *eventBasedInputReconciler) deleteLocalEventObjs(typedObjs map[schema.GroupVersionKind][]ezkube.ResourceId) {
	r.localLock.Lock()
	defer r.localLock.Unlock()

	for gvk, objs := range typedObjs {
		cachedObjs, ok := r.localEventCache[gvk]

		if !ok {
			contextutils.LoggerFrom(r.ctx).DPanicf("deleting object with GVK that does not exist in cache: %v", gvk)
			continue
		}

		for _, obj := range objs {
			cachedObjs.Delete(obj)
		}
	}
}

// deletes the remote objects specified by its key from the event cache, signifying successful processing of those objects
func (r *eventBasedInputReconciler) deleteRemoteEventObjs(typedObjs map[schema.GroupVersionKind][]ezkube.ClusterResourceId) {
	r.remoteLock.Lock()
	defer r.remoteLock.Unlock()

	for gvk, objs := range typedObjs {
		cachedObjs, ok := r.remoteEventCache[gvk]

		if !ok {
			contextutils.LoggerFrom(r.ctx).DPanicf("deleting object with GVK that does not exist in cache: %v", gvk)
			continue
		}

		for _, obj := range objs {
			cachedObjs.Delete(obj)
		}
	}
}
