package kubeconfig

import (
	"os"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// default KubeLoader
func NewKubeLoader(timeout time.Duration) KubeLoader {
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
	clientCfg, err := clientcmd.NewClientConfigFromBytes(config)
	if err != nil {
		return nil, err
	}
	return clientCfg.ClientConfig()
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
