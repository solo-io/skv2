package register_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	mock_k8s_core_clients "github.com/solo-io/skv2/pkg/generated/kubernetes/mocks/core/v1"
	mock_clientcmd "github.com/solo-io/skv2/pkg/generated/mocks/k8s/clientcmd"
	mock_auth "github.com/solo-io/skv2/pkg/multicluster/auth/mocks"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	mock_kubeconfig "github.com/solo-io/skv2/pkg/multicluster/kubeconfig/mocks"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Registrant", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		clusterAuthClient *mock_auth.MockClusterAuthorization
		clientset         *mock_k8s_core_clients.MockClientset
		secretClient      *mock_k8s_core_clients.MockSecretClient
		loader            *mock_kubeconfig.MockKubeLoader
		clientConfig      *mock_clientcmd.MockClientConfig

		remoteCfg, remoteCtx, clusterName, namespace = "cfg-path", "context", "cluster-name", "namespace"
		testErr                                      = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		clientset = mock_k8s_core_clients.NewMockClientset(ctrl)
		secretClient = mock_k8s_core_clients.NewMockSecretClient(ctrl)
		clusterAuthClient = mock_auth.NewMockClusterAuthorization(ctrl)
		loader = mock_kubeconfig.NewMockKubeLoader(ctrl)
		clientConfig = mock_clientcmd.NewMockClientConfig(ctrl)

		clientset.EXPECT().Secrets().Return(secretClient)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will fail if client config cannot be fetched", func() {
		clusterRegistrant := register.NewClusterRegistrant(loader, clusterAuthClient, clientset)

		loader.EXPECT().
			GetClientConfigForContext(remoteCfg, remoteCtx).
			Return(nil, testErr)

		err := clusterRegistrant.RegisterCluster(ctx, remoteCfg, remoteCtx, clusterName, namespace)
		Expect(err).To(HaveOccurred())
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})

	It("will fail if remote bearer token cannot be built", func() {
		clusterRegistrant := register.NewClusterRegistrant(loader, clusterAuthClient, clientset)

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
			BuildRemoteBearerToken(ctx, cfg, clusterName, namespace).
			Return("", testErr)

		err := clusterRegistrant.RegisterCluster(ctx, remoteCfg, remoteCtx, clusterName, namespace)
		Expect(err).To(HaveOccurred())
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})

	It("will create secret if not found", func() {
		clusterRegistrant := register.NewClusterRegistrant(loader, clusterAuthClient, clientset)

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
			BuildRemoteBearerToken(ctx, restCfg, clusterName, namespace).
			Return(token, nil)

		clientConfig.EXPECT().
			RawConfig().
			Return(apiCfg, nil)

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

		secretClient.EXPECT().CreateSecret(ctx, secret).Return(nil)

		err = clusterRegistrant.RegisterCluster(ctx, remoteCfg, remoteCtx, clusterName, namespace)
		Expect(err).NotTo(HaveOccurred())
	})

})
