package kubeconfig

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

//go:generate mockgen -destination ./mocks/mock_interfaces.go -source ./interfaces.go

type KubeLoader interface {
	GetRestConfigForContext(path string, context string) (*rest.Config, error)
	GetRawConfigForContext(path, context string) (clientcmdapi.Config, error)
	GetClientConfigForContext(path, context string) (clientcmd.ClientConfig, error)
	GetRestConfigFromBytes(config []byte) (*rest.Config, error)
}
