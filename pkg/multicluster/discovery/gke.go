package discovery

import (
	"context"
	"encoding/base64"

	"github.com/solo-io/skv2/pkg/multicluster/discovery/cloud"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewGkeConfigBuilder(gkeClient cloud.GkeClient) GkeConfigBuilder {
	return &gkeConfigBuilder{gkeClient: gkeClient}
}

type gkeConfigBuilder struct {
	gkeClient cloud.GkeClient
}

func (g *gkeConfigBuilder) ConfigForCluster(ctx context.Context, cluster *containerpb.Cluster) (clientcmd.ClientConfig, error) {
	token, err := g.gkeClient.Token(ctx)
	if err != nil {
		return nil, err
	}
	ca, err := base64.StdEncoding.DecodeString(cluster.GetMasterAuth().GetClusterCaCertificate())
	if err != nil {
		return nil, err
	}

	cfg := kubeconfig.BuildRemoteCfg(
		&clientcmdapi.Cluster{
			Server:                   cluster.GetEndpoint(),
			CertificateAuthorityData: ca,
		},
		&clientcmdapi.Context{
			Cluster: cluster.GetName(),
		},
		cluster.GetName(),
		token.AccessToken,
	)

	return clientcmd.NewDefaultClientConfig(cfg, &clientcmd.ConfigOverrides{}), nil
}
