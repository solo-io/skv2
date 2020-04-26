package register

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/aws/aws-sdk-go/service/eks"

	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)


type awsClusterConfigBuilder struct {
}

func (a *awsClusterConfigBuilder) ConfigForCluster(cluster *eks.Cluster) (clientcmd.ClientConfig, error) {
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return nil, err
	}
	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(cluster.Name),
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		return nil, err
	}
	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(cluster.CertificateAuthority.Data))
	if err != nil {
		return nil, err
	}

	cfg := buildRemoteCfg(
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

	return clientcmd.NewDefaultClientConfig(cfg, &clientcmd.ConfigOverrides{}), nil
}

func buildSession() {
	name := "wonderful-outfit-1583362361"
	region := "us-east-2"
	sess := session.Must(session.NewSession(&aws.Config{

		Region: aws.String(region),
	}))
	eksSvc := eks.New(sess)

	input := &eks.DescribeClusterInput{
		Name: aws.String(name),
	}

	eksSvc.ListClusters()
	result, err := eksSvc.DescribeCluster(input)
	if err != nil {
		log.Fatalf("Error calling DescribeCluster: %v", err)
	}
}
