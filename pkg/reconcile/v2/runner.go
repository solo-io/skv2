package reconcile_v2

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/verifier"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/utils"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type Request = reconcile.Request
type Result = reconcile.Result

type Reconciler[T client.Object] interface {
	// reconcile an object
	// requeue the object if returning an error, or a non-zero "requeue-after" duration
	Reconcile(ctx context.Context, object T) (Result, error)
}

type DeletionReconciler interface {
	// we received a reconcile request for an object that was removed from the cache
	// requeue the object if returning an error
	ReconcileDeletion(ctx context.Context, request Request) error
}

type ReconcileFuncs[T client.Object] struct {
	ReconcileFunc          func(ctx context.Context, object T) (Result, error)
	DeletionReconcilerFunc func(ctx context.Context, request Request) error
}

func (r *ReconcileFuncs[T]) Reconcile(ctx context.Context, object T) (Result, error) {
	if r.ReconcileFunc == nil {
		return Result{}, nil
	}
	return r.ReconcileFunc(ctx, object)
}

func (r *ReconcileFuncs[T]) DeletionReconciler(ctx context.Context, request Request) error {
	if r.DeletionReconcilerFunc == nil {
		return nil
	}
	return r.DeletionReconcilerFunc(ctx, request)
}

type FinalizingReconciler[T client.Object] interface {
	Reconciler[T]

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	FinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	Finalize(ctx context.Context, object T) error
}

// a Reconcile Loop runs resource reconcilers until the context gets cancelled
type Loop[T client.Object] interface {
	RunReconciler(ctx context.Context, reconciler Reconciler[T], predicates ...predicate.Predicate) error
}

type runner[T client.Object] struct {
	name    string
	mgr     manager.Manager
	options Options
	cluster string
	t       T
}

type Options struct {
	// If true will wait for cache sync before returning from RunReconcile
	WaitForCacheSync bool

	// If provided, attempt to verify the resource before beginning the reconcile loop
	Verifier verifier.ServerResourceVerifier
}

func NewLoop[T client.Object](
	name, cluster string,
	mgr manager.Manager,
	t T,
	options Options,
) *runner[T] {
	return &runner[T]{
		name:    name,
		mgr:     mgr,
		cluster: cluster,
		options: options,
		t:       t,
	}
}

type runnerReconciler[T client.Object] struct {
	mgr        manager.Manager
	reconciler Reconciler[T]
	t          T
}

func (r *runner[T]) RunReconciler(
	ctx context.Context,
	reconciler Reconciler[T],
	predicates ...predicate.Predicate,
) error {
	obj := r.t.DeepCopyObject().(T)
	gvk, err := apiutil.GVKForObject(obj, r.mgr.GetScheme())
	if err != nil {
		return err
	}

	if r.options.Verifier != nil {
		// Always local cluster
		if resourceRegistered, err := r.options.Verifier.VerifyServerResource("", r.mgr.GetConfig(), gvk); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	rec := &runnerReconciler[T]{
		mgr:        r.mgr,
		reconciler: reconciler,
		t:          obj,
	}

	ctl, err := controller.New(
		r.name, r.mgr, controller.Options{
			Reconciler: rec,
		},
	)
	if err != nil {
		return err
	}

	// send us watch events
	if err := ctl.Watch(&source.Kind{Type: obj}, &handler.EnqueueRequestForObject{}, predicates...); err != nil {
		return err
	}

	// Only wait for cache sync if specified in options
	if r.options.WaitForCacheSync {
		contextutils.LoggerFrom(ctx).Debug("waiting for cache sync...")
		if synced := r.mgr.GetCache().WaitForCacheSync(ctx); !synced {
			return errors.Errorf("waiting for cache sync failed")
		}
	}

	return nil
}

func (r *runnerReconciler[T]) Reconcile(ctx context.Context, request Request) (reconcile.Result, error) {
	contextutils.LoggerFrom(ctx).Debug("handling event", "event", request)

	// get the object from our cache
	restClient := ezkube.NewRestClient(r.mgr)

	obj := r.t.DeepCopyObject().(T)
	obj.SetName(request.Name)
	obj.SetNamespace(request.Namespace)
	if err := restClient.Get(ctx, obj); err != nil {
		if err := client.IgnoreNotFound(err); err != nil {
			return reconcile.Result{}, err
		}
		contextutils.LoggerFrom(ctx).Debug(fmt.Sprintf("unable to fetch %T", obj), "request", request, "err", err)

		// call OnDelete
		if deletionReconciler, ok := r.reconciler.(DeletionReconciler); ok {
			return reconcile.Result{}, deletionReconciler.ReconcileDeletion(ctx, request)
		}

		return reconcile.Result{}, nil
	}

	// if the handler is a finalizer, check if we need to finalize
	if finalizer, ok := r.reconciler.(FinalizingReconciler[T]); ok {
		finalizers := obj.GetFinalizers()
		finalizerName := finalizer.FinalizerName()
		if obj.GetDeletionTimestamp().IsZero() {
			// The object is not being deleted, so if it does not have our finalizer,
			// then lets add the finalizer and update the object. This is equivalent to
			// registering our finalizer.

			if !utils.ContainsString(finalizers, finalizerName) {
				obj.SetFinalizers(
					append(
						finalizers,
						finalizerName,
					),
				)
				if err := restClient.Update(ctx, obj); err != nil {
					return reconcile.Result{}, err
				}
			}
		} else {

			// The object is being deleted
			if utils.ContainsString(finalizers, finalizerName) {
				// our finalizer is present, so lets handle any external dependency
				if err := finalizer.Finalize(ctx, obj); err != nil {
					// if fail to delete the external dependency here, return with error
					// so that it can be retried
					return reconcile.Result{}, err
				}

				// remove our finalizer from the list and update it.
				obj.SetFinalizers(utils.RemoveString(finalizers, finalizerName))
				if err := restClient.Update(ctx, obj); err != nil {
					return reconcile.Result{}, err
				}
			}

			// We have already finalized the object, so return early to skip reconcile.
			return reconcile.Result{}, nil
		}
	}

	result, err := r.reconciler.Reconcile(ctx, obj)
	if err != nil {
		return result, eris.Wrap(err, "handler error. retrying")
	}
	contextutils.LoggerFrom(ctx).Debug("handler success.", "result", result)

	return result, nil
}
