package verifier

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

type ClearableCache interface {
	ResetCache(ctx context.Context)
}

type Factory interface {
	// NewVerifier returns a new verifier with the given options
	NewOutputResourceVerifier(
		ctx context.Context,
		discoveryClient discovery.DiscoveryInterface,
		options map[schema.GroupVersionKind]ServerVerifyOption,
	) OutputResourceVerifier

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
	serverVerifiers []*verifier
	outputVerifiers []*outputVerifier
}

func (f *factory) ResetAllCaches(ctx context.Context) {
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
	newVerifier := NewOutputVerifier(ctx, discoveryClient, options)
	f.outputVerifiers = append(f.outputVerifiers, newVerifier)
	return newVerifier
}

func (f *factory) NewServerResourceVerifier(
	ctx context.Context,
	options map[schema.GroupVersionKind]ServerVerifyOption,
) ServerResourceVerifier {
	newVerifier := NewServerResourceVerifier(ctx, options)
	f.serverVerifiers = append(f.serverVerifiers, newVerifier)
	return newVerifier
}
