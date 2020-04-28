package cloud

import (
	"context"
	"fmt"

	container "cloud.google.com/go/container/apiv1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func NewGkeClient(ctx context.Context, projectId string) (GkeClient, error) {
	creds, err := google.FindDefaultCredentials(ctx, container.DefaultAuthScopes()...)
	if err != nil {
		return nil, err
	}
	return &gkeClient{
		creds:     creds,
		projectId: projectId,
	}, nil
}

type gkeClient struct {
	projectId string
	creds     *google.Credentials
}

func (g *gkeClient) Token(ctx context.Context) (*oauth2.Token, error) {
	return g.creds.TokenSource.Token()
}

func (g *gkeClient) ListClusters(ctx context.Context) ([]*containerpb.Cluster, error) {
	client, err := container.NewClusterManagerClient(ctx, option.WithCredentials(g.creds))
	if err != nil {
		return nil, err
	}
	resp, err := client.ListClusters(ctx, &containerpb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", g.projectId),
	})
	if err != nil {
		return nil, err
	}
	return resp.GetClusters(), nil
}
