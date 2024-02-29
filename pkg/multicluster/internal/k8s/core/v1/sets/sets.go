// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets



import (
    v1 "k8s.io/api/core/v1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type SecretSet interface {
	// Get the set stored keys
    Keys() sets.Set[uint64]
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*v1.Secret) bool) []*v1.Secret
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*v1.Secret) bool) []*v1.Secret
    // Return the Set as a map of key to resource.
    Map() map[uint64]*v1.Secret
    // Insert a resource into the set.
    Insert(secret ...*v1.Secret)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(secretSet SecretSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(secret ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(secret  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set SecretSet) SecretSet
    // Return the difference with the provided set
    Difference(set SecretSet) SecretSet
    // Return the intersection with the provided set
    Intersection(set SecretSet) SecretSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*v1.Secret, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another SecretSet
    Delta(newSet SecretSet) sksets.ResourceDelta
    // Create a deep copy of the current SecretSet
    Clone() SecretSet
}

func makeGenericSecretSet(secretList []*v1.Secret) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range secretList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type secretSet struct {
    set sksets.ResourceSet
}

func NewSecretSet(secretList ...*v1.Secret) SecretSet {
    return &secretSet{set: makeGenericSecretSet(secretList)}
}

func NewSecretSetFromList(secretList *v1.SecretList) SecretSet {
    list := make([]*v1.Secret, 0, len(secretList.Items))
    for idx := range secretList.Items {
        list = append(list, &secretList.Items[idx])
    }
    return &secretSet{set: makeGenericSecretSet(list)}
}

func (s *secretSet) Keys() sets.Set[uint64] {
	if s == nil {
		return sets.Set[uint64]{}
    }
    return s.Generic().Keys()
}

func (s *secretSet) List(filterResource ... func(*v1.Secret) bool) []*v1.Secret {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Secret))
        })
    }

    objs := s.Generic().List(genericFilters...)
    secretList := make([]*v1.Secret, 0, len(objs))
    for _, obj := range objs {
        secretList = append(secretList, obj.(*v1.Secret))
    }
    return secretList
}

func (s *secretSet) UnsortedList(filterResource ... func(*v1.Secret) bool) []*v1.Secret {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Secret))
        })
    }

    var secretList []*v1.Secret
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        secretList = append(secretList, obj.(*v1.Secret))
    }
    return secretList
}

func (s *secretSet) Map() map[uint64]*v1.Secret {
    if s == nil {
        return nil
    }

    newMap := map[uint64]*v1.Secret{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*v1.Secret)
    }
    return newMap
}

func (s *secretSet) Insert(
        secretList ...*v1.Secret,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range secretList {
        s.Generic().Insert(obj)
    }
}

func (s *secretSet) Has(secret ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(secret)
}

func (s *secretSet) Equal(
        secretSet SecretSet,
) bool {
    if s == nil {
        return secretSet == nil
    }
    return s.Generic().Equal(secretSet.Generic())
}

func (s *secretSet) Delete(Secret ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(Secret)
}

func (s *secretSet) Union(set SecretSet) SecretSet {
    if s == nil {
        return set
    }
    return &secretMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *secretSet) Difference(set SecretSet) SecretSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &secretSet{set: newSet}
}

func (s *secretSet) Intersection(set SecretSet) SecretSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var secretList []*v1.Secret
    for _, obj := range newSet.List() {
        secretList = append(secretList, obj.(*v1.Secret))
    }
    return NewSecretSet(secretList...)
}


func (s *secretSet) Find(id ezkube.ResourceId) (*v1.Secret, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Secret %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&v1.Secret{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*v1.Secret), nil
}

