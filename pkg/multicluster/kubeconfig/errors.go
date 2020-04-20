package kubeconfig

import (
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	/*
		KubeConfig Secret errors.
	*/

	FailedToConvertKubeConfigToSecret = func(err error) error {
		return eris.Wrap(err, "Could not serialize KubeConfig to yaml while generating secret.")
	}
	FailedToConvertSecretToKubeConfig = func(err error) error {
		return eris.Wrapf(err, "Could not deserialize string to KubeConfig while generating KubeConfig")
	}
	FailedToConvertSecretToClientConfig = func(err error) error {
		return eris.Wrap(err, "Could not convert config to ClientConfig")
	}
	SecretHasNoKubeConfig = func(meta metav1.ObjectMeta) error {
		return eris.Errorf("kube config secret %s.%s has no KubeConfig value for key %v",
			meta.Namespace, meta.Name, Key)
	}
	FailedToConvertSecretToRestConfig = func(err error) error {
		return eris.Wrap(err, "Could not convert config to *rest.Config")
	}
)
