package register

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
)

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