package model

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/solo-io/skv2/codegen/model/values"

	"github.com/iancoleman/strcase"
	"github.com/solo-io/skv2/codegen/doc"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

type Chart struct {
	Operators []Operator

	// filter out the template based on its output name
	FilterTemplate func(outPath string) bool

	// outPath: content template map
	CustomTemplates CustomTemplates

	Values interface{}

	// goes into the chart.yaml
	Data Data

	// inline string of custom _helpers.tpl which can be provided for the generated chart
	HelpersTpl string

	// Only generate the chart
	ChartOnly bool

	// if specified, generate reference docs for the chart values to the provided filename
	ValuesReferenceDocs ValuesReferenceDocs

	// if specificed, generate inline documentation for the values in chart's values.yaml files
	ValuesInlineDocs *ValuesInlineDocs

	// if specificed, values.schema.json will be generated with a JSON Schema that
	// imposes structure on the values.yaml file
	JsonSchema *JsonSchema
}

type ValuesReferenceDocs struct {
	Title    string
	Filename string
}

type ValuesInlineDocs struct {
	// if specified, inline field documentation comments will be wrapped at many characters
	LineLengthLimit int
}

type JsonSchema struct {
	// (Optional) will be called to override the default json schema mapping
	// for the type. This is useful for types that also override default json/yaml
	// serialization behaviour. It accepts the json schema as a map and is
	// expected to return a value that can serialize to the json schema or null if
	// there is no custom mapping for this type
	CustomTypeMapper func(reflect.Type, map[string]interface{}) interface{}
}

type Operator struct {
	Name string

	// (Optional) To change the name referenced in the values file. If not specified a camelcase version of name is used
	ValuesFileNameOverride string

	// (Optional) For nesting operators in values API (e.g. a value of "agent" would place an operator at "agent.<operatorName>" )
	ValuePath string

	// deployment config
	Deployment Deployment

	// these populate the generated ClusterRole for the operator
	Rbac []rbacv1.PolicyRule

	// if at least one port is defined, create a Service for it
	Service Service

	// Custom values to include at operator level
	Values interface{}

	// (Optional) If this operator depends on another operator being enabled,
	// the name of the other operator can be included in this list. This operator
	// will not be provisioned unless both are enabled (by having values.enabled = true)
	EnabledDependsOn []string
}

func (o Operator) FormattedName() string {
	formattedName := strcase.ToLowerCamel(o.Name)
	if o.ValuesFileNameOverride != "" {
		formattedName = strcase.ToLowerCamel(o.ValuesFileNameOverride)
	}
	return formattedName
}

// values for Deployment template
type Deployment struct {
	// TODO support use of a DaemonSet instead of a Deployment
	UseDaemonSet bool
	Container
	Sidecars                    []Sidecar
	Volumes                     []v1.Volume
	CustomPodLabels             map[string]string
	CustomPodAnnotations        map[string]string
	CustomDeploymentLabels      map[string]string
	CustomDeploymentAnnotations map[string]string
}

// values for a container
type Container struct {
	// not configurable via helm values
	Args           []string
	VolumeMounts   []v1.VolumeMount
	ReadinessProbe *v1.Probe
	LivenessProbe  *v1.Probe

	Image           Image
	Env             []v1.EnvVar
	Resources       *v1.ResourceRequirements
	SecurityContext *v1.SecurityContext
}

// sidecars require a container config and a unique name
type Sidecar struct {
	Container
	Name string
}

// values for struct template
type Service struct {
	Type              v1.ServiceType
	Ports             []ServicePort
	CustomLabels      map[string]string
	CustomAnnotations map[string]string
}

type ServicePort struct {
	// The name of this port within the service.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" desc:"The name of this port within the service."`

	// The default port that will be exposed by this service.
	DefaultPort int32 `json:"port" protobuf:"varint,3,opt,name=port" desc:"The default port that will be exposed by this service."`
}

type Image = values.Image

