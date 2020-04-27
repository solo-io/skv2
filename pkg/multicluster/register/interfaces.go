package register

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

type ClusterInfo struct {
	ClusterName string
	Namespace   string
	LocalClusterDomainOverride string
}

type ClusterRegistrant interface {
	RegisterCluster(
		ctx context.Context,
		info ClusterInfo,
		remoteCfg, remoteCtx string,
	) error
	RegisterClusterFromConfig(
		ctx context.Context,
		config clientcmd.ClientConfig,
		info ClusterInfo,
	) error
}
