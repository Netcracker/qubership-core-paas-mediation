package v2

import (
	"strings"
	"time"

	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	pmTypes "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/types"
)

type AppVersionData struct {
	AppName    string    `json:"appName"`
	AppVersion string    `json:"appVersion"`
	DeployTime time.Time `json:"deployTime"`
}

type AnnotationResource struct {
	ResourceName    string `json:"resourceName"`
	Namespace       string `json:"namespace"`
	AnnotationValue string `json:"annotationValue"`
}

type AnnotationKey struct {
	Annotation   string `json:"annotation"`
	ResourceType string `json:"resourceType"`
}

type RolloutDeploymentBody struct {
	DeploymentNames []string `json:"deployment_names"`
	Parallel        bool     `json:"parallel,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Metadata struct {
	Kind            string            `json:"kind"`
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace,omitempty"`
	UID             string            `json:"uid,omitempty"`
	Generation      int64             `json:"generation,omitempty"`
	ResourceVersion string            `json:"resourceVersion,omitempty"`
	Annotations     map[string]string `json:"annotations,omitempty" swaggertype:"object,string"`
	Labels          map[string]string `json:"labels,omitempty" swaggertype:"object,string"`
}

type Namespace struct {
	Metadata `json:"metadata"`
}

type ConfigMap struct {
	Metadata `json:"metadata"`
	Data     map[string]string `json:"data" swaggertype:"object,string"`
}

type (
	Route struct {
		Metadata `json:"metadata"`
		Spec     RouteSpec `json:"spec"`
	}

	RouteSpec struct {
		Host             string    `json:"host"`
		PathType         string    `json:"pathType"`
		Path             string    `json:"path"`
		Service          Target    `json:"to"`
		Port             RoutePort `json:"port"`
		IngressClassName *string   `json:"ingressClassName"`
	}

	RoutePort struct {
		TargetPort int32 `json:"targetPort"`
	}

	Target struct {
		Name string `json:"name"`
	}
)

type (
	Service struct {
		Metadata `json:"metadata"`
		Spec     ServiceSpec `json:"spec"`
	}

	ServiceSpec struct {
		Ports     []ServicePort     `json:"ports"`
		Selector  map[string]string `json:"selector" swaggertype:"object,string"`
		ClusterIP string            `json:"clusterIP"`
		Type      string            `json:"type"`
	}

	ServicePort struct {
		Name       string `json:"name"`
		Protocol   string `json:"protocol"`
		Port       int32  `json:"port"`
		TargetPort int32  `json:"targetPort"`
		NodePort   int32  `json:"nodePort,omitempty"`
	}
)

type (
	Deployment struct {
		Metadata `json:"metadata"`
		Spec     DeploymentSpec   `json:"spec"`
		Status   DeploymentStatus `json:"status"`
	}

	DeploymentSpec struct {
		Replicas             *int32             `json:"replicas"`
		RevisionHistoryLimit *int32             `json:"revisionHistoryLimit"`
		Template             PodTemplateSpec    `json:"template"`
		Strategy             DeploymentStrategy `json:"strategy"`
	}

	DeploymentStrategy struct {
		Type string `json:"type"`
	}

	PodTemplateSpec struct {
		Metadata TemplateMetadata `json:"metadata"`
		Spec     PodSpec          `json:"spec"`
	}

	TemplateMetadata struct {
		Labels map[string]string `json:"labels,omitempty" swaggertype:"object,string"`
	}

	DeploymentStatus struct {
		AvailableReplicas  int32                 `json:"availableReplicas"`
		Conditions         []DeploymentCondition `json:"conditions"`
		ObservedGeneration int64                 `json:"observedGeneration"`
		ReadyReplicas      int32                 `json:"readyReplicas"`
		Replicas           int32                 `json:"replicas"`
		UpdatedReplicas    int32                 `json:"updatedReplicas"`
	}

	DeploymentCondition struct {
		LastTransitionTime *string `json:"lastTransitionTime"`
		LastUpdateTime     *string `json:"lastUpdateTime"`
		Message            string  `json:"message"`
		Reason             string  `json:"reason"`
		Status             string  `json:"status"`
		Type               string  `json:"type"`
	}
)

