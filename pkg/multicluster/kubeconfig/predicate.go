package kubeconfig

import (
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

var Predicate = predicate.Funcs{
	CreateFunc: func(e event.CreateEvent) bool {
		return isKubeConfigSecret(e.Object)
	},
	DeleteFunc: func(e event.DeleteEvent) bool {
		return isKubeConfigSecret(e.Object)
	},
	UpdateFunc: func(e event.UpdateEvent) bool {
		if isKubeConfigSecret(e.ObjectNew) {
			// Ignore metadata changes.
			oldSecret, newSecret := e.ObjectOld.(*corev1.Secret), e.ObjectNew.(*corev1.Secret)
			return !reflect.DeepEqual(oldSecret.Data, newSecret.Data)
		}
		return false
	},
	GenericFunc: func(e event.GenericEvent) bool {
		return isKubeConfigSecret(e.Object)
	},
}

func isKubeConfigSecret(obj client.Object) bool {
	if s, ok := obj.(*corev1.Secret); ok {
		return s.Type == SecretType
	}
	return false
}
