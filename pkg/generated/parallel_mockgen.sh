#!/bin/sh

# K8s config
mockgen -package mock_clientcmd -destination ./mocks/k8s/clientcmd/config.go k8s.io/client-go/tools/clientcmd ClientConfig &

# K8s clients
mockgen -package mock_k8s_core_clients -destination ./kubernetes/mocks/core/v1/clients.go github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1 Clientset,ServiceClient,PodClient,NamespaceClient,NodeClient,ServiceAccountClient,SecretClient,ConfigMapClient &
mockgen -package mock_k8s_rbac_clients -destination ./kubernetes/mocks/rbac.authorization.k8s.io/v1/clients.go github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1 Clientset,ClusterRoleBindingClient &

wait