func (s *secretSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *secretSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *secretSet) Delta(newSet SecretSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *secretSet) Clone() SecretSet {
	if s == nil {
		return nil
	}
	return &secretMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type secretMergedSet struct {
    sets []sksets.ResourceSet
}

func NewSecretMergedSet(secretList ...*v1.Secret) SecretSet {
    return &secretMergedSet{sets: []sksets.ResourceSet{makeGenericSecretSet(secretList)}}
}

func NewSecretMergedSetFromList(secretList *v1.SecretList) SecretSet {
    list := make([]*v1.Secret, 0, len(secretList.Items))
    for idx := range secretList.Items {
        list = append(list, &secretList.Items[idx])
    }
    return &secretMergedSet{sets: []sksets.ResourceSet{makeGenericSecretSet(list)}}
}

func (s *secretMergedSet) Keys() sets.Set[uint64] {
	if s == nil {
		return sets.Set[uint64]{}
    }
    toRet := sets.Set[uint64]{}
	for _ , set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *secretMergedSet) List(filterResource ... func(*v1.Secret) bool) []*v1.Secret {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Secret))
        })
    }
   secretList := []*v1.Secret{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			secretList = append(secretList, obj.(*v1.Secret))
		}
	}
    return secretList
}

func (s *secretMergedSet) UnsortedList(filterResource ... func(*v1.Secret) bool) []*v1.Secret {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Secret))
        })
    }

    secretList := []*v1.Secret{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			secretList = append(secretList, obj.(*v1.Secret))
		}
	}
    return secretList
}

func (s *secretMergedSet) Map() map[uint64]*v1.Secret {
    if s == nil {
        return nil
    }

    newMap := map[uint64]*v1.Secret{}
    for _, set := range s.sets {
		for k, v := range set.Map() {
            newMap[k] = v.(*v1.Secret)
        }
    }
    return newMap
}

func (s *secretMergedSet) Insert(
        secretList ...*v1.Secret,
) {
    if s == nil {
    }
    if len(s.sets) == 0 {
        s.sets = append(s.sets, makeGenericSecretSet(secretList))
    }
    for _, obj := range secretList {
        s.sets[0].Insert(obj)
    }
}

func (s *secretMergedSet) Has(secret ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    for _, set := range s.sets {
		if set.Has(secret) {
			return true
		}
	}
    return false
}

func (s *secretMergedSet) Equal(
        secretSet SecretSet,
) bool {
    panic("unimplemented")
}

func (s *secretMergedSet) Delete(Secret ezkube.ResourceId) {
    panic("unimplemented")
}

func (s *secretMergedSet) Union(set SecretSet) SecretSet {
    return &secretMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *secretMergedSet) Difference(set SecretSet) SecretSet {
    panic("unimplemented")
}

func (s *secretMergedSet) Intersection(set SecretSet) SecretSet {
    panic("unimplemented")
}

func (s *secretMergedSet) Find(id ezkube.ResourceId) (*v1.Secret, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Secret %v", sksets.Key(id))
    }

    var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&v1.Secret{}, id)
		if err == nil {
			return obj.(*v1.Secret), nil
		}
	}

    return nil, err
}

func (s *secretMergedSet) Length() int {
    if s == nil {
        return 0
    }
    totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
    return totalLen
}

func (s *secretMergedSet) Generic() sksets.ResourceSet {
    panic("unimplemented")
}

func (s *secretMergedSet) Delta(newSet SecretSet) sksets.ResourceDelta {
    panic("unimplemented")
}

func (s *secretMergedSet) Clone() SecretSet {
	if s == nil {
		return nil
	}
	return &secretMergedSet{sets: s.sets[:]}
}

