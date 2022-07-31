package verifier

import (
	"context"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

type ClearableCache interface {
	ResetCache(ctx context.Context)
}


type Factory interface {
	// NewOutputResourceVerifier returns a new output verifier with the given options
	NewOutputResourceVerifier(
		ctx context.Context,
		discoveryClient discovery.DiscoveryInterface,
		options map[schema.GroupVersionKind]ServerVerifyOption,
	) OutputResourceVerifier
	// NewServerResourceVerifier returns a new server verifier with the given options
	NewServerResourceVerifier(
		ctx context.Context,
		options map[schema.GroupVersionKind]ServerVerifyOption,
	) ServerResourceVerifier
	// ResetAllCaches invalidates the caches of each verifier created by this factory
	ResetAllCaches(
		ctx context.Context,
	)
}

func NewVerifierFactory() Factory {
	return &factory{}
}

type factory struct {
	lock sync.RWMutex
	serverVerifiers []*verifier
	outputVerifiers []*outputVerifier
}

func (f *factory) ResetAllCaches(ctx context.Context) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	for _, cache := range f.serverVerifiers {
		cache.ResetCache(ctx)
	}
	for _, cache := range f.outputVerifiers {
		cache.ResetCache(ctx)
	}
}

func (f *factory) NewOutputResourceVerifier(
	ctx context.Context,
	discoveryClient discovery.DiscoveryInterface,
	options map[schema.GroupVersionKind]ServerVerifyOption,
) OutputResourceVerifier {
	f.lock.Lock()
	defer f.lock.Unlock()
	newVerifier := NewOutputVerifier(ctx, discoveryClient, options)
	f.outputVerifiers = append(f.outputVerifiers, newVerifier)
	return newVerifier
}

func (f *factory) NewServerResourceVerifier(
	ctx context.Context,
	options map[schema.GroupVersionKind]ServerVerifyOption,
) ServerResourceVerifier {
	f.lock.Lock()
	defer f.lock.Unlock()
	newVerifier := NewVerifier(ctx, options)
	f.serverVerifiers = append(f.serverVerifiers, newVerifier)
	return newVerifier
}
