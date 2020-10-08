package kubeconfig

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// Fetch ClientConfig. If kubeConfigPath is not specified, retrieve the kubeconfig from environment in which this is invoked.
// Override the API Server URL and current context if specified.
// Copied and modified from https://github.com/kubernetes-sigs/controller-runtime/blob/cb7f85860a8cde7259b35bb84af1fdcb02c098f2/pkg/client/config/config.go#L135
func GetClientConfigWithContext(kubeConfigPath, kubeContext, apiServerUrl string) (clientcmd.ClientConfig, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	if kubeConfigPath != "" {
		loadingRules.ExplicitPath = kubeConfigPath
	} else {
		// Fetch kubeconfig from environment in which this is invoked
		if _, ok := os.LookupEnv("HOME"); !ok {
			u, err := user.Current()
			if err != nil {
				return nil, fmt.Errorf("could not get current user: %v", err)
			}
			loadingRules.Precedence = append(loadingRules.Precedence, path.Join(u.HomeDir, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName))
		}
	}

	overrides := &clientcmd.ConfigOverrides{}
	if kubeContext != "" {
		overrides.CurrentContext = kubeContext
	}
	if apiServerUrl != "" {
		overrides.ClusterInfo = clientcmdapi.Cluster{
			Server: apiServerUrl,
		}
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, overrides), nil
}

// Fetch rest.Config for environment in which this is invoked, override the API Server URL and current context if specified.
func GetRestConfigWithContext(kubeConfigPath, kubeContext, apiServerUrl string) (*rest.Config, error) {
	clientConfig, err := GetClientConfigWithContext(kubeConfigPath, kubeContext, apiServerUrl)
	if err != nil {
		return nil, err
	}
	return clientConfig.ClientConfig()
}

// Fetch raw Config for environment in which this is invoked, override the API Server URL and current context if specified.
func GetRawConfigWithContext(kubeConfigPath, kubeContext, apiServerUrl string) (clientcmdapi.Config, error) {
	clientConfig, err := GetClientConfigWithContext(kubeConfigPath, kubeContext, apiServerUrl)
	if err != nil {
		return clientcmdapi.Config{}, err
	}
	return clientConfig.RawConfig()
}
