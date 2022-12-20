package ezkube_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/solo-io/skv2/pkg/ezkube"
)

var _ = Describe("ResourceId conversion", func() {
	DescribeTable("creating keys from resourceids",
		func(name string, namespace string, cluster string, deprecated bool, separator string, expectedKey string) {
			var id ezkube.ResourceId
			if cluster != "" {
				if deprecated {
					id = testDeprecatedClusterResourceId{
						name:      name,
						namespace: namespace,
						cluster:   cluster,
					}
				} else {
					id = testClusterResourceId{
						name:      name,
						namespace: namespace,
						annotations: map[string]string{
							ezkube.ClusterAnnotation: cluster,
						},
					}
				}
			} else {
				id = testResourceId{
					name:      name,
					namespace: namespace,
				}
			}
			key := ezkube.KeyWithSeparator(id, separator)
			Expect(key).To(Equal(expectedKey))
		},
		Entry("can create key for resource id", "a", "b", "", false, "/", "a/b/"),
		Entry("can create key for cluster resource id", "a", "b", "c", false, "/", "a/b/c"),
		Entry("can create key for deprecated cluster resource id", "a", "b", "c", true, "/", "a/b/c"),
	)
	DescribeTable("converting keys to resourceids",
		func(key string, separator string, expectedName string, expectedNamespace string, expectedCluster string, expectedError string) {
			resource, err := ezkube.ResourceIdFromKeyWithSeparator(key, separator)
			if expectedError != "" {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(expectedError))
			} else {
				Expect(err).NotTo(HaveOccurred())
				Expect(resource.GetName()).To(Equal(expectedName))
				Expect(resource.GetNamespace()).To(Equal(expectedNamespace))
				if expectedCluster != "" {
					clusterResource, ok := resource.(ezkube.ClusterResourceId)
					Expect(ok).To(BeTrue())
					Expect(ezkube.GetClusterName(clusterResource)).To(Equal(expectedCluster))
				}
			}
		},
		Entry("not enough parts", "a", "/", "", "", "", "could not convert key a with separator / into resource id; unexpected number of parts: 1"),
		Entry("too many parts", "a/b/c/d", "/", "", "", "", "could not convert key a/b/c/d with separator / into resource id; unexpected number of parts: 4"),
		Entry("can convert key to resource id", "a/b", "/", "a", "b", "", ""),
		Entry("can convert key with trailing separator to resource id", "a/b/", "/", "a", "b", "", ""),
		Entry("can convert key to cluster resource id", "a/b/c", "/", "a", "b", "c", ""),
	)
})

type testResourceId struct {
	name, namespace string
}

func (id testResourceId) GetName() string {
	return id.name
}

func (id testResourceId) GetNamespace() string {
	return id.namespace
}

type testClusterResourceId struct {
	name, namespace string
	annotations     map[string]string
}

func (id testClusterResourceId) GetName() string {
	return id.name
}

func (id testClusterResourceId) GetNamespace() string {
	return id.namespace
}

func (id testClusterResourceId) GetAnnotations() map[string]string {
	return id.annotations
}

type testDeprecatedClusterResourceId struct {
	name, namespace, cluster string
}

func (id testDeprecatedClusterResourceId) GetName() string {
	return id.name
}

func (id testDeprecatedClusterResourceId) GetNamespace() string {
	return id.namespace
}

func (id testDeprecatedClusterResourceId) GetClusterName() string {
	return id.cluster
}