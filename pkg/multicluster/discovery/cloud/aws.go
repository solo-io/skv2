package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func NewEksClient(region string, creds *credentials.Credentials) (EksClient, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})
	if err != nil {
		return nil, err
	}
	return &awsClient{
		sess: sess,
	}, nil
}

// Construct client from preconfigured session.
func NewEksClientFromSession(sess *session.Session) EksClient {
	return &awsClient{
		sess: sess,
	}
}

type awsClient struct {
	sess *session.Session
}

func (a *awsClient) Session() *session.Session {
	return a.sess
}

func (a *awsClient) DescribeCluster(ctx context.Context, name string) (*eks.Cluster, error) {
	eksSvc := eks.New(a.sess)
	input := &eks.DescribeClusterInput{
		Name: aws.String(name),
	}
	resp, err := eksSvc.DescribeClusterWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return resp.Cluster, nil
}

func (a *awsClient) ListClusters(ctx context.Context, fn func(*eks.ListClustersOutput)) error {
	eksSvc := eks.New(a.sess)
	return eksSvc.ListClustersPagesWithContext(ctx, &eks.ListClustersInput{}, func(page *eks.ListClustersOutput, isLast bool) bool {
		if page == nil {
			return !isLast
		}
		fn(page)
		return !isLast
	})
}

func (a *awsClient) Token(ctx context.Context, name string) (token.Token, error) {
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return token.Token{}, err
	}
	opts := &token.GetTokenOptions{
		ClusterID: name,
		Session:   a.sess,
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		return token.Token{}, err
	}
	return tok, nil
}
