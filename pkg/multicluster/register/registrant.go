package register

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	v1alpha1_providers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/providers"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1"
	k8s_core_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/providers"
	rbac_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/providers"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/register/internal"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	k8s_errs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// visible for testing
	SecretTokenKey = "token"

	DefaultClusterDomain = "cluster.local"
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

/*
	NewClusterRegistrant returns an implementation of ClusterRegistrant.

	localAPIServerAddress is optional. When passed in, it will overwrite the Api Server endpoint in
	the kubeconfig before it is written. This is primarily useful when running multi cluster KinD environments
	on a mac as  the local IP needs to be re-written to `host.docker.internal` so that the local instance
	knows to hit localhost.
*/
func NewClusterRegistrant(
	localAPIServerAddress string,
	clusterRBAC internal.ClusterRBACBinderFactory,
	secretClient k8s_core_v1.SecretClient,
	secretClientFactory k8s_core_v1_providers.SecretClientFromConfigFactory,
	nsClientFactory k8s_core_v1_providers.NamespaceClientFromConfigFactory,
	saClientFactory k8s_core_v1_providers.ServiceAccountClientFromConfigFactory,
	clusterRoleClientFactory rbac_v1_providers.ClusterRoleClientFromConfigFactory,
	roleClientFactory rbac_v1_providers.RoleClientFromConfigFactory,
	kubeClusterFactory v1alpha1_providers.KubernetesClusterClientFromConfigFactory,
) ClusterRegistrant {
	return &clusterRegistrant{
		localAPIServerAddress:    localAPIServerAddress,
		clusterRBACBinderFactory: clusterRBAC,
		secretClient:             secretClient,
		nsClientFactory:          nsClientFactory,
		secretClientFactory:      secretClientFactory,
		saClientFactory:          saClientFactory,
		clusterRoleClientFactory: clusterRoleClientFactory,
		roleClientFactory:        roleClientFactory,
		kubeClusterFactory:       kubeClusterFactory,
	}
}

type clusterRegistrant struct {
	clusterRBACBinderFactory internal.ClusterRBACBinderFactory
	secretClient             k8s_core_v1.SecretClient
	secretClientFactory      k8s_core_v1_providers.SecretClientFromConfigFactory
	nsClientFactory          k8s_core_v1_providers.NamespaceClientFromConfigFactory
	saClientFactory          k8s_core_v1_providers.ServiceAccountClientFromConfigFactory
	clusterRoleClientFactory rbac_v1_providers.ClusterRoleClientFromConfigFactory
	roleClientFactory        rbac_v1_providers.RoleClientFromConfigFactory
	kubeClusterFactory       v1alpha1_providers.KubernetesClusterClientFromConfigFactory

	localAPIServerAddress string
}

