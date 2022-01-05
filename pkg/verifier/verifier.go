package verifier

import (
	"context"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

// ServerVerifyOption is used to specify one of several options on
// whether and how to verify the resource's availabiltiy on the API server
type ServerVerifyOption int

const (
	// skip verifying whether a resource is  for a kind before reading it from the API server
	ServerVerifyOption_SkipVerify ServerVerifyOption = iota

	// return an error if the resource does not exist for a kind before reading it from the API server
	ServerVerifyOption_ErrorIfNotPresent

	// log a warning (and continue) if the resource does not exist for a kind before reading it from the API server
	// the reconcile loop will not be started if the server resource is not supported.
	ServerVerifyOption_WarnIfNotPresent

	// write a debug log (and continue) if the resource does not exist for a kind before reading it from the API server
	// the reconcile loop will not be started if the server resource is not supported.
	ServerVerifyOption_LogDebugIfNotPresent

	// ignore error (and continue) if the resource does not exist for a kind before reading it from the API server.
	// the reconcile loop will not be started if the server resource is not supported.
	ServerVerifyOption_IgnoreIfNotPresent
)

// ServerResourceVerifier verifies whether a given cluster server supports a given resource.
type ServerResourceVerifier interface {
	// VerifyServerResource verifies whether the API server for the given rest.Config supports the resource with the given GVK.
	// For the "local/management" cluster, set cluster to ""
	// Note that once a resource has been verified, the result will be cached for subsequent calls.
	// Returns true if the resource is registered, false otherwise.
	// an error will be returned if the ErrorIfNotPresent option is selected for the given GVK
	VerifyServerResource(cluster string, cfg *rest.Config, gvk schema.GroupVersionKind) (bool, error)
}

type verifier struct {
	// used for logging
	ctx context.Context

	// set of per-resource-type verify options; default is Skip.
	options     map[schema.GroupVersionKind]ServerVerifyOption
	optionsLock sync.RWMutex

	// set when a resource is successfully verified for the first time.
	// future calls to VerifyServerResource will return quickly if the
	// resources have already been verified for the cluster.
	cachedVerificationResponses *cachedVerificationResponses
}

type cachedVerificationResponses struct {
	verifiedClusterResources map[string]map[schema.GroupVersionKind]bool
	lock                     sync.RWMutex
}

func newCachedVerificationResponses() *cachedVerificationResponses {
	return &cachedVerificationResponses{verifiedClusterResources: map[string]map[schema.GroupVersionKind]bool{}}
}

// returns <is response cached>, <is resource supported>
func (c *cachedVerificationResponses) getCachedResponse(cluster string, gvk schema.GroupVersionKind) (bool, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	verifiedResources, ok := c.verifiedClusterResources[cluster]
	if !ok {
		return false, false
	}
	resourceRegistered, ok := verifiedResources[gvk]
	if !ok {
		return false, false
	}
	return true, resourceRegistered
}

// returns <is response cached>, <is resource supported>
func (c *cachedVerificationResponses) setCachedResponse(cluster string, gvk schema.GroupVersionKind, registered bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	verifiedResources, ok := c.verifiedClusterResources[cluster]
	if !ok {
		verifiedResources = map[schema.GroupVersionKind]bool{}
	}
	verifiedResources[gvk] = registered
}

func NewVerifier(
	ctx context.Context,
	options map[schema.GroupVersionKind]ServerVerifyOption,
) *verifier {
	if options == nil {
		options = map[schema.GroupVersionKind]ServerVerifyOption{}
	}
	return &verifier{
		ctx:                         ctx,
		options:                     options,
		cachedVerificationResponses: newCachedVerificationResponses(),
	}
}

// Verify whether the API server for the given rest.Config supports the resource with the given GVK.
func (v *verifier) VerifyServerResource(cluster string, cfg *rest.Config, gvk schema.GroupVersionKind) (bool, error) {
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
	if verifyFailed, err := verifyServerResource(cfg, gvk); err != nil {
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

// The boolean return return argument indicates whether a resulting error was caused by
// a failure to run the verification verify.
func verifyServerResource(
	cfg *rest.Config,
	gvk schema.GroupVersionKind,
) (bool, error) {
	disc, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return true, err
	}
	rss, err := disc.ServerResourcesForGroupVersion(gvk.GroupVersion().String())
	if err != nil {
		verifyFailed := !apierrors.IsNotFound(err)
		return verifyFailed, err
	}
	for _, resource := range rss.APIResources {
		if resource.Kind == gvk.Kind {
			// success
			return false, nil
		}
	}
	return false, eris.Errorf("resource %v not supported by API Server", gvk.String())
}
