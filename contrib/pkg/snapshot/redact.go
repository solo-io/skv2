package snapshot

import (
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	redactedString              = "redacted"
)

// RedactSecretData returns a copy with sensitive information redacted
func RedactSecretData(obj client.Object) client.Object {
	if sec, ok := obj.(*v1.Secret); ok {
		redacted := sec.DeepCopyObject().(*v1.Secret)
		for k := range redacted.Data {
			redacted.Data[k] = []byte(redactedString)
		}

		// Also need to check for kubectl apply, last applied config.
		// Secret data can be found there as well if that's how the secret is created
		for key, _ := range redacted.Annotations {
			if key == v1.LastAppliedConfigAnnotation {
				redacted.Annotations[key] = redactedString
			}
		}
		return redacted
	}
	return obj
}
