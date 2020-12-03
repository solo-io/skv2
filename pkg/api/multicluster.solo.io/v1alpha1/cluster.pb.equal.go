// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/multicluster/v1alpha1/cluster.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *KubernetesClusterSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesClusterSpec)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetSecretName(), target.GetSecretName()) != 0 {
		return false
	}

	if strings.Compare(m.GetClusterDomain(), target.GetClusterDomain()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetProviderInfo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProviderInfo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProviderInfo(), target.GetProviderInfo()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *KubernetesClusterStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesClusterStatus)
	if !ok {
		that2, ok := that.(KubernetesClusterStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetStatus()) != len(target.GetStatus()) {
		return false
	}
	for idx, v := range m.GetStatus() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatus()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetStatus()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if len(m.GetPolicyRules()) != len(target.GetPolicyRules()) {
		return false
	}
	for idx, v := range m.GetPolicyRules() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPolicyRules()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPolicyRules()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *PolicyRule) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PolicyRule)
	if !ok {
		that2, ok := that.(PolicyRule)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetVerbs()) != len(target.GetVerbs()) {
		return false
	}
	for idx, v := range m.GetVerbs() {

		if strings.Compare(v, target.GetVerbs()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetApiGroups()) != len(target.GetApiGroups()) {
		return false
	}
	for idx, v := range m.GetApiGroups() {

		if strings.Compare(v, target.GetApiGroups()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetResources()) != len(target.GetResources()) {
		return false
	}
	for idx, v := range m.GetResources() {

		if strings.Compare(v, target.GetResources()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetResourceNames()) != len(target.GetResourceNames()) {
		return false
	}
	for idx, v := range m.GetResourceNames() {

		if strings.Compare(v, target.GetResourceNames()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetNonResourceUrls()) != len(target.GetNonResourceUrls()) {
		return false
	}
	for idx, v := range m.GetNonResourceUrls() {

		if strings.Compare(v, target.GetNonResourceUrls()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *KubernetesClusterSpec_ProviderInfo) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesClusterSpec_ProviderInfo)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_ProviderInfo)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.ProviderInfoType.(type) {

	case *KubernetesClusterSpec_ProviderInfo_Eks:

		if h, ok := interface{}(m.GetEks()).(equality.Equalizer); ok {
			if !h.Equal(target.GetEks()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetEks(), target.GetEks()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *KubernetesClusterSpec_Eks) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesClusterSpec_Eks)
	if !ok {
		that2, ok := that.(KubernetesClusterSpec_Eks)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetArn(), target.GetArn()) != 0 {
		return false
	}

	if strings.Compare(m.GetAccountId(), target.GetAccountId()) != 0 {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	return true
}
