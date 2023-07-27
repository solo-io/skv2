// Code generated by skv2. DO NOT EDIT.

package v1



import (
    things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

    "k8s.io/client-go/rest"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for PaintClient from Clientset
func PaintClientFromClientsetProvider(clients things_test_io_v1.Clientset) things_test_io_v1.PaintClient {
    return clients.Paints()
}

// Provider for Paint Client from Client
func PaintClientProvider(client client.Client) things_test_io_v1.PaintClient {
    return things_test_io_v1.NewPaintClient(client)
}

type PaintClientFactory func(client client.Client) things_test_io_v1.PaintClient

func PaintClientFactoryProvider() PaintClientFactory {
    return PaintClientProvider
}

type PaintClientFromConfigFactory func(cfg *rest.Config) (things_test_io_v1.PaintClient, error)

func PaintClientFromConfigFactoryProvider() PaintClientFromConfigFactory {
    return func(cfg *rest.Config) (things_test_io_v1.PaintClient, error) {
        clients, err := things_test_io_v1.NewClientsetFromConfig(cfg)
        if err != nil {
            return nil, err
        }
        return clients.Paints(), nil
    }
}

// Provider for ClusterResourceClient from Clientset
func ClusterResourceClientFromClientsetProvider(clients things_test_io_v1.Clientset) things_test_io_v1.ClusterResourceClient {
    return clients.ClusterResources()
}

// Provider for ClusterResource Client from Client
func ClusterResourceClientProvider(client client.Client) things_test_io_v1.ClusterResourceClient {
    return things_test_io_v1.NewClusterResourceClient(client)
}

type ClusterResourceClientFactory func(client client.Client) things_test_io_v1.ClusterResourceClient

func ClusterResourceClientFactoryProvider() ClusterResourceClientFactory {
    return ClusterResourceClientProvider
}

type ClusterResourceClientFromConfigFactory func(cfg *rest.Config) (things_test_io_v1.ClusterResourceClient, error)

func ClusterResourceClientFromConfigFactoryProvider() ClusterResourceClientFromConfigFactory {
    return func(cfg *rest.Config) (things_test_io_v1.ClusterResourceClient, error) {
        clients, err := things_test_io_v1.NewClientsetFromConfig(cfg)
        if err != nil {
            return nil, err
        }
        return clients.ClusterResources(), nil
    }
}