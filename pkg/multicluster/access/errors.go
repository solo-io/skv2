package registration

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
	DuplicateClusterName = func(repeatedClusterName string) error {
		return eris.Errorf("Error converting KubeConfigs to secret, duplicate cluster name found: %s", repeatedClusterName)
	}
	FailedToConvertSecretToKubeConfig = func(err error) error {
		return eris.Wrapf(err, "Could not deserialize string to KubeConfig while generating KubeConfig")
	}
	NoDataInKubeConfigSecret = func(meta metav1.ObjectMeta) error {
		return eris.Errorf("No data in kube config secret %s.%s", meta.Namespace, meta.Name)
	}
	FailedToConvertSecretToClientConfig = func(err error) error {
		return eris.Wrap(err, "Could not convert config to ClientConfig")
	}
	SecretHasMultipleKeys = func(meta metav1.ObjectMeta) error {
		return eris.Errorf("kube config secret %s.%s has multiple keys in its data, this is unexpected",
			meta.Namespace, meta.Name)
	}
	FailedToConvertSecretToRestConfig = func(err error) error {
		return eris.Wrap(err, "Could not convert config to *rest.Config")
	}
)
