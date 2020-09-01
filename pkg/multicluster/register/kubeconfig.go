package register

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Represents a Kubeconfig as either on disk (string path) or in memory (clientcmd.ClientConfig)
// Implementation follows the Golang representation of protobuf oneofs.
type KubeCfg struct {
	KubeCfgType isKubeCfgType
}

func (k *KubeCfg) getKubeCfgType() isKubeCfgType {
	if k != nil {
		return k.KubeCfgType
	}
	return nil
}

func (k *KubeCfg) getKubeCfgDisk() string {
	if x, ok := k.getKubeCfgType().(*KubeCfgDisk); ok {
		return x.KubeConfigPath
	}
	return ""
}

func (k *KubeCfg) getClientConfig() clientcmd.ClientConfig {
	if x, ok := k.getKubeCfgType().(*KubeCfgClientConfig); ok {
		return x.ClientConfig
	}
	return nil
}

func (k *KubeCfg) getRestConfig() *rest.Config {
	if x, ok := k.getKubeCfgType().(*KubeCfgRestConfig); ok {
		return x.RestConfig
	}
	return nil
}

type isKubeCfgType interface {
	isKubeCfgType()
}

type KubeCfgDisk struct {
	KubeConfigPath string
}

func (k *KubeCfgDisk) isKubeCfgType() {}

type KubeCfgClientConfig struct {
	ClientConfig clientcmd.ClientConfig
}

func (k *KubeCfgClientConfig) isKubeCfgType() {}

type KubeCfgRestConfig struct {
	RestConfig *rest.Config
}

func (k *KubeCfgRestConfig) isKubeCfgType() {}
