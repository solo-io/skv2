package register

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/rotisserie/eris"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
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

const (
	// visible for testing
	SecretTokenKey = "token"
)

var (
	MalformedSecret = eris.New("service account secret does not contain a bearer token")
	SecretNotReady  = func(err error) error {
		return eris.Wrap(err, "secret for service account not ready yet")
	}
	// exponential backoff retry with an initial period of 0.1s for 7 iterations, which will mean a cumulative retry period of ~6s
	// visible for testing
	SecretLookupOpts = []retry.Option{
		retry.Delay(time.Millisecond * 100),
		retry.Attempts(7),
		retry.DelayType(retry.BackOffDelay),
	}
)

func NewClusterRegistrant(
	authorization ClusterRBACBinderFactory,
	secretClient k8s_core_v1.SecretClient,
	secretClientFactory k8s_core_v1.SecretClientFromConfigFactory,
	nsClientFactory k8s_core_v1.NamespaceClientFromConfigFactory,
	saClientFactory k8s_core_v1.ServiceAccountClientFromConfigFactory,
	clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory,
	roleClientFactory rbac_v1.RoleClientFromConfigFactory,
) ClusterRegistrant {
	return &clusterRegistrant{
		clusterRBACBinderFactory: authorization,
		secretClient:             secretClient,
		nsClientFactory:          nsClientFactory,
		secretClientFactory:      secretClientFactory,
		clusterRoleClientFactory: clusterRoleClientFactory,
		roleClientFactory:        roleClientFactory,
		saClientFactory:          saClientFactory,
	}
}

/*
	This option should be used mostly for testing.
	When passed in, it will overwrite the Api Server endpoint in the the kubeconfig before it is written.
	This is primarily useful when running multi cluster KinD environments on a mac as  the local IP needs
	to be re-written to `host.docker.internal` so that the local instance knows to hit localhost.
*/
func NewTestingRegistrant(
	localClusterDomainOverride string,
	clusterRBAC ClusterRBACBinderFactory,
	secretClient k8s_core_v1.SecretClient,
	secretClientFactory k8s_core_v1.SecretClientFromConfigFactory,
	nsClientFactory k8s_core_v1.NamespaceClientFromConfigFactory,
	saClientFactory k8s_core_v1.ServiceAccountClientFromConfigFactory,
	clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory,
	roleClientFactory rbac_v1.RoleClientFromConfigFactory,
) ClusterRegistrant {
	return &clusterRegistrant{
		localClusterDomainOverride: localClusterDomainOverride,
		clusterRBACBinderFactory:   clusterRBAC,
		secretClient:               secretClient,
		nsClientFactory:            nsClientFactory,
		secretClientFactory:        secretClientFactory,
		saClientFactory:            saClientFactory,
		clusterRoleClientFactory:   clusterRoleClientFactory,
		roleClientFactory:          roleClientFactory,
	}
}

type clusterRegistrant struct {
	clusterRBACBinderFactory ClusterRBACBinderFactory
	secretClient             k8s_core_v1.SecretClient
	secretClientFactory      k8s_core_v1.SecretClientFromConfigFactory
	nsClientFactory          k8s_core_v1.NamespaceClientFromConfigFactory
	saClientFactory          k8s_core_v1.ServiceAccountClientFromConfigFactory
	clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory
	roleClientFactory        rbac_v1.RoleClientFromConfigFactory

	localClusterDomainOverride string
}