type (
	Pod struct {
		Metadata `json:"metadata"`
		Spec     PodSpec   `json:"spec"`
		Status   PodStatus `json:"status"`
	}

	PodSpec struct {
		Volumes                       []SpecVolume    `json:"volumes,omitempty"`
		Containers                    []SpecContainer `json:"containers"`
		RestartPolicy                 string          `json:"restartPolicy,omitempty"`
		TerminationGracePeriodSeconds int64           `json:"terminationGracePeriodSeconds,omitempty"`
		DnsPolicy                     string          `json:"dnsPolicy,omitempty"`
		NodeName                      string          `json:"nodeName,omitempty"`
	}

	SpecVolume struct {
		Name   string         `json:"name"`
		Secret *VolumesSecret `json:"secret,omitempty"`
	}

	VolumesSecret struct {
		SecretName  string `json:"secretName,omitempty"`
		DefaultMode int32  `json:"defaultMode,omitempty"`
	}

	SpecContainer struct {
		Name            string                 `json:"name"`
		Image           string                 `json:"image,omitempty"`
		Ports           []ContainerPort        `json:"ports,omitempty"`
		Env             []ContainerEnv         `json:"env,omitempty"`
		Resources       ContainerResources     `json:"resources,omitempty"`
		VolumeMounts    []ContainerVolumeMount `json:"volumeMounts,omitempty"`
		ImagePullPolicy string                 `json:"imagePullPolicy,omitempty"`
		Args            []string               `json:"args"`
	}

	ContainerPort struct {
		ContainerPort int32  `json:"containerPort"`
		Protocol      string `json:"protocol,omitempty"`
		Name          string `json:"name"`
	}

	ContainerEnv struct {
		Name      string     `json:"name"`
		Value     string     `json:"value,omitempty"`
		ValueFrom *ValueFrom `json:"valueFrom,omitempty"`
	}

	ValueFrom struct {
		FieldRef     *FieldRef     `json:"fieldRef,omitempty"`
		SecretKeyRef *SecretKeyRef `json:"secretKeyRef,omitempty"`
	}

	SecretKeyRef struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}

	FieldRef struct {
		APIVersion string `json:"apiVersion"`
		FieldPath  string `json:"fieldPath"`
	}

	ContainerResources struct {
		Limits   CpuMemoryResource `json:"limits,omitempty"`
		Requests CpuMemoryResource `json:"requests,omitempty"`
	}

	CpuMemoryResource struct {
		Cpu    string `json:"cpu"`
		Memory string `json:"memory"`
	}

	ContainerVolumeMount struct {
		Name      string `json:"name"`
		MountPath string `json:"mountPath"`
		ReadOnly  bool   `json:"readOnly,omitempty"`
	}

	PodStatus struct {
		Phase             string            `json:"phase,omitempty"`
		Conditions        []StatusCondition `json:"conditions,omitempty"`
		HostIP            string            `json:"hostIP,omitempty"`
		PodIP             string            `json:"podIP,omitempty"`
		StartTime         *string           `json:"startTime,omitempty"`
		ContainerStatuses []ContainerStatus `json:"containerStatuses"`
	}

	StatusCondition struct {
		Type               string  `json:"type"`
		Status             string  `json:"status"`
		LastProbeTime      *string `json:"lastProbeTime,omitempty"`
		LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	}

	ContainerStatus struct {
		Name         string         `json:"name"`
		State        ContainerState `json:"state,omitempty"`
		LastState    ContainerState `json:"lastState,omitempty"`
		Ready        bool           `json:"ready"`
		RestartCount int32          `json:"restartCount"`
		Image        string         `json:"image"`
		ImageID      string         `json:"imageID"`
		ContainerID  string         `json:"containerID,omitempty"`
	}

	ContainerState struct {
		Running    *ContainerStateRunning    `json:"running,omitempty"`
		Terminated *ContainerStateTerminated `json:"terminated,omitempty"`
		Waiting    *ContainerStateWaiting    `json:"waiting,omitempty"`
	}

	ContainerStateRunning struct {
		StartedAt *string `json:"startedAt,omitempty"`
	}

	ContainerStateTerminated struct {
		ContainerID string  `json:"containerID,omitempty"`
		ExitCode    int32   `json:"exitCode"`
		FinishedAt  *string `json:"finishedAt,omitempty"`
		Reason      string  `json:"reason,omitempty"`
		StartedAt   *string `json:"startedAt,omitempty"`
	}

	ContainerStateWaiting struct {
		Message string `json:"message,omitempty"`
		Reason  string `json:"reason,omitempty"`
	}
)

type DeploymentFamilyVersion struct {
	AppName          string `json:"app_name"`
	AppVersion       string `json:"app_version"`
	Name             string `json:"name"`
	FamilyName       string `json:"family_name"`
	BlueGreenVersion string `json:"bluegreen_version"`
	Version          string `json:"version"`
	State            string `json:"state"`
}

