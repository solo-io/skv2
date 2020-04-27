package register

import (
	"context"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_errs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClusterRegistrant(
	loader kubeconfig.KubeLoader,
	authorization auth.ClusterAuthorization,
	v1Clientset k8s_core_v1.Clientset,
) ClusterRegistrant {
	return &clusterRegistrant{
		clusterAuthClient: authorization,
		secretClient:      v1Clientset.Secrets(),
		kubeLoader:        loader,
	}
}

type clusterRegistrant struct {
	clusterAuthClient auth.ClusterAuthorization
	secretClient      k8s_core_v1.SecretClient
	kubeLoader        kubeconfig.KubeLoader
}

func (c *clusterRegistrant) RegisterCluster(
	ctx context.Context,
	remoteCfg, remoteCtx, clusterName, namespace string,
) error {

	clientCfg, err := c.kubeLoader.GetClientConfigForContext(remoteCfg, remoteCtx)
	if err != nil {
		return err
	}

	return c.RegisterClusterFromConfig(ctx, clientCfg, clusterName, namespace)
}

func (c *clusterRegistrant) RegisterClusterFromConfig(
	ctx context.Context,
	clientCfg clientcmd.ClientConfig,
	clusterName, namespace string,
) error {
	cfg, err := clientCfg.ClientConfig()
	if err != nil {
		return err
	}

	token, err := c.clusterAuthClient.BuildRemoteBearerToken(ctx, cfg, clusterName, namespace)
	if err != nil {
		return err
	}

	rawRemoteCfg, err := clientCfg.RawConfig()
	if err != nil {
		return err
	}

	remoteContextName := rawRemoteCfg.CurrentContext
	if remoteContextName != "" {
		remoteContextName = remoteContextName
	}

	remoteContext := rawRemoteCfg.Contexts[remoteContextName]
	remoteCluster := rawRemoteCfg.Clusters[remoteContext.Cluster]

	secret, err := kubeconfig.ToSecret(
		namespace,
		clusterName,
		buildRemoteCfg(remoteCluster, remoteContext, clusterName, token),
	)
	if err != nil {
		return err
	}

	if err = c.upsertSecretData(ctx, secret); err != nil {
		return err
	}

	return nil
}

func buildRemoteCfg(
	remoteCluster *api.Cluster,
	remoteCtx *api.Context,
	clusterName, token string,
) api.Config {
	return api.Config{
		Kind:        "Secret",
		APIVersion:  "kubernetes_core",
		Preferences: api.Preferences{},
		Clusters: map[string]*api.Cluster{
			clusterName: remoteCluster,
		},
		AuthInfos: map[string]*api.AuthInfo{
			clusterName: {
				Token: token,
			},
		},
		Contexts: map[string]*api.Context{
			clusterName: {
				LocationOfOrigin: remoteCtx.LocationOfOrigin,
				Cluster:          clusterName,
				AuthInfo:         clusterName,
				Namespace:        remoteCtx.Namespace,
				Extensions:       remoteCtx.Extensions,
			},
		},
		CurrentContext: clusterName,
	}
}

func (c *clusterRegistrant) upsertSecretData(
	ctx context.Context,
	secret *k8s_core_types.Secret,
) error {
	existing, err := c.secretClient.GetSecret(ctx, client.ObjectKey{Name: secret.Name, Namespace: secret.Namespace})
	if err != nil {
		if k8s_errs.IsNotFound(err) {
			return c.secretClient.CreateSecret(ctx, secret)
		}
		return err
	}
	existing.Data = secret.Data
	existing.StringData = secret.StringData
	return c.secretClient.UpdateSecret(ctx, existing)
}
