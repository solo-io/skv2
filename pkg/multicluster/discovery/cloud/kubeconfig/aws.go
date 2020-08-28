package kubeconfig

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/clients"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewEksConfigBuilder(eksClient clients.EksClient) EksConfigBuilder {
	return &awsClusterConfigBuilder{
		eksClient: eksClient,
	}
}

type awsClusterConfigBuilder struct {
	eksClient clients.EksClient
}

func (a *awsClusterConfigBuilder) ConfigForCluster(ctx context.Context, cluster *eks.Cluster) (clientcmd.ClientConfig, error) {
	tok, err := a.eksClient.Token(ctx, aws.StringValue(cluster.Name))
	if err != nil {
		return nil, err
	}
	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(cluster.CertificateAuthority.Data))
	if err != nil {
		return nil, err
	}

	cfg := kubeconfig.BuildRemoteCfg(
		&clientcmdapi.Cluster{
			Server:                   aws.StringValue(cluster.Endpoint),
			CertificateAuthorityData: ca,
		},
		&clientcmdapi.Context{
			Cluster: aws.StringValue(cluster.Name),
		},
		aws.StringValue(cluster.Name),
		tok.Token,
	)

	return clientcmd.NewDefaultClientConfig(cfg, &clientcmd.ConfigOverrides{}), nil
}
