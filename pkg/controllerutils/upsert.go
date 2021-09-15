package controllerutils

import (
	"context"
	"reflect"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// TransitionFunc performs a comparison of the the existing object with the desired object before a desired object is Upserted to kube storage.
type TransitionFunc func(existing, desired runtime.Object) error

// Easily Upsert a desired object to the cluster.
//
// If the object exists, the provided TransitionFuncs will be called.
// Resource version for the desired object will be set automatically.
//
// If the desired object (after applying transition funcs) is semantically equal
// to the existing object, the update is skipped.
func Upsert(
	ctx context.Context,
	c client.Client,
	obj client.Object,
	// objCreator runtime.ObjectCreater,
	transitionFuncs ...TransitionFunc,
) (controllerutil.OperationResult, error) {
	// Always valid because obj is client.Object
	existing := obj.DeepCopyObject().(client.Object)
	return upsert(ctx, c, obj, existing, transitionFuncs...)
}

func UpsertNoCopy(
	ctx context.Context,
	c client.Client,
	obj client.Object,
	scheme *runtime.Scheme,
	transitionFuncs ...TransitionFunc,
) (controllerutil.OperationResult, error) {
	gvk, err := apiutil.GVKForObject(obj, scheme)
	if err != nil {
		return controllerutil.OperationResultNone, err
	}
	newObj, err := scheme.New(gvk)
	if err != nil {
		return controllerutil.OperationResultNone, err
	}
	// Always valid because obj is client.Object
	existing := newObj.(client.Object)

	return upsert(ctx, c, obj, existing, transitionFuncs...)
}

func transition(existing, desired runtime.Object, transitionFuncs []TransitionFunc) error {
	for _, txFunc := range transitionFuncs {
		if err := txFunc(existing, desired); err != nil {
			return err
		}
	}

	if existingMeta, ok := existing.(metav1.Object); ok {
		desired.(metav1.Object).SetResourceVersion(existingMeta.GetResourceVersion())
	}

	return nil
}

func upsert(
	ctx context.Context,
	c client.Client,
	obj, existing client.Object,
	transitionFuncs ...TransitionFunc,
) (controllerutil.OperationResult, error) {

	key := client.ObjectKeyFromObject(obj)
	if err := c.Get(ctx, key, existing); err != nil {
		if !errors.IsNotFound(err) {
			return controllerutil.OperationResultNone, err
		}
		if err := c.Create(ctx, obj); err != nil {
			return controllerutil.OperationResultNone, err
		}
		return controllerutil.OperationResultCreated, nil
	}

	if err := transition(existing, obj, transitionFuncs); err != nil {
		return controllerutil.OperationResultNone, err
	}

	if ObjectsEqual(existing, obj) {
		return controllerutil.OperationResultNone, nil
	}

	if err := c.Update(ctx, obj); err != nil {
		return controllerutil.OperationResultNone, err
	}
	return controllerutil.OperationResultUpdated, nil
}

// Easily Update the Status of a desired object in the cluster.
// Requires that the object already exists (will only attempt to update).
//
// If the desired object status is semantically equal
// to the existing object status, the update is skipped.
func UpdateStatus(
	ctx context.Context,
	c client.Client,
	obj client.Object,
) (controllerutil.OperationResult, error) {
	updateNeeded, err := needUpdate(ctx, c, obj)
	if err != nil || !updateNeeded {
		return controllerutil.OperationResultNone, err
	}

	return update(ctx, c, obj)
}

// Easily Update the Status of a desired object in the cluster.
// Requires that the object already exists (will only attempt to update).
//
// If the desired object status is semantically equal
// to the existing object status, the update is skipped.
// Unlike the method above, this method does not update obj.
func UpdateStatusImmutable(
	ctx context.Context,
	c client.Client,
	obj client.Object,
) (controllerutil.OperationResult, error) {
	updateNeeded, err := needUpdate(ctx, c, obj)
	if err != nil || !updateNeeded {
		return controllerutil.OperationResultNone, err
	}

	// Always valid because obj is client.Object
	copyOfObj := obj.DeepCopyObject().(client.Object)
	return update(ctx, c, copyOfObj)
}

func needUpdate(ctx context.Context, c client.Client, obj client.Object) (bool, error) {
	key := client.ObjectKeyFromObject(obj)

	// create empty object of the same type so that Get will work
	existing := reflect.New(reflect.TypeOf(obj).Elem()).Interface().(client.Object)
	if err := c.Get(ctx, key, existing); err != nil {
		return false, err
	}

	if ObjectStatusesEqual(existing, obj) {
		return false, nil
	}
	return true, nil
}

func update(
	ctx context.Context,
	c client.Client,
	obj client.Object,
) (controllerutil.OperationResult, error) {
	if err := c.Status().Update(ctx, obj); err != nil {
		return controllerutil.OperationResultNone, err
	}
	return controllerutil.OperationResultUpdated, nil
}