type DeploymentResponse struct {
	Deployments        []map[string]DeploymentRollout `json:"deployments"`
	PodStatusWebsocket string                         `json:"pod_status_websocket"`
}

type DeploymentRollout struct {
	Kind    string `json:"kind"`
	Active  string `json:"active"`
	Rolling string `json:"rolling"`
}

func ToMetadata(metadata entity.Metadata) Metadata {
	return Metadata{
		Kind:            metadata.Kind,
		Name:            metadata.Name,
		Namespace:       metadata.Namespace,
		UID:             metadata.UID,
		Generation:      metadata.Generation,
		ResourceVersion: metadata.ResourceVersion,
		Annotations:     metadata.Annotations,
		Labels:          metadata.Labels,
	}
}

func FromMetadata(metadata Metadata) entity.Metadata {
	return entity.Metadata{
		Kind:            metadata.Kind,
		Name:            metadata.Name,
		Namespace:       metadata.Namespace,
		UID:             metadata.UID,
		Generation:      metadata.Generation,
		ResourceVersion: metadata.ResourceVersion,
		Annotations:     metadata.Annotations,
		Labels:          metadata.Labels,
	}
}

func ToNamespace(resource entity.Namespace) Namespace {
	return Namespace{Metadata: ToMetadata(resource.Metadata)}
}

func ToConfigMap(resource entity.ConfigMap) ConfigMap {
	return ConfigMap{
		Metadata: ToMetadata(resource.Metadata),
		Data:     resource.Data,
	}
}

func FromConfigMap(resource ConfigMap) entity.ConfigMap {
	return entity.ConfigMap{
		Metadata: FromMetadata(resource.Metadata),
		Data:     resource.Data,
	}
}

func Same[T any](resource T) T {
	return resource
}

func ToRoute(resource entity.Route) Route {
	return Route{
		Metadata: ToMetadata(resource.Metadata),
		Spec: RouteSpec{
			Host:             resource.Spec.Host,
			PathType:         resource.Spec.PathType,
			Path:             resource.Spec.Path,
			Service:          Target{Name: resource.Spec.Service.Name},
			Port:             RoutePort{TargetPort: resource.Spec.Port.TargetPort},
			IngressClassName: resource.Spec.IngressClassName,
		},
	}
}

func FromRoute(resource Route) entity.Route {
	return entity.Route{
		Metadata: FromMetadata(resource.Metadata),
		Spec: entity.RouteSpec{
			Host:             resource.Spec.Host,
			PathType:         resource.Spec.PathType,
			Path:             resource.Spec.Path,
			Service:          entity.Target{Name: resource.Spec.Service.Name},
			Port:             entity.RoutePort{TargetPort: resource.Spec.Port.TargetPort},
			IngressClassName: resource.Spec.IngressClassName,
		},
	}
}

func ToService(resource entity.Service) Service {
	var servicePorts []ServicePort
	for _, port := range resource.Spec.Ports {
		servicePorts = append(servicePorts, ServicePort{
			Name:       port.Name,
			Protocol:   port.Protocol,
			Port:       port.Port,
			TargetPort: port.TargetPort,
			NodePort:   port.NodePort,
		})
	}
	return Service{
		Metadata: ToMetadata(resource.Metadata),
		Spec: ServiceSpec{
			Ports:     servicePorts,
			Selector:  resource.Spec.Selector,
			ClusterIP: resource.Spec.ClusterIP,
			Type:      resource.Spec.Type,
		},
	}
}

func FromService(resource Service) entity.Service {
	var servicePorts []entity.Port
	for _, port := range resource.Spec.Ports {
		servicePorts = append(servicePorts, entity.Port{
			Name:       port.Name,
			Protocol:   port.Protocol,
			Port:       port.Port,
			TargetPort: port.TargetPort,
			NodePort:   port.NodePort,
		})
	}
	return entity.Service{
		Metadata: FromMetadata(resource.Metadata),
		Spec: entity.ServiceSpec{
			Ports:     servicePorts,
			Selector:  resource.Spec.Selector,
			ClusterIP: resource.Spec.ClusterIP,
			Type:      resource.Spec.Type,
		},
	}
}