func (c *clusterRegistrant) EnsureRemoteServiceAccount(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	opts Options,
) (*corev1.ServiceAccount, error) {

	if err := (&opts).validate(); err != nil {
		return nil, err
	}

	saToCreate := &corev1.ServiceAccount{
		ObjectMeta: serviceAccountObjMeta(opts),
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

func (c *clusterRegistrant) DeleteRemoteServiceAccount(
	ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	opts Options,
) error {
	if err := (&opts).validate(); err != nil {
		return err
	}

	saToDelete := &corev1.ServiceAccount{
		ObjectMeta: serviceAccountObjMeta(opts),
	}

	restCfg, err := remoteClientCfg.ClientConfig()
	if err != nil {
		return err
	}

	saClient, err := c.saClientFactory(restCfg)
	if err != nil {
		return err
	}

	err = saClient.DeleteServiceAccount(ctx, client.ObjectKey{
		Namespace: saToDelete.Namespace,
		Name:      saToDelete.Name,
	})
	if err != nil && !k8s_errs.IsNotFound(err) {
		return err
	}
	return nil
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

func (c *clusterRegistrant) DeleteRemoteAccessResources(ctx context.Context,
	remoteClientCfg clientcmd.ClientConfig,
	opts RbacOptions,
) error {
	saObjMeta := serviceAccountObjMeta(opts.Options)
	sa := client.ObjectKey{Name: saObjMeta.Name, Namespace: saObjMeta.Namespace}

	remoteCfg, err := remoteClientCfg.ClientConfig()
	if err != nil {
		return err
	}
	var multierr *multierror.Error
	// Delete Roles
	if err := c.deleteRoles(ctx, remoteCfg, opts.Roles); err != nil {
		multierr = multierror.Append(multierr, err)
	}
	// Delete ClusterRoles
	if err := c.deleteClusterRoles(ctx, remoteCfg, opts.ClusterRoles); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	rbacBinder, err := c.clusterRBACBinderFactory(remoteClientCfg)
	if err != nil {
		return err
	}

	// Delete RoleBindings
	roleBindings := make([]client.ObjectKey, 0, len(opts.Roles)+len(opts.RoleBindings))
	roleBindings = append(roleBindings, opts.RoleBindings...)
	for _, v := range opts.Roles {
		roleBindings = append(roleBindings, client.ObjectKey{
			Name:      v.GetName(),
			Namespace: v.GetNamespace(),
		})
	}
	if err = rbacBinder.DeleteRoleBindings(ctx, sa, roleBindings); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	// Delete ClusterRoleBindings
	clusterRoleBindings := make([]client.ObjectKey, 0, len(opts.ClusterRoles)+len(opts.ClusterRoleBindings))
	clusterRoleBindings = append(clusterRoleBindings, opts.ClusterRoleBindings...)
	for _, v := range opts.ClusterRoles {
		clusterRoleBindings = append(clusterRoleBindings, client.ObjectKey{
			Name: v.GetName(),
		})
	}
	if err = rbacBinder.DeleteClusterRoleBindings(ctx, sa, clusterRoleBindings); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	return multierr.ErrorOrNil()
}

func (c *clusterRegistrant) RegisterClusterWithToken(
	ctx context.Context,
	masterClusterCfg *rest.Config,
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
	if remoteContextName == "" {
		remoteContextName = rawRemoteCfg.CurrentContext
	}
	remoteContext := rawRemoteCfg.Contexts[remoteContextName]
	remoteCluster := rawRemoteCfg.Clusters[remoteContext.Cluster]

	// hacky step for running locally in KIND
	if err = c.hackClusterConfigForLocalTestingInKIND(remoteCluster, remoteContextName); err != nil {
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

	kubeCluster := buildKubeClusterResource(kcSecret, opts.ClusterDomain)

	kubeClusterClient, err := c.kubeClusterFactory(masterClusterCfg)
	if err != nil {
		return err
	}

	return kubeClusterClient.UpsertKubernetesCluster(ctx, kubeCluster)
}

func (c *clusterRegistrant) DeregisterCluster(
	ctx context.Context,
	masterClusterCfg *rest.Config,
	opts Options,
) error {
	if err := (&opts).validate(); err != nil {
		return err
	}

	kcSecretObjMeta := kubeconfig.SecretObjMeta(opts.Namespace, opts.ClusterName)
	kubeClusterObjMeta := kubeClusterObjMeta(kcSecretObjMeta.Name, kcSecretObjMeta.Namespace)
	kubeClusterClient, err := c.kubeClusterFactory(masterClusterCfg)
	if err != nil {
		return err
	}

	var multierr *multierror.Error
	// Delete local KubernetesCluster
	if err = kubeClusterClient.DeleteKubernetesCluster(ctx, client.ObjectKey{Name: kubeClusterObjMeta.Name, Namespace: kubeClusterObjMeta.Namespace}); err != nil {
		multierr = multierror.Append(multierr, err)
	}
	// Delete remote secret
	if err = c.deleteSecret(ctx, client.ObjectKey{Name: kcSecretObjMeta.Name, Namespace: kcSecretObjMeta.Namespace}); err != nil {
		multierr = multierror.Append(multierr, err)
	}
	return multierr.ErrorOrNil()
}

func buildKubeClusterResource(secret *corev1.Secret, clusterDomain string) *v1alpha1.KubernetesCluster {
	if clusterDomain == "" {
		clusterDomain = DefaultClusterDomain
	}
	return &v1alpha1.KubernetesCluster{
		ObjectMeta: kubeClusterObjMeta(secret.Name, secret.Namespace),
		Spec: v1alpha1.KubernetesClusterSpec{
			SecretName:    secret.Name,
			ClusterDomain: clusterDomain,
		},
	}
}

func kubeClusterObjMeta(secretName, secretNamespace string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      secretName,
		Namespace: secretNamespace,
	}
}

func (c *clusterRegistrant) ensureRemoteNamespace(ctx context.Context, writeNamespace string, cfg *rest.Config) error {
	nsClient, err := c.nsClientFactory(cfg)
	if err != nil {
		return err
	}
	_, err = nsClient.GetNamespace(ctx, writeNamespace)
	if k8s_errs.IsNotFound(err) {
		return nsClient.CreateNamespace(ctx, &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
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
	secret *corev1.Secret,
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

func (c *clusterRegistrant) deleteSecret(
	ctx context.Context,
	secretObjKey client.ObjectKey,
) error {
	err := c.secretClient.DeleteSecret(ctx, secretObjKey)
	if err != nil && !k8s_errs.IsNotFound(err) {
		return err
	}
	return nil
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

func (c *clusterRegistrant) deleteRoles(ctx context.Context,
	cfg *rest.Config,
	roles []*rbacv1.Role,
) error {
	roleClient, err := c.roleClientFactory(cfg)
	if err != nil {
		return err
	}
	for _, v := range roles {
		if err = roleClient.DeleteRole(ctx, client.ObjectKey{Name: v.Name, Namespace: v.Namespace}); err != nil {
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

func (c *clusterRegistrant) deleteClusterRoles(
	ctx context.Context,
	cfg *rest.Config,
	roles []*rbacv1.ClusterRole,
) error {
	clusterRoleClient, err := c.clusterRoleClientFactory(cfg)
	if err != nil {
		return err
	}
	for _, v := range roles {
		if err = clusterRoleClient.DeleteClusterRole(ctx, v.Name); err != nil {
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
	remoteSecretClient, err := c.secretClientFactory(cfg)
	if err != nil {
		return "", err
	}
	var foundSecret *corev1.Secret
	if err = retry.Do(func() error {
		sa, err := saClient.GetServiceAccount(ctx, saRef)
		if err != nil {
			return err
		}
		if len(sa.Secrets) == 0 {
			return eris.Errorf(
				"service account %s.%s does not have a token secret associated with it",
				saRef.Name,
				saRef.Namespace,
			)
		}
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
		c.localAPIServerAddress != "" {

		port := serverUrl.Port()
		if host, newPort, err := net.SplitHostPort(c.localAPIServerAddress); err == nil {
			c.localAPIServerAddress = host
			port = newPort
		}
		remoteCluster.Server = fmt.Sprintf("https://%s:%s", c.localAPIServerAddress, port)
		remoteCluster.InsecureSkipTLSVerify = true
		remoteCluster.CertificateAuthority = ""
		remoteCluster.CertificateAuthorityData = []byte("")
	}

	return nil
}

func serviceAccountObjMeta(opts Options) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      opts.ClusterName,
		Namespace: opts.RemoteNamespace,
	}
}
