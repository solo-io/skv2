// Code generated by skv2. DO NOT EDIT.

package v1

import (
	rbac_authorization_k8s_io_v1 "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/rbac.authorization.k8s.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for Role Client from Client
func RoleClientProvider(client client.Client) rbac_authorization_k8s_io_v1.RoleClient {
	return rbac_authorization_k8s_io_v1.NewRoleClient(client)
}

type RoleClientFactory func(client client.Client) rbac_authorization_k8s_io_v1.RoleClient

func RoleClientFactoryProvider() RoleClientFactory {
	return RoleClientProvider
}

type RoleClientFromConfigFactory func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.RoleClient, error)

func RoleClientFromConfigFactoryProvider() RoleClientFromConfigFactory {
	return func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.RoleClient, error) {
		clients, err := rbac_authorization_k8s_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Roles(), nil
	}
}

// Provider for RoleBinding Client from Client
func RoleBindingClientProvider(client client.Client) rbac_authorization_k8s_io_v1.RoleBindingClient {
	return rbac_authorization_k8s_io_v1.NewRoleBindingClient(client)
}

type RoleBindingClientFactory func(client client.Client) rbac_authorization_k8s_io_v1.RoleBindingClient

func RoleBindingClientFactoryProvider() RoleBindingClientFactory {
	return RoleBindingClientProvider
}

type RoleBindingClientFromConfigFactory func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.RoleBindingClient, error)

func RoleBindingClientFromConfigFactoryProvider() RoleBindingClientFromConfigFactory {
	return func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.RoleBindingClient, error) {
		clients, err := rbac_authorization_k8s_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RoleBindings(), nil
	}
}

// Provider for ClusterRole Client from Client
func ClusterRoleClientProvider(client client.Client) rbac_authorization_k8s_io_v1.ClusterRoleClient {
	return rbac_authorization_k8s_io_v1.NewClusterRoleClient(client)
}

type ClusterRoleClientFactory func(client client.Client) rbac_authorization_k8s_io_v1.ClusterRoleClient

func ClusterRoleClientFactoryProvider() ClusterRoleClientFactory {
	return ClusterRoleClientProvider
}

type ClusterRoleClientFromConfigFactory func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.ClusterRoleClient, error)

func ClusterRoleClientFromConfigFactoryProvider() ClusterRoleClientFromConfigFactory {
	return func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.ClusterRoleClient, error) {
		clients, err := rbac_authorization_k8s_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ClusterRoles(), nil
	}
}

// Provider for ClusterRoleBinding Client from Client
func ClusterRoleBindingClientProvider(client client.Client) rbac_authorization_k8s_io_v1.ClusterRoleBindingClient {
	return rbac_authorization_k8s_io_v1.NewClusterRoleBindingClient(client)
}

type ClusterRoleBindingClientFactory func(client client.Client) rbac_authorization_k8s_io_v1.ClusterRoleBindingClient

func ClusterRoleBindingClientFactoryProvider() ClusterRoleBindingClientFactory {
	return ClusterRoleBindingClientProvider
}

type ClusterRoleBindingClientFromConfigFactory func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.ClusterRoleBindingClient, error)

func ClusterRoleBindingClientFromConfigFactoryProvider() ClusterRoleBindingClientFromConfigFactory {
	return func(cfg *rest.Config) (rbac_authorization_k8s_io_v1.ClusterRoleBindingClient, error) {
		clients, err := rbac_authorization_k8s_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ClusterRoleBindings(), nil
	}
}
