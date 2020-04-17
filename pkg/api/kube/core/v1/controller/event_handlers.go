// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	core_v1 "k8s.io/api/core/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Secret Resource
// DEPRECATED: Prefer reconciler pattern.
type SecretEventHandler interface {
	CreateSecret(obj *core_v1.Secret) error
	UpdateSecret(old, new *core_v1.Secret) error
	DeleteSecret(obj *core_v1.Secret) error
	GenericSecret(obj *core_v1.Secret) error
}

type SecretEventHandlerFuncs struct {
	OnCreate  func(obj *core_v1.Secret) error
	OnUpdate  func(old, new *core_v1.Secret) error
	OnDelete  func(obj *core_v1.Secret) error
	OnGeneric func(obj *core_v1.Secret) error
}

func (f *SecretEventHandlerFuncs) CreateSecret(obj *core_v1.Secret) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *SecretEventHandlerFuncs) DeleteSecret(obj *core_v1.Secret) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *SecretEventHandlerFuncs) UpdateSecret(objOld, objNew *core_v1.Secret) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *SecretEventHandlerFuncs) GenericSecret(obj *core_v1.Secret) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type SecretEventWatcher interface {
	AddEventHandler(ctx context.Context, h SecretEventHandler, predicates ...predicate.Predicate) error
}

type secretEventWatcher struct {
	watcher events.EventWatcher
}

func NewSecretEventWatcher(name string, mgr manager.Manager) SecretEventWatcher {
	return &secretEventWatcher{
		watcher: events.NewWatcher(name, mgr, &core_v1.Secret{}),
	}
}

func (c *secretEventWatcher) AddEventHandler(ctx context.Context, h SecretEventHandler, predicates ...predicate.Predicate) error {
	handler := genericSecretHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericSecretHandler implements a generic events.EventHandler
type genericSecretHandler struct {
	handler SecretEventHandler
}

func (h genericSecretHandler) Create(object runtime.Object) error {
	obj, ok := object.(*core_v1.Secret)
	if !ok {
		return errors.Errorf("internal error: Secret handler received event for %T", object)
	}
	return h.handler.CreateSecret(obj)
}

func (h genericSecretHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*core_v1.Secret)
	if !ok {
		return errors.Errorf("internal error: Secret handler received event for %T", object)
	}
	return h.handler.DeleteSecret(obj)
}

func (h genericSecretHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*core_v1.Secret)
	if !ok {
		return errors.Errorf("internal error: Secret handler received event for %T", old)
	}
	objNew, ok := new.(*core_v1.Secret)
	if !ok {
		return errors.Errorf("internal error: Secret handler received event for %T", new)
	}
	return h.handler.UpdateSecret(objOld, objNew)
}

func (h genericSecretHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*core_v1.Secret)
	if !ok {
		return errors.Errorf("internal error: Secret handler received event for %T", object)
	}
	return h.handler.GenericSecret(obj)
}

// Handle events for the ConfigMap Resource
// DEPRECATED: Prefer reconciler pattern.
type ConfigMapEventHandler interface {
	CreateConfigMap(obj *core_v1.ConfigMap) error
	UpdateConfigMap(old, new *core_v1.ConfigMap) error
	DeleteConfigMap(obj *core_v1.ConfigMap) error
	GenericConfigMap(obj *core_v1.ConfigMap) error
}

type ConfigMapEventHandlerFuncs struct {
	OnCreate  func(obj *core_v1.ConfigMap) error
	OnUpdate  func(old, new *core_v1.ConfigMap) error
	OnDelete  func(obj *core_v1.ConfigMap) error
	OnGeneric func(obj *core_v1.ConfigMap) error
}

func (f *ConfigMapEventHandlerFuncs) CreateConfigMap(obj *core_v1.ConfigMap) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *ConfigMapEventHandlerFuncs) DeleteConfigMap(obj *core_v1.ConfigMap) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *ConfigMapEventHandlerFuncs) UpdateConfigMap(objOld, objNew *core_v1.ConfigMap) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *ConfigMapEventHandlerFuncs) GenericConfigMap(obj *core_v1.ConfigMap) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type ConfigMapEventWatcher interface {
	AddEventHandler(ctx context.Context, h ConfigMapEventHandler, predicates ...predicate.Predicate) error
}

type configMapEventWatcher struct {
	watcher events.EventWatcher
}

func NewConfigMapEventWatcher(name string, mgr manager.Manager) ConfigMapEventWatcher {
	return &configMapEventWatcher{
		watcher: events.NewWatcher(name, mgr, &core_v1.ConfigMap{}),
	}
}

func (c *configMapEventWatcher) AddEventHandler(ctx context.Context, h ConfigMapEventHandler, predicates ...predicate.Predicate) error {
	handler := genericConfigMapHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericConfigMapHandler implements a generic events.EventHandler
type genericConfigMapHandler struct {
	handler ConfigMapEventHandler
}

func (h genericConfigMapHandler) Create(object runtime.Object) error {
	obj, ok := object.(*core_v1.ConfigMap)
	if !ok {
		return errors.Errorf("internal error: ConfigMap handler received event for %T", object)
	}
	return h.handler.CreateConfigMap(obj)
}

func (h genericConfigMapHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*core_v1.ConfigMap)
	if !ok {
		return errors.Errorf("internal error: ConfigMap handler received event for %T", object)
	}
	return h.handler.DeleteConfigMap(obj)
}

func (h genericConfigMapHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*core_v1.ConfigMap)
	if !ok {
		return errors.Errorf("internal error: ConfigMap handler received event for %T", old)
	}
	objNew, ok := new.(*core_v1.ConfigMap)
	if !ok {
		return errors.Errorf("internal error: ConfigMap handler received event for %T", new)
	}
	return h.handler.UpdateConfigMap(objOld, objNew)
}

func (h genericConfigMapHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*core_v1.ConfigMap)
	if !ok {
		return errors.Errorf("internal error: ConfigMap handler received event for %T", object)
	}
	return h.handler.GenericConfigMap(obj)
}
