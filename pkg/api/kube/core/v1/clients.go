package v1

import (
	"context"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "k8s.io/api/core/v1"
)

// clienset for the core/v1 APIs
type Clientset interface {
	// clienset for the core/v1/v1 APIs
	Secrets() SecretClient
	// clienset for the core/v1/v1 APIs
	ConfigMaps() ConfigMapClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (*clientSet, error) {
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

func NewClientset(client client.Client) *clientSet {
	return &clientSet{client: client}
}

// clienset for the core/v1/v1 APIs
func (c *clientSet) Secrets() SecretClient {
	return NewSecretClient(c.client)
}

// clienset for the core/v1/v1 APIs
func (c *clientSet) ConfigMaps() ConfigMapClient {
	return NewConfigMapClient(c.client)
}

// Reader knows how to read and list Secrets.
type SecretReader interface {
	// Get retrieves a Secret for the given object key
	GetSecret(ctx context.Context, key client.ObjectKey) (*Secret, error)

	// List retrieves list of Secrets for a given namespace and list options.
	ListSecret(ctx context.Context, opts ...client.ListOption) (*SecretList, error)
}

// Writer knows how to create, delete, and update Secrets.
type SecretWriter interface {
	// Create saves the Secret object.
	CreateSecret(ctx context.Context, obj *Secret, opts ...client.CreateOption) error

	// Delete deletes the Secret object.
	DeleteSecret(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Secret object.
	UpdateSecret(ctx context.Context, obj *Secret, opts ...client.UpdateOption) error

	// Patch patches the given Secret object.
	PatchSecret(ctx context.Context, obj *Secret, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Secret objects matching the given options.
	DeleteAllOfSecret(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a Secret object.
type SecretStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Secret object.
	UpdateSecretStatus(ctx context.Context, obj *Secret, opts ...client.UpdateOption) error

	// Patch patches the given Secret object's subresource.
	PatchSecretStatus(ctx context.Context, obj *Secret, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Secrets.
type SecretClient interface {
	SecretReader
	SecretWriter
	SecretStatusWriter
}

type secretClient struct {
	client client.Client
}

func NewSecretClient(client client.Client) *secretClient {
	return &secretClient{client: client}
}

func (c *secretClient) GetSecret(ctx context.Context, key client.ObjectKey) (*Secret, error) {
	obj := &Secret{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *secretClient) ListSecret(ctx context.Context, opts ...client.ListOption) (*SecretList, error) {
	list := &SecretList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *secretClient) CreateSecret(ctx context.Context, obj *Secret, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *secretClient) DeleteSecret(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Secret{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *secretClient) UpdateSecret(ctx context.Context, obj *Secret, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *secretClient) PatchSecret(ctx context.Context, obj *Secret, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *secretClient) DeleteAllOfSecret(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Secret{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *secretClient) UpdateSecretStatus(ctx context.Context, obj *Secret, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *secretClient) PatchSecretStatus(ctx context.Context, obj *Secret, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ConfigMaps.
type ConfigMapReader interface {
	// Get retrieves a ConfigMap for the given object key
	GetConfigMap(ctx context.Context, key client.ObjectKey) (*ConfigMap, error)

	// List retrieves list of ConfigMaps for a given namespace and list options.
	ListConfigMap(ctx context.Context, opts ...client.ListOption) (*ConfigMapList, error)
}

// Writer knows how to create, delete, and update ConfigMaps.
type ConfigMapWriter interface {
	// Create saves the ConfigMap object.
	CreateConfigMap(ctx context.Context, obj *ConfigMap, opts ...client.CreateOption) error

	// Delete deletes the ConfigMap object.
	DeleteConfigMap(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ConfigMap object.
	UpdateConfigMap(ctx context.Context, obj *ConfigMap, opts ...client.UpdateOption) error

	// Patch patches the given ConfigMap object.
	PatchConfigMap(ctx context.Context, obj *ConfigMap, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ConfigMap objects matching the given options.
	DeleteAllOfConfigMap(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a ConfigMap object.
type ConfigMapStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ConfigMap object.
	UpdateConfigMapStatus(ctx context.Context, obj *ConfigMap, opts ...client.UpdateOption) error

	// Patch patches the given ConfigMap object's subresource.
	PatchConfigMapStatus(ctx context.Context, obj *ConfigMap, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ConfigMaps.
type ConfigMapClient interface {
	ConfigMapReader
	ConfigMapWriter
	ConfigMapStatusWriter
}

type configMapClient struct {
	client client.Client
}

func NewConfigMapClient(client client.Client) *configMapClient {
	return &configMapClient{client: client}
}

func (c *configMapClient) GetConfigMap(ctx context.Context, key client.ObjectKey) (*ConfigMap, error) {
	obj := &ConfigMap{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *configMapClient) ListConfigMap(ctx context.Context, opts ...client.ListOption) (*ConfigMapList, error) {
	list := &ConfigMapList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *configMapClient) CreateConfigMap(ctx context.Context, obj *ConfigMap, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *configMapClient) DeleteConfigMap(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ConfigMap{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *configMapClient) UpdateConfigMap(ctx context.Context, obj *ConfigMap, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *configMapClient) PatchConfigMap(ctx context.Context, obj *ConfigMap, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *configMapClient) DeleteAllOfConfigMap(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ConfigMap{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *configMapClient) UpdateConfigMapStatus(ctx context.Context, obj *ConfigMap, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *configMapClient) PatchConfigMapStatus(ctx context.Context, obj *ConfigMap, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
