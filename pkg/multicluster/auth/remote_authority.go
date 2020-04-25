package auth

import (
	"context"

	k8s_core_v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/core/v1"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_rbac_types "k8s.io/api/rbac/v1"
	k8s_errs "k8s.io/apimachinery/pkg/api/errors"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewRemoteAuthorityManager(
	serviceAccountClient k8s_core_v1.ServiceAccountClient,
	rbacClient RbacBinder,
) RemoteAuthorityManager {
	return &remoteAuthorityManager{
		serviceAccountClient: serviceAccountClient,
		rbacClient:           rbacClient,
	}
}

type remoteAuthorityManager struct {
	serviceAccountClient k8s_core_v1.ServiceAccountClient
	rbacClient           RbacBinder
}

func (r *remoteAuthorityManager) ApplyRemoteServiceAccount(
	ctx context.Context,
	name, namespace string,
	roles []*k8s_rbac_types.ClusterRole,
) (*k8s_core_types.ServiceAccount, error) {

	saToCreate := &k8s_core_types.ServiceAccount{
		ObjectMeta: k8s_meta.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	err := r.serviceAccountClient.CreateServiceAccount(ctx, saToCreate)
	if err != nil {
		if k8s_errs.IsAlreadyExists(err) {
			err = r.serviceAccountClient.UpdateServiceAccount(ctx, saToCreate)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	err = r.rbacClient.BindClusterRolesToServiceAccount(ctx, saToCreate, roles)
	if err != nil {
		return nil, err
	}

	return saToCreate, nil
}
