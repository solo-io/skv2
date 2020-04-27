package kubeconfig

import (
	"k8s.io/client-go/tools/clientcmd/api"
)

func BuildRemoteCfg(
	remoteCluster *api.Cluster,
	remoteCtx *api.Context,
	clusterName, token string,
) api.Config {
	return api.Config{
		Kind:        "Secret",
		APIVersion:  "kubernetes_core",
		Preferences: api.Preferences{},
		Clusters: map[string]*api.Cluster{
			clusterName: remoteCluster,
		},
		AuthInfos: map[string]*api.AuthInfo{
			clusterName: {
				Token: token,
			},
		},
		Contexts: map[string]*api.Context{
			clusterName: {
				LocationOfOrigin: remoteCtx.LocationOfOrigin,
				Cluster:          clusterName,
				AuthInfo:         clusterName,
				Namespace:        remoteCtx.Namespace,
				Extensions:       remoteCtx.Extensions,
			},
		},
		CurrentContext: clusterName,
	}
}

