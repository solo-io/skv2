package register

import (
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
		return x.kubeConfigPath
	}
	return ""
}

func (k *KubeCfg) getClientConfig() clientcmd.ClientConfig {
	if x, ok := k.getKubeCfgType().(*KubeCfgMemory); ok {
		return x.clientConfig
	}
	return nil
}

type isKubeCfgType interface {
	isKubeCfgType()
}

type KubeCfgDisk struct {
	kubeConfigPath string
}

func (k *KubeCfgDisk) isKubeCfgType() {}

type KubeCfgMemory struct {
	clientConfig clientcmd.ClientConfig
}

func (k *KubeCfgMemory) isKubeCfgType() {}
