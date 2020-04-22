package kubeconfig

import (
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

const (
	// SecretType is used to indicate which kubernetes secrets contain a KubeConfig.
	SecretType = "solo.io/kubeconfig"

	// Key is the KubeConfig's key in the KubeConfig secret data.
	Key = "kubeconfig"
)

// ToSecret converts a kubernetes api.Config to a secret with the provided name and namespace.
func ToSecret(namespace string, cluster string, kc api.Config) (*kubev1.Secret, error) {
	rawKubeConfig, err := clientcmd.Write(kc)
	if err != nil {
		return nil, FailedToConvertKubeConfigToSecret(err)
	}

	return &kubev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster,
			Namespace: namespace,
		},
		Type: SecretType,
		Data: map[string][]byte{Key: rawKubeConfig},
	}, nil
}

// Config contains various KubeConfig formats for convenience.
type Config struct {
	ClientConfig clientcmd.ClientConfig
	ApiConfig    *api.Config
	RestConfig   *rest.Config
}

// SecretToConfig extracts the cluster name and *Config from a KubeConfig secret.
// If the provided secret is not a KubeConfig secret, an error is returned.
func SecretToConfig(secret *kubev1.Secret) (clusterName string, config *Config, err error) {
	clusterName = secret.Name
	kubeConfigBytes, ok := secret.Data[Key]
	if !ok {
		return clusterName, nil, SecretHasNoKubeConfig(secret.ObjectMeta)
	}

	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigBytes)
	if err != nil {
		return clusterName, nil, FailedToConvertSecretToClientConfig(err)
	}

	apiConfig, err := clientcmd.Load(kubeConfigBytes)
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
