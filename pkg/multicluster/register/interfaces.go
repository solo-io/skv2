package register

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

type ClusterRegistrant interface {
	RegisterCluster(
		ctx context.Context,
		remoteCfg, remoteCtx, clusterName, namespace string,
	) error
	RegisterClusterFromConfig(
		ctx context.Context,
		config clientcmd.ClientConfig,
		clusterName, namespace string,
	) error
}
