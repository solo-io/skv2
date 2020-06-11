package main

import (
	"log"
	"path/filepath"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// K8s config
//go:generate  mockgen -package mock_clientcmd -destination ./mocks/k8s/clientcmd/config.go k8s.io/client-go/tools/clientcmd ClientConfig

// K8s clients
//go:generate mockgen -package mock_k8s_core_clients -destination ./kubernetes/mocks/core/v1/clients.go github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1 Clientset,ServiceClient,PodClient,NamespaceClient,NodeClient,ServiceAccountClient,SecretClient,ConfigMapClient
//go:generate mockgen -package mock_k8s_rbac_clients -destination ./kubernetes/mocks/rbac.authorization.k8s.io/v1/clients.go github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1 Clientset,ClusterRoleBindingClient,RoleBindingClient,ClusterRoleClient,RoleClient

const (
	generatedPackageName = "pkg/generated"
)

var (
	kubeGeneratedPackage = filepath.Join(generatedPackageName, "kubernetes")
)

func main() {
	log.Println("starting kube client generation")
	skv2Cmd := codegen.Command{
		Groups: []model.Group{
			{
				GroupVersion: schema.GroupVersion{
					Group:   "core",
					Version: "v1",
				},
				Module: "k8s.io/api",
				Resources: []model.Resource{
					{
						Kind: "Secret",
					},
					{
						Kind: "ServiceAccount",
					},
					{
						Kind: "ConfigMap",
					},
					{
						Kind: "Service",
					},
					{
						Kind: "Pod",
					},
					{
						Kind:          "Namespace",
						ClusterScoped: true,
					},
					{
						Kind:          "Node",
						ClusterScoped: true,
					},
				},
				RenderController:      true,
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/core/v1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "apps",
					Version: "v1",
				},
				Module: "k8s.io/api",
				Resources: []model.Resource{
					{
						Kind: "Deployment",
					},
					{
						Kind: "ReplicaSet",
					},
					{
						Kind: "DaemonSet",
					},
				},
				RenderController:      true,
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/apps/v1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "batch",
					Version: "v1",
				},
				Module: "k8s.io/api",
				Resources: []model.Resource{
					{
						Kind: "Job",
					},
				},
				RenderController:      true,
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/batch/v1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "admissionregistration.k8s.io",
					Version: "v1",
				},
				Module: "k8s.io/apiextensions-apiserver",
				Resources: []model.Resource{
					{
						Kind: "ValidatingWebhookConfiguration",
					},
				},
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/admissionregistration/v1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "rbac.authorization.k8s.io",
					Version: "v1",
				},
				Resources: []model.Resource{
					{
						Kind: "Role",
					},
					{
						Kind: "RoleBinding",
					},
					{
						Kind:          "ClusterRole",
						ClusterScoped: true,
					},
					{
						Kind:          "ClusterRoleBinding",
						ClusterScoped: true,
					},
				},
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/rbac/v1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "certificates.k8s.io",
					Version: "v1beta1",
				},
				Resources: []model.Resource{
					{
						Kind: "CertificateSigningRequest",
					},
				},
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/certificates/v1beta1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "apiextensions.k8s.io",
					Version: "v1beta1",
				},
				Module: "k8s.io/apiextensions-apiserver",
				Resources: []model.Resource{
					{
						Kind:          "CustomResourceDefinition",
						ClusterScoped: true,
					},
				},
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
				CustomTemplates:       contrib.AllCustomTemplates,
				ApiRoot:               kubeGeneratedPackage,
			},
		},
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating kube clients\n")
}
