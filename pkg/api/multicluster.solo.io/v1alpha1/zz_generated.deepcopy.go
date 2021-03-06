// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for multicluster.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for KubernetesCluster

func (in *KubernetesCluster) DeepCopyInto(out *KubernetesCluster) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *KubernetesCluster) DeepCopy() *KubernetesCluster {
	if in == nil {
		return nil
	}
	out := new(KubernetesCluster)
	in.DeepCopyInto(out)
	return out
}

func (in *KubernetesCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *KubernetesClusterList) DeepCopyInto(out *KubernetesClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubernetesCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *KubernetesClusterList) DeepCopy() *KubernetesClusterList {
	if in == nil {
		return nil
	}
	out := new(KubernetesClusterList)
	in.DeepCopyInto(out)
	return out
}

func (in *KubernetesClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
