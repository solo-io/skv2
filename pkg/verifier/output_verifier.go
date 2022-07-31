package verifier

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// ServerResourceVerifier verifies whether a given cluster server supports a given resource.
type OutputResourceVerifier interface {
	ClearableCache
	// VerifyServerResource verifies whether the API server for the given rest.Config supports the resource with the given GVK.
	// For the "local/management" cluster, set cluster to ""
	// Note that once a resource has been verified, the result will be cached for subsequent calls.
	// Returns true if the resource is registered, false otherwise.
	// an error will be returned if the ErrorIfNotPresent option is selected for the given GVK
	VerifyServerResource(cluster string, gvk schema.GroupVersionKind) (bool, error)
}

type outputVerifier struct {
	// used for logging
	ctx context.Context

	// set of per-resource-type verify options; default is Skip.
	options     map[schema.GroupVersionKind]ServerVerifyOption
	optionsLock sync.RWMutex

	// set when a resource is successfully verified for the first time.
	// future calls to VerifyServerResource will return quickly if the
	// resources have already been verified for the cluster.
	cachedVerificationResponses *cachedVerificationResponses

	discClient discovery.DiscoveryInterface
}

func NewOutputVerifier(
	ctx context.Context,
	discClient discovery.DiscoveryInterface,
	options map[schema.GroupVersionKind]ServerVerifyOption,
) *outputVerifier {
	if options == nil {
		options = map[schema.GroupVersionKind]ServerVerifyOption{}
	}
	return &outputVerifier{
		ctx:                         ctx,
		discClient:                  discClient,
		options:                     options,
		cachedVerificationResponses: newCachedVerificationResponses(),
	}
}

func (v *outputVerifier) ResetCache(_ context.Context) {
	v.cachedVerificationResponses.clearedCachedResponses()
}

// Verify whether the API server for the given rest.Config supports the resource with the given GVK.
func (v *outputVerifier) VerifyServerResource(cluster string, gvk schema.GroupVersionKind) (bool, error) {
	responseCached, resourceVerified := v.cachedVerificationResponses.getCachedResponse(cluster, gvk)
	if responseCached {
		return resourceVerified, nil
	}

	// get option for this gvk
	// defaults to skip
	v.optionsLock.RLock()
	verifyOption := v.options[gvk]
	v.optionsLock.RUnlock()

	if verifyOption == ServerVerifyOption_SkipVerify {
		return true, nil
	}

	var resourceRegistered bool
	if verifyFailed, err := verifyServerResource(v.discClient, gvk); err != nil {
		if verifyFailed {
			return false, eris.Wrap(err, "resource verify failed")
		}

		switch verifyOption {
		case ServerVerifyOption_ErrorIfNotPresent:
			return false, err
		case ServerVerifyOption_WarnIfNotPresent:
			contextutils.LoggerFrom(v.ctx).Warnf("%v not registered (fetch err: %v)", gvk, err)
			resourceRegistered = false
		case ServerVerifyOption_LogDebugIfNotPresent:
			contextutils.LoggerFrom(v.ctx).Debugf("%v not registered (fetch err: %v)", gvk, err)
			resourceRegistered = false
		case ServerVerifyOption_IgnoreIfNotPresent:
			resourceRegistered = false
		}
	} else {
		resourceRegistered = true
	}

	// update cache
	v.cachedVerificationResponses.setCachedResponse(cluster, gvk, resourceRegistered)

	return resourceRegistered, nil
}