func ToDeployment(resource entity.Deployment) Deployment {
	var conditions []DeploymentCondition
	for _, condition := range resource.Status.Conditions {
		conditions = append(conditions, DeploymentCondition{
			LastTransitionTime: condition.LastTransitionTime,
			LastUpdateTime:     condition.LastUpdateTime,
			Message:            condition.Message,
			Reason:             condition.Reason,
			Status:             condition.Status,
			Type:               condition.Type,
		})
	}
	return Deployment{
		Metadata: ToMetadata(resource.Metadata),
		Spec: DeploymentSpec{
			Replicas:             resource.Spec.Replicas,
			RevisionHistoryLimit: resource.Spec.RevisionHistoryLimit,
			Template: PodTemplateSpec{
				Metadata: TemplateMetadata{Labels: resource.Spec.Template.Metadata.Labels},
				Spec:     ToPodSpec(resource.Spec.Template.Spec),
			},
			Strategy: DeploymentStrategy{Type: resource.Spec.Strategy.Type},
		},
		Status: DeploymentStatus{
			AvailableReplicas:  resource.Status.AvailableReplicas,
			Conditions:         conditions,
			ObservedGeneration: resource.Status.ObservedGeneration,
			ReadyReplicas:      resource.Status.ReadyReplicas,
			Replicas:           resource.Status.Replicas,
			UpdatedReplicas:    resource.Status.UpdatedReplicas,
		},
	}
}

func ToPod(resource entity.Pod) Pod {
	var conditions []StatusCondition
	for _, condition := range resource.Status.Conditions {
		conditions = append(conditions, StatusCondition{
			Type:               condition.Type,
			Status:             condition.Status,
			LastProbeTime:      condition.LastProbeTime,
			LastTransitionTime: condition.LastTransitionTime,
		})
	}
	var containerStatuses []ContainerStatus
	for _, containerStatus := range resource.Status.ContainerStatuses {
		containerStatuses = append(containerStatuses, ContainerStatus{
			Name:         containerStatus.Name,
			State:        ToContainerStatus(containerStatus.State),
			LastState:    ToContainerStatus(containerStatus.LastState),
			Ready:        containerStatus.Ready,
			RestartCount: containerStatus.RestartCount,
			Image:        containerStatus.Image,
			ImageID:      containerStatus.ImageID,
			ContainerID:  containerStatus.ContainerID,
		})
	}
	return Pod{
		Metadata: ToMetadata(resource.Metadata),
		Spec:     ToPodSpec(resource.Spec),
		Status: PodStatus{
			Phase:             resource.Status.Phase,
			Conditions:        conditions,
			HostIP:            resource.Status.HostIP,
			PodIP:             resource.Status.PodIP,
			StartTime:         resource.Status.StartTime,
			ContainerStatuses: containerStatuses,
		},
	}
}

func ToContainerStatus(containerState entity.ContainerState) ContainerState {
	var running *ContainerStateRunning
	if containerState.Running != nil {
		running = &ContainerStateRunning{StartedAt: containerState.Running.StartedAt}
	}
	var terminated *ContainerStateTerminated
	if containerState.Terminated != nil {
		terminated = &ContainerStateTerminated{
			ContainerID: containerState.Terminated.ContainerID,
			ExitCode:    containerState.Terminated.ExitCode,
			FinishedAt:  containerState.Terminated.FinishedAt,
			Reason:      containerState.Terminated.Reason,
			StartedAt:   containerState.Terminated.StartedAt,
		}
	}
	var waiting *ContainerStateWaiting
	if containerState.Waiting != nil {
		waiting = &ContainerStateWaiting{
			Message: containerState.Waiting.Message,
			Reason:  containerState.Waiting.Reason,
		}
	}
	return ContainerState{
		Running:    running,
		Terminated: terminated,
		Waiting:    waiting,
	}
}

