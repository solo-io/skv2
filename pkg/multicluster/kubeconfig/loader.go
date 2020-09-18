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

// Fetch ClientConfig for environment in which this is invoked, override the API Server URL and current context.
// Lifted from https://github.com/kubernetes-sigs/controller-runtime/blob/cb7f85860a8cde7259b35bb84af1fdcb02c098f2/pkg/client/config/config.go#L135
func GetClientConfigWithContext(apiServerUrl, kubeContext string) (clientcmd.ClientConfig, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if _, ok := os.LookupEnv("HOME"); !ok {
		u, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("could not get current user: %v", err)
		}
		loadingRules.Precedence = append(loadingRules.Precedence, path.Join(u.HomeDir, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName))
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{
			ClusterInfo: clientcmdapi.Cluster{
				Server: apiServerUrl,
			},
			CurrentContext: kubeContext,
		}), nil
}

// Fetch rest.Config for environment in which this is invoked, override the API Server URL and current context.
func GetRestConfigWithContext(apiServerUrl, kubeContext string) (*rest.Config, error) {
	clientConfig, err := GetClientConfigWithContext(apiServerUrl, kubeContext)
	if err != nil {
		return nil, err
	}
	return clientConfig.ClientConfig()
}

// Fetch raw Config for environment in which this is invoked, override the API Server URL and current context.
func GetRawConfigWithContext(apiServerUrl, kubeContext string) (clientcmdapi.Config, error) {
	clientConfig, err := GetClientConfigWithContext(apiServerUrl, kubeContext)
	if err != nil {
		return clientcmdapi.Config{}, err
	}
	return clientConfig.RawConfig()
}
