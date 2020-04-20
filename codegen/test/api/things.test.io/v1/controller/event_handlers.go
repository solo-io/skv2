// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Paint Resource
type PaintEventHandler interface {
	CreatePaint(obj *things_test_io_v1.Paint) error
	UpdatePaint(old, new *things_test_io_v1.Paint) error
	DeletePaint(obj *things_test_io_v1.Paint) error
	GenericPaint(obj *things_test_io_v1.Paint) error
}

type PaintEventHandlerFuncs struct {
	OnCreate  func(obj *things_test_io_v1.Paint) error
	OnUpdate  func(old, new *things_test_io_v1.Paint) error
	OnDelete  func(obj *things_test_io_v1.Paint) error
	OnGeneric func(obj *things_test_io_v1.Paint) error
}

func (f *PaintEventHandlerFuncs) CreatePaint(obj *things_test_io_v1.Paint) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *PaintEventHandlerFuncs) DeletePaint(obj *things_test_io_v1.Paint) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *PaintEventHandlerFuncs) UpdatePaint(objOld, objNew *things_test_io_v1.Paint) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *PaintEventHandlerFuncs) GenericPaint(obj *things_test_io_v1.Paint) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type PaintEventWatcher interface {
	AddEventHandler(ctx context.Context, h PaintEventHandler, predicates ...predicate.Predicate) error
}

type paintEventWatcher struct {
	watcher events.EventWatcher
}

func NewPaintEventWatcher(name string, mgr manager.Manager) PaintEventWatcher {
	return &paintEventWatcher{
		watcher: events.NewWatcher(name, mgr, &things_test_io_v1.Paint{}),
	}
}

func (c *paintEventWatcher) AddEventHandler(ctx context.Context, h PaintEventHandler, predicates ...predicate.Predicate) error {
	handler := genericPaintHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericPaintHandler implements a generic events.EventHandler
type genericPaintHandler struct {
	handler PaintEventHandler
}

func (h genericPaintHandler) Create(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return h.handler.CreatePaint(obj)
}

func (h genericPaintHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return h.handler.DeletePaint(obj)
}

func (h genericPaintHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", old)
	}
	objNew, ok := new.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", new)
	}
	return h.handler.UpdatePaint(objOld, objNew)
}

func (h genericPaintHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return h.handler.GenericPaint(obj)
}

// Handle events for the ClusterResource Resource
type ClusterResourceEventHandler interface {
	CreateClusterResource(obj *things_test_io_v1.ClusterResource) error
	UpdateClusterResource(old, new *things_test_io_v1.ClusterResource) error
	DeleteClusterResource(obj *things_test_io_v1.ClusterResource) error
	GenericClusterResource(obj *things_test_io_v1.ClusterResource) error
}

type ClusterResourceEventHandlerFuncs struct {
	OnCreate  func(obj *things_test_io_v1.ClusterResource) error
	OnUpdate  func(old, new *things_test_io_v1.ClusterResource) error
	OnDelete  func(obj *things_test_io_v1.ClusterResource) error
	OnGeneric func(obj *things_test_io_v1.ClusterResource) error
}

func (f *ClusterResourceEventHandlerFuncs) CreateClusterResource(obj *things_test_io_v1.ClusterResource) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ClusterResourceEventHandlerFuncs) DeleteClusterResource(obj *things_test_io_v1.ClusterResource) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ClusterResourceEventHandlerFuncs) UpdateClusterResource(objOld, objNew *things_test_io_v1.ClusterResource) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ClusterResourceEventHandlerFuncs) GenericClusterResource(obj *things_test_io_v1.ClusterResource) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ClusterResourceEventWatcher interface {
	AddEventHandler(ctx context.Context, h ClusterResourceEventHandler, predicates ...predicate.Predicate) error
}

type clusterResourceEventWatcher struct {
	watcher events.EventWatcher
}

func NewClusterResourceEventWatcher(name string, mgr manager.Manager) ClusterResourceEventWatcher {
	return &clusterResourceEventWatcher{
		watcher: events.NewWatcher(name, mgr, &things_test_io_v1.ClusterResource{}),
	}
}

func (c *clusterResourceEventWatcher) AddEventHandler(ctx context.Context, h ClusterResourceEventHandler, predicates ...predicate.Predicate) error {
	handler := genericClusterResourceHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericClusterResourceHandler implements a generic events.EventHandler
type genericClusterResourceHandler struct {
	handler ClusterResourceEventHandler
}

func (h genericClusterResourceHandler) Create(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.ClusterResource)
	if !ok {
		return errors.Errorf("internal error: ClusterResource handler received event for %T", object)
	}
	return h.handler.CreateClusterResource(obj)
}

func (h genericClusterResourceHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.ClusterResource)
	if !ok {
		return errors.Errorf("internal error: ClusterResource handler received event for %T", object)
	}
	return h.handler.DeleteClusterResource(obj)
}

func (h genericClusterResourceHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*things_test_io_v1.ClusterResource)
	if !ok {
		return errors.Errorf("internal error: ClusterResource handler received event for %T", old)
	}
	objNew, ok := new.(*things_test_io_v1.ClusterResource)
	if !ok {
		return errors.Errorf("internal error: ClusterResource handler received event for %T", new)
	}
	return h.handler.UpdateClusterResource(objOld, objNew)
}

func (h genericClusterResourceHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.ClusterResource)
	if !ok {
		return errors.Errorf("internal error: ClusterResource handler received event for %T", object)
	}
	return h.handler.GenericClusterResource(obj)
}
