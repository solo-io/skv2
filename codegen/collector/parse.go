package collector

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

type DescriptorWithPath struct {
	*descriptor.FileDescriptorProto
	ProtoFilePath string
}

func filterDuplicateDescriptors(descriptors []*DescriptorWithPath) []*DescriptorWithPath {
	var uniqueDescriptors []*DescriptorWithPath
	for _, desc := range descriptors {
		unique, matchingDesc := isUnique(desc, uniqueDescriptors)
		// if this proto file first came in as an import, but later as a solo-kit project proto,
		// ensure the original proto gets updated with the correct proto file path
		if !unique && matchingDesc.ProtoFilePath == "" {
			matchingDesc.ProtoFilePath = desc.ProtoFilePath
		}
		if unique {
			uniqueDescriptors = append(uniqueDescriptors, desc)
		}
	}
	return uniqueDescriptors
}

// If it finds a matching proto, also returns the matching proto's file descriptor
func isUnique(desc *DescriptorWithPath, descriptors []*DescriptorWithPath) (bool, *DescriptorWithPath) {
	for _, existing := range descriptors {
		existingCopy := proto.Clone(existing.FileDescriptorProto).(*descriptor.FileDescriptorProto)
		existingCopy.Name = desc.Name
		if proto.Equal(existingCopy, desc.FileDescriptorProto) {
			return false, existing
		}
	}
	return true, nil
}
