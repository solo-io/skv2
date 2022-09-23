// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the multicluster.solo.io/v1alpha1 APIs
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

// clienset for the multicluster.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the multicluster.solo.io/v1alpha1/v1alpha1 APIs
	KubernetesClusters() KubernetesClusterClient
}

type clientSet struct {
	client client.Client
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the multicluster.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) KubernetesClusters() KubernetesClusterClient {
	return NewKubernetesClusterClient(c.client)
}

// Reader knows how to read and list KubernetesClusters.
type KubernetesClusterReader interface {
	// Get retrieves a KubernetesCluster for the given object key
	GetKubernetesCluster(ctx context.Context, key client.ObjectKey) (*KubernetesCluster, error)

	// List retrieves list of KubernetesClusters for a given namespace and list options.
	ListKubernetesCluster(ctx context.Context, opts ...client.ListOption) (*KubernetesClusterList, error)
}

// KubernetesClusterTransitionFunction instructs the KubernetesClusterWriter how to transition between an existing
// KubernetesCluster object and a desired on an Upsert
type KubernetesClusterTransitionFunction func(existing, desired *KubernetesCluster) error

// Writer knows how to create, delete, and update KubernetesClusters.
type KubernetesClusterWriter interface {
	// Create saves the KubernetesCluster object.
	CreateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.CreateOption) error

	// Delete deletes the KubernetesCluster object.
	DeleteKubernetesCluster(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given KubernetesCluster object.
	UpdateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error

	// Patch patches the given KubernetesCluster object.
	PatchKubernetesCluster(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all KubernetesCluster objects matching the given options.
	DeleteAllOfKubernetesCluster(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the KubernetesCluster object.
	UpsertKubernetesCluster(ctx context.Context, obj *KubernetesCluster, transitionFuncs ...KubernetesClusterTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a KubernetesCluster object.
type KubernetesClusterStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given KubernetesCluster object.
	UpdateKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error

	// Patch patches the given KubernetesCluster object's subresource.
	PatchKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on KubernetesClusters.
type KubernetesClusterClient interface {
	KubernetesClusterReader
	KubernetesClusterWriter
	KubernetesClusterStatusWriter
}

type kubernetesClusterClient struct {
	client client.Client
}

func NewKubernetesClusterClient(client client.Client) *kubernetesClusterClient {
	return &kubernetesClusterClient{client: client}
}

func (c *kubernetesClusterClient) GetKubernetesCluster(ctx context.Context, key client.ObjectKey) (*KubernetesCluster, error) {
	obj := &KubernetesCluster{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *kubernetesClusterClient) ListKubernetesCluster(ctx context.Context, opts ...client.ListOption) (*KubernetesClusterList, error) {
	list := &KubernetesClusterList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *kubernetesClusterClient) CreateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) DeleteKubernetesCluster(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &KubernetesCluster{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) UpdateKubernetesCluster(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) PatchKubernetesCluster(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *kubernetesClusterClient) DeleteAllOfKubernetesCluster(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &KubernetesCluster{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) UpsertKubernetesCluster(ctx context.Context, obj *KubernetesCluster, transitionFuncs ...KubernetesClusterTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*KubernetesCluster), desired.(*KubernetesCluster)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *kubernetesClusterClient) UpdateKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *kubernetesClusterClient) PatchKubernetesClusterStatus(ctx context.Context, obj *KubernetesCluster, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides KubernetesClusterClients for multiple clusters.
type MulticlusterKubernetesClusterClient interface {
	// Cluster returns a KubernetesClusterClient for the given cluster
	Cluster(cluster string) (KubernetesClusterClient, error)
}

type multiclusterKubernetesClusterClient struct {
	client multicluster.Client
}

func NewMulticlusterKubernetesClusterClient(client multicluster.Client) MulticlusterKubernetesClusterClient {
	return &multiclusterKubernetesClusterClient{client: client}
}

func (m *multiclusterKubernetesClusterClient) Cluster(cluster string) (KubernetesClusterClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewKubernetesClusterClient(client), nil
}
