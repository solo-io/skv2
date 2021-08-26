package snapshot

import (
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RedactSecretData returns a copy with sensitive information redacted
func RedactSecretData(obj client.Object) client.Object {
	if sec, ok := obj.(*v1.Secret); ok {
		redacted := sec.DeepCopyObject().(*v1.Secret)
		for k := range redacted.Data {
			redacted.Data[k] = []byte("*****")
		}
		return redacted
	}
	return obj
}
