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
    apiextensions_k8s_io_v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
    
)

// MulticlusterClientset for the apiextensions.k8s.io/v1 APIs
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

// clienset for the apiextensions.k8s.io/v1 APIs
type Clientset interface {
    // clienset for the apiextensions.k8s.io/v1/v1 APIs
    CustomResourceDefinitions() CustomResourceDefinitionClient
}

type clientSet struct {
    client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
    scheme := scheme.Scheme
    if err := apiextensions_k8s_io_v1.SchemeBuilder.AddToScheme(scheme); err != nil{
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

// clienset for the apiextensions.k8s.io/v1/v1 APIs
func (c *clientSet) CustomResourceDefinitions() CustomResourceDefinitionClient {
    return NewCustomResourceDefinitionClient(c.client)
}

// Reader knows how to read and list CustomResourceDefinitions.
type CustomResourceDefinitionReader interface {
    // Get retrieves a CustomResourceDefinition for the given object key
    GetCustomResourceDefinition(ctx context.Context, name string) (*apiextensions_k8s_io_v1.CustomResourceDefinition, error)

    // List retrieves list of CustomResourceDefinitions for a given namespace and list options.
    ListCustomResourceDefinition(ctx context.Context, opts ...client.ListOption) (*apiextensions_k8s_io_v1.CustomResourceDefinitionList, error)
}

// CustomResourceDefinitionTransitionFunction instructs the CustomResourceDefinitionWriter how to transition between an existing
// CustomResourceDefinition object and a desired on an Upsert
type CustomResourceDefinitionTransitionFunction func(existing, desired *apiextensions_k8s_io_v1.CustomResourceDefinition) error

// Writer knows how to create, delete, and update CustomResourceDefinitions.
type CustomResourceDefinitionWriter interface {
    // Create saves the CustomResourceDefinition object.
    CreateCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, opts ...client.CreateOption) error

    // Delete deletes the CustomResourceDefinition object.
    DeleteCustomResourceDefinition(ctx context.Context, name string, opts ...client.DeleteOption) error

    // Update updates the given CustomResourceDefinition object.
    UpdateCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, opts ...client.UpdateOption) error

    // Patch patches the given CustomResourceDefinition object.
    PatchCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, patch client.Patch, opts ...client.PatchOption) error

    // DeleteAllOf deletes all CustomResourceDefinition objects matching the given options.
    DeleteAllOfCustomResourceDefinition(ctx context.Context, opts ...client.DeleteAllOfOption) error

    // Create or Update the CustomResourceDefinition object.
    UpsertCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, transitionFuncs ...CustomResourceDefinitionTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a CustomResourceDefinition object.
type CustomResourceDefinitionStatusWriter interface {
    // Update updates the fields corresponding to the status subresource for the
    // given CustomResourceDefinition object.
    UpdateCustomResourceDefinitionStatus(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, opts ...client.SubResourceUpdateOption) error

    // Patch patches the given CustomResourceDefinition object's subresource.
    PatchCustomResourceDefinitionStatus(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, patch client.Patch, opts ...client.SubResourcePatchOption) error
}

// Client knows how to perform CRUD operations on CustomResourceDefinitions.
type CustomResourceDefinitionClient interface {
    CustomResourceDefinitionReader
    CustomResourceDefinitionWriter
    CustomResourceDefinitionStatusWriter
}

type customResourceDefinitionClient struct {
    client client.Client
}

func NewCustomResourceDefinitionClient(client client.Client) *customResourceDefinitionClient {
    return &customResourceDefinitionClient{client: client}
}


func (c *customResourceDefinitionClient) GetCustomResourceDefinition(ctx context.Context, name string) (*apiextensions_k8s_io_v1.CustomResourceDefinition, error) {
    obj := &apiextensions_k8s_io_v1.CustomResourceDefinition{}
    key := client.ObjectKey{
        Name: name,
    }
    if err := c.client.Get(ctx, key, obj); err != nil {
        return nil, err
    }
    return obj, nil
}

func (c *customResourceDefinitionClient) ListCustomResourceDefinition(ctx context.Context, opts ...client.ListOption) (*apiextensions_k8s_io_v1.CustomResourceDefinitionList, error) {
    list := &apiextensions_k8s_io_v1.CustomResourceDefinitionList{}
    if err := c.client.List(ctx, list, opts...); err != nil {
        return nil, err
    }
    return list, nil
}

func (c *customResourceDefinitionClient) CreateCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, opts ...client.CreateOption) error {
    return c.client.Create(ctx, obj, opts...)
}


func (c *customResourceDefinitionClient) DeleteCustomResourceDefinition(ctx context.Context, name string, opts ...client.DeleteOption) error {
    obj := &apiextensions_k8s_io_v1.CustomResourceDefinition{}
    obj.SetName(name)
    return c.client.Delete(ctx, obj, opts...)
}

func (c *customResourceDefinitionClient) UpdateCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, opts ...client.UpdateOption) error {
    return c.client.Update(ctx, obj, opts...)
}

func (c *customResourceDefinitionClient) PatchCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, patch client.Patch, opts ...client.PatchOption) error {
    return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *customResourceDefinitionClient) DeleteAllOfCustomResourceDefinition(ctx context.Context, opts ...client.DeleteAllOfOption) error {
    obj := &apiextensions_k8s_io_v1.CustomResourceDefinition{}
    return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *customResourceDefinitionClient) UpsertCustomResourceDefinition(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, transitionFuncs ...CustomResourceDefinitionTransitionFunction) error {
    genericTxFunc := func(existing, desired runtime.Object) error {
        for _, txFunc := range transitionFuncs {
            if err := txFunc(existing.(*apiextensions_k8s_io_v1.CustomResourceDefinition), desired.(*apiextensions_k8s_io_v1.CustomResourceDefinition)); err != nil {
                return err
            }
        }
        return nil
    }
    _, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
    return err
}

func (c *customResourceDefinitionClient) UpdateCustomResourceDefinitionStatus(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, opts ...client.SubResourceUpdateOption) error {
    return c.client.Status().Update(ctx, obj, opts...)
}

func (c *customResourceDefinitionClient) PatchCustomResourceDefinitionStatus(ctx context.Context, obj *apiextensions_k8s_io_v1.CustomResourceDefinition, patch client.Patch, opts ...client.SubResourcePatchOption) error {
    return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides CustomResourceDefinitionClients for multiple clusters.
type MulticlusterCustomResourceDefinitionClient interface {
    // Cluster returns a CustomResourceDefinitionClient for the given cluster
    Cluster(cluster string) (CustomResourceDefinitionClient, error)
}

type multiclusterCustomResourceDefinitionClient struct {
    client multicluster.Client
}

func NewMulticlusterCustomResourceDefinitionClient(client multicluster.Client) MulticlusterCustomResourceDefinitionClient {
    return &multiclusterCustomResourceDefinitionClient{client: client}
}

func (m *multiclusterCustomResourceDefinitionClient) Cluster(cluster string) (CustomResourceDefinitionClient, error) {
    client, err := m.client.Cluster(cluster)
    if err != nil {
        return nil, err
    }
    return  NewCustomResourceDefinitionClient(client), nil
}