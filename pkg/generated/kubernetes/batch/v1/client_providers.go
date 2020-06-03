// Code generated by skv2. DO NOT EDIT.

package v1

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.
*/

// Provider for JobClient from Client
func JobClientProvider(client client.Client) JobClient {
	return NewJobClient(client)
}

type JobClientFactory func(client client.Client) JobClient

func JobClientFactoryProvider() JobClientFactory {
	return JobClientProvider
}

type JobClientFromConfigFactory func(cfg *rest.Config) (JobClient, error)

func JobClientFromConfigFactoryProvider() JobClientFromConfigFactory {
	return func(cfg *rest.Config) (JobClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Jobs(), nil
	}
}
