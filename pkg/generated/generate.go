package main

import (
	"log"
	"path/filepath"

	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//go:generate go run generate.go

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
				},
				RenderController:      true,
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/apps/v1",
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
				ApiRoot:               kubeGeneratedPackage,
			},
			{
				GroupVersion: schema.GroupVersion{
					Group:   "certificates.k8s.io/v1beta1",
					Version: "v1beta1",
				},
				Resources: []model.Resource{
					{
						Kind: "CertificateSigningRequest",
					},
				},
				RenderClients:         true,
				CustomTypesImportPath: "k8s.io/api/certificates/v1beta1",
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
				ApiRoot:               kubeGeneratedPackage,
			},
		},
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating kube clients\n")
}
