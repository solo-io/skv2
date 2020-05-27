package register

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	k8s_core_types "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
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
	secretClient k8s_core_v1.SecretClient,
	nsClientFactory k8s_core_v1.NamespaceClientFromConfigFactory,
	clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory,
	roleClientFactory rbac_v1.RoleClientFromConfigFactory,
) ClusterRegistrant {
	return &clusterRegistrant{
		clusterAuthClientFactory: authorization,
		secretClient:             secretClient,
		nsClientFactory:          nsClientFactory,
		clusterRoleClientFactory: clusterRoleClientFactory,
		roleClientFactory:        roleClientFactory,
		kubeLoader:               loader,
	}
}

/*
	This option should be used mostly for testing.
	When passed in, it will overwrite the Api Server endpoint in the the kubeconfig before it is written.
	This is primarily useful when running multi cluster KinD environments on a mac as  the local IP needs
	to be re-written to `host.docker.internal` so that the local instance knows to hit localhost.
*/
func NewMacTestingRegistrant(
	localClusterDomainOverride string,
	loader kubeconfig.KubeLoader,
	authorization auth.ClusterAuthorizationFactory,
	secretClient k8s_core_v1.SecretClient,
	nsClientFactory k8s_core_v1.NamespaceClientFromConfigFactory,
	clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory,
	roleClientFactory rbac_v1.RoleClientFromConfigFactory,
) ClusterRegistrant {
	return &clusterRegistrant{
		clusterAuthClientFactory: authorization,
		secretClient:             secretClient,
		nsClientFactory:          nsClientFactory,
		clusterRoleClientFactory: clusterRoleClientFactory,
		roleClientFactory:        roleClientFactory,
		kubeLoader:               loader,
	}
}
}
type clusterRegistrant struct {
	clusterAuthClientFactory auth.ClusterAuthorizationFactory
	secretClient             k8s_core_v1.SecretClient
	nsClientFactory          k8s_core_v1.NamespaceClientFromConfigFactory
	saClientFactory          k8s_core_v1.ServiceAccountClientFromConfigFactory
	clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory
	roleClientFactory        rbac_v1.RoleClientFromConfigFactory
	kubeLoader               kubeconfig.KubeLoader

	localClusterDomainOverride string
}

func (c *clusterRegistrant) EnsureRemoteServiceAccount(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	info Options,
) (*k8s_core_types.ServiceAccount, error) {
	saToCreate := &k8s_core_types.ServiceAccount{
		ObjectMeta: k8s_meta_types.ObjectMeta{
			Name:      info.ClusterName,
			Namespace: info.Namespace,
		},
	}

	restCfg, err := remoteClientCfg.ClientConfig()
	if err != nil {
		return nil, err
	}

	saClient, err := c.saClientFactory(restCfg)
	if err != nil {
		return nil, err
	}

	existing, err := saClient.GetServiceAccount(ctx, client.ObjectKey{
		Namespace: saToCreate.Namespace,
		Name:      saToCreate.Name,
	})
	if err != nil {
		if k8s_errs.IsNotFound(err) {
			if err = saClient.CreateServiceAccount(ctx, saToCreate); err != nil {
				return nil, err
			}
			return saToCreate, nil
		}
		return nil, err
	}
	return existing, nil

}

