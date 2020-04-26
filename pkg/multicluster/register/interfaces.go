package register

import (
	"context"

	"github.com/aws/aws-sdk-go/service/eks"
	"k8s.io/client-go/tools/clientcmd"
)

type AwsClusterConfigBuilder interface {
	ConfigForCluster(cluster *eks.Cluster) (clientcmd.ClientConfig, error)
}

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