package client

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reader[T client.Object, L client.ObjectList] interface {

	// Get retrieves an obj for the given object key from the Kubernetes Cluster.
	// obj must be a struct pointer so that obj can be updated with the response
	// returned by the Server.
	Get(ctx context.Context, key client.ObjectKey) (T, error)

	// List retrieves list of objects for a given namespace and list options. On a
	// successful call, Items field in the list will be populated with the
	// result returned from the server.
	List(ctx context.Context, opts ...client.ListOption) (L, error)
}

type TransitionFunction[T client.Object] func(existing, desired T) error

type Writer[T client.Object] interface {
	// Create saves the object obj in the Kubernetes cluster.
	Create(ctx context.Context, obj T, opts ...client.CreateOption) error

	// Delete deletes the given obj from Kubernetes cluster.
	Delete(ctx context.Context, obj T, opts ...client.DeleteOption) error

	// Update updates the given obj in the Kubernetes cluster. obj must be a
	// struct pointer so that obj can be updated with the content returned by the Server.
	Update(ctx context.Context, obj T, opts ...client.UpdateOption) error

	// Patch patches the given obj in the Kubernetes cluster. obj must be a
	// struct pointer so that obj can be updated with the content returned by the Server.
	Patch(ctx context.Context, obj T, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all objects of the given type matching the given options.
	DeleteAllOf(ctx context.Context, obj T, opts ...client.DeleteAllOfOption) error

	// Create or Update the passed in object.
	Upsert(ctx context.Context, obj T, transitionFuncs ...TransitionFunction[T]) error
}

// StatusClient knows how to create a client which can update status subresource
// for kubernetes objects.
type StatusClient[T client.Object] interface {
	Status() StatusWriter[T]
}

// StatusWriter knows how to update status subresource of a Kubernetes object.
type StatusWriter[T client.Object] interface {
	// Update updates the fields corresponding to the status subresource for the
	// given obj. obj must be a struct pointer so that obj can be updated
	// with the content returned by the Server.
	Update(ctx context.Context, obj T, opts ...client.UpdateOption) error

	// Patch patches the given object's subresource. obj must be a struct
	// pointer so that obj can be updated with the content returned by the
	// Server.
	Patch(ctx context.Context, obj T, patch client.Patch, opts ...client.PatchOption) error
}

type GenericClient[T client.Object, L client.ObjectList] interface {
	Reader[T, L]
	Writer[T]
	StatusClient[T]

	// Scheme returns the scheme this client is using.
	Scheme() *runtime.Scheme
	// RESTMapper returns the rest this client is using.
	RESTMapper() meta.RESTMapper
}

func NewGenericClient[T client.Object, L client.ObjectList](
	cli client.Client,
	t T,
	l L,
) GenericClient[T, L] {
	return &genericClient[T, L]{
		genericClient: cli,
		l: l,
		t: t,
	}
}

func NewGenericClientFromConfig[T client.Object, L client.ObjectList](
	cfg *rest.Config,
	t T,
	l L,
) (GenericClient[T, L], error) {
	cli, err := client.New(cfg, client.Options{})
	if err != nil {
		return nil, err
	}
	return NewGenericClient[T, L](cli, t, l), nil
}

type genericClient[T client.Object, L client.ObjectList] struct {
	t             T
	l             L
	genericClient client.Client
}

func (g *genericClient[T, L]) Get(ctx context.Context, key client.ObjectKey) (T, error) {
	obj := g.t.DeepCopyObject().(T)
	if err := g.genericClient.Get(ctx, key, obj); err != nil {
		return obj, err
	}
	return obj, nil
}

func (g *genericClient[T, L]) List(ctx context.Context, opts ...client.ListOption) (L, error) {
	list := g.l.DeepCopyObject().(L)
	if err := g.genericClient.List(ctx, list, opts...); err != nil {
		return list, err
	}
	return list, nil
}

func (g *genericClient[T, L]) Create(ctx context.Context, obj T, opts ...client.CreateOption) error {
	return g.genericClient.Create(ctx, obj, opts...)
}

func (g *genericClient[T, L]) Delete(ctx context.Context, obj T, opts ...client.DeleteOption) error {
	return g.genericClient.Delete(ctx, obj, opts...)
}

func (g *genericClient[T, L]) Update(ctx context.Context, obj T, opts ...client.UpdateOption) error {
	return g.genericClient.Update(ctx, obj, opts...)
}

func (g *genericClient[T, L]) Patch(ctx context.Context, obj T, patch client.Patch, opts ...client.PatchOption) error {
	return g.genericClient.Patch(ctx, obj, patch, opts...)
}

func (g *genericClient[T, L]) DeleteAllOf(ctx context.Context, obj T, opts ...client.DeleteAllOfOption) error {
	return g.genericClient.DeleteAllOf(ctx, obj, opts...)
}

func (g *genericClient[T, L]) Upsert(ctx context.Context, obj T, transitionFuncs ...TransitionFunction[T]) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(T), desired.(T)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, g.genericClient, obj, genericTxFunc)
	return err
}

func (g *genericClient[T, L]) Status() StatusWriter[T] {
	return &statusWriter[T]{
		statusWriter: g.genericClient.Status(),
	}
}

func (g *genericClient[T, L]) Scheme() *runtime.Scheme {
	return g.genericClient.Scheme()
}

func (g *genericClient[T, L]) RESTMapper() meta.RESTMapper {
	return g.genericClient.RESTMapper()
}

type statusWriter[T client.Object] struct {
	statusWriter client.StatusWriter
}

func (s *statusWriter[T]) Update(ctx context.Context, obj T, opts ...client.UpdateOption) error {
	return s.statusWriter.Update(ctx, obj, opts...)
}

func (s *statusWriter[T]) Patch(ctx context.Context, obj T, patch client.Patch, opts ...client.PatchOption) error {
	return s.statusWriter.Patch(ctx, obj, patch, opts...)
}
