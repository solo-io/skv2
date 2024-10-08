package handler

import (
	"context"

	"github.com/solo-io/skv2/pkg/request"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ handler.EventHandler = &MultiClusterRequestTracker{}

// MultiClusterRequestTracker tracks reconcile requests across clusters
// It is used to map requests for input resources (in any cluster)
// back to the original
type MultiClusterRequestTracker struct {
	Cluster  string
	Requests *request.MultiClusterRequests
}

func (h *MultiClusterRequestTracker) Create(ctx context.Context, evt event.CreateEvent, queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	h.Requests.Append(h.Cluster, RequestForObject(evt.Object))
}

func (h *MultiClusterRequestTracker) Delete(ctx context.Context, evt event.DeleteEvent, queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	h.Requests.Remove(h.Cluster, RequestForObject(evt.Object))
}

func (h *MultiClusterRequestTracker) Update(context.Context, event.UpdateEvent, workqueue.TypedRateLimitingInterface[reconcile.Request]) {
}

func (h *MultiClusterRequestTracker) Generic(context.Context, event.GenericEvent, workqueue.TypedRateLimitingInterface[reconcile.Request]) {
}

func RequestForObject(meta v1.Object) reconcile.Request {
	return reconcile.Request{NamespacedName: types.NamespacedName{
		Name:      meta.GetName(),
		Namespace: meta.GetNamespace(),
	}}
}
