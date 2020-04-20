package kubeconfig

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

var _ predicate.Predicate = SecretPredicate{}

// SecretPredicate is a controller-runtime predicate that filters for KubeConfig secrets.
type SecretPredicate struct{}

func (p SecretPredicate) Create(e event.CreateEvent) bool {
	return isKubeConfigSecret(e.Object)
}

func (p SecretPredicate) Delete(e event.DeleteEvent) bool {
	return isKubeConfigSecret(e.Object)
}

func (p SecretPredicate) Update(e event.UpdateEvent) bool {
	return isKubeConfigSecret(e.ObjectNew)
}

func (p SecretPredicate) Generic(e event.GenericEvent) bool {
	return isKubeConfigSecret(e.Object)
}

func isKubeConfigSecret(obj runtime.Object) bool {
	if s, ok := obj.(*corev1.Secret); ok {
		return s.Type == SecretType
	}
	return false
}