type ServiceAccountSet interface {
	// Get the set stored keys
    Keys() sets.Set[uint64]
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*v1.ServiceAccount) bool) []*v1.ServiceAccount
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*v1.ServiceAccount) bool) []*v1.ServiceAccount
    // Return the Set as a map of key to resource.
    Map() map[uint64]*v1.ServiceAccount
    // Insert a resource into the set.
    Insert(serviceAccount ...*v1.ServiceAccount)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(serviceAccountSet ServiceAccountSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(serviceAccount ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(serviceAccount  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set ServiceAccountSet) ServiceAccountSet
    // Return the difference with the provided set
    Difference(set ServiceAccountSet) ServiceAccountSet
    // Return the intersection with the provided set
    Intersection(set ServiceAccountSet) ServiceAccountSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*v1.ServiceAccount, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another ServiceAccountSet
    Delta(newSet ServiceAccountSet) sksets.ResourceDelta
    // Create a deep copy of the current ServiceAccountSet
    Clone() ServiceAccountSet
}

func makeGenericServiceAccountSet(serviceAccountList []*v1.ServiceAccount) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range serviceAccountList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type serviceAccountSet struct {
    set sksets.ResourceSet
}

func NewServiceAccountSet(serviceAccountList ...*v1.ServiceAccount) ServiceAccountSet {
    return &serviceAccountSet{set: makeGenericServiceAccountSet(serviceAccountList)}
}

func NewServiceAccountSetFromList(serviceAccountList *v1.ServiceAccountList) ServiceAccountSet {
    list := make([]*v1.ServiceAccount, 0, len(serviceAccountList.Items))
    for idx := range serviceAccountList.Items {
        list = append(list, &serviceAccountList.Items[idx])
    }
    return &serviceAccountSet{set: makeGenericServiceAccountSet(list)}
}

func (s *serviceAccountSet) Keys() sets.Set[uint64] {
	if s == nil {
		return sets.Set[uint64]{}
    }
    return s.Generic().Keys()
}

func (s *serviceAccountSet) List(filterResource ... func(*v1.ServiceAccount) bool) []*v1.ServiceAccount {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.ServiceAccount))
        })
    }

    objs := s.Generic().List(genericFilters...)
    serviceAccountList := make([]*v1.ServiceAccount, 0, len(objs))
    for _, obj := range objs {
        serviceAccountList = append(serviceAccountList, obj.(*v1.ServiceAccount))
    }
    return serviceAccountList
}

func (s *serviceAccountSet) UnsortedList(filterResource ... func(*v1.ServiceAccount) bool) []*v1.ServiceAccount {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.ServiceAccount))
        })
    }

    var serviceAccountList []*v1.ServiceAccount
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        serviceAccountList = append(serviceAccountList, obj.(*v1.ServiceAccount))
    }
    return serviceAccountList
}

func (s *serviceAccountSet) Map() map[uint64]*v1.ServiceAccount {
    if s == nil {
        return nil
    }

    newMap := map[uint64]*v1.ServiceAccount{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*v1.ServiceAccount)
    }
    return newMap
}

func (s *serviceAccountSet) Insert(
        serviceAccountList ...*v1.ServiceAccount,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range serviceAccountList {
        s.Generic().Insert(obj)
    }
}

func (s *serviceAccountSet) Has(serviceAccount ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(serviceAccount)
}

func (s *serviceAccountSet) Equal(
        serviceAccountSet ServiceAccountSet,
) bool {
    if s == nil {
        return serviceAccountSet == nil
    }
    return s.Generic().Equal(serviceAccountSet.Generic())
}

func (s *serviceAccountSet) Delete(ServiceAccount ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(ServiceAccount)
}

func (s *serviceAccountSet) Union(set ServiceAccountSet) ServiceAccountSet {
    if s == nil {
        return set
    }
    return &serviceAccountMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *serviceAccountSet) Difference(set ServiceAccountSet) ServiceAccountSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &serviceAccountSet{set: newSet}
}

func (s *serviceAccountSet) Intersection(set ServiceAccountSet) ServiceAccountSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var serviceAccountList []*v1.ServiceAccount
    for _, obj := range newSet.List() {
        serviceAccountList = append(serviceAccountList, obj.(*v1.ServiceAccount))
    }
    return NewServiceAccountSet(serviceAccountList...)
}


func (s *serviceAccountSet) Find(id ezkube.ResourceId) (*v1.ServiceAccount, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find ServiceAccount %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&v1.ServiceAccount{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*v1.ServiceAccount), nil
}

