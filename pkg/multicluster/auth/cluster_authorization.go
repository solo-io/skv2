package auth

import (
	"context"

	k8s_rbac_types "k8s.io/api/rbac/v1"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

var (
	// visible for testing
	ServiceAccountRoles = []*k8s_rbac_types.ClusterRole{{
		ObjectMeta: k8s_meta_types.ObjectMeta{Name: "cluster-admin"},
	}}
)

type clusterAuthorization struct {
	configCreator          RemoteAuthorityConfigCreator
	remoteAuthorityManager RemoteAuthorityManager
}

func NewClusterAuthorization(
	configCreator RemoteAuthorityConfigCreator,
	remoteAuthorityManager RemoteAuthorityManager) ClusterAuthorization {
	return &clusterAuthorization{configCreator, remoteAuthorityManager}
}

func (c *clusterAuthorization) BuildRemoteBearerToken(
	ctx context.Context,
	targetClusterCfg *rest.Config,
	name, namespace string,
) (bearerToken string, err error) {
	_, err = c.remoteAuthorityManager.ApplyRemoteServiceAccount(ctx, name, namespace, ServiceAccountRoles)
	if err != nil {
		return "", err
	}

	saConfig, err := c.configCreator.ConfigFromRemoteServiceAccount(ctx, targetClusterCfg, name, namespace)
	if err != nil {
		return "", err
	}

	// we only want the bearer token for that service account
	return saConfig.BearerToken, nil
}
