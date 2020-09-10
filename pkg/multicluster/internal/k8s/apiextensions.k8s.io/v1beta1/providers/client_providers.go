// Code generated by skv2. DO NOT EDIT.

package v1beta1

import (
	apiextensions_k8s_io_v1beta1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/apiextensions.k8s.io/v1beta1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for CustomResourceDefinitionClient from Clientset
func CustomResourceDefinitionClientFromClientsetProvider(clients apiextensions_k8s_io_v1beta1.Clientset) apiextensions_k8s_io_v1beta1.CustomResourceDefinitionClient {
	return clients.CustomResourceDefinitions()
}

// Provider for CustomResourceDefinition Client from Client
func CustomResourceDefinitionClientProvider(client client.Client) apiextensions_k8s_io_v1beta1.CustomResourceDefinitionClient {
	return apiextensions_k8s_io_v1beta1.NewCustomResourceDefinitionClient(client)
}

type CustomResourceDefinitionClientFactory func(client client.Client) apiextensions_k8s_io_v1beta1.CustomResourceDefinitionClient

func CustomResourceDefinitionClientFactoryProvider() CustomResourceDefinitionClientFactory {
	return CustomResourceDefinitionClientProvider
}

type CustomResourceDefinitionClientFromConfigFactory func(cfg *rest.Config) (apiextensions_k8s_io_v1beta1.CustomResourceDefinitionClient, error)

func CustomResourceDefinitionClientFromConfigFactoryProvider() CustomResourceDefinitionClientFromConfigFactory {
	return func(cfg *rest.Config) (apiextensions_k8s_io_v1beta1.CustomResourceDefinitionClient, error) {
		clients, err := apiextensions_k8s_io_v1beta1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.CustomResourceDefinitions(), nil
	}
}