func (s *serviceAccountSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *serviceAccountSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *serviceAccountSet) Delta(newSet ServiceAccountSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *serviceAccountSet) Clone() ServiceAccountSet {
	if s == nil {
		return nil
	}
	return &serviceAccountMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type serviceAccountMergedSet struct {
    sets []sksets.ResourceSet
}

func NewServiceAccountMergedSet(serviceAccountList ...*v1.ServiceAccount) ServiceAccountSet {
    return &serviceAccountMergedSet{sets: []sksets.ResourceSet{makeGenericServiceAccountSet(serviceAccountList)}}
}

func NewServiceAccountMergedSetFromList(serviceAccountList *v1.ServiceAccountList) ServiceAccountSet {
    list := make([]*v1.ServiceAccount, 0, len(serviceAccountList.Items))
    for idx := range serviceAccountList.Items {
        list = append(list, &serviceAccountList.Items[idx])
    }
    return &serviceAccountMergedSet{sets: []sksets.ResourceSet{makeGenericServiceAccountSet(list)}}
}

func (s *serviceAccountMergedSet) Keys() sets.Set[uint64] {
	if s == nil {
		return sets.Set[uint64]{}
    }
    toRet := sets.Set[uint64]{}
	for _ , set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *serviceAccountMergedSet) List(filterResource ... func(*v1.ServiceAccount) bool) []*v1.ServiceAccount {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.ServiceAccount))
        })
    }
   serviceAccountList := []*v1.ServiceAccount{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			serviceAccountList = append(serviceAccountList, obj.(*v1.ServiceAccount))
		}
	}
    return serviceAccountList
}

func (s *serviceAccountMergedSet) UnsortedList(filterResource ... func(*v1.ServiceAccount) bool) []*v1.ServiceAccount {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.ServiceAccount))
        })
    }

    serviceAccountList := []*v1.ServiceAccount{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			serviceAccountList = append(serviceAccountList, obj.(*v1.ServiceAccount))
		}
	}
    return serviceAccountList
}

func (s *serviceAccountMergedSet) Map() map[uint64]*v1.ServiceAccount {
    if s == nil {
        return nil
    }

    newMap := map[uint64]*v1.ServiceAccount{}
    for _, set := range s.sets {
		for k, v := range set.Map() {
            newMap[k] = v.(*v1.ServiceAccount)
        }
    }
    return newMap
}

func (s *serviceAccountMergedSet) Insert(
        serviceAccountList ...*v1.ServiceAccount,
) {
    if s == nil {
    }
    if len(s.sets) == 0 {
        s.sets = append(s.sets, makeGenericServiceAccountSet(serviceAccountList))
    }
    for _, obj := range serviceAccountList {
        s.sets[0].Insert(obj)
    }
}

func (s *serviceAccountMergedSet) Has(serviceAccount ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    for _, set := range s.sets {
		if set.Has(serviceAccount) {
			return true
		}
	}
    return false
}

func (s *serviceAccountMergedSet) Equal(
        serviceAccountSet ServiceAccountSet,
) bool {
    panic("unimplemented")
}

func (s *serviceAccountMergedSet) Delete(ServiceAccount ezkube.ResourceId) {
    panic("unimplemented")
}

func (s *serviceAccountMergedSet) Union(set ServiceAccountSet) ServiceAccountSet {
    return &serviceAccountMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *serviceAccountMergedSet) Difference(set ServiceAccountSet) ServiceAccountSet {
    panic("unimplemented")
}

func (s *serviceAccountMergedSet) Intersection(set ServiceAccountSet) ServiceAccountSet {
    panic("unimplemented")
}

func (s *serviceAccountMergedSet) Find(id ezkube.ResourceId) (*v1.ServiceAccount, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find ServiceAccount %v", sksets.Key(id))
    }

    var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&v1.ServiceAccount{}, id)
		if err == nil {
			return obj.(*v1.ServiceAccount), nil
		}
	}

    return nil, err
}

func (s *serviceAccountMergedSet) Length() int {
    if s == nil {
        return 0
    }
    totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
    return totalLen
}

func (s *serviceAccountMergedSet) Generic() sksets.ResourceSet {
    panic("unimplemented")
}

func (s *serviceAccountMergedSet) Delta(newSet ServiceAccountSet) sksets.ResourceDelta {
    panic("unimplemented")
}

func (s *serviceAccountMergedSet) Clone() ServiceAccountSet {
	if s == nil {
		return nil
	}
	return &serviceAccountMergedSet{sets: s.sets[:]}
}

