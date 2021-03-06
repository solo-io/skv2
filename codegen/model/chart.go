package model

import (
	"fmt"

	"github.com/solo-io/skv2/codegen/model/values"

	"github.com/iancoleman/strcase"
	"github.com/solo-io/skv2/codegen/doc"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// Chart provides the input data structure for generating Helm charts from the skv2 chart "meta-templates" (templates whose outputs are templates and other files used by generated Helm charts)
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
}

type Operator struct {
	Name string

	// deployment config
	Deployment Deployment

	// these populate the generated ClusterRole for the operator
	Rbac []rbacv1.PolicyRule

	// add a manifest for each configmap
	ConfigMaps []v1.ConfigMap

	// if at least one port is defined, create a service for the
	Service Service
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
	Args            []string
	VolumeMounts    []v1.VolumeMount
	ReadinessProbe  *v1.Probe
	LivenessProbe   *v1.Probe
	SecurityContext *v1.SecurityContext

	Image     Image
	Env       []v1.EnvVar
	Resources *v1.ResourceRequirements
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
		Image:     c.Image,
		Env:       c.Env,
		Resources: c.Resources,
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
			sidecars[sidecar.Name] = makeContainerDocs(sidecar.Container)
		}

		helmValues.Operators = append(helmValues.Operators, values.UserOperatorValues{
			Name: operator.Name,
			Values: values.UserValues{
				UserContainerValues: makeContainerDocs(operator.Deployment.Container),
				Sidecars:            sidecars,
				ServiceType:         operator.Service.Type,
				ServicePorts:        servicePorts,
			},
		})
	}

	return helmValues
}

func (c Chart) GenerateHelmDoc(title string) string {
	helmValues := c.BuildChartValues()

	// generate documentation for custom values
	helmValuesForDoc := doc.GenerateHelmValuesDoc(helmValues.CustomValues, "", "")

	// generate documentation for operator values
	for _, operatorWithValues := range helmValues.Operators {
		name := strcase.ToLowerCamel(operatorWithValues.Name)
		values := operatorWithValues.Values

		// clear image tag so it doesn't show build time commit hashes
		values.Image.Tag = ""
		for name, container := range values.Sidecars {
			container.Image.Tag = ""
			values.Sidecars[name] = container
		}

		helmValuesForDoc = append(helmValuesForDoc, doc.GenerateHelmValuesDoc(values, name, fmt.Sprintf("Configuration for the %s deployment.", name))...)
	}

	return helmValuesForDoc.ToMarkdown(title)
}
