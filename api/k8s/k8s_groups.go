package k8s

import (
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

const (
	k8sApiRoot = "pkg/multicluster/internal/k8s"
	k8sModule  = "k8s.io/api"
)

func Groups() []model.Group {
	groups := []model.Group{
		{
			GroupVersion: corev1.SchemeGroupVersion,
			Module:       k8sModule,
			Resources: []model.Resource{
				{
					Kind: "Secret",
				},
				{
					Kind: "ServiceAccount",
				},
				{
					Kind:          "Namespace",
					ClusterScoped: true,
				},
			},
			CustomTypesImportPath: "k8s.io/api/core/v1",
			ApiRoot:               k8sApiRoot + "/core",
		},
		{
			GroupVersion: admissionregistrationv1.SchemeGroupVersion,
			Module:       "k8s.io/apiextensions-apiserver",
			Resources: []model.Resource{
				{
					Kind: "ValidatingWebhookConfiguration",
				},
			},
			CustomTypesImportPath: "k8s.io/api/admissionregistration/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: rbacv1.SchemeGroupVersion,
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
			CustomTypesImportPath: "k8s.io/api/rbac/v1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: certificatesv1beta1.SchemeGroupVersion,
			Resources: []model.Resource{
				{
					Kind: "CertificateSigningRequest",
				},
			},
			CustomTypesImportPath: "k8s.io/api/certificates/v1beta1",
			ApiRoot:               k8sApiRoot,
		},
		{
			GroupVersion: apiextensionsv1beta1.SchemeGroupVersion,
			Module:       "k8s.io/apiextensions-apiserver",
			Resources: []model.Resource{
				{
					Kind:          "CustomResourceDefinition",
					ClusterScoped: true,
				},
			},
			CustomTypesImportPath: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
			ApiRoot:               k8sApiRoot,
		},
	}
	for i, group := range groups {
		group.RenderClients = true
		group.RenderController = true
		group.MockgenDirective = true
		group.CustomTemplates = contrib.AllGroupCustomTemplates
		groups[i] = group
	}
	return groups
}
