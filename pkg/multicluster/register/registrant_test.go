package register_test

import (
	"context"
	"fmt"
	"time"

	"github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	mock_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/mocks"
	v1alpha1_providers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/providers"

	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/mocks"
	rbac_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1"
	mock_k8s_rbac_clients "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/mocks"
	"github.com/solo-io/skv2/pkg/multicluster/register/mock_clientcmd"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1"
	k8s_core_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/providers"
	rbac_v1_providers "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1/providers"

	"github.com/avast/retry-go"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	"github.com/solo-io/skv2/pkg/multicluster/register/internal"
	mock_internal "github.com/solo-io/skv2/pkg/multicluster/register/internal/mocks"
	k8s_core_types "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
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

		clusterRBACBinder        *mock_internal.MockClusterRBACBinder
		clusterRbacBinderFactory internal.ClusterRBACBinderFactory
		secretClient             *mock_k8s_core_clients.MockSecretClient
		secretClientFactory      k8s_core_v1_providers.SecretClientFromConfigFactory
		nsClient                 *mock_k8s_core_clients.MockNamespaceClient
		nsClientFactory          k8s_core_v1_providers.NamespaceClientFromConfigFactory
		saClient                 *mock_k8s_core_clients.MockServiceAccountClient
		saClientFactory          k8s_core_v1_providers.ServiceAccountClientFromConfigFactory
		clusterRoleClient        *mock_k8s_rbac_clients.MockClusterRoleClient
		clusterRoleClientFactory rbac_v1_providers.ClusterRoleClientFromConfigFactory
		roleClient               *mock_k8s_rbac_clients.MockRoleClient
		roleClientFactory        rbac_v1_providers.RoleClientFromConfigFactory
		kubeClusterClient        *mock_v1alpha1.MockKubernetesClusterClient
		kubeClusterClientFactory v1alpha1_providers.KubernetesClusterClientFromConfigFactory
		clientConfig             *mock_clientcmd.MockClientConfig

		_, remoteCtx, clusterName, namespace = "cfg-path", "context", "cluster-name", "namespace"
		testErr                              = eris.New("hello")

		saName = "sa-name"

		saObjectKey = func() client.ObjectKey {
			return client.ObjectKey{
				Namespace: namespace,
				Name:      saName,
			}
		}
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		secretClient = mock_k8s_core_clients.NewMockSecretClient(ctrl)
		secretClientFactory = func(_ *rest.Config) (k8s_core_v1.SecretClient, error) {
			return secretClient, nil
		}
		nsClient = mock_k8s_core_clients.NewMockNamespaceClient(ctrl)
		nsClientFactory = func(_ *rest.Config) (k8s_core_v1.NamespaceClient, error) {
			return nsClient, nil
		}
		saClient = mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		saClientFactory = func(_ *rest.Config) (k8s_core_v1.ServiceAccountClient, error) {
			return saClient, nil
		}
		clusterRoleClient = mock_k8s_rbac_clients.NewMockClusterRoleClient(ctrl)
		clusterRoleClientFactory = func(_ *rest.Config) (rbac_v1.ClusterRoleClient, error) {
			return clusterRoleClient, nil
		}
		roleClient = mock_k8s_rbac_clients.NewMockRoleClient(ctrl)
		roleClientFactory = func(_ *rest.Config) (rbac_v1.RoleClient, error) {
			return roleClient, nil
		}

		kubeClusterClient = mock_v1alpha1.NewMockKubernetesClusterClient(ctrl)
		kubeClusterClientFactory = func(_ *rest.Config) (v1alpha1.KubernetesClusterClient, error) {
			return kubeClusterClient, nil
		}

		clusterRBACBinder = mock_internal.NewMockClusterRBACBinder(ctrl)
		clusterRbacBinderFactory = func(_ clientcmd.ClientConfig) (internal.ClusterRBACBinder, error) {
			return clusterRBACBinder, nil
		}
		clientConfig = mock_clientcmd.NewMockClientConfig(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("EnsureRemoteServiceAccount", func() {

		It("will create if does not exist", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			clientConfig.EXPECT().
				ClientConfig().
				Return(nil, nil)

			opts := register.Options{
				ClusterName:     clusterName,
				Namespace:       namespace,
				RemoteNamespace: "remote",
				RemoteCtx:       remoteCtx,
			}

			saClient.EXPECT().
				GetServiceAccount(ctx, client.ObjectKey{
					Namespace: opts.RemoteNamespace,
					Name:      opts.ClusterName,
				}).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, ""))

			expected := &k8s_core_types.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: opts.RemoteNamespace,
					Name:      opts.ClusterName,
				},
			}

			saClient.EXPECT().
				CreateServiceAccount(ctx, expected).
				Return(nil)

			sa, err := clusterRegistrant.EnsureRemoteServiceAccount(ctx, clientConfig, opts)
			Expect(err).NotTo(HaveOccurred())
			Expect(sa).To(Equal(expected))
		})

		It("will delete remote service account", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			clientConfig.EXPECT().
				ClientConfig().
				Return(nil, nil)

			opts := register.Options{
				ClusterName:     clusterName,
				Namespace:       namespace,
				RemoteNamespace: "remote",
				RemoteCtx:       remoteCtx,
			}

			saClient.EXPECT().
				DeleteServiceAccount(ctx, client.ObjectKey{
					Namespace: opts.RemoteNamespace,
					Name:      opts.ClusterName,
				}).
				Return(nil)

			err := clusterRegistrant.DeleteRemoteServiceAccount(ctx, clientConfig, opts)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("CreateRemoteAccessToken", func() {

		It("can successfully upsert all roles, and cluster roles", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			// Set secret lookup opts to reduce testing time
			register.SecretLookupOpts = []retry.Option{
				retry.Delay(time.Nanosecond),
				retry.Attempts(2),
				retry.DelayType(retry.FixedDelay),
			}

			sa := saObjectKey()

			clientConfig.EXPECT().
				ClientConfig().
				Return(nil, nil)

			opts := register.RbacOptions{
				Options: register.Options{
					ClusterName: clusterName,
					Namespace:   namespace,
					RemoteCtx:   remoteCtx,
				},
				Roles: []*rbacv1.Role{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "r-1",
							Namespace: namespace,
						},
					},
				},
				ClusterRoles: []*rbacv1.ClusterRole{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "cr-1",
						},
					},
				},
				RoleBindings: []client.ObjectKey{
					{
						Namespace: "namespace",
						Name:      "rb-1",
					},
				},
				ClusterRoleBindings: []client.ObjectKey{
					{
						Namespace: "",
						Name:      "crb-1",
					},
				},
			}

			roleClient.EXPECT().
				UpsertRole(ctx, opts.Roles[0]).
				Return(nil)

			clusterRBACBinder.EXPECT().
				BindRoles(ctx, sa, append(opts.RoleBindings, client.ObjectKey{
					Name:      opts.Roles[0].GetName(),
					Namespace: opts.Roles[0].GetNamespace(),
				})).
				Return(nil)

			clusterRoleClient.EXPECT().
				UpsertClusterRole(ctx, opts.ClusterRoles[0]).
				Return(nil)

			clusterRBACBinder.EXPECT().
				BindClusterRoles(ctx, sa, append(opts.ClusterRoleBindings, client.ObjectKey{
					Name: opts.ClusterRoles[0].GetName(),
				})).Return(nil)

			token := "hello"
			saSecret := &k8s_core_types.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "sa-secret",
					Namespace: namespace,
				},
				Data: map[string][]byte{
					register.SecretTokenKey: []byte(token),
				},
			}
			expectedSa := &k8s_core_types.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sa.Name,
					Namespace: sa.Namespace,
				},
				Secrets: []k8s_core_types.ObjectReference{
					{
						Namespace: namespace,
						Name:      saSecret.GetName(),
					},
				},
			}

			saClient.EXPECT().
				GetServiceAccount(ctx, sa).
				Return(expectedSa, nil).
				Times(2)

			secretClient.EXPECT().
				GetSecret(ctx, client.ObjectKey{
					Namespace: saSecret.GetNamespace(),
					Name:      saSecret.GetName(),
				}).
				Return(nil, testErr)

			secretClient.EXPECT().
				GetSecret(ctx, client.ObjectKey{
					Namespace: saSecret.GetNamespace(),
					Name:      saSecret.GetName(),
				}).Return(saSecret, nil)

			returnedToken, err := clusterRegistrant.CreateRemoteAccessToken(ctx, clientConfig, sa, opts)
			Expect(err).NotTo(HaveOccurred())
			Expect(returnedToken).To(Equal(token))
		})

		It("will delete remote access resources", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			sa := client.ObjectKey{Name: clusterName, Namespace: "remote-namespace"}

			clientConfig.EXPECT().
				ClientConfig().
				Return(nil, nil)

			opts := register.RbacOptions{
				Options: register.Options{
					ClusterName:     clusterName,
					Namespace:       namespace,
					RemoteCtx:       remoteCtx,
					RemoteNamespace: "remote-namespace",
				},
				Roles: []*rbacv1.Role{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "r-1",
							Namespace: namespace,
						},
					},
				},
				ClusterRoles: []*rbacv1.ClusterRole{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "cr-1",
						},
					},
				},
				RoleBindings: []client.ObjectKey{
					{
						Namespace: "namespace",
						Name:      "rb-1",
					},
				},
				ClusterRoleBindings: []client.ObjectKey{
					{
						Namespace: "",
						Name:      "crb-1",
					},
				},
			}

			roleClient.EXPECT().
				DeleteRole(ctx, client.ObjectKey{Name: opts.Roles[0].Name, Namespace: opts.Roles[0].Namespace}).
				Return(nil)

			clusterRBACBinder.EXPECT().
				DeleteRoleBindings(ctx, sa, append(opts.RoleBindings, client.ObjectKey{
					Name:      opts.Roles[0].GetName(),
					Namespace: opts.Roles[0].GetNamespace(),
				})).
				Return(nil)

			clusterRoleClient.EXPECT().
				DeleteClusterRole(ctx, opts.ClusterRoles[0].Name).
				Return(nil)

			clusterRBACBinder.EXPECT().
				DeleteClusterRoleBindings(ctx, sa, append(opts.ClusterRoleBindings, client.ObjectKey{
					Name: opts.ClusterRoles[0].GetName(),
				})).Return(nil)

			err := clusterRegistrant.DeleteRemoteAccessResources(ctx, clientConfig, opts)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("RegisterClusterWithToken", func() {

		var (
			token = "token"
		)

		It("works", func() {

			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			opts := register.Options{
				ClusterName: clusterName,
				Namespace:   namespace,
				RemoteCtx:   remoteCtx,
			}

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

			clientConfig.EXPECT().
				ClientConfig().
				Return(restCfg, nil)

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

			kubeClusterClient.EXPECT().
				UpsertKubernetesCluster(ctx, &v1alpha1.KubernetesCluster{
					ObjectMeta: secret.ObjectMeta,
					Spec: v1alpha1.KubernetesClusterSpec{
						SecretName:    secret.Name,
						ClusterDomain: "cluster.local",
					},
				}).Return(nil)

			err = clusterRegistrant.RegisterClusterWithToken(ctx, restCfg, clientConfig, token, opts)

			Expect(err).NotTo(HaveOccurred())
		})

		It("works with provider info", func() {

			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			opts := register.Options{
				ClusterName: clusterName,
				Namespace:   namespace,
				RemoteCtx:   remoteCtx,
			}

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

			clientConfig.EXPECT().
				ClientConfig().
				Return(restCfg, nil)

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

			providerInfo := &v1alpha1.KubernetesClusterSpec_ProviderInfo{
				ProviderInfoType: &v1alpha1.KubernetesClusterSpec_ProviderInfo_Eks{
					Eks: &v1alpha1.KubernetesClusterSpec_Eks{
						Arn:       "arn",
						AccountId: "accountId",
						Region:    "region",
						Name:      "name",
					},
				},
			}

			kubeClusterClient.EXPECT().
				UpsertKubernetesCluster(ctx, &v1alpha1.KubernetesCluster{
					ObjectMeta: secret.ObjectMeta,
					Spec: v1alpha1.KubernetesClusterSpec{
						SecretName:    secret.Name,
						ClusterDomain: "cluster.local",
						ProviderInfo:  providerInfo,
					},
				}).Return(nil)

			err = clusterRegistrant.RegisterProviderClusterWithToken(ctx, restCfg, clientConfig, token, opts, providerInfo)

			Expect(err).NotTo(HaveOccurred())
		})

		It("can override local cluster domain", func() {
			apiServerAddress := "test-override"

			clusterRegistrant := register.NewClusterRegistrant(
				apiServerAddress,
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			opts := register.Options{
				ClusterName: clusterName,
				Namespace:   namespace,
				RemoteCtx:   "kind-test",
			}

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
					opts.RemoteCtx: {
						Cluster: clusterName,
					},
				},
				CurrentContext: opts.RemoteCtx,
			}

			clientConfig.EXPECT().
				ClientConfig().
				Return(restCfg, nil)

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
			overwrittenApiConfig.Clusters[clusterName].Server = fmt.Sprintf("https://%s:9080", apiServerAddress)
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

			kubeClusterClient.EXPECT().
				UpsertKubernetesCluster(ctx, &v1alpha1.KubernetesCluster{
					ObjectMeta: secret.ObjectMeta,
					Spec: v1alpha1.KubernetesClusterSpec{
						SecretName:    secret.Name,
						ClusterDomain: "cluster.local",
					},
				}).Return(nil)

			err = clusterRegistrant.RegisterClusterWithToken(ctx, restCfg, clientConfig, token, opts)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should deregister cluster", func() {
			clusterRegistrant := register.NewClusterRegistrant(
				"",
				clusterRbacBinderFactory,
				secretClient,
				secretClientFactory,
				nsClientFactory,
				saClientFactory,
				clusterRoleClientFactory,
				roleClientFactory,
				kubeClusterClientFactory,
			)

			opts := register.Options{
				ClusterName: clusterName,
				Namespace:   namespace,
				RemoteCtx:   remoteCtx,
			}

			restCfg := &rest.Config{
				Host: "mock-host",
			}

			secretClient.EXPECT().
				DeleteSecret(ctx, client.ObjectKey{Name: clusterName, Namespace: namespace}).
				Return(nil)

			kubeClusterClient.EXPECT().
				DeleteKubernetesCluster(ctx, client.ObjectKey{Name: clusterName, Namespace: namespace}).Return(nil)

			err := clusterRegistrant.DeregisterCluster(ctx, restCfg, opts)
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
