package auth

import (
	"context"
	"time"

	"github.com/avast/retry-go"
	"github.com/rotisserie/eris"
	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	k8s_core_types "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// visible for testing
	SecretTokenKey = "token"
)

var (
	// exponential backoff retry with an initial period of 0.1s for 7 iterations, which will mean a cumulative retry period of ~6s
	// visible for testing
	SecretLookupOpts = []retry.Option{
		retry.Delay(time.Millisecond * 100),
		retry.Attempts(7),
		retry.DelayType(retry.BackOffDelay),
	}
)

func NewRemoteAuthorityConfigCreator(
	secretClient k8s_core_v1.SecretClient,
	serviceAccountClient k8s_core_v1.ServiceAccountClient,
) RemoteAuthorityConfigCreator {
	return &remoteAuthorityConfigCreator{
		serviceAccountClient: serviceAccountClient,
		secretClient:         secretClient,
	}
}

type remoteAuthorityConfigCreator struct {
	secretClient         k8s_core_v1.SecretClient
	serviceAccountClient k8s_core_v1.ServiceAccountClient
}

func (r *remoteAuthorityConfigCreator) ConfigFromRemoteServiceAccount(
	ctx context.Context,
	targetClusterCfg *rest.Config,
	name, namespace string,
) (*rest.Config, error) {

	tokenSecret, err := r.waitForSecret(ctx, name, namespace)
	if err != nil {
		return nil, SecretNotReady(err)
	}

	serviceAccountToken, ok := tokenSecret.Data[SecretTokenKey]
	if !ok {
		return nil, MalformedSecret
	}

	// make a copy of the config we were handed, with all user credentials removed
	// https://github.com/kubernetes/client-go/blob/9bbcc2938d41daa40d3080a1b6524afbe4e27bd9/rest/config.go#L542
	newCfg := rest.AnonymousClientConfig(targetClusterCfg)

	// authorize ourselves as the service account we were given
	newCfg.BearerToken = string(serviceAccountToken)

	return newCfg, nil
}

func (r *remoteAuthorityConfigCreator) waitForSecret(
	ctx context.Context,
	name, namespace string,
) (*k8s_core_types.Secret, error) {

	var foundSecret *k8s_core_types.Secret
	if err := retry.Do(func() error {
		serviceAccount, err := r.serviceAccountClient.GetServiceAccount(
			ctx,
			client.ObjectKey{Name: name, Namespace: namespace},
		)
		if err != nil {
			return err
		}

		if len(serviceAccount.Secrets) == 0 {
			return eris.Errorf("service account %s.%s does not have a token secret associated with it", name, namespace)
		}

		secretName := serviceAccount.Secrets[0].Name

		secret, err := r.secretClient.GetSecret(ctx, client.ObjectKey{Name: secretName, Namespace: namespace})
		if err != nil {
			return err
		}

		foundSecret = secret
		return nil
	}, SecretLookupOpts...); err != nil {
		return nil, err
	}

	return foundSecret, nil
}
