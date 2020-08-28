package kubeconfig_test

import (
	"context"
	"encoding/base64"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	mock_cloud "github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/clients/mocks"
	kubeconfig2 "github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"golang.org/x/oauth2"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var _ = Describe("Gke", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		gkeClient *mock_cloud.MockGkeClient

		testErr = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		gkeClient = mock_cloud.NewMockGkeClient(ctrl)
	})

	It("will fail if token cannot be found", func() {
		configBuilder := kubeconfig2.NewGkeConfigBuilder(gkeClient)

		cluster := &containerpb.Cluster{}
		gkeClient.EXPECT().
			Token(ctx).
			Return(nil, testErr)

		_, err := configBuilder.ConfigForCluster(ctx, cluster)
		Expect(err).To(HaveOccurred())
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})

	It("will Create ClientConfig if token can be found", func() {
		configBuilder := kubeconfig2.NewGkeConfigBuilder(gkeClient)
		caData := []byte("fake-ca-data")
		token := &oauth2.Token{
			AccessToken: "new-token-who-dis",
		}
		cluster := &containerpb.Cluster{
			Endpoint: "fake-endpoint",
			Name:     "fake-name",
			MasterAuth: &containerpb.MasterAuth{
				ClusterCaCertificate: base64.StdEncoding.EncodeToString(caData),
			},
		}
		gkeClient.EXPECT().
			Token(ctx).
			Return(token, nil)

		cfg, err := configBuilder.ConfigForCluster(ctx, cluster)
		Expect(err).NotTo(HaveOccurred())

		ca, err := base64.StdEncoding.DecodeString(cluster.GetMasterAuth().GetClusterCaCertificate())
		Expect(err).NotTo(HaveOccurred())
		newCfg := kubeconfig.BuildRemoteCfg(
			&clientcmdapi.Cluster{
				Server:                   cluster.Endpoint,
				CertificateAuthorityData: ca,
			},
			&clientcmdapi.Context{
				Cluster: cluster.Name,
			},
			cluster.Name,
			token.AccessToken,
		)

		rawCfg, err := cfg.RawConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(rawCfg).To(Equal(newCfg))

	})

})
