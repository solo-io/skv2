package clients

import (
	"context"

	"github.com/aws/aws-sdk-go/service/eks"
	"golang.org/x/oauth2"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

type EksClient interface {
	DescribeCluster(ctx context.Context, name string) (*eks.Cluster, error)
	ListClusters(ctx context.Context, fn func(*eks.ListClustersOutput)) error
	Token(ctx context.Context, name string) (token.Token, error)
}

type GkeClient interface {
	Token(ctx context.Context) (*oauth2.Token, error)
	ListClusters(ctx context.Context) ([]*containerpb.Cluster, error)
}
