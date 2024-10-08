package handler

import (
	"context"

	"github.com/solo-io/skv2/pkg/request"
	skqueue "github.com/solo-io/skv2/pkg/workqueue"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var enqueueMultiClusterLog = log.Log.WithName("eventhandler").WithName("BroadcastRequests")

var _ handler.EventHandler = &BroadcastRequests{}

// BroadcastRequests enqueues statically defined requests across clusters
// whenever an event is received. Use this to propagate a list of requests
// to queues shared across cluster managers.
// This is used by SKv2 to enqueueRequestsAllClusters requests for a primary level resource
// whenever a watched input resource changes, regardless of the cluster the primary resource lives in.
type BroadcastRequests struct {
	// the set of all requests to enqueueRequestsAllClusters by the target cluster (where the primary resource lives)
	RequestsToEnqueue *request.MultiClusterRequests

	// use this to queue requests to controllers registered to another manager
	WorkQueues *skqueue.MultiClusterQueues
}

// Create implements EventHandler
func (e *BroadcastRequests) Create(ctx context.Context, evt event.CreateEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	if evt.Object == nil {
		enqueueMultiClusterLog.Error(nil, "CreateEvent received with no metadata", "event", evt)
		return
	}
	e.enqueueRequestsAllClusters()
}

// Update implements EventHandler
func (e *BroadcastRequests) Update(ctx context.Context, evt event.UpdateEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	if evt.ObjectOld != nil {
		e.enqueueRequestsAllClusters()
	} else {
		enqueueMultiClusterLog.Error(nil, "UpdateEvent received with no old metadata", "event", evt)
	}

	if evt.ObjectNew != nil {
		e.enqueueRequestsAllClusters()
	} else {
		enqueueMultiClusterLog.Error(nil, "UpdateEvent received with no new metadata", "event", evt)
	}
}

// Delete implements EventHandler
func (e *BroadcastRequests) Delete(ctx context.Context, evt event.DeleteEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	if evt.Object == nil {
		enqueueMultiClusterLog.Error(nil, "DeleteEvent received with no metadata", "event", evt)
		return
	}
	e.enqueueRequestsAllClusters()
}

// Generic implements EventHandler
func (e *BroadcastRequests) Generic(ctx context.Context, evt event.GenericEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	if evt.Object == nil {
		enqueueMultiClusterLog.Error(nil, "GenericEvent received with no metadata", "event", evt)
		return
	}
	e.enqueueRequestsAllClusters()
}

func (e *BroadcastRequests) enqueueRequestsAllClusters() {
	e.RequestsToEnqueue.Each(func(cluster string, i reconcile.Request) {
		q := e.WorkQueues.Get(cluster)
		if q == nil {
			enqueueMultiClusterLog.Error(nil, "Cannot enqueue request, no queue registered for cluster", "request", i, "cluster", cluster)
			return
		}
		q.Add(i)
	})
}
