package model

import (
	"fmt"

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
}

type Operator struct {
	Name string

	// deployment config
	Deployment Deployment

	// these populate the generated ClusterRole for the operator
	Rbac []rbacv1.PolicyRule
	// these populate the generated ClusterRole for the operator
	Volumes []v1.Volume
	// mount these volumes to the operator container
	VolumeMounts []v1.VolumeMount
	// set these environment variables on the operator container
	Env []v1.EnvVar
	// add a manifest for each configmap
	ConfigMaps []v1.ConfigMap

	// if at least one port is defined, create a service for the
	Service Service

	// args for the container
	Args []string
}

// values for Deployment template
type Deployment struct {
	// use a DaemonSet instead of a Deployment
	UseDaemonSet bool                     `json:"-"`
	Image        Image                    `json:"image,omitempty"`
	Resources    *v1.ResourceRequirements `json:"resources,omitempty"`
}

// values for struct template
type Service struct {
	Type  v1.ServiceType
	Ports []ServicePort
}
type ServicePort struct {
	// The name of this port within the service.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// The default port that will be exposed by this service.
	DefaultPort int32 `json:"port" protobuf:"varint,3,opt,name=port"`
}

type Image struct {
	Tag        string        `json:"tag,omitempty"  desc:"tag for the container"`
	Repository string        `json:"repository,omitempty"  desc:"image name (repository) for the container."`
	Registry   string        `json:"registry,omitempty" desc:"image prefix/registry e.g. (quay.io/solo-io)"`
	PullPolicy v1.PullPolicy `json:"pullPolicy,omitempty"  desc:"image pull policy for the container"`
	PullSecret string        `json:"pullSecret,omitempty" desc:"image pull policy for the container "`
}

// Helm chart dependency
type Dependency struct {
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
	Repository string `json:"repository,omitempty"`
	Condition  string `json:"condition,omitempty"`
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

// fields exposed as Helm values
type HelmValues struct {
	Operators    map[string]OperatorHelmValues
	CustomValues interface{}
}

type OperatorHelmValues struct {
	Image        Image                    `json:"image" desc:"Specify the deployment image"`
	Resources    *v1.ResourceRequirements `json:"resources" desc:"Specify deployment resource requirements"`
	ServiceType  v1.ServiceType           `json:"serviceType" desc:"Specify the service type"`
	ServicePorts []ServicePort            `json:"ports" desc:"Specify service ports"`
	Env          []v1.EnvVar              `json:"env" desc:"Specify environment variables for the deployment"`
}

func (c Chart) BuildChartValues() HelmValues {
	values := HelmValues{}
	values.CustomValues = c.Values
	values.Operators = map[string]OperatorHelmValues{}

	for _, operator := range c.Operators {
		values.Operators[operator.Name] = OperatorHelmValues{
			Image:        operator.Deployment.Image,
			Resources:    operator.Deployment.Resources,
			ServiceType:  operator.Service.Type,
			ServicePorts: operator.Service.Ports,
			Env:          operator.Env,
		}
	}

	return values
}

func (c Chart) GenerateHelmDoc(title string) string {
	helmValues := c.BuildChartValues()

	// generate documentation for custom values
	helmValuesForDoc := doc.GenerateHelmValuesDoc(helmValues.CustomValues, "", "")

	// generate documentation for operator values
	for name, values := range helmValues.Operators {
		name = strcase.ToLowerCamel(name)

		// clear image tag so it doesn't show build time commit hashes
		values.Image.Tag = ""

		helmValuesForDoc = append(helmValuesForDoc, doc.GenerateHelmValuesDoc(values, name, fmt.Sprintf("Configuration for the %s deployment.", name))...)
	}

	return helmValuesForDoc.ToMarkdown(title)
}
