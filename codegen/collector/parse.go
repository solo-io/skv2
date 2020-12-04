package collector

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type DescriptorWithPath struct {
	*descriptor.FileDescriptorProto
	ProtoFilePath string
}

func (file *DescriptorWithPath) GetMessage(typeName string) *descriptor.DescriptorProto {
	for _, msg := range file.GetMessageType() {
		if msg.GetName() == typeName {
			return msg
		}
		// nes := file.GetNestedMessage(msg, strings.TrimPrefix(typeName, msg.GetName()+"."))
		// if nes != nil {
		// 	return nes
		// }
	}
	return nil
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
