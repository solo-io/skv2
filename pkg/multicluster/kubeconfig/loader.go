package kubeconfig

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// Fetch ClientConfig. If kubeConfigPath is not specified, retrieve the kubeconfig from environment in which this is invoked.
// Override the API Server URL and current context if specified.
func GetClientConfigWithContext(kubeConfigPath, kubeContext, apiServerUrl string) (clientcmd.ClientConfig, error) {

	// default loading rules checks for KUBECONFIG env var
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// also check recommended default kubeconfig file locations
	loadingRules.Precedence = append(loadingRules.Precedence, clientcmd.RecommendedHomeFile)

	// explicit path overrides all loading rules, will error if not found
	if kubeConfigPath != "" {
		loadingRules.ExplicitPath = kubeConfigPath
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