type NamespaceSet interface {
	// Get the set stored keys
    Keys() sets.Set[uint64]
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    List(filterResource ... func(*v1.Namespace) bool) []*v1.Namespace
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    // The filter function should return false to keep the resource, true to drop it.
    UnsortedList(filterResource ... func(*v1.Namespace) bool) []*v1.Namespace
    // Return the Set as a map of key to resource.
    Map() map[uint64]*v1.Namespace
    // Insert a resource into the set.
    Insert(namespace ...*v1.Namespace)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(namespaceSet NamespaceSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(namespace ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(namespace  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set NamespaceSet) NamespaceSet
    // Return the difference with the provided set
    Difference(set NamespaceSet) NamespaceSet
    // Return the intersection with the provided set
    Intersection(set NamespaceSet) NamespaceSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*v1.Namespace, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another NamespaceSet
    Delta(newSet NamespaceSet) sksets.ResourceDelta
    // Create a deep copy of the current NamespaceSet
    Clone() NamespaceSet
}

func makeGenericNamespaceSet(namespaceList []*v1.Namespace) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range namespaceList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type namespaceSet struct {
    set sksets.ResourceSet
}

func NewNamespaceSet(namespaceList ...*v1.Namespace) NamespaceSet {
    return &namespaceSet{set: makeGenericNamespaceSet(namespaceList)}
}

func NewNamespaceSetFromList(namespaceList *v1.NamespaceList) NamespaceSet {
    list := make([]*v1.Namespace, 0, len(namespaceList.Items))
    for idx := range namespaceList.Items {
        list = append(list, &namespaceList.Items[idx])
    }
    return &namespaceSet{set: makeGenericNamespaceSet(list)}
}

func (s *namespaceSet) Keys() sets.Set[uint64] {
	if s == nil {
		return sets.Set[uint64]{}
    }
    return s.Generic().Keys()
}

func (s *namespaceSet) List(filterResource ... func(*v1.Namespace) bool) []*v1.Namespace {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Namespace))
        })
    }

    objs := s.Generic().List(genericFilters...)
    namespaceList := make([]*v1.Namespace, 0, len(objs))
    for _, obj := range objs {
        namespaceList = append(namespaceList, obj.(*v1.Namespace))
    }
    return namespaceList
}

func (s *namespaceSet) UnsortedList(filterResource ... func(*v1.Namespace) bool) []*v1.Namespace {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Namespace))
        })
    }

    var namespaceList []*v1.Namespace
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        namespaceList = append(namespaceList, obj.(*v1.Namespace))
    }
    return namespaceList
}

func (s *namespaceSet) Map() map[uint64]*v1.Namespace {
    if s == nil {
        return nil
    }

    newMap := map[uint64]*v1.Namespace{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*v1.Namespace)
    }
    return newMap
}

func (s *namespaceSet) Insert(
        namespaceList ...*v1.Namespace,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range namespaceList {
        s.Generic().Insert(obj)
    }
}

func (s *namespaceSet) Has(namespace ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(namespace)
}

func (s *namespaceSet) Equal(
        namespaceSet NamespaceSet,
) bool {
    if s == nil {
        return namespaceSet == nil
    }
    return s.Generic().Equal(namespaceSet.Generic())
}

func (s *namespaceSet) Delete(Namespace ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(Namespace)
}

func (s *namespaceSet) Union(set NamespaceSet) NamespaceSet {
    if s == nil {
        return set
    }
    return &namespaceMergedSet{sets: []sksets.ResourceSet{s.Generic(), set.Generic()}}
}

func (s *namespaceSet) Difference(set NamespaceSet) NamespaceSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &namespaceSet{set: newSet}
}

func (s *namespaceSet) Intersection(set NamespaceSet) NamespaceSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var namespaceList []*v1.Namespace
    for _, obj := range newSet.List() {
        namespaceList = append(namespaceList, obj.(*v1.Namespace))
    }
    return NewNamespaceSet(namespaceList...)
}


func (s *namespaceSet) Find(id ezkube.ResourceId) (*v1.Namespace, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Namespace %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&v1.Namespace{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*v1.Namespace), nil
}

func (s *namespaceSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *namespaceSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *namespaceSet) Delta(newSet NamespaceSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}

func (s *namespaceSet) Clone() NamespaceSet {
	if s == nil {
		return nil
	}
	return &namespaceMergedSet{sets: []sksets.ResourceSet{s.Generic()}}
}

type namespaceMergedSet struct {
    sets []sksets.ResourceSet
}

func NewNamespaceMergedSet(namespaceList ...*v1.Namespace) NamespaceSet {
    return &namespaceMergedSet{sets: []sksets.ResourceSet{makeGenericNamespaceSet(namespaceList)}}
}

