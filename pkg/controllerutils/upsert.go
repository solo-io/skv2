package controllerutils

import (
	"context"

	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
func Upsert(ctx context.Context, c client.Client, obj runtime.Object, transitionFuncs ...TransitionFunc) (controllerutil.OperationResult, error) {
	key, err := client.ObjectKeyFromObject(obj)
	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	existing := obj.DeepCopyObject()

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

	if equality.Semantic.DeepEqual(existing, obj) {
		return controllerutil.OperationResultNone, nil
	}

	if err := c.Update(ctx, obj); err != nil {
		return controllerutil.OperationResultNone, err
	}
	return controllerutil.OperationResultUpdated, nil
}

func transition(existing, desired runtime.Object, transitionFuncs []TransitionFunc) error {
	if existingMeta, ok := existing.(metav1.Object); ok {
		desired.(metav1.Object).SetResourceVersion(existingMeta.GetResourceVersion())
	}
	for _, txFunc := range transitionFuncs {
		if err := txFunc(existing, desired); err != nil {
			return err
		}
	}
	return nil
}
