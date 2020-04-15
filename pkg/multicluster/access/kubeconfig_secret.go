package registration

import (
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// KubeConfigSecretLabel is applied to a kubernetes secret to indicate that the secret contains a kubeconfig.
const KubeConfigSecretLabel = "solo.io/kubeconfig"

// KubeConfig wraps kubernetes KubeConfigs and also indicates the name of the cluster the KubeConfig belongs to.
// Used to persist KubeConfigs to kubernetes secrets.
type KubeConfig struct {
	// the actual KubeConfig
	Config api.Config
	// expected to be used as an identifier string for a cluster
	// stored as the key for the KubeConfig data in a kubernetes secret
	Cluster string
}

// KubeConfigToSecret converts a single KubeConfig to a secret with the provided name and namespace.
func KubeConfigToSecret(name string, namespace string, kc *KubeConfig) (*kubev1.Secret, error) {
	return KubeConfigsToSecret(name, namespace, []*KubeConfig{kc})
}

// KubeConfigsToSecret converts a list of KubeConfigs to a KubeConfig secret with the provided name and namespace.
func KubeConfigsToSecret(name string, namespace string, kcs []*KubeConfig) (*kubev1.Secret, error) {
	secretData := map[string][]byte{}
	for _, kc := range kcs {
		rawKubeConfig, err := clientcmd.Write(kc.Config)
		if err != nil {
			return nil, FailedToConvertKubeConfigToSecret(err)
		}
		if _, exists := secretData[kc.Cluster]; exists {
			return nil, DuplicateClusterName(kc.Cluster)
		}
		secretData[kc.Cluster] = rawKubeConfig
	}
	return &kubev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Labels:    map[string]string{KubeConfigSecretLabel: "true"},
			Name:      name,
			Namespace: namespace,
		},
		Type: kubev1.SecretTypeOpaque,
		Data: secretData,
	}, nil
}

// Config contains various KubeConfig formats for convenience.
type Config struct {
	ClientConfig clientcmd.ClientConfig
	ApiConfig    *api.Config
	RestConfig   *rest.Config
}

// SecretToConfigConverter functions extract the cluster name and *Config from a KubeConfig secret.
// If the provided secret is not a KubeConfig secret, an error is returned.
type SecretToConfigConverter func(secret *kubev1.Secret) (clusterName string, config *Config, err error)

// SecretToConfigConverterProvider returns an implemented SecretToConfigConverter
func SecretToConfigConverterProvider() SecretToConfigConverter {
	return SecretToConfig
}

// SecretToKubeConfig is an implementation of SecretToConfigConverter.
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
