package cloud

import (
	"github.com/aws/aws-sdk-go/service/eks"
)

type AwsClient interface {

	DescribeCluster(name string) (*eks.Cluster, error)
	ListClusters(input *eks.ListClustersInput) (*eks.ListClustersOutput, error)

}

