package auth

import (
	"context"
	"fmt"

	v1 "github.com/solo-io/skv2/pkg/generated/kubernetes/rbac.authorization.k8s.io/v1"
	k8sapiv1 "k8s.io/api/core/v1"
	rbactypes "k8s.io/api/rbac/v1"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewRbacBinder(kubeClients v1.Clientset) RbacBinder {
	return &defaultRbacBinder{clusterRoleBindingClient: kubeClients.ClusterRoleBindings()}
}

type defaultRbacBinder struct {
	clusterRoleBindingClient v1.ClusterRoleBindingClient
}

func (r *defaultRbacBinder) BindClusterRolesToServiceAccount(
	ctx context.Context,
	targetServiceAccount *k8sapiv1.ServiceAccount,
	roles []*rbactypes.ClusterRole,
) error {

	for _, role := range roles {
		crbToCreate := &rbactypes.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-%s-clusterrole-binding", targetServiceAccount.GetName(), role.GetName()),
			},
			TypeMeta: metav1.TypeMeta{
				Kind:       "ClusterRoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			Subjects: []rbactypes.Subject{{
				Kind:      "ServiceAccount",
				Name:      targetServiceAccount.GetName(),
				Namespace: targetServiceAccount.GetNamespace(),
			}},
			RoleRef: rbactypes.RoleRef{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
				Name:     role.GetName(),
			},
		}
		err := r.clusterRoleBindingClient.CreateClusterRoleBinding(ctx, crbToCreate)
		if err != nil {
			if kubeerrs.IsAlreadyExists(err) {
				err = r.clusterRoleBindingClient.UpdateClusterRoleBinding(ctx, crbToCreate)
				if err != nil {
					return err
				}
				continue
			}
			return err
		}
	}

	return nil
}
