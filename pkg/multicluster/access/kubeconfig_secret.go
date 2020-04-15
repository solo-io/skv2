package access

import (
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// KubeConfigSecretType is used to indicate which kubernetes secrets contain a KubeConfig.
const KubeConfigSecretType = "solo.io/kubeconfig"

// KubeConfigToSecret converts a KubeConfig to a secret with the provided name and namespace.
func KubeConfigToSecret(name string, namespace string, cluster string, kc *api.Config) (*kubev1.Secret, error) {
	rawKubeConfig, err := clientcmd.Write(kc)
	if err != nil {
		return nil, FailedToConvertKubeConfigToSecret(err)
	}

	return &kubev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Type: KubeConfigSecretType,
		Data: map[string][]byte{cluster: rawKubeConfig},
	}, nil
}

// Config contains various KubeConfig formats for convenience.
type Config struct {
	ClientConfig clientcmd.ClientConfig
	ApiConfig    *api.Config
	RestConfig   *rest.Config
}

// SecretToKubeConfig extracts the cluster name and *Config from a KubeConfig secret.
// If the provided secret is not a KubeConfig secret, an error is returned.
func SecretToConfig(secret *kubev1.Secret) (clusterName string, config *Config, err error) {
	if len(secret.Data) > 1 {
		return "", nil, SecretHasMultipleKeys(secret.ObjectMeta)
	}
	for clusterName, dataEntry := range secret.Data {
		clientConfig, err := clientcmd.NewClientConfigFromBytes(dataEntry)
		if err != nil {
			return clusterName, nil, FailedToConvertSecretToClientConfig(err)
		}

		apiConfig, err := clientcmd.Load(dataEntry)
		if err != nil {
			return clusterName, nil, FailedToConvertSecretToKubeConfig(err)
		}

		restConfig, err := clientConfig.ClientConfig()
		if err != nil {
			return clusterName, nil, FailedToConvertSecretToRestConfig(err)
		}
		return clusterName, &Config{
			ClientConfig: clientConfig,
			RestConfig:   restConfig,
			ApiConfig:    apiConfig,
		}, nil
	}

	return "", nil, NoDataInKubeConfigSecret(secret.ObjectMeta)
}