func NewNamespaceMergedSetFromList(namespaceList *v1.NamespaceList) NamespaceSet {
    list := make([]*v1.Namespace, 0, len(namespaceList.Items))
    for idx := range namespaceList.Items {
        list = append(list, &namespaceList.Items[idx])
    }
    return &namespaceMergedSet{sets: []sksets.ResourceSet{makeGenericNamespaceSet(list)}}
}

func (s *namespaceMergedSet) Keys() sets.Set[uint64] {
	if s == nil {
		return sets.Set[uint64]{}
    }
    toRet := sets.Set[uint64]{}
	for _ , set := range s.sets {
		toRet = toRet.Union(set.Keys())
	}
	return toRet
}

func (s *namespaceMergedSet) List(filterResource ... func(*v1.Namespace) bool) []*v1.Namespace {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Namespace))
        })
    }
   namespaceList := []*v1.Namespace{}
	for _, set := range s.sets {
		for _, obj := range set.List(genericFilters...) {
			namespaceList = append(namespaceList, obj.(*v1.Namespace))
		}
	}
    return namespaceList
}

func (s *namespaceMergedSet) UnsortedList(filterResource ... func(*v1.Namespace) bool) []*v1.Namespace {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        filter := filter
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*v1.Namespace))
        })
    }

    namespaceList := []*v1.Namespace{}
	for _, set := range s.sets {
		for _, obj := range set.UnsortedList(genericFilters...) {
			namespaceList = append(namespaceList, obj.(*v1.Namespace))
		}
	}
    return namespaceList
}

func (s *namespaceMergedSet) Map() map[uint64]*v1.Namespace {
    if s == nil {
        return nil
    }

    newMap := map[uint64]*v1.Namespace{}
    for _, set := range s.sets {
		for k, v := range set.Map() {
            newMap[k] = v.(*v1.Namespace)
        }
    }
    return newMap
}

func (s *namespaceMergedSet) Insert(
        namespaceList ...*v1.Namespace,
) {
    if s == nil {
    }
    if len(s.sets) == 0 {
        s.sets = append(s.sets, makeGenericNamespaceSet(namespaceList))
    }
    for _, obj := range namespaceList {
        s.sets[0].Insert(obj)
    }
}

func (s *namespaceMergedSet) Has(namespace ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    for _, set := range s.sets {
		if set.Has(namespace) {
			return true
		}
	}
    return false
}

func (s *namespaceMergedSet) Equal(
        namespaceSet NamespaceSet,
) bool {
    panic("unimplemented")
}

func (s *namespaceMergedSet) Delete(Namespace ezkube.ResourceId) {
    panic("unimplemented")
}

func (s *namespaceMergedSet) Union(set NamespaceSet) NamespaceSet {
    return &namespaceMergedSet{sets: append(s.sets, set.Generic())}
}

func (s *namespaceMergedSet) Difference(set NamespaceSet) NamespaceSet {
    panic("unimplemented")
}

func (s *namespaceMergedSet) Intersection(set NamespaceSet) NamespaceSet {
    panic("unimplemented")
}

func (s *namespaceMergedSet) Find(id ezkube.ResourceId) (*v1.Namespace, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find Namespace %v", sksets.Key(id))
    }

    var err error
	for _, set := range s.sets {
		var obj ezkube.ResourceId
		obj, err = set.Find(&v1.Namespace{}, id)
		if err == nil {
			return obj.(*v1.Namespace), nil
		}
	}

    return nil, err
}

func (s *namespaceMergedSet) Length() int {
    if s == nil {
        return 0
    }
    totalLen := 0
	for _, set := range s.sets {
		totalLen += set.Length()
	}
    return totalLen
}

func (s *namespaceMergedSet) Generic() sksets.ResourceSet {
    panic("unimplemented")
}

func (s *namespaceMergedSet) Delta(newSet NamespaceSet) sksets.ResourceDelta {
    panic("unimplemented")
}

func (s *namespaceMergedSet) Clone() NamespaceSet {
	if s == nil {
		return nil
	}
	return &namespaceMergedSet{sets: s.sets[:]}
}
