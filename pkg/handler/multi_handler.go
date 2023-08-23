package handler

import (
	"context"

	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

// MultiHandler wraps and calls multiple event handlers as a single handler.Handler
type MultiHandler struct {
	Handlers []handler.EventHandler
}

func (h *MultiHandler) Create(ctx context.Context, evt event.CreateEvent, queue workqueue.RateLimitingInterface) {
	for _, hl := range h.Handlers {
		hl.Create(ctx, evt, queue)
	}
}

func (h *MultiHandler) Update(ctx context.Context, evt event.UpdateEvent, queue workqueue.RateLimitingInterface) {
	for _, hl := range h.Handlers {
		hl.Update(ctx, evt, queue)
	}
}

func (h *MultiHandler) Delete(ctx context.Context, evt event.DeleteEvent, queue workqueue.RateLimitingInterface) {
	for _, hl := range h.Handlers {
		hl.Delete(ctx, evt, queue)
	}
}

func (h *MultiHandler) Generic(ctx context.Context, evt event.GenericEvent, queue workqueue.RateLimitingInterface) {
	for _, hl := range h.Handlers {
		hl.Generic(ctx, evt, queue)
	}
}