func (c *clusterRegistrant) EnsureRemoteServiceAccount(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	opts Options,
) (*k8s_core_types.ServiceAccount, error) {

	if err := (&opts).validate(); err != nil {
		return nil, err
	}

	saToCreate := &k8s_core_types.ServiceAccount{
		ObjectMeta: k8s_meta_types.ObjectMeta{
			Name:      opts.ClusterName,
			Namespace: opts.RemoteNamespace,
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
	sa client.ObjectKey,
	opts RbacOptions,
) (token string, err error) {

	if err = (&opts).validate(); err != nil {
		return "", err
	}

	remoteCfg, err := remoteClientCfg.ClientConfig()
	if err != nil {
		return "", err
	}

	rbacBinder, err := c.clusterRBACBinderFactory(remoteClientCfg)
	if err != nil {
		return "", err
	}

	roleBindings := make([]client.ObjectKey, 0, len(opts.Roles)+len(opts.RoleBindings))
	roleBindings = append(roleBindings, opts.RoleBindings...)
	if len(opts.Roles) != 0 {
		if err = c.upsertRoles(ctx, remoteCfg, opts.Roles); err != nil {
			return "", err
		}
		for _, v := range opts.Roles {
			roleBindings = append(roleBindings, client.ObjectKey{
				Namespace: v.GetNamespace(),
				Name:      v.GetName(),
			})
		}
	}
	if len(roleBindings) > 0 {
		if err = rbacBinder.BindRoles(ctx, sa, roleBindings); err != nil {
			return "", err
		}
	}

	clusterRoleBindings := make([]client.ObjectKey, 0, len(opts.ClusterRoles)+len(opts.ClusterRoleBindings))
	clusterRoleBindings = append(clusterRoleBindings, opts.ClusterRoleBindings...)
	if len(opts.ClusterRoles) != 0 {
		if err = c.upsertClusterRoles(ctx, remoteCfg, opts.ClusterRoles); err != nil {
			return "", err
		}
		for _, v := range opts.ClusterRoles {
			clusterRoleBindings = append(clusterRoleBindings, client.ObjectKey{
				Name: v.GetName(),
			})
		}
	}
	if len(clusterRoleBindings) > 0 {
		if err = rbacBinder.BindClusterRoles(ctx, sa, clusterRoleBindings); err != nil {
			return "", err
		}
	}

	return c.getTokenForSa(ctx, remoteCfg, sa)
}

func (c *clusterRegistrant) RegisterClusterWithToken(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	token string,
	opts Options,
) error {

	if err := (&opts).validate(); err != nil {
		return err
	}

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

func (c *clusterRegistrant) getTokenForSa(
	ctx context.Context,
	cfg *rest.Config,
	saRef client.ObjectKey,
) (string, error) {
	saClient, err := c.saClientFactory(cfg)
	if err != nil {
		return "", err
	}
	sa, err := saClient.GetServiceAccount(ctx, saRef)
	if err != nil {
		return "", err
	}
	if len(sa.Secrets) == 0 {
		return "", eris.Errorf(
			"service account %s.%s does not have a token secret associated with it",
			saRef.Name,
			saRef.Namespace,
		)
	}
	remoteSecretClient, err := c.secretClientFactory(cfg)
	if err != nil {
		return "", err
	}
	var foundSecret *k8s_core_types.Secret
	if err = retry.Do(func() error {
		secretName := sa.Secrets[0].Name
		secret, err := remoteSecretClient.GetSecret(ctx, client.ObjectKey{Name: secretName, Namespace: saRef.Namespace})
		if err != nil {
			return err
		}

		foundSecret = secret
		return nil
	}, SecretLookupOpts...); err != nil {
		return "", SecretNotReady(err)
	}

	serviceAccountToken, ok := foundSecret.Data[SecretTokenKey]
	if !ok {
		return "", MalformedSecret
	}

	return string(serviceAccountToken), nil
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

		port := serverUrl.Port()
		if host, newPort, err := net.SplitHostPort(c.localClusterDomainOverride); err == nil {
			c.localClusterDomainOverride = host
			port = newPort
		}
		remoteCluster.Server = fmt.Sprintf("https://%s:%s", c.localClusterDomainOverride, port)
		remoteCluster.InsecureSkipTLSVerify = true
		remoteCluster.CertificateAuthority = ""
		remoteCluster.CertificateAuthorityData = []byte("")
	}

	return nil
}
