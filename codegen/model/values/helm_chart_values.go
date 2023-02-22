package values

import (
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

// used to document the values structure of the generated Helm Chart
type UserHelmValues struct {
	Operators        []UserOperatorValues
	CustomValues     interface{}
	ValuesInlineDocs *UserValuesInlineDocs
	JsonSchema       UserJsonSchema
}

type UserOperatorValues struct {
	Name         string
	ValuePath    string
	Values       UserValues
	CustomValues interface{}
}

type UserValuesInlineDocs struct {
	LineLengthLimit int
}

func (u *UserValuesInlineDocs) Enabled() bool {
	return u != nil
}

func (u *UserValuesInlineDocs) LineLength() int {
	if u == nil {
		return 0
	}
	return u.LineLengthLimit
}

type UserJsonSchema struct {
	CustomTypeMapper func(reflect.Type, map[string]interface{}) interface{}
}

// document values structure for an operator
type UserValues struct {
	UserContainerValues `json:",inline"`

	// Required to have an interface value in order to use the `index` function in the template
	Sidecars       map[string]UserContainerValues `json:"sidecars" desc:"Configuration for the deployed containers."`
	FloatingUserID bool                           `json:"floatingUserId" desc:"Allow the pod to be assigned a dynamic user ID."`
	RunAsUser      uint32                         `json:"runAsUser" desc:"Static user ID to run the containers as. Unused if floatingUserId is 'true'."`
	ServiceType    v1.ServiceType                 `json:"serviceType" desc:"Specify the service type. Can be either \"ClusterIP\", \"NodePort\", \"LoadBalancer\", or \"ExternalName\"."`
	ServicePorts   map[string]uint32              `json:"ports" desc:"Specify service ports as a map from port name to port number."`

	// Overrides which can be set by the user
	DeploymentOverrides *appsv1.Deployment `json:"deploymentOverrides,omitempty" desc:"Provide arbitrary overrides for the component's [deployment template](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/)" omitChildren:"true"`
	ServiceOverrides    *v1.Service        `json:"serviceOverrides,omitempty" desc:"Provide arbitrary overrides for the component's [service template](https://kubernetes.io/docs/reference/kubernetes-api/service-resources/service-v1/)." omitChildren:"true"`

	Enabled bool `json:"enabled" desc:"Enables or disables creation of the operator deployment/service"`
}

// document values structure for a container
type UserContainerValues struct {
	Image           Image                    `json:"image" desc:"Specify the container image"`
	Env             []v1.EnvVar              `json:"env" desc:"Specify environment variables for the container. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#envvarsource-v1-core) for specification details." omitChildren:"true"`
	Resources       *v1.ResourceRequirements `json:"resources,omitempty" desc:"Specify container resource requirements. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#resourcerequirements-v1-core) for specification details."`
	SecurityContext *v1.SecurityContext      `json:"securityContext,omitempty" desc:"Specify container security context. Set to 'false' to omit the security context entirely. See the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#securitycontext-v1-core) for specification details."`
}

// this image type is shared with the skv2 Chart model
type Image struct {
	Tag        string        `json:"tag,omitempty"  desc:"Tag for the container."`
	Repository string        `json:"repository,omitempty"  desc:"Image name (repository)."`
	Registry   string        `json:"registry,omitempty" desc:"Image registry."`
	PullPolicy v1.PullPolicy `json:"pullPolicy,omitempty"  desc:"Image pull policy."`
	PullSecret string        `json:"pullSecret,omitempty" desc:"Image pull secret."`
}