// Helm chart dependency
type Dependency struct {
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
	Repository string `json:"repository,omitempty"`
	Condition  string `json:"condition,omitempty"`
	Alias      string `json:"alias,omitempty"`
}

type Data struct {
	ApiVersion   string       `json:"apiVersion,omitempty"`
	Description  string       `json:"description,omitempty"`
	Name         string       `json:"name,omitempty"`
	Version      string       `json:"version,omitempty"`
	Home         string       `json:"home,omitempty"`
	Icon         string       `json:"icon,omitempty"`
	Sources      []string     `json:"sources,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

func makeContainerDocs(c Container) values.UserContainerValues {
	return values.UserContainerValues{
		Image:           c.Image,
		Env:             c.Env,
		Resources:       c.Resources,
		SecurityContext: c.SecurityContext,
	}
}

func (c Chart) BuildChartValues() values.UserHelmValues {
	helmValues := values.UserHelmValues{CustomValues: c.Values}

	for _, operator := range c.Operators {
		servicePorts := map[string]uint32{}
		for _, port := range operator.Service.Ports {
			servicePorts[port.Name] = uint32(port.DefaultPort)
		}
		sidecars := map[string]values.UserContainerValues{}
		for _, sidecar := range operator.Deployment.Sidecars {
			sidecars[strcase.ToLowerCamel(sidecar.Name)] = makeContainerDocs(sidecar.Container)
		}

		helmValues.Operators = append(helmValues.Operators, values.UserOperatorValues{
			Name:                   operator.Name,
			ValuesFileNameOverride: operator.ValuesFileNameOverride,
			ValuePath:              operator.ValuePath,
			Values: values.UserValues{
				UserContainerValues: makeContainerDocs(operator.Deployment.Container),
				Sidecars:            sidecars,
				FloatingUserID:      false,
				RunAsUser:           10101,
				ServiceType:         operator.Service.Type,
				ServicePorts:        servicePorts,
				Enabled:             true,
			},
			CustomValues: operator.Values,
		})
	}

	if c.ValuesInlineDocs != nil {
		helmValues.ValuesInlineDocs = &values.UserValuesInlineDocs{
			LineLengthLimit: c.ValuesInlineDocs.LineLengthLimit,
		}
	}

	if c.JsonSchema != nil {
		helmValues.JsonSchema.CustomTypeMapper = c.JsonSchema.CustomTypeMapper
	}

	return helmValues
}

func (c Chart) GenerateHelmDoc() string {
	helmValues := c.BuildChartValues()

	// generate documentation for custom values
	helmValuesForDoc := doc.GenerateHelmValuesDoc(helmValues.CustomValues, "", "")

	// generate documentation for operator values
	for _, operatorWithValues := range helmValues.Operators {

		name := operatorWithValues.FormattedName()
		values := operatorWithValues.Values

		// clear image tag so it doesn't show build time commit hashes
		values.Image.Tag = ""
		for name, container := range values.Sidecars {
			container.Image.Tag = ""
			values.Sidecars[name] = container
		}

		keyPath := name
		if operatorWithValues.ValuePath != "" {
			keyPath = fmt.Sprintf("%s.%s", operatorWithValues.ValuePath, name)
		}

		helmValuesForDoc = append(helmValuesForDoc, doc.GenerateHelmValuesDoc(operatorWithValues.CustomValues, keyPath, fmt.Sprintf("Configuration for the %s deployment.", name))...)
		helmValuesForDoc = append(helmValuesForDoc, doc.GenerateHelmValuesDoc(values, keyPath, fmt.Sprintf("Configuration for the %s deployment.", name))...)
	}

	helmValuesForDoc = doc.removeDuplicates(helmValuesForDoc)

	// alphabetize all values
	sort.Slice(helmValuesForDoc, func(i, j int) bool {
		return helmValuesForDoc[i].Key < helmValuesForDoc[j].Key
	})

	return helmValuesForDoc.ToMarkdown(c.ValuesReferenceDocs.Title)
}