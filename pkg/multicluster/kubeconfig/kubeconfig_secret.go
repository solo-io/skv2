package kubeconfig

import (
	"io/ioutil"

	"github.com/rotisserie/eris"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

const (
	// SecretType is used to indicate which kubernetes secrets contain a KubeConfig.
	SecretType = "solo.io/kubeconfig"

	// Key is the KubeConfig's key in the KubeConfig secret data.
	Key = "kubeconfig"
)

var (
	FailedToReadCAFile = func(err error, fileName string) error {
		return eris.Wrapf(err, "Failed to read kubeconfig CA file: %s", fileName)
	}
)

// TODO settle on how to canonicalize cluster names: https://github.com/solo-io/skv2/issues/15

// ToSecret converts a kubernetes api.Config to a secret with the provided name and namespace.
func ToSecret(namespace string, cluster string, kc api.Config) (*kubev1.Secret, error) {
	err := convertCertFilesToInline(kc)
	if err != nil {
		return nil, err
	}

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

/*
	https://github.com/solo-io/service-mesh-hub/issues/590
	If the user has a cert authority file set instead of the raw bytes in their kubeconfig, then
	we'll fail later when the pods in-cluster try to read that file path.
	We need to read the file right now, in a CLI context, and manually shuffle the bytes over to the CA data field

	This is necessary if we are reading a kubeconfig on a user's local machine which has it's CertificateAuthority
	set to a file, rather than the raw bytes. When a KubeConfig file is built it will first attempt to use the bytes,
	but if they are missing, it will use the file instead. Since we are copying this kubeconfig into the cluster,
	the file will no longer be available, therefore we have to read the contents of the file, and store them instead
	as bytes.
*/
func convertCertFilesToInline(cfg api.Config) error {
	currentCluster := cfg.Clusters[cfg.Contexts[cfg.CurrentContext].Cluster]
	if len(currentCluster.CertificateAuthority) > 0 {
		fileContent, err := ioutil.ReadFile(currentCluster.CertificateAuthority)
		if err != nil {
			return FailedToReadCAFile(err, currentCluster.CertificateAuthority)
		}

		currentCluster.CertificateAuthorityData = fileContent
		currentCluster.CertificateAuthority = "" // dont need to record the filename in the config; we have the data present
	}

	return nil
}

// SecretToConfig extracts the cluster name and *Config from a KubeConfig secret.
// If the provided secret is not a KubeConfig secret, an error is returned.
func SecretToConfig(secret *kubev1.Secret) (clusterName string, config clientcmd.ClientConfig, err error) {
	clusterName = secret.Name
	kubeConfigBytes, ok := secret.Data[Key]
	if !ok {
		return clusterName, nil, SecretHasNoKubeConfig(secret.ObjectMeta)
	}

	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigBytes)
	if err != nil {
		return clusterName, nil, FailedToConvertSecretToClientConfig(err)
	}
	return clusterName, clientConfig, nil
}
