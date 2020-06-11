// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the things.test.io/v1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the things.test.io/v1 APIs
type Clientset interface {
	// clienset for the things.test.io/v1/v1 APIs
	Paints() PaintClient
	// clienset for the things.test.io/v1/v1 APIs
	ClusterResources() ClusterResourceClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the things.test.io/v1/v1 APIs
func (c *clientSet) Paints() PaintClient {
	return NewPaintClient(c.client)
}

// clienset for the things.test.io/v1/v1 APIs
func (c *clientSet) ClusterResources() ClusterResourceClient {
	return NewClusterResourceClient(c.client)
}

// Reader knows how to read and list Paints.
type PaintReader interface {
	// Get retrieves a Paint for the given object key
	GetPaint(ctx context.Context, key client.ObjectKey) (*Paint, error)

	// List retrieves list of Paints for a given namespace and list options.
	ListPaint(ctx context.Context, opts ...client.ListOption) (*PaintList, error)
}

// PaintTransitionFunction instructs the PaintWriter how to transition between an existing
// Paint object and a desired on an Upsert
type PaintTransitionFunction func(existing, desired *Paint) error

// Writer knows how to create, delete, and update Paints.
type PaintWriter interface {
	// Create saves the Paint object.
	CreatePaint(ctx context.Context, obj *Paint, opts ...client.CreateOption) error

	// Delete deletes the Paint object.
	DeletePaint(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Paint object.
	UpdatePaint(ctx context.Context, obj *Paint, opts ...client.UpdateOption) error

	// Patch patches the given Paint object.
	PatchPaint(ctx context.Context, obj *Paint, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Paint objects matching the given options.
	DeleteAllOfPaint(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Paint object.
	UpsertPaint(ctx context.Context, obj *Paint, transitionFuncs ...PaintTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Paint object.
type PaintStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Paint object.
	UpdatePaintStatus(ctx context.Context, obj *Paint, opts ...client.UpdateOption) error

	// Patch patches the given Paint object's subresource.
	PatchPaintStatus(ctx context.Context, obj *Paint, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Paints.
type PaintClient interface {
	PaintReader
	PaintWriter
	PaintStatusWriter
}

type paintClient struct {
	client client.Client
}

func NewPaintClient(client client.Client) *paintClient {
	return &paintClient{client: client}
}

func (c *paintClient) GetPaint(ctx context.Context, key client.ObjectKey) (*Paint, error) {
	obj := &Paint{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *paintClient) ListPaint(ctx context.Context, opts ...client.ListOption) (*PaintList, error) {
	list := &PaintList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *paintClient) CreatePaint(ctx context.Context, obj *Paint, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *paintClient) DeletePaint(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Paint{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *paintClient) UpdatePaint(ctx context.Context, obj *Paint, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *paintClient) PatchPaint(ctx context.Context, obj *Paint, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *paintClient) DeleteAllOfPaint(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Paint{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *paintClient) UpsertPaint(ctx context.Context, obj *Paint, transitionFuncs ...PaintTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Paint), desired.(*Paint)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *paintClient) UpdatePaintStatus(ctx context.Context, obj *Paint, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *paintClient) PatchPaintStatus(ctx context.Context, obj *Paint, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ClusterResources.
type ClusterResourceReader interface {
	// Get retrieves a ClusterResource for the given object key
	GetClusterResource(ctx context.Context, name string) (*ClusterResource, error)

	// List retrieves list of ClusterResources for a given namespace and list options.
	ListClusterResource(ctx context.Context, opts ...client.ListOption) (*ClusterResourceList, error)
}

// ClusterResourceTransitionFunction instructs the ClusterResourceWriter how to transition between an existing
// ClusterResource object and a desired on an Upsert
type ClusterResourceTransitionFunction func(existing, desired *ClusterResource) error

// Writer knows how to create, delete, and update ClusterResources.
type ClusterResourceWriter interface {
	// Create saves the ClusterResource object.
	CreateClusterResource(ctx context.Context, obj *ClusterResource, opts ...client.CreateOption) error

	// Delete deletes the ClusterResource object.
	DeleteClusterResource(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given ClusterResource object.
	UpdateClusterResource(ctx context.Context, obj *ClusterResource, opts ...client.UpdateOption) error

	// Patch patches the given ClusterResource object.
	PatchClusterResource(ctx context.Context, obj *ClusterResource, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ClusterResource objects matching the given options.
	DeleteAllOfClusterResource(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ClusterResource object.
	UpsertClusterResource(ctx context.Context, obj *ClusterResource, transitionFuncs ...ClusterResourceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ClusterResource object.
type ClusterResourceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ClusterResource object.
	UpdateClusterResourceStatus(ctx context.Context, obj *ClusterResource, opts ...client.UpdateOption) error

	// Patch patches the given ClusterResource object's subresource.
	PatchClusterResourceStatus(ctx context.Context, obj *ClusterResource, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ClusterResources.
type ClusterResourceClient interface {
	ClusterResourceReader
	ClusterResourceWriter
	ClusterResourceStatusWriter
}

type clusterResourceClient struct {
	client client.Client
}

func NewClusterResourceClient(client client.Client) *clusterResourceClient {
	return &clusterResourceClient{client: client}
}

func (c *clusterResourceClient) GetClusterResource(ctx context.Context, name string) (*ClusterResource, error) {
	obj := &ClusterResource{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *clusterResourceClient) ListClusterResource(ctx context.Context, opts ...client.ListOption) (*ClusterResourceList, error) {
	list := &ClusterResourceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *clusterResourceClient) CreateClusterResource(ctx context.Context, obj *ClusterResource, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *clusterResourceClient) DeleteClusterResource(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &ClusterResource{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *clusterResourceClient) UpdateClusterResource(ctx context.Context, obj *ClusterResource, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *clusterResourceClient) PatchClusterResource(ctx context.Context, obj *ClusterResource, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *clusterResourceClient) DeleteAllOfClusterResource(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ClusterResource{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *clusterResourceClient) UpsertClusterResource(ctx context.Context, obj *ClusterResource, transitionFuncs ...ClusterResourceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*ClusterResource), desired.(*ClusterResource)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *clusterResourceClient) UpdateClusterResourceStatus(ctx context.Context, obj *ClusterResource, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *clusterResourceClient) PatchClusterResourceStatus(ctx context.Context, obj *ClusterResource, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
