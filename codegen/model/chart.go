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
	Args         []string         `json:"-"`
	VolumeMounts []v1.VolumeMount `json:"-"`

	Image     Image                    `json:"image" desc:"Specify the container image"`
	Env       []v1.EnvVar              `json:"env" desc:"Specify environment variables for the container. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#envvarsource-v1-core) for specification details." omitChildren:"true"`
	Resources *v1.ResourceRequirements `json:"resources,omitempty" desc:"Specify container resource requirements. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#resourcerequirements-v1-core) for specification details."`
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

type Image struct {
	Tag        string        `json:"tag,omitempty"  desc:"Tag for the container."`
	Repository string        `json:"repository,omitempty"  desc:"Image name (repository)."`
	Registry   string        `json:"registry,omitempty" desc:"Image registry."`
	PullPolicy v1.PullPolicy `json:"pullPolicy,omitempty"  desc:"Image pull policy."`
	PullSecret string        `json:"pullSecret,omitempty" desc:"Image pull secret."`
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
	Operators    []OperatorValues
	CustomValues interface{}
}

type OperatorValues struct {
	Name   string
	Values Values
}

type Values struct {
	Container `json:",inline"`
	// Required to have an interface value in order to use the `index` function in the template
	Sidecars     map[string]Container `json:"sidecars" desc:"Configuration for the deployed containers."`
	ServiceType  v1.ServiceType       `json:"serviceType" desc:"Specify the service type. Can be either \"ClusterIP\", \"NodePort\", \"LoadBalancer\", or \"ExternalName\"."`
	ServicePorts map[string]uint32    `json:"ports" desc:"Specify service ports as a map from port name to port number."`

	ExtraPodLabels             map[string]string `json:"extraPodLabels,omitempty" desc:"Custom labels for the pod"`
	ExtraPodAnnotations        map[string]string `json:"extraPodAnnotations,omitempty" desc:"Custom annotations for the pod"`
	ExtraDeploymentLabels      map[string]string `json:"extraDeploymentLabels,omitempty" desc:"Custom labels for the deployment"`
	ExtraDeploymentAnnotations map[string]string `json:"extraDeploymentAnnotations,omitempty" desc:"Custom annotations for the deployment"`
	ExtraServiceLabels         map[string]string `json:"extraServiceLabels,omitempty" desc:"Custom labels for the service"`
	ExtraServiceAnnotations    map[string]string `json:"extraServiceAnnotations,omitempty" desc:"Custom annotations for the service"`
}

func (c Chart) BuildChartValues() HelmValues {
	values := HelmValues{CustomValues: c.Values}

	for _, operator := range c.Operators {
		servicePorts := map[string]uint32{}
		for _, port := range operator.Service.Ports {
			servicePorts[port.Name] = uint32(port.DefaultPort)
		}
		sidecars := map[string]Container{}
		for _, sidecar := range operator.Deployment.Sidecars {
			sidecars[sidecar.Name] = sidecar.Container
		}

		values.Operators = append(values.Operators, OperatorValues{
			Name: operator.Name,
			Values: Values{
				Container:                  operator.Deployment.Container,
				Sidecars:                   sidecars,
				ServiceType:                operator.Service.Type,
				ServicePorts:               servicePorts,
				ExtraPodLabels:             operator.Deployment.CustomPodLabels,
				ExtraPodAnnotations:        operator.Deployment.CustomPodAnnotations,
				ExtraDeploymentLabels:      operator.Deployment.CustomDeploymentLabels,
				ExtraDeploymentAnnotations: operator.Deployment.CustomDeploymentAnnotations,
				ExtraServiceLabels:         operator.Service.CustomLabels,
				ExtraServiceAnnotations:    operator.Service.CustomAnnotations,
			},
		})
	}

	return values
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
