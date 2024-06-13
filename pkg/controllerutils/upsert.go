package controllerutils

import (
	"context"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"
)

// TransitionFunc performs a comparison of the the existing object with the desired object before a desired object is Upserted to kube storage.
type TransitionFunc func(existing, desired runtime.Object) error

// Upsert a desired object to the cluster.
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
	transitionFuncs ...TransitionFunc,
) (controllerutil.OperationResult, error) {
	if _, ok := obj.GetLabels()["ssa"]; ok {
		yml, _ := serializeToYAML(obj)
		fmt.Printf("UUU called upsert. serialized yaml is %s\n", yml)
	}

	return upsert(ctx, c, obj, transitionFuncs...)
}

// UpsertImmutable functions similarly to it's Upsert counterpart,
// but it will copy obj before saving it.
func UpsertImmutable(
	ctx context.Context,
	c client.Client,
	obj client.Object,
	transitionFuncs ...TransitionFunc,
) (controllerutil.OperationResult, error) {
	return upsert(ctx, c, obj.DeepCopyObject().(client.Object), transitionFuncs...)
}

func upsert(
	ctx context.Context,
	c client.Client,
	obj client.Object,
	transitionFuncs ...TransitionFunc,
) (controllerutil.OperationResult, error) {

	gvk, err := apiutil.GVKForObject(obj, c.Scheme())
	if err != nil {
		return controllerutil.OperationResultNone, err
	}
	newObj, err := c.Scheme().New(gvk)
	if err != nil {
		return controllerutil.OperationResultNone, err
	}
	// Always valid because obj is client.Object
	existing := newObj.(client.Object)
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
		if _, ok := obj.GetLabels()["ssa"]; ok {
			fmt.Printf("UUU objects are equal\n")
		}

		return controllerutil.OperationResultNone, nil
	}

	if _, ok := obj.GetLabels()["ssa"]; ok {
		yml, _ := serializeToYAML(obj)
		fmt.Printf("UUU updating. serialized yaml is %s\n", yml)
	}

	err = c.Patch(ctx, obj, client.Apply, client.ForceOwnership, client.FieldOwner("foo"))
	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	return controllerutil.OperationResultUpdated, nil
}

func serializeToYAML(obj client.Object) (string, error) {
	// Use reflection to get the underlying value of the interface
	value := reflect.ValueOf(obj)

	// Ensure the value is not nil and is a pointer
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return "", fmt.Errorf("expected a non-nil pointer to a client.Object")
	}

	// Convert the value to a YAML byte slice
	yamlData, err := yaml.Marshal(value.Interface())
	if err != nil {
		return "", fmt.Errorf("failed to serialize object to YAML: %v", err)
	}

	return string(yamlData), nil
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
	existing, updateNeeded, err := needUpdate(ctx, c, obj)
	if err != nil || !updateNeeded {
		return controllerutil.OperationResultNone, err
	}

	// https://github.com/solo-io/skv2/issues/344
	obj.SetUID(existing.GetUID())
	obj.SetCreationTimestamp(existing.GetCreationTimestamp())
	obj.SetResourceVersion(existing.GetResourceVersion())

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
	existing, updateNeeded, err := needUpdate(ctx, c, obj)
	if err != nil || !updateNeeded {
		return controllerutil.OperationResultNone, err
	}

	// Always valid because obj is client.Object
	copyOfObj := obj.DeepCopyObject().(client.Object)

	// https://github.com/solo-io/skv2/issues/344
	copyOfObj.SetUID(existing.GetUID())
	copyOfObj.SetCreationTimestamp(existing.GetCreationTimestamp())
	copyOfObj.SetResourceVersion(existing.GetResourceVersion())

	return update(ctx, c, copyOfObj)
}

func needUpdate(ctx context.Context, c client.Client, obj client.Object) (client.Object, bool, error) {
	key := client.ObjectKeyFromObject(obj)

	// create empty object of the same type so that Get will work
	existing := reflect.New(reflect.TypeOf(obj).Elem()).Interface().(client.Object)
	if err := c.Get(ctx, key, existing); err != nil {
		return nil, false, err
	}

	if ObjectStatusesEqual(existing, obj) {
		return nil, false, nil
	}
	return existing, true, nil
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
