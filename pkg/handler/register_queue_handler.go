package handler

import (
	"context"
	"sync"

	apqueue "github.com/solo-io/skv2/pkg/workqueue"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// QueueRegisteringHandler registers the queue on the first create event
// it receives to a multi cluster queue registry.
func QueueRegisteringHandler(cluster string, queues *apqueue.MultiClusterQueues) handler.EventHandler {
	do := &sync.Once{}
	return &handler.Funcs{
		CreateFunc: func(ctx context.Context, _ event.CreateEvent, queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
			do.Do(func() {
				queues.Set(cluster, queue)
			})
		},
	}
}
