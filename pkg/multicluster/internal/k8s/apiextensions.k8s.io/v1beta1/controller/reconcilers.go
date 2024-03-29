// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	apiextensions_k8s_io_v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CustomResourceDefinition Resource.
// implemented by the user
type CustomResourceDefinitionReconciler interface {
	ReconcileCustomResourceDefinition(obj *apiextensions_k8s_io_v1beta1.CustomResourceDefinition) (reconcile.Result, error)
}

// Reconcile deletion events for the CustomResourceDefinition Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CustomResourceDefinitionDeletionReconciler interface {
	ReconcileCustomResourceDefinitionDeletion(req reconcile.Request) error
}

type CustomResourceDefinitionReconcilerFuncs struct {
	OnReconcileCustomResourceDefinition         func(obj *apiextensions_k8s_io_v1beta1.CustomResourceDefinition) (reconcile.Result, error)
	OnReconcileCustomResourceDefinitionDeletion func(req reconcile.Request) error
}

func (f *CustomResourceDefinitionReconcilerFuncs) ReconcileCustomResourceDefinition(obj *apiextensions_k8s_io_v1beta1.CustomResourceDefinition) (reconcile.Result, error) {
	if f.OnReconcileCustomResourceDefinition == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCustomResourceDefinition(obj)
}

func (f *CustomResourceDefinitionReconcilerFuncs) ReconcileCustomResourceDefinitionDeletion(req reconcile.Request) error {
	if f.OnReconcileCustomResourceDefinitionDeletion == nil {
		return nil
	}
	return f.OnReconcileCustomResourceDefinitionDeletion(req)
}

// Reconcile and finalize the CustomResourceDefinition Resource
// implemented by the user
type CustomResourceDefinitionFinalizer interface {
	CustomResourceDefinitionReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CustomResourceDefinitionFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCustomResourceDefinition(obj *apiextensions_k8s_io_v1beta1.CustomResourceDefinition) error
}

type CustomResourceDefinitionReconcileLoop interface {
	RunCustomResourceDefinitionReconciler(ctx context.Context, rec CustomResourceDefinitionReconciler, predicates ...predicate.Predicate) error
}

type customResourceDefinitionReconcileLoop struct {
	loop reconcile.Loop
}

func NewCustomResourceDefinitionReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CustomResourceDefinitionReconcileLoop {
	return &customResourceDefinitionReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &apiextensions_k8s_io_v1beta1.CustomResourceDefinition{}, options),
	}
}

func (c *customResourceDefinitionReconcileLoop) RunCustomResourceDefinitionReconciler(ctx context.Context, reconciler CustomResourceDefinitionReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCustomResourceDefinitionReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CustomResourceDefinitionFinalizer); ok {
		reconcilerWrapper = genericCustomResourceDefinitionFinalizer{
			genericCustomResourceDefinitionReconciler: genericReconciler,
			finalizingReconciler:                      finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCustomResourceDefinitionHandler implements a generic reconcile.Reconciler
type genericCustomResourceDefinitionReconciler struct {
	reconciler CustomResourceDefinitionReconciler
}

func (r genericCustomResourceDefinitionReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CustomResourceDefinition handler received event for %T", object)
	}
	return r.reconciler.ReconcileCustomResourceDefinition(obj)
}

func (r genericCustomResourceDefinitionReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CustomResourceDefinitionDeletionReconciler); ok {
		return deletionReconciler.ReconcileCustomResourceDefinitionDeletion(request)
	}
	return nil
}

// genericCustomResourceDefinitionFinalizer implements a generic reconcile.FinalizingReconciler
type genericCustomResourceDefinitionFinalizer struct {
	genericCustomResourceDefinitionReconciler
	finalizingReconciler CustomResourceDefinitionFinalizer
}

func (r genericCustomResourceDefinitionFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CustomResourceDefinitionFinalizerName()
}

func (r genericCustomResourceDefinitionFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*apiextensions_k8s_io_v1beta1.CustomResourceDefinition)
	if !ok {
		return errors.Errorf("internal error: CustomResourceDefinition handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCustomResourceDefinition(obj)
}
