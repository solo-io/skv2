package schemagen

type ValidationSchemaOptions struct {
	// Whether to remove descriptions from validation schemas
	// Default: false
	//
	// NOTE: I'd prefer a positive field name (ie includeDescriptions)
	//	but I wanted to avoid changing the default behavior
	RemoveDescriptionsFromSchema bool

	// Whether to assign Enum fields the `x-kubernetes-int-or-string` property
	// which allows the value to either be an integer or a string
	// If this is false, only string values are allowed
	// Default: false
	EnumAsIntOrString bool

	// A list of messages (core.solo.io.Status) whose validation schema should
	// not be generated
	MessagesWithEmptySchema []string
}

// prevent k8s from validating proto.Any fields (since it's unstructured)
func removeProtoAnyValidation(d map[string]interface{}, propertyField string) {
	for _, v := range d {
		values, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		desc, ok := values["properties"]
		properties, isObj := desc.(map[string]interface{})
		// detect proto.Any field from presence of [propertyField] as field under "properties"
		if !ok || !isObj || properties[propertyField] == nil {
			removeProtoAnyValidation(values, propertyField)
			continue
		}
		// remove "properties" value
		delete(values, "properties")
		// remove "required" value
		delete(values, "required")
		// x-kubernetes-preserve-unknown-fields allows for unknown fields from a particular node
		// see https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#specifying-a-structural-schema
		values["x-kubernetes-preserve-unknown-fields"] = true
	}
}
