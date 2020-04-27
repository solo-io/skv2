package auth

import (
	"context"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	rbac_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

func NewClusterAuthorizationFactory() ClusterAuthorizationFactory {
	return func(cfg clientcmd.ClientConfig) (ClusterAuthorization, error) {
		restCfg, err := cfg.ClientConfig()
		if err != nil {
			return nil, err
		}
		v1Clientset, err := k8s_core_v1.NewClientsetFromConfig(restCfg)
		if err != nil {
			return nil, err
		}
		rbacClientset, err := rbac_v1.NewClientsetFromConfig(restCfg)
		if err != nil {
			return nil, err
		}
		cfgCreator := NewRemoteAuthorityConfigCreator(v1Clientset.Secrets(), v1Clientset.ServiceAccounts())
		rbacBinder := NewRbacBinder(rbacClientset)
		authorityManager := NewRemoteAuthorityManager(v1Clientset.ServiceAccounts(), rbacBinder)
		return NewClusterAuthorization(cfgCreator, authorityManager), nil
	}
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
