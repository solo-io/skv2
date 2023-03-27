// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for things.test.io/v1 resources

package v1

import (
    runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for CueBug

func (in *CueBug) DeepCopyInto(out *CueBug) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *CueBug) DeepCopy() *CueBug {
    if in == nil {
        return nil
    }
    out := new(CueBug)
    in.DeepCopyInto(out)
    return out
}

func (in *CueBug) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *CueBugList) DeepCopyInto(out *CueBugList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]CueBug, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *CueBugList) DeepCopy() *CueBugList {
    if in == nil {
        return nil
    }
    out := new(CueBugList)
    in.DeepCopyInto(out)
    return out
}

func (in *CueBugList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

