package kubeconfig_test

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	mock_cloud "github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/clients/mocks"
	kubeconfig2 "github.com/solo-io/skv2/pkg/multicluster/discovery/cloud/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

var _ = Describe("Aws", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		eksClient *mock_cloud.MockEksClient

		testErr = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())

		eksClient = mock_cloud.NewMockEksClient(ctrl)
	})

	It("will fail if token cannot be found", func() {
		configBuilder := kubeconfig2.NewEksConfigBuilder(eksClient)

		cluster := &eks.Cluster{
			Name: aws.String("cluster-name"),
		}
		eksClient.EXPECT().
			Token(ctx, aws.StringValue(cluster.Name)).
			Return(token.Token{}, testErr)

		_, err := configBuilder.ConfigForCluster(ctx, cluster)
		Expect(err).To(HaveOccurred())
		Expect(err).To(testutils.HaveInErrorChain(testErr))
	})

	It("will Create ClientConfig if token can be found", func() {
		configBuilder := kubeconfig2.NewEksConfigBuilder(eksClient)
		caData := []byte("fake-ca-data")
		tok := token.Token{Token: "new-token-who-dis"}
		cluster := &eks.Cluster{
			Name: aws.String("cluster-name"),
			CertificateAuthority: &eks.Certificate{
				Data: aws.String(base64.StdEncoding.EncodeToString(caData)),
			},
			Endpoint: aws.String("fake-endpoint"),
		}
		eksClient.EXPECT().
			Token(ctx, aws.StringValue(cluster.Name)).
			Return(tok, nil)

		cfg, err := configBuilder.ConfigForCluster(ctx, cluster)
		Expect(err).NotTo(HaveOccurred())

		ca, err := base64.StdEncoding.DecodeString(aws.StringValue(cluster.CertificateAuthority.Data))
		Expect(err).NotTo(HaveOccurred())
		newCfg := kubeconfig.BuildRemoteCfg(
			&clientcmdapi.Cluster{
				Server:                   aws.StringValue(cluster.Endpoint),
				CertificateAuthorityData: ca,
			},
			&clientcmdapi.Context{
				Cluster: aws.StringValue(cluster.Name),
			},
			aws.StringValue(cluster.Name),
			tok.Token,
		)

		rawCfg, err := cfg.RawConfig()
		Expect(err).NotTo(HaveOccurred())
		Expect(rawCfg).To(Equal(newCfg))

	})

})