func ToPodSpec(podSpec entity.PodSpec) PodSpec {
	var volumes []SpecVolume
	for _, volume := range podSpec.Volumes {
		var secret *VolumesSecret
		if volume.Secret != nil {
			secret = &VolumesSecret{
				SecretName:  volume.Secret.SecretName,
				DefaultMode: volume.Secret.DefaultMode,
			}
		}
		volumes = append(volumes, SpecVolume{
			Name:   volume.Name,
			Secret: secret,
		})
	}
	var containers []SpecContainer
	for _, container := range podSpec.Containers {
		var ports []ContainerPort
		for _, port := range container.Ports {
			ports = append(ports, ContainerPort{
				ContainerPort: port.ContainerPort,
				Protocol:      port.Protocol,
				Name:          port.Name,
			})
		}
		var envs []ContainerEnv
		for _, env := range container.Env {
			var valFrom *ValueFrom
			if env.ValueFrom != nil {
				var fieldRef *FieldRef
				if env.ValueFrom.FieldRef != nil {
					fieldRef = &FieldRef{
						APIVersion: env.ValueFrom.FieldRef.APIVersion,
						FieldPath:  env.ValueFrom.FieldRef.FieldPath,
					}
				}
				var secretKeyRef *SecretKeyRef
				if env.ValueFrom.SecretKeyRef != nil {
					secretKeyRef = &SecretKeyRef{
						Key:  env.ValueFrom.SecretKeyRef.Key,
						Name: env.ValueFrom.SecretKeyRef.Name,
					}
				}
				valFrom = &ValueFrom{
					FieldRef:     fieldRef,
					SecretKeyRef: secretKeyRef,
				}
			}
			envs = append(envs, ContainerEnv{
				Name:      env.Name,
				Value:     env.Value,
				ValueFrom: valFrom,
			})
		}
		var volumeMounts []ContainerVolumeMount
		for _, volumeMount := range container.VolumeMounts {
			volumeMounts = append(volumeMounts, ContainerVolumeMount{
				Name:      volumeMount.Name,
				MountPath: volumeMount.MountPath,
				ReadOnly:  volumeMount.ReadOnly,
			})
		}
		containers = append(containers, SpecContainer{
			Name:  container.Name,
			Image: container.Image,
			Ports: ports,
			Env:   envs,
			Resources: ContainerResources{
				Limits: CpuMemoryResource{
					Cpu:    container.Resources.Limits.Cpu,
					Memory: container.Resources.Limits.Memory,
				},
				Requests: CpuMemoryResource{
					Cpu:    container.Resources.Requests.Cpu,
					Memory: container.Resources.Requests.Memory,
				},
			},
			VolumeMounts:    volumeMounts,
			ImagePullPolicy: container.ImagePullPolicy,
			Args:            container.Args,
		})
	}
	return PodSpec{
		Volumes:                       volumes,
		Containers:                    containers,
		RestartPolicy:                 podSpec.RestartPolicy,
		TerminationGracePeriodSeconds: podSpec.TerminationGracePeriodSeconds,
		DnsPolicy:                     podSpec.DnsPolicy,
		NodeName:                      podSpec.NodeName,
	}
}

func ToDeploymentFamilyVersion(resource []entity.DeploymentFamilyVersion) []DeploymentFamilyVersion {
	result := make([]DeploymentFamilyVersion, 0)
	for _, version := range resource {
		result = append(result, DeploymentFamilyVersion{
			AppName:          version.AppName,
			AppVersion:       version.AppVersion,
			Name:             version.Name,
			FamilyName:       version.FamilyName,
			BlueGreenVersion: version.BlueGreenVersion,
			Version:          version.Version,
			State:            version.State,
		})
	}
	return result
}

// ToDeploymentResponse TODO in /v3 API get rid of []map[string]DeploymentRollout and use []DeploymentRollout (with Name field inside DeploymentRollout)
func ToDeploymentResponse(namespace string, resource *entity.DeploymentResponse) DeploymentResponse {
	deployments := make([]map[string]DeploymentRollout, 0)
	for _, deployment := range resource.Deployments {
		rolloutsMap := map[string]DeploymentRollout{
			deployment.Name: {
				Kind:    deployment.Kind,
				Active:  deployment.Active,
				Rolling: deployment.Rolling,
			},
		}
		deployments = append(deployments, rolloutsMap)
	}
	watchUrl := toWatchUrl(namespace, resource)
	return DeploymentResponse{
		Deployments:        deployments,
		PodStatusWebsocket: watchUrl,
	}
}

func toWatchUrl(namespace string, resource *entity.DeploymentResponse) string {
	var urlBuilder strings.Builder
	urlBuilder.WriteString("/watchapi/v2/paas-mediation/namespaces/" + namespace + "/rollout-status?replicas=")
	var replicationControllers []string
	var replicaSets []string
	for _, d := range resource.Deployments {
		if d.Kind == pmTypes.ReplicationController {
			replicationControllers = append(replicationControllers, d.Rolling)
		} else if d.Kind == pmTypes.ReplicaSet {
			replicaSets = append(replicaSets, d.Rolling)
		}
	}
	notEmptyCheck := func(reps []string) bool { return len(reps) > 0 }
	if notEmptyCheck(replicationControllers) {
		// todo in next major release change 'replication-controller'
		urlBuilder.WriteString("replication-controller:" + strings.Join(replicationControllers, ","))
	}
	if notEmptyCheck(replicationControllers) && notEmptyCheck(replicaSets) {
		urlBuilder.WriteString("%3B")
	}
	if notEmptyCheck(replicaSets) {
		// todo in next major release change 'replica-set'
		urlBuilder.WriteString("replica-set:" + strings.Join(replicaSets, ","))
	}
	return urlBuilder.String()
}
