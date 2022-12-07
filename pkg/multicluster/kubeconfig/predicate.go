package kubeconfig

import (
	"reflect"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

func BuildPredicate(watchNamespaces []string) predicate.Funcs {
	return predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			if !isNamespaceWatched(watchNamespaces, e.Object.GetNamespace()) {
				return false
			}
			return isKubeConfigSecret(e.Object)
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			if !isNamespaceWatched(watchNamespaces, e.Object.GetNamespace()) {
				return false
			}
			return isKubeConfigSecret(e.Object)
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			if !isNamespaceWatched(watchNamespaces, e.ObjectNew.GetNamespace()) {
				return false
			}
			if isKubeConfigSecret(e.ObjectNew) {
				// Ignore metadata changes.
				oldSecret, newSecret := e.ObjectOld.(*corev1.Secret), e.ObjectNew.(*corev1.Secret)
				return !reflect.DeepEqual(oldSecret.Data, newSecret.Data)
			}
			return false
		},
		GenericFunc: func(e event.GenericEvent) bool {
			if !isNamespaceWatched(watchNamespaces, e.Object.GetNamespace()) {
				return false
			}
			return isKubeConfigSecret(e.Object)
		},
	}
}

func isKubeConfigSecret(obj client.Object) bool {
	if s, ok := obj.(*corev1.Secret); ok {
		return s.Type == SecretType
	}
	return false
}

// If watchNamespaces is empty, then watch all namespaces. Otherwise, watch
// only the specified namespaces.
func isNamespaceWatched(watchNamespaces []string, namespace string) bool {
	if len(watchNamespaces) == 0 {
		return true
	}

	for _, ns := range watchNamespaces {
		if ns == namespace {
			return true
		}
	}
	return false
}
