package register

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_errs "k8s.io/apimachinery/pkg/api/errors"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClusterRegistrant(
	loader kubeconfig.KubeLoader,
	authorization auth.ClusterAuthorizationFactory,
	v1Clientset k8s_core_v1.Clientset,
	nsClientFactory k8s_core_v1.NamespaceClientFromConfigFactory,
) ClusterRegistrant {
	return &clusterRegistrant{
		clusterAuthClientFactory: authorization,
		secretClient:             v1Clientset.Secrets(),
		nsClientFactory:          nsClientFactory,
		kubeLoader:               loader,
	}
}

type clusterRegistrant struct {
	clusterAuthClientFactory auth.ClusterAuthorizationFactory
	secretClient             k8s_core_v1.SecretClient
	nsClientFactory          k8s_core_v1.NamespaceClientFromConfigFactory
	kubeLoader               kubeconfig.KubeLoader
}

func (c *clusterRegistrant) RegisterCluster(
	ctx context.Context,
	info ClusterInfo,
	remoteCfg, remoteCtx string,
) error {

	clientCfg, err := c.kubeLoader.GetClientConfigForContext(remoteCfg, remoteCtx)
	if err != nil {
		return err
	}

	return c.RegisterClusterFromConfig(ctx, clientCfg, info)
}

func (c *clusterRegistrant) RegisterClusterFromConfig(
	ctx context.Context,
	clientCfg clientcmd.ClientConfig,
	info ClusterInfo,
) error {
	cfg, err := clientCfg.ClientConfig()
	if err != nil {
		return err
	}

	authClient, err := c.clusterAuthClientFactory(clientCfg)
	if err != nil {
		return err
	}

	token, err := authClient.BuildClusterScopedRemoteBearerToken(ctx, cfg, info.ClusterName, info.Namespace)
	if err != nil {
		return err
	}

	rawRemoteCfg, err := clientCfg.RawConfig()
	if err != nil {
		return err
	}

	remoteContextName := rawRemoteCfg.CurrentContext
	remoteContext := rawRemoteCfg.Contexts[remoteContextName]
	remoteCluster := rawRemoteCfg.Clusters[remoteContext.Cluster]

	// hacky step for running locally in KIND
	if err = c.hackClusterConfigForLocalTestingInKIND(remoteCluster, remoteContextName, info.LocalClusterDomainOverride); err != nil {
		return err
	}

	if err = c.ensureRemoteNamespace(ctx, info.Namespace, cfg); err != nil {
		return err
	}

	secret, err := kubeconfig.ToSecret(
		info.Namespace,
		info.ClusterName,
		c.buildRemoteCfg(remoteCluster, remoteContext, info.ClusterName, token),
	)
	if err != nil {
		return err
	}

	if err = c.upsertSecretData(ctx, secret); err != nil {
		return err
	}

	return nil
}

func (c *clusterRegistrant) ensureRemoteNamespace(ctx context.Context, writeNamespace string, cfg *rest.Config) error {
	nsClient, err := c.nsClientFactory(cfg)
	if err != nil {
		return err
	}
	_, err = nsClient.GetNamespace(ctx, writeNamespace)
	if k8s_errs.IsNotFound(err) {
		return nsClient.CreateNamespace(ctx, &k8s_core_types.Namespace{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name: writeNamespace,
			},
		})
	} else if err != nil {
		return err
	}
	return nil
}

func (c *clusterRegistrant) buildRemoteCfg(
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

// if:
//   * we are operating against a context named "kind-", AND
//   * the server appears to point to localhost, AND
//   * the --local-cluster-domain-override flag is populated with a value
//
// then we rewrite the server config to communicate over the value of `--local-cluster-domain-override`, which
// resolves to the host machine of docker. We also need to skip TLS verification
// and zero-out the cert data, because the cert on the remote cluster's API server wasn't
// issued for the domain contained in the value of `--local-cluster-domain-override`.
//
// this function call is a no-op if those conditions are not met
func (c *clusterRegistrant) hackClusterConfigForLocalTestingInKIND(
	remoteCluster *api.Cluster,
	remoteContextName, clusterDomainOverride string,
) error {
	serverUrl, err := url.Parse(remoteCluster.Server)
	if err != nil {
		return err
	}

	if strings.HasPrefix(remoteContextName, "kind-") &&
		(serverUrl.Hostname() == "127.0.0.1" || serverUrl.Hostname() == "localhost") &&
		clusterDomainOverride != "" {

		remoteCluster.Server = fmt.Sprintf("https://%s:%s", clusterDomainOverride, serverUrl.Port())
		remoteCluster.InsecureSkipTLSVerify = true
		remoteCluster.CertificateAuthority = ""
		remoteCluster.CertificateAuthorityData = []byte("")
	}

	return nil
}