func (c *clusterRegistrant) CreateRemoteAccessToken(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	sa *k8s_core_types.ServiceAccount,
	opts RbacOptions,
) (token string, err error) {

	if err = (&opts).validate(); err != nil {
		return err
	}

	remoteCfg, err := remoteClientCfg.ClientConfig()
	if err != nil {
		return "", err
	}

	authClient, err := c.clusterAuthClientFactory(remoteClientCfg)
	if err != nil {
		return "", err
	}

	if len(opts.Roles) != 0 {
		if err = c.upsertRoles(ctx, remoteCfg, opts.Roles); err != nil {
			return "", err
		}
		for _, v := range opts.Roles {
			opts.RoleBindings = append(opts.RoleBindings, client.ObjectKey{
				Namespace: v.GetNamespace(),
				Name:      v.GetName(),
			})
		}
		token, err = authClient.BuildRemoteBearerToken(
			ctx,
			remoteCfg,
			opts.ClusterName,
			opts.Namespace,
			opts.RoleBindings,
		)
		if err != nil {
			return "", err
		}

	}

	if len(opts.ClusterRoles) != 0 {
		if err = c.upsertClusterRoles(ctx, remoteCfg, opts.ClusterRoles); err != nil {
			return "", err
		}
		for _, v := range opts.ClusterRoles {
			opts.ClusterRoleBindings = append(opts.ClusterRoleBindings, client.ObjectKey{
				Namespace: v.GetNamespace(),
				Name:      v.GetName(),
			})
		}
		token, err = authClient.BuildClusterScopedRemoteBearerToken(
			ctx,
			remoteCfg,
			opts.ClusterName,
			opts.Namespace,
			opts.ClusterRoleBindings,
		)
		if err != nil {
			return "", err
		}
	}
	return token, nil
}

func (c *clusterRegistrant) RegisterClusterWithToken(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	token string,
	opts Options,
) error {
	remoteCfg, err := remoteClientCfg.ClientConfig()
	if err != nil {
		return err
	}

	rawRemoteCfg, err := remoteClientCfg.RawConfig()
	if err != nil {
		return err
	}

	remoteContextName := opts.RemoteCtx
	if opts.RemoteCtx == "" {
		remoteContextName = rawRemoteCfg.CurrentContext
	}
	remoteContext := rawRemoteCfg.Contexts[remoteContextName]
	remoteCluster := rawRemoteCfg.Clusters[remoteContext.Cluster]

	// hacky step for running locally in KIND
	if err = c.hackClusterConfigForLocalTestingInKIND(remoteCluster, opts.RemoteCtx); err != nil {
		return err
	}

	if err = c.ensureRemoteNamespace(ctx, opts.Namespace, remoteCfg); err != nil {
		return err
	}

	kcSecret, err := kubeconfig.ToSecret(
		opts.Namespace,
		opts.ClusterName,
		c.buildRemoteCfg(remoteCluster, remoteContext, opts.ClusterName, token),
	)
	if err != nil {
		return err
	}

	if err = c.upsertSecretData(ctx, kcSecret); err != nil {
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

func (c *clusterRegistrant) upsertRoles(
	ctx context.Context,
	cfg *rest.Config,
	roles []*rbacv1.Role,
) error {
	roleClient, err := c.roleClientFactory(cfg)
	if err != nil {
		return err
	}
	for _, v := range roles {
		if err = roleClient.UpsertRole(ctx, v); err != nil {
			return err
		}
	}
	return nil
}

func (c *clusterRegistrant) upsertClusterRoles(
	ctx context.Context,
	cfg *rest.Config,
	roles []*rbacv1.ClusterRole,
) error {
	clusterRoleClient, err := c.clusterRoleClientFactory(cfg)
	if err != nil {
		return err
	}
	for _, v := range roles {
		if err = clusterRoleClient.UpsertClusterRole(ctx, v); err != nil {
			return err
		}
	}
	return nil
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
	remoteContextName string,
) error {
	serverUrl, err := url.Parse(remoteCluster.Server)
	if err != nil {
		return err
	}

	if strings.HasPrefix(remoteContextName, "kind-") &&
		(serverUrl.Hostname() == "127.0.0.1" || serverUrl.Hostname() == "localhost") &&
		c.localClusterDomainOverride != "" {

		remoteCluster.Server = fmt.Sprintf("https://%s:%s", c.localClusterDomainOverride, serverUrl.Port())
		remoteCluster.InsecureSkipTLSVerify = true
		remoteCluster.CertificateAuthority = ""
		remoteCluster.CertificateAuthorityData = []byte("")
	}

	return nil
}
