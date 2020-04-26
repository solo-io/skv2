package kubeconfig

import (
	"os"
	"time"

	"github.com/solo-io/go-utils/kubeutils"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// given a path to a kube config file, convert it into either creds for hitting the API server of the cluster it points to,
// or return the contexts/clusters it is aware of
//go:generate mockgen -destination ./mocks/mock_kube_loader.go -source ./loader.go

type KubeLoader interface {
	GetRestConfigForContext(path string, context string) (*rest.Config, error)
	GetRawConfigForContext(path, context string) (clientcmdapi.Config, error)
	GetClientConfigForContext(path, context string) (clientcmd.ClientConfig, error)
	GetRestConfigFromBytes(config []byte) (*rest.Config, error)
}

// only the pieces from a kube config that we need to operate on
// mainly just used to simplify from the complexity of the actual object
type KubeContext struct {
	CurrentContext string
	Contexts       map[string]*api.Context
	Clusters       map[string]*api.Cluster
}

// default KubeLoader
func DefaultKubeLoaderProvider(timeout *time.Duration) KubeLoader {
	return &kubeLoader{
		timeout: timeout.String(),
	}
}

type kubeLoader struct {
	timeout string
}

func (k *kubeLoader) GetClientConfigForContext(path, context string) (clientcmd.ClientConfig, error) {
	return k.getConfigWithContext("", path, context)
}

func (k *kubeLoader) GetRestConfigForContext(path string, context string) (*rest.Config, error) {
	cfg, err := k.getConfigWithContext("", path, context)
	if err != nil {
		return nil, err
	}

	return cfg.ClientConfig()
}

func (k *kubeLoader) GetRestConfigFromBytes(config []byte) (*rest.Config, error) {
	return clientcmd.RESTConfigFromKubeConfig(config)
}

func (k *kubeLoader) getConfigWithContext(masterURL, kubeconfigPath, context string) (clientcmd.ClientConfig, error) {
	verifiedKubeConfigPath := clientcmd.RecommendedHomeFile
	if kubeconfigPath != "" {
		verifiedKubeConfigPath = kubeconfigPath
	}

	if err := assertKubeConfigExists(verifiedKubeConfigPath); err != nil {
		return nil, err
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = verifiedKubeConfigPath
	configOverrides := &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterURL}}

	if context != "" {
		configOverrides.CurrentContext = context
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides), nil
}

// expects `path` to be nonempty
func assertKubeConfigExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	return nil
}

func (k *kubeLoader) GetRawConfigForContext(path, context string) (clientcmdapi.Config, error) {
	cfg, err := k.getConfigWithContext("", path, context)
	if err != nil {
		return clientcmdapi.Config{}, err
	}

	return cfg.RawConfig()
}

func (k *kubeLoader) ParseContext(path string) (*KubeContext, error) {
	cfg, err := kubeutils.GetKubeConfig("", path)
	if err != nil {
		return nil, err
	}

	return &KubeContext{
		CurrentContext: cfg.CurrentContext,
		Contexts:       cfg.Contexts,
		Clusters:       cfg.Clusters,
	}, nil
}
