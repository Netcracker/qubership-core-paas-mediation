package v2

import (
	"errors"
	"fmt"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

// run these test to make sure all fields from entity structs are processed by To... functions

func TestToNamespace(t *testing.T) {
	testConversion(t, testNamespaceEntity, ToNamespace)
}

func TestToConfigMap(t *testing.T) {
	testConversion(t, testConfigMapEntity, ToConfigMap)
}

func TestToRoute(t *testing.T) {
	testConversion(t, testRouteEntity, ToRoute)
}

func TestToService(t *testing.T) {
	testConversion(t, testServiceEntity, ToService)
}

func TestToDeployment(t *testing.T) {
	testConversion(t, testDeploymentEntity, ToDeployment)
}

func TestToPod(t *testing.T) {
	testConversion(t, testPodEntity, ToPod)
}

func TestFromConfigMap(t *testing.T) {
	testConversion(t, testConfigMapModel, FromConfigMap)
}

func TestFromRoute(t *testing.T) {
	testConversion(t, testRouteModel, FromRoute)
}

func TestFromervice(t *testing.T) {
	testConversion(t, testServiceModel, FromService)
}

var (
	testMetadataModel = Metadata{
		Kind:            "namespace",
		Name:            "test",
		Namespace:       "test",
		UID:             "123e4567-e89b-12d3-a456-426614174000",
		Generation:      1,
		ResourceVersion: "1",
		Annotations:     map[string]string{"test": "test"},
		Labels:          map[string]string{"test": "test"},
	}
	testConfigMapModel = ConfigMap{Metadata: testMetadataModel,
		Data: map[string]string{"test": "test"}}

	testIngressClassName = "test-ingress-class-name"

	testRouteModel = Route{Metadata: testMetadataModel,
		Spec: RouteSpec{
			Host:             "test.host",
			PathType:         "type-test",
			Path:             "path-test",
			Service:          Target{Name: "test"},
			Port:             RoutePort{TargetPort: 8080},
			IngressClassName: &testIngressClassName,
		}}

	testServiceModel = Service{Metadata: testMetadataModel,
		Spec: ServiceSpec{
			Ports: []ServicePort{{
				Name:       "test-name",
				Protocol:   "test-protocol",
				Port:       8081,
				TargetPort: 8082,
				NodePort:   8083,
			}},
			Selector:  map[string]string{"test": "test"},
			ClusterIP: "127.0.0.1",
			Type:      "ClusterIP",
		},
	}

	testMetadataEntity = entity.Metadata{
		Kind:            "namespace",
		Name:            "test",
		Namespace:       "test",
		UID:             "123e4567-e89b-12d3-a456-426614174000",
		Generation:      1,
		ResourceVersion: "1",
		Annotations:     map[string]string{"test": "test"},
		Labels:          map[string]string{"test": "test"},
	}

	testNamespaceEntity = entity.Namespace{Metadata: testMetadataEntity}

	testConfigMapEntity = entity.ConfigMap{Metadata: testMetadataEntity,
		Data: map[string]string{"test": "test"}}

	testRouteEntity = entity.Route{Metadata: testMetadataEntity,
		Spec: entity.RouteSpec{
			Host:             "test.host",
			PathType:         "type-test",
			Path:             "path-test",
			Service:          entity.Target{Name: "test"},
			Port:             entity.RoutePort{TargetPort: 8080},
			IngressClassName: &testIngressClassName,
		}}

	testServiceEntity = entity.Service{Metadata: testMetadataEntity,
		Spec: entity.ServiceSpec{
			Ports: []entity.Port{{
				Name:       "test-name",
				Protocol:   "test-protocol",
				Port:       8081,
				TargetPort: 8082,
				NodePort:   8083,
			}},
			Selector:  map[string]string{"test": "test"},
			ClusterIP: "127.0.0.1",
			Type:      "ClusterIP",
		},
	}

	int32Val = int32(1)
	timeStr  = "2022-12-29T18:21:18Z"

	testPodSpecEntity = entity.PodSpec{
		Volumes: []entity.SpecVolume{{
			Name: "test-volume",
			Secret: &entity.VolumesSecret{
				SecretName:  "test-secret",
				DefaultMode: 600,
			},
		}},
		Containers: []entity.SpecContainer{{
			Name:  "test-container",
			Image: "test-image",
			Ports: []entity.ContainerPort{{
				ContainerPort: 8080,
				Protocol:      "http",
				Name:          "test-port",
			}},
			Env: []entity.ContainerEnv{{
				Name:  "test-env",
				Value: "env",
				ValueFrom: &entity.ValueFrom{
					FieldRef: &entity.FieldRef{
						APIVersion: "test-version",
						FieldPath:  "test-path",
					},
					SecretKeyRef: &entity.SecretKeyRef{
						Key:  "test-secret-key",
						Name: "test-secret",
					},
				},
			}},
			Resources: entity.ContainerResources{
				Limits: entity.CpuMemoryResource{
					Cpu:    "2",
					Memory: "2Gi",
				},
				Requests: entity.CpuMemoryResource{
					Cpu:    "1",
					Memory: "1Gi",
				},
			},
			VolumeMounts: []entity.ContainerVolumeMount{{
				Name:      "test-volume-mount",
				MountPath: "test-mount-path",
				ReadOnly:  true,
			}},
			ImagePullPolicy: "Always",
			Args:            []string{"arg"},
		},
		},
		RestartPolicy:                 "Always",
		TerminationGracePeriodSeconds: 30,
		DnsPolicy:                     "ClusterFirst",
		NodeName:                      "test-node",
	}
	testDeploymentEntity = entity.Deployment{
		Metadata: testMetadataEntity,
		Spec: entity.DeploymentSpec{
			Replicas:             &int32Val,
			RevisionHistoryLimit: &int32Val,
			Template: entity.PodTemplateSpec{
				Metadata: entity.TemplateMetadata{Labels: map[string]string{"test": "test"}},
				Spec:     testPodSpecEntity,
			},
			Strategy: entity.DeploymentStrategy{Type: "test-type"},
		},
		Status: entity.DeploymentStatus{
			AvailableReplicas: 1,
			Conditions: []entity.DeploymentCondition{{
				LastTransitionTime: &timeStr,
				LastUpdateTime:     &timeStr,
				Message:            "test-message",
				Reason:             "test-reason",
				Status:             "test-status",
				Type:               "test-type",
			}},
			ObservedGeneration: 1,
			ReadyReplicas:      2,
			Replicas:           3,
			UpdatedReplicas:    4,
		},
	}

	testPodEntity = entity.Pod{
		Metadata: testMetadataEntity,
		Spec:     testPodSpecEntity,
		Status: entity.PodStatus{
			Phase: "test-phase",
			Conditions: []entity.StatusCondition{{
				Type:               "test-type",
				Status:             "test-status",
				LastProbeTime:      &timeStr,
				LastTransitionTime: &timeStr,
			}},
			HostIP:    "127.0.0.1",
			PodIP:     "127.0.0.2",
			StartTime: &timeStr,
			ContainerStatuses: []entity.ContainerStatus{{
				Name: "test-status",
				State: entity.ContainerState{
					Running: &entity.ContainerStateRunning{StartedAt: &timeStr},
					Terminated: &entity.ContainerStateTerminated{
						ContainerID: "1",
						ExitCode:    143,
						FinishedAt:  &timeStr,
						Reason:      "test-terminating-reason",
						StartedAt:   &timeStr,
					},
					Waiting: &entity.ContainerStateWaiting{
						Message: "test-message",
						Reason:  "test-waiting-reason",
					},
				},
				LastState: entity.ContainerState{
					Running: &entity.ContainerStateRunning{StartedAt: &timeStr},
					Terminated: &entity.ContainerStateTerminated{
						ContainerID: "1",
						ExitCode:    143,
						FinishedAt:  &timeStr,
						Reason:      "test-last-terminating-reason",
						StartedAt:   &timeStr,
					},
					Waiting: &entity.ContainerStateWaiting{
						Message: "test-last-message",
						Reason:  "test-last-waiting-reason",
					},
				},
				Ready:        true,
				RestartCount: 1,
				Image:        "test-image",
				ImageID:      "1234",
				ContainerID:  "5678",
			},
			},
		},
	}
)

func testConversion[T any, R any](t *testing.T, from T, convert func(T) R) {
	assertions := require.New(t)
	assertions.Nil(validateStruct(from))
	assertions.Nil(validateStruct(convert(from)))
}

func validateStruct(s any, path ...string) (err error) {
	structType := reflect.TypeOf(s)
	if structType.Kind() != reflect.Struct {
		return errors.New("input param should be a struct")
	}
	structVal := reflect.ValueOf(s)
	fieldNum := structVal.NumField()

	var prefix string
	if len(path) > 0 {
		prefix = path[0]
	} else {
		prefix = structType.Name()
	}

	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		if !isSet {
			if err != nil {
				err = fmt.Errorf("%w %s in not set; ", err, prefix+"."+fieldName)
			} else {
				err = fmt.Errorf("%s in not set; ", prefix+"."+fieldName)
			}
		} else {
			if field.Type().Kind() == reflect.Struct {
				if inErr := validateStruct(field.Interface(), prefix+"."+fieldName); inErr != nil {
					err = inErr
				}
			}
		}
	}
	return err
}
