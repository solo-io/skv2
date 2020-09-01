package register

import (
	"github.com/rotisserie/eris"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// Represents a Kubeconfig as either on disk (string path) or in memory as either clientcmd.ClientConfig or rest.Config.
// Implementation follows the Golang representation of protobuf oneofs.
type KubeCfg struct {
	KubeCfgType isKubeCfgType
}

func NewDiskKubeCfg(kubeConfigPath, kubeContext string) *KubeCfg {
	return &KubeCfg{
		KubeCfgType: &KubeCfgDisk_{
			KubeCfgDisk: KubeCfgDisk{
				KubeConfigPath: kubeConfigPath,
				KubeContext:    kubeContext,
			},
		},
	}
}

func NewClientKubeConfig(clientConfig clientcmd.ClientConfig) *KubeCfg {
	return &KubeCfg{
		KubeCfgType: &KubeCfgClientConfig{
			ClientConfig: clientConfig,
		},
	}
}

func NewRestKubeConfig(restConfig *rest.Config) *KubeCfg {
	return &KubeCfg{
		KubeCfgType: &KubeCfgRestConfig{
			RestConfig: restConfig,
		},
	}
}

func (k *KubeCfg) GetKubeCfgType() isKubeCfgType {
	if k != nil {
		return k.KubeCfgType
	}
	return nil
}

func (k *KubeCfg) ConstructRestConfig() (*rest.Config, error) {
	switch k.GetKubeCfgType().(type) {
	case *KubeCfgDisk_:
		kubeCfgDisk := k.GetKubeCfgDisk()
		if kubeCfgDisk.KubeConfigPath == "" {
			kubeCfgDisk.KubeConfigPath = clientcmd.RecommendedHomeFile
		}
		masterCfg, err := getClientConfigWithContext(kubeCfgDisk.ApiServerUrl, kubeCfgDisk.KubeConfigPath, kubeCfgDisk.KubeContext)
		if err != nil {
			return nil, err
		}
		return masterCfg.ClientConfig()
	case *KubeCfgClientConfig:
		return k.GetClientConfig().ClientConfig()
	case *KubeCfgRestConfig:
		return k.GetRestConfig(), nil
	}
	return nil, eris.New("No kubeconfig data found.")
}

func (k *KubeCfg) ConstructClientConfig() (clientcmd.ClientConfig, error) {
	switch k.GetKubeCfgType().(type) {
	case *KubeCfgDisk_:
		kubeCfgDisk := k.GetKubeCfgDisk()
		if kubeCfgDisk.KubeConfigPath == "" {
			kubeCfgDisk.KubeConfigPath = clientcmd.RecommendedHomeFile
		}
		masterCfg, err := getClientConfigWithContext(kubeCfgDisk.ApiServerUrl, kubeCfgDisk.KubeConfigPath, kubeCfgDisk.KubeContext)
		if err != nil {
			return nil, err
		}
		return masterCfg, nil
	case *KubeCfgClientConfig:
		return k.GetClientConfig(), nil
	case *KubeCfgRestConfig:
		return nil, eris.New("Cannot convert rest.Config into clientcmd.ClientConfig.")
	}
	return nil, eris.New("No kubeconfig data found.")
}

// Attempts to load a Client KubeConfig from a default list of sources.
func getClientConfigWithContext(serverUrl, kubeCfgPath, context string) (clientcmd.ClientConfig, error) {
	verifiedKubeConfigPath := clientcmd.RecommendedHomeFile
	if kubeCfgPath != "" {
		verifiedKubeConfigPath = kubeCfgPath
	}

	if err := assertKubeConfigExists(verifiedKubeConfigPath); err != nil {
		return nil, err
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = verifiedKubeConfigPath
	configOverrides := &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: serverUrl}}

	if context != "" {
		configOverrides.CurrentContext = context
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides), nil
}

func (k *KubeCfg) GetKubeCfgDisk() *KubeCfgDisk {
	if x, ok := k.GetKubeCfgType().(*KubeCfgDisk_); ok {
		return &x.KubeCfgDisk
	}
	return nil
}

func (k *KubeCfg) GetClientConfig() clientcmd.ClientConfig {
	if x, ok := k.GetKubeCfgType().(*KubeCfgClientConfig); ok {
		return x.ClientConfig
	}
	return nil
}

func (k *KubeCfg) GetRestConfig() *rest.Config {
	if x, ok := k.GetKubeCfgType().(*KubeCfgRestConfig); ok {
		return x.RestConfig
	}
	return nil
}

type isKubeCfgType interface {
	isKubeCfgType()
}

type KubeCfgDisk_ struct {
	KubeCfgDisk KubeCfgDisk
}

type KubeCfgDisk struct {
	KubeConfigPath string
	// override the context to use from the local kubeconfig. if unset, use current context
	KubeContext string
	// override the URL of the k8s server
	ApiServerUrl string
}

func (k *KubeCfgDisk_) isKubeCfgType() {}

type KubeCfgClientConfig struct {
	ClientConfig clientcmd.ClientConfig
}

func (k *KubeCfgClientConfig) isKubeCfgType() {}

type KubeCfgRestConfig struct {
	RestConfig *rest.Config
}

func (k *KubeCfgRestConfig) isKubeCfgType() {}
