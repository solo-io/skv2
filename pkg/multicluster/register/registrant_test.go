package register_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/core/v1"
	mock_k8s_rbac_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/rbac.authorization.k8s.io/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	mock_clientcmd "github.com/solo-io/skv2/pkg/generated/mocks/k8s/clientcmd"
	"github.com/solo-io/skv2/pkg/multicluster/auth"
	mock_auth "github.com/solo-io/skv2/pkg/multicluster/auth/mocks"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	mock_kubeconfig "github.com/solo-io/skv2/pkg/multicluster/kubeconfig/mocks"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	k8s_core_types "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Registrant", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		clusterAuthClient        *mock_auth.MockClusterAuthorization
		clusterAuthClientFactory auth.ClusterAuthorizationFactory
		secretClient             *mock_k8s_core_clients.MockSecretClient
		nsClient                 *mock_k8s_core_clients.MockNamespaceClient
		nsClientFactory          k8s_core_v1.NamespaceClientFromConfigFactory
		clusterRoleClient        *mock_k8s_rbac_clients.MockClusterRoleClient
		clusterRoleClientFactory rbac_v1.ClusterRoleClientFromConfigFactory
		roleClient               *mock_k8s_rbac_clients.MockRoleClient
		roleClientFactory        rbac_v1.RoleClientFromConfigFactory
		loader                   *mock_kubeconfig.MockKubeLoader
		clientConfig             *mock_clientcmd.MockClientConfig

		remoteCfg, remoteCtx, clusterName, namespace = "cfg-path", "context", "cluster-name", "namespace"
		testErr                                      = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		secretClient = mock_k8s_core_clients.NewMockSecretClient(ctrl)

		nsClient = mock_k8s_core_clients.NewMockNamespaceClient(ctrl)
		nsClientFactory = func(_ *rest.Config) (k8s_core_v1.NamespaceClient, error) {
			return nsClient, nil
		}
		clusterRoleClient = mock_k8s_rbac_clients.NewMockClusterRoleClient(ctrl)
		clusterRoleClientFactory = func(_ *rest.Config) (rbac_v1.ClusterRoleClient, error) {
			return clusterRoleClient, nil
		}
		roleClient = mock_k8s_rbac_clients.NewMockRoleClient(ctrl)
		roleClientFactory = func(_ *rest.Config) (rbac_v1.RoleClient, error) {
			return roleClient, nil
		}

		clusterAuthClient = mock_auth.NewMockClusterAuthorization(ctrl)
		clusterAuthClientFactory = func(_ clientcmd.ClientConfig) (auth.ClusterAuthorization, error) {
			return clusterAuthClient, nil
		}
		loader = mock_kubeconfig.NewMockKubeLoader(ctrl)
		clientConfig = mock_clientcmd.NewMockClientConfig(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will fail if client config cannot be fetched", func() {
		clusterRegistrant := register.NewClusterRegistrant(
			loader, clusterAuthClientFactory, secretClient, nsClientFactory, clusterRoleClientFactory, roleClientFactory,
		)

		loader.EXPECT().
			GetClientConfigForContext(remoteCfg, remoteCtx).
			Return(nil, testErr)

		err := clusterRegistrant.RegisterCluster(ctx, remoteCfg, register.Options{
			ClusterName: clusterName,
			Namespace:   namespace,
			RemoteCtx:   remoteCtx,
		})
		Expect(err).To(HaveOccurred())
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})

	It("will fail if remote bearer token cannot be built", func() {
		clusterRegistrant := register.NewClusterRegistrant(
			loader, clusterAuthClientFactory, secretClient, nsClientFactory, clusterRoleClientFactory, roleClientFactory,
		)

		cfg := &rest.Config{
			Host: "mock-host",
		}
		loader.EXPECT().
			GetClientConfigForContext(remoteCfg, remoteCtx).
			Return(clientConfig, nil)

		clientConfig.EXPECT().
			ClientConfig().
			Return(cfg, nil)

		clusterAuthClient.EXPECT().
			BuildClusterScopedRemoteBearerToken(ctx, cfg, clusterName, namespace, auth.ServiceAccountClusterAdminRoles).
			Return("", testErr)

		err := clusterRegistrant.RegisterCluster(ctx, remoteCfg, register.Options{
			ClusterName:  clusterName,
			Namespace:    namespace,
			RemoteCtx:    remoteCtx,
			ClusterRoles: auth.ServiceAccountClusterAdminRoles,
		})
		Expect(err).To(HaveOccurred())
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})

	It("will create secret if not found", func() {
		clusterRegistrant := register.NewClusterRegistrant(
			loader, clusterAuthClientFactory, secretClient, nsClientFactory, clusterRoleClientFactory, roleClientFactory,
		)

		restCfg := &rest.Config{
			Host: "mock-host",
		}
		apiCfg := api.Config{
			Clusters: map[string]*api.Cluster{
				clusterName: {
					Server:                   "fake-server",
					CertificateAuthorityData: []byte("fake-ca-data"),
				},
			},
			Contexts: map[string]*api.Context{
				remoteCtx: {
					Cluster: clusterName,
				},
			},
			CurrentContext: remoteCtx,
		}
		token := "token"
		loader.EXPECT().
			GetClientConfigForContext(remoteCfg, remoteCtx).
			Return(clientConfig, nil)

		clientConfig.EXPECT().
			ClientConfig().
			Return(restCfg, nil)

		clusterAuthClient.EXPECT().
			BuildClusterScopedRemoteBearerToken(ctx, restCfg, clusterName, namespace, auth.ServiceAccountClusterAdminRoles).
			Return(token, nil)

		clientConfig.EXPECT().
			RawConfig().
			Return(apiCfg, nil)

		nsClient.EXPECT().
			GetNamespace(ctx, namespace).
			Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

		nsClient.EXPECT().
			CreateNamespace(ctx, &k8s_core_types.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				},
			}).Return(nil)

		secretClient.EXPECT().
			GetSecret(ctx, client.ObjectKey{
				Namespace: namespace,
				Name:      clusterName,
			}).
			Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

		secret, err := kubeconfig.ToSecret(namespace, clusterName, api.Config{
			Kind:        "Secret",
			APIVersion:  "kubernetes_core",
			Preferences: api.Preferences{},
			Clusters: map[string]*api.Cluster{
				clusterName: apiCfg.Clusters[clusterName],
			},
			AuthInfos: map[string]*api.AuthInfo{
				clusterName: {
					Token: token,
				},
			},
			Contexts: map[string]*api.Context{
				clusterName: {
					Cluster:  clusterName,
					AuthInfo: clusterName,
				},
			},
			CurrentContext: clusterName,
		})
		Expect(err).NotTo(HaveOccurred())

		secretClient.EXPECT().
			CreateSecret(ctx, secret).
			Return(nil)

		err = clusterRegistrant.RegisterCluster(ctx, remoteCfg, register.Options{
			ClusterName:  clusterName,
			Namespace:    namespace,
			RemoteCtx:    remoteCtx,
			ClusterRoles: auth.ServiceAccountClusterAdminRoles,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("will overwrite the client config if a LocalClusterDomainOverride is passed in", func() {
		clusterRegistrant := register.NewClusterRegistrant(
			loader, clusterAuthClientFactory, secretClient, nsClientFactory, clusterRoleClientFactory, roleClientFactory,
		)
		clusterDomainOverride := "test-override"
		remoteCtx = "kind-test"
		restCfg := &rest.Config{
			Host: "mock-host",
		}
		apiCfg := api.Config{
			Clusters: map[string]*api.Cluster{
				clusterName: {
					Server:                   "http://localhost:9080",
					CertificateAuthorityData: []byte("fake-ca-data"),
				},
			},
			Contexts: map[string]*api.Context{
				remoteCtx: {
					Cluster: clusterName,
				},
			},
			CurrentContext: remoteCtx,
		}
		token := "token"
		loader.EXPECT().
			GetClientConfigForContext(remoteCfg, remoteCtx).
			Return(clientConfig, nil)

		clientConfig.EXPECT().
			ClientConfig().
			Return(restCfg, nil)

		clusterAuthClient.EXPECT().
			BuildClusterScopedRemoteBearerToken(ctx, restCfg, clusterName, namespace, auth.ServiceAccountClusterAdminRoles).
			Return(token, nil)

		clientConfig.EXPECT().
			RawConfig().
			Return(apiCfg, nil)

		nsClient.EXPECT().
			GetNamespace(ctx, namespace).
			Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

		nsClient.EXPECT().
			CreateNamespace(ctx, &k8s_core_types.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespace,
				},
			}).Return(nil)

		secretClient.EXPECT().
			GetSecret(ctx, client.ObjectKey{
				Namespace: namespace,
				Name:      clusterName,
			}).
			Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

		overwrittenApiConfig := apiCfg.DeepCopy()
		overwrittenApiConfig.Clusters[clusterName].Server = fmt.Sprintf("https://%s:9080", clusterDomainOverride)
		overwrittenApiConfig.Clusters[clusterName].InsecureSkipTLSVerify = true
		overwrittenApiConfig.Clusters[clusterName].CertificateAuthority = ""
		overwrittenApiConfig.Clusters[clusterName].CertificateAuthorityData = []byte("")

		secret, err := kubeconfig.ToSecret(namespace, clusterName, api.Config{
			Kind:        "Secret",
			APIVersion:  "kubernetes_core",
			Preferences: api.Preferences{},
			Clusters: map[string]*api.Cluster{
				clusterName: overwrittenApiConfig.Clusters[clusterName],
			},
			AuthInfos: map[string]*api.AuthInfo{
				clusterName: {
					Token: token,
				},
			},
			Contexts: map[string]*api.Context{
				clusterName: {
					Cluster:  clusterName,
					AuthInfo: clusterName,
				},
			},
			CurrentContext: clusterName,
		})
		Expect(err).NotTo(HaveOccurred())

		secretClient.EXPECT().
			CreateSecret(ctx, secret).
			Return(nil)

		err = clusterRegistrant.RegisterCluster(ctx, remoteCfg, register.Options{
			ClusterName:                clusterName,
			Namespace:                  namespace,
			LocalClusterDomainOverride: clusterDomainOverride,
			RemoteCtx:                  remoteCtx,
			ClusterRoles:               auth.ServiceAccountClusterAdminRoles,
		})
		Expect(err).NotTo(HaveOccurred())
	})

})
