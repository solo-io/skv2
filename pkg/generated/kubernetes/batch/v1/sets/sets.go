// Code generated by skv2. DO NOT EDIT.

package v1sets

import (
	. "k8s.io/api/batch/v1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type JobSet interface {
	Keys() sets.String
	List() []*Job
	Map() map[string]*Job
	Insert(job ...*Job)
	Equal(jobSet JobSet) bool
	Has(job *Job) bool
	Delete(job *Job)
	Union(set JobSet) JobSet
	Difference(set JobSet) JobSet
	Intersection(set JobSet) JobSet
}

func makeGenericJobSet(jobList []*Job) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range jobList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type jobSet struct {
	set sksets.ResourceSet
}

func NewJobSet(jobList ...*Job) JobSet {
	return &jobSet{set: makeGenericJobSet(jobList)}
}

func (s jobSet) Keys() sets.String {
	return s.set.Keys()
}

func (s jobSet) List() []*Job {
	var jobList []*Job
	for _, obj := range s.set.List() {
		jobList = append(jobList, obj.(*Job))
	}
	return jobList
}

func (s jobSet) Map() map[string]*Job {
	newMap := map[string]*Job{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*Job)
	}
	return newMap
}

func (s jobSet) Insert(
	jobList ...*Job,
) {
	for _, obj := range jobList {
		s.set.Insert(obj)
	}
}

func (s jobSet) Has(job *Job) bool {
	return s.set.Has(job)
}

func (s jobSet) Equal(
	jobSet JobSet,
) bool {
	return s.set.Equal(makeGenericJobSet(jobSet.List()))
}

func (s jobSet) Delete(Job *Job) {
	s.set.Delete(Job)
}

func (s jobSet) Union(set JobSet) JobSet {
	return NewJobSet(append(s.List(), set.List()...)...)
}

func (s jobSet) Difference(set JobSet) JobSet {
	newSet := s.set.Difference(makeGenericJobSet(set.List()))
	return jobSet{set: newSet}
}

func (s jobSet) Intersection(set JobSet) JobSet {
	newSet := s.set.Intersection(makeGenericJobSet(set.List()))
	var jobList []*Job
	for _, obj := range newSet.List() {
		jobList = append(jobList, obj.(*Job))
	}
	return NewJobSet(jobList...)
}
