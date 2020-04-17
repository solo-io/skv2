// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	core_v1 "k8s.io/api/core/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Secret Resource.
// implemented by the user
type SecretReconciler interface {
	ReconcileSecret(obj *core_v1.Secret) (reconcile.Result, error)
}

// Reconcile deletion events for the Secret Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type SecretDeletionReconciler interface {
	ReconcileSecretDeletion(req reconcile.Request)
}

type SecretReconcilerFuncs struct {
	OnReconcileSecret         func(obj *core_v1.Secret) (reconcile.Result, error)
	OnReconcileSecretDeletion func(req reconcile.Request)
}

func (f *SecretReconcilerFuncs) ReconcileSecret(obj *core_v1.Secret) (reconcile.Result, error) {
	if f.OnReconcileSecret == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSecret(obj)
}

func (f *SecretReconcilerFuncs) ReconcileSecretDeletion(req reconcile.Request) {
	if f.OnReconcileSecretDeletion == nil {
		return
	}
	f.OnReconcileSecretDeletion(req)
}

// Reconcile and finalize the Secret Resource
// implemented by the user
type SecretFinalizer interface {
	SecretReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	SecretFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeSecret(obj *core_v1.Secret) error
}

type SecretReconcileLoop interface {
	RunSecretReconciler(ctx context.Context, rec SecretReconciler, predicates ...predicate.Predicate) error
}

type secretReconcileLoop struct {
	loop reconcile.Loop
}

func NewSecretReconcileLoop(name string, mgr manager.Manager) SecretReconcileLoop {
	return &secretReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &core_v1.Secret{}),
	}
}

func (c *secretReconcileLoop) RunSecretReconciler(ctx context.Context, reconciler SecretReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericSecretReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(SecretFinalizer); ok {
		reconcilerWrapper = genericSecretFinalizer{
			genericSecretReconciler: genericReconciler,
			finalizingReconciler:    finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericSecretHandler implements a generic reconcile.Reconciler
type genericSecretReconciler struct {
	reconciler SecretReconciler
}

func (r genericSecretReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.Secret)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Secret handler received event for %T", object)
	}
	return r.reconciler.ReconcileSecret(obj)
}

func (r genericSecretReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(SecretDeletionReconciler); ok {
		deletionReconciler.ReconcileSecretDeletion(request)
	}
}

// genericSecretFinalizer implements a generic reconcile.FinalizingReconciler
type genericSecretFinalizer struct {
	genericSecretReconciler
	finalizingReconciler SecretFinalizer
}

func (r genericSecretFinalizer) FinalizerName() string {
	return r.finalizingReconciler.SecretFinalizerName()
}

func (r genericSecretFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*core_v1.Secret)
	if !ok {
		return errors.Errorf("internal error: Secret handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeSecret(obj)
}

// Reconcile Upsert events for the ConfigMap Resource.
// implemented by the user
type ConfigMapReconciler interface {
	ReconcileConfigMap(obj *core_v1.ConfigMap) (reconcile.Result, error)
}

// Reconcile deletion events for the ConfigMap Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ConfigMapDeletionReconciler interface {
	ReconcileConfigMapDeletion(req reconcile.Request)
}

type ConfigMapReconcilerFuncs struct {
	OnReconcileConfigMap         func(obj *core_v1.ConfigMap) (reconcile.Result, error)
	OnReconcileConfigMapDeletion func(req reconcile.Request)
}

func (f *ConfigMapReconcilerFuncs) ReconcileConfigMap(obj *core_v1.ConfigMap) (reconcile.Result, error) {
	if f.OnReconcileConfigMap == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileConfigMap(obj)
}

func (f *ConfigMapReconcilerFuncs) ReconcileConfigMapDeletion(req reconcile.Request) {
	if f.OnReconcileConfigMapDeletion == nil {
		return
	}
	f.OnReconcileConfigMapDeletion(req)
}

// Reconcile and finalize the ConfigMap Resource
// implemented by the user
type ConfigMapFinalizer interface {
	ConfigMapReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ConfigMapFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeConfigMap(obj *core_v1.ConfigMap) error
}

type ConfigMapReconcileLoop interface {
	RunConfigMapReconciler(ctx context.Context, rec ConfigMapReconciler, predicates ...predicate.Predicate) error
}

type configMapReconcileLoop struct {
	loop reconcile.Loop
}

func NewConfigMapReconcileLoop(name string, mgr manager.Manager) ConfigMapReconcileLoop {
	return &configMapReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &core_v1.ConfigMap{}),
	}
}

func (c *configMapReconcileLoop) RunConfigMapReconciler(ctx context.Context, reconciler ConfigMapReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericConfigMapReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ConfigMapFinalizer); ok {
		reconcilerWrapper = genericConfigMapFinalizer{
			genericConfigMapReconciler: genericReconciler,
			finalizingReconciler:       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericConfigMapHandler implements a generic reconcile.Reconciler
type genericConfigMapReconciler struct {
	reconciler ConfigMapReconciler
}

func (r genericConfigMapReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*core_v1.ConfigMap)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ConfigMap handler received event for %T", object)
	}
	return r.reconciler.ReconcileConfigMap(obj)
}

func (r genericConfigMapReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(ConfigMapDeletionReconciler); ok {
		deletionReconciler.ReconcileConfigMapDeletion(request)
	}
}

// genericConfigMapFinalizer implements a generic reconcile.FinalizingReconciler
type genericConfigMapFinalizer struct {
	genericConfigMapReconciler
	finalizingReconciler ConfigMapFinalizer
}

func (r genericConfigMapFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ConfigMapFinalizerName()
}

func (r genericConfigMapFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*core_v1.ConfigMap)
	if !ok {
		return errors.Errorf("internal error: ConfigMap handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeConfigMap(obj)
}
