
---
title: "Name Override Chart"
description: Reference for Helm values.
weight: 2
---

|Option|Type|Default Value|Description|
|------|----|-----------|-------------|
|overrideName|struct| |Configuration for the overrideName deployment.|
|overrideName|struct| ||
|overrideName.image|struct| |Container image.|
|overrideName.image.tag|string| |Version tag for the container image.|
|overrideName.image.repository|string|painter|Image name (repository).|
|overrideName.image.registry|string|quay.io/solo-io|Image registry.|
|overrideName.image.pullPolicy|string|IfNotPresent|Image pull policy.|
|overrideName.image.pullSecret|string| |Image pull secret.|
|overrideName.env[]|slice|null|Environment variables for the container. For more info, see the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#envvarsource-v1-core).|
|overrideName.extraEnvs|struct| |Extra environment variables for the container|
|overrideName.resources|struct| |Container resource requirements. For more info, see the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#resourcerequirements-v1-core).|
|overrideName.securityContext|struct| |Container security context. Set to 'false' to omit the security context entirely. For more info, see the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core).|
|overrideName.sidecars|map[string, struct]|{}|Optional configuration for the deployed containers.|
|overrideName.sidecars.<MAP_KEY>|struct| |Optional configuration for the deployed containers.|
|overrideName.sidecars.<MAP_KEY>.image|struct| |Container image.|
|overrideName.sidecars.<MAP_KEY>.image.tag|string| |Version tag for the container image.|
|overrideName.sidecars.<MAP_KEY>.image.repository|string| |Image name (repository).|
|overrideName.sidecars.<MAP_KEY>.image.registry|string| |Image registry.|
|overrideName.sidecars.<MAP_KEY>.image.pullPolicy|string| |Image pull policy.|
|overrideName.sidecars.<MAP_KEY>.image.pullSecret|string| |Image pull secret.|
|overrideName.sidecars.<MAP_KEY>.env[]|slice| |Environment variables for the container. For more info, see the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#envvarsource-v1-core).|
|overrideName.sidecars.<MAP_KEY>.extraEnvs|struct| |Extra environment variables for the container|
|overrideName.sidecars.<MAP_KEY>.resources|struct| |Container resource requirements. For more info, see the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#resourcerequirements-v1-core).|
|overrideName.sidecars.<MAP_KEY>.securityContext|struct| |Container security context. Set to 'false' to omit the security context entirely. For more info, see the [Kubernetes documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#securitycontext-v1-core).|
|overrideName.floatingUserId|bool|false|Allow the pod to be assigned a dynamic user ID. Required for OpenShift installations.|
|overrideName.runAsUser|uint32|10101|Static user ID to run the containers as. Unused if floatingUserId is 'true'.|
|overrideName.serviceType|string| |Kubernetes service type. Can be either "ClusterIP", "NodePort", "LoadBalancer", or "ExternalName".|
|overrideName.ports|map[string, uint32]|{}|Service ports as a map from port name to port number.|
|overrideName.ports.<MAP_KEY>|uint32| |Service ports as a map from port name to port number.|
|overrideName.deploymentOverrides|struct| |Arbitrary overrides for the component's [deployment template](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/deployment-v1/)|
|overrideName.serviceOverrides|struct| |Arbitrary overrides for the component's [service template](https://kubernetes.io/docs/reference/kubernetes-api/service-resources/service-v1/).|
|overrideName.enabled|bool|true|Enable creation of the deployment/service.|
