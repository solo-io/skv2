package multicluster

import (
	"github.com/rotisserie/eris"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// KubeConfigSecretLabel is applied to a kubernetes secret to indicate that the secret contains a kubeconfig.
const KubeConfigSecretLabel = "solo.io/kubeconfig"

// Kubeconfig wraps kubernetes kubeconfigs and also indicates the name of the cluster the kubeconfig belongs to.
type KubeConfig struct {
	// the actual kubeconfig
	Config api.Config
	// expected to be used as an identifier string for a cluster
	// stored as the key for the kubeconfig data in a kubernetes secret
	Cluster string
}

var (
	DuplicateClusterName = func(repeatedClusterName string) error {
		return eris.Errorf("Error converting KubeConfigs to secret, duplicate cluster name found: %s", repeatedClusterName)
	}
	FailedToConvertSecretToKubeConfig = func(err error) error {
		return eris.Wrapf(err, "Could not deserialize string to KubeConfig while generating KubeConfig")
	}
	NoDataInKubeConfigSecret = func(secret *kubev1.Secret) error {
		return eris.Errorf("No data in kube config secret %s.%s", secret.ObjectMeta.Name, secret.ObjectMeta.Namespace)
	}
	FailedToConvertSecretToClientConfig = func(err error) error {
		return eris.Wrap(err, "Could not convert config to ClientConfig")
	}
)

// KubeConfigToSecret converts a single kubeconfig to a secret with the provided name and namespace.
func KubeConfigToSecret(name string, namespace string, kc *KubeConfig) (*kubev1.Secret, error) {
	return KubeConfigsToSecret(name, namespace, []*KubeConfig{kc})
}

// KubeConfigsToSecret converts a list of kubeconfigs to a kubeconfig secret with the provided name and namespace.
func KubeConfigsToSecret(name string, namespace string, kcs []*KubeConfig) (*kubev1.Secret, error) {
	secretData := map[string][]byte{}
	for _, kc := range kcs {
		rawKubeConfig, err := clientcmd.Write(kc.Config)
		if err != nil {
			return nil, eris.Wrap(err, "Could not serialize KubeConfig to yaml while generating secret.")
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

// Config contains various structures of kubeconfig.
type Config struct {
	ClientConfig clientcmd.ClientConfig
	ApiConfig    *api.Config
	RestConfig   *rest.Config
}

// SecretToConfigConverter functions extract the cluster name and *Config from a kubeconfig secret.
// If the provided secret is not a kubeconfig secret, an error is returned.
type SecretToConfigConverter func(secret *kubev1.Secret) (clusterName string, config *Config, err error)

// SecretToConfigConverterProvider returns an implemented SecretToConfigConverter
func SecretToConfigConverterProvider() SecretToConfigConverter {
	return SecretToConfig
}

// SecretToKubeConfig is an implementation of SecretToConfigConverter.
func SecretToConfig(secret *kubev1.Secret) (clusterName string, config *Config, err error) {
	if len(secret.Data) > 1 {
		return "", nil, eris.Errorf("kube config secret %s.%s has multiple keys in its data, this is unexpected",
			secret.ObjectMeta.Name, secret.ObjectMeta.Namespace)
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
			return clusterName, nil, eris.Wrapf(err, "Failed to convert secret %s.%s to rest config",
				secret.ObjectMeta.Name, secret.ObjectMeta.Namespace)
		}
		return clusterName, &Config{
			ClientConfig: clientConfig,
			RestConfig:   restConfig,
			ApiConfig:    apiConfig,
		}, nil
	}

	return "", nil, NoDataInKubeConfigSecret(secret)
}
