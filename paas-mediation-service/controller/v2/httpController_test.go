package v2

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang/mock/gomock"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	pmTypes "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/types"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/docs"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"testing"
)

func Test_Secret(t *testing.T) {
	resourceType := types.Secret
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v1/namespaces/%s/%ss", testNamespace, resourceType), 404},
			respBody: map[string]string{"error": "Cannot GET /api/v1/namespaces/test-namespace/secrets"},
		},
		{
			rest:     r{"GET", url("/api/v1/namespaces/%s/%ss/%s", testNamespace, resourceType, "test-1"), 404},
			respBody: map[string]string{"error": "Cannot GET /api/v1/namespaces/test-namespace/secrets/test-1"},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func TestParamsToAnnotationKey(t *testing.T) {
	assertions := require.New(t)
	queryArgs := fasthttp.Args{}
	queryArgs.Parse("annotation=owner&resourceType=SERVICE")
	result, _ := getParamsToAnnotationKey(&queryArgs)
	assertions.Equal("SERVICE", result.ResourceType)
	assertions.Equal("owner", result.Annotation)
}

func TestGetAnnotationResourcesFromMetadata(t *testing.T) {
	assertions := require.New(t)
	metaData := []entity.Metadata{
		{Name: metadataName1, Annotations: map[string]string{"owner": "someOwner"}},
		{Name: metadataName2, Annotations: map[string]string{"notOwner": "nothing"}},
		{Name: metadataName3, Annotations: map[string]string{"owner": "anotherOwner"}},
	}
	resources := getAnnotationResourcesFromMetadata(metaData, "owner", testNamespace)
	annotationResource := AnnotationResource{ResourceName: metadataName1, Namespace: testNamespace, AnnotationValue: "someOwner"}
	annotationResource2 := AnnotationResource{ResourceName: metadataName3, Namespace: testNamespace, AnnotationValue: "anotherOwner"}

	assertions.Equal(annotationResource, resources[0])
	assertions.Equal(annotationResource2, resources[1])
}

func TestExtractVersionsFromConfigMap(t *testing.T) {
	assertions := require.New(t)
	versionsMap := entity.ConfigMap{
		Data: map[string]string{
			"cloud-core":                                                 "release-2023.1-7.6.0-20221222.090546-0-RELEASE",
			"cloud-core.2023-01-26-11-00-00-883":                         "release-2023.1-7.6.0-20221222.090546-1-RELEASE",
			"cloud-core.2023-01-26-11-05-00-883.Ivan_Ivanov":             "release-2023.1-7.6.0-20221222.090546-2-RELEASE",
			"cloud-core.2023-01-26-11-10-00-883.Ivan_Ivanov":             "release-2023.1-7.6.0-20221222.090546-3-RELEASE",
			"oss.common.umbrella-ui.2023-01-26-11-44-51-309.Ivan_Ivanov": "release-2022.4-20221224.080630-163-RELEASE",
		},
	}
	versions := extractVersionsFromConfigMap(context.Background(), &versionsMap)
	assertions.Equal(2, len(versions))
	cloudCoreVersion := versions["cloud-core"]
	assertions.Equal("cloud-core", cloudCoreVersion.AppName)
	assertions.Equal("release-2023.1-7.6.0-20221222.090546-3-RELEASE", cloudCoreVersion.AppVersion)
	assertions.Equal("2023-01-26 11:10:00.883 +0000 UTC", cloudCoreVersion.DeployTime.String())
	umbrellaUIVersion := versions["oss.common.umbrella-ui"]
	assertions.Equal("oss.common.umbrella-ui", umbrellaUIVersion.AppName)
	assertions.Equal("release-2022.4-20221224.080630-163-RELEASE", umbrellaUIVersion.AppVersion)
	assertions.Equal("2023-01-26 11:44:51.309 +0000 UTC", umbrellaUIVersion.DeployTime.String())
}

func Test_GetAnnotationResource_GeneralErrors(t *testing.T) {
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations", testNamespace), 400},
			respBody: map[string]string{"error": "failed to parse request parameters: 'resourceType' param must be a non empty string"},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, "invalid", "all"), 400},
			respBody: map[string]string{"error": "Unsupported resource type 'invalid'"},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_GetAnnotationResourceConfigMap(t *testing.T) {
	name := "test-1"
	annotationKey := "all"
	annotationValue := "all-value"
	testConfigMap := &entity.ConfigMap{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	testRoute := &entity.Route{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	testService := &entity.Service{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	testDeployment := &entity.Deployment{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	testPod := &entity.Pod{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	testServiceAccount := &entity.ServiceAccount{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	testNS := &entity.Namespace{Metadata: entity.Metadata{Name: name, Annotations: map[string]string{annotationKey: annotationValue}}}
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.ConfigMap, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace, filter.Meta{}).Return([]entity.ConfigMap{*testConfigMap}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.ConfigMap, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace, filter.Meta{}).Return(nil, errors.New("test error"))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Route, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace, filter.Meta{}).Return([]entity.Route{*testRoute}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Route, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace, filter.Meta{}).Return(nil, errors.New("test error"))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Service, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetServiceList(gomock.Any(), testNamespace, filter.Meta{}).Return([]entity.Service{*testService}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Service, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetServiceList(gomock.Any(), testNamespace, filter.Meta{}).Return(nil, errors.New("test error"))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Secret, "all"), 400},
			respBody: map[string]string{"error": "Unsupported resource type 'secret'"},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Deployment, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace, filter.Meta{}).Return([]entity.Deployment{*testDeployment}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Deployment, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace, filter.Meta{}).Return(nil, errors.New("test error"))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Pod, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetPodList(gomock.Any(), testNamespace, filter.Meta{}).Return([]entity.Pod{*testPod}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Pod, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetPodList(gomock.Any(), testNamespace, filter.Meta{}).Return(nil, errors.New("test error"))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.ServiceAccount, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetServiceAccountList(gomock.Any(), testNamespace, filter.Meta{}).Return([]entity.ServiceAccount{*testServiceAccount}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.ServiceAccount, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetServiceAccountList(gomock.Any(), testNamespace, filter.Meta{}).Return(nil, errors.New("test error"))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Namespace, "all"), 200},
			respBody: []AnnotationResource{{name, testNamespace, annotationValue}},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetNamespaces(gomock.Any(), gomock.Any()).Return([]entity.Namespace{*testNS}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/annotations?resourceType=%s&annotation=%s", testNamespace, types.Namespace, "all"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetNamespaces(gomock.Any(), gomock.Any()).Return(nil, errors.New("test error"))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_GetNamespaces(t *testing.T) {
	namespace := &entity.Namespace{Metadata: entity.Metadata{Name: testNamespace, Annotations: map[string]string{"annotation-1": "value-1", "all": "all"}}}
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces"), 200},
			respBody: []entity.Namespace{*namespace},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetNamespaces(gomock.Any(), gomock.Any()).Return([]entity.Namespace{*namespace}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces"), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetNamespaces(gomock.Any(), gomock.Any()).Return(nil, errors.New("test error"))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_GetDeploymentFamilyVersions(t *testing.T) {
	deploymentFamilyVersion := &entity.DeploymentFamilyVersion{
		AppName:          "test",
		AppVersion:       "v1",
		Name:             "family-name",
		FamilyName:       "family-name",
		BlueGreenVersion: "v1",
		Version:          "v1",
		State:            "active",
	}
	expectedDeploymentFamilyVersion := &DeploymentFamilyVersion{
		AppName:          "test",
		AppVersion:       "v1",
		Name:             "family-name",
		FamilyName:       "family-name",
		BlueGreenVersion: "v1",
		Version:          "v1",
		State:            "active",
	}

	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/deployment-family/%s", testNamespace, deploymentFamilyVersion.Name), 200},
			respBody: []DeploymentFamilyVersion{},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentFamilyVersions(gomock.Any(), deploymentFamilyVersion.Name, testNamespace).
					Return([]entity.DeploymentFamilyVersion{}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/deployment-family/%s", testNamespace, deploymentFamilyVersion.Name), 200},
			respBody: []DeploymentFamilyVersion{*expectedDeploymentFamilyVersion},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentFamilyVersions(gomock.Any(), deploymentFamilyVersion.Name, testNamespace).
					Return([]entity.DeploymentFamilyVersion{*deploymentFamilyVersion}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/deployment-family/%s", testNamespace, deploymentFamilyVersion.Name), 500},
			respBody: map[string]string{"error": "test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentFamilyVersions(gomock.Any(), deploymentFamilyVersion.Name, testNamespace).
					Return(nil, errors.New("test error"))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_RestartDeployment(t *testing.T) {
	name1 := "test-1"
	name2 := "test-2"
	tests := []*testCase{
		{
			rest: r{"POST", url("/api/v2/namespaces/%s/rollout/%s", testNamespace, name1), 200},
			respBody: &DeploymentResponse{
				Deployments: []map[string]DeploymentRollout{
					{name1: DeploymentRollout{Kind: pmTypes.ReplicaSet, Active: "controller-1", Rolling: "controller-2"}},
				},
				PodStatusWebsocket: fmt.Sprintf("/watchapi/v2/paas-mediation/namespaces/%s/rollout-status?replicas=replica-set:%s",
					testNamespace, "controller-2")},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().RolloutDeployments(gomock.Any(), testNamespace, []string{name1}).
					Return(&entity.DeploymentResponse{Deployments: []entity.DeploymentRollout{
						{Name: name1, Kind: pmTypes.ReplicaSet, Active: "controller-1", Rolling: "controller-2"},
					}}, nil)
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/rollout", testNamespace), 200},
			reqBody:  RolloutDeploymentBody{DeploymentNames: []string{name1, name2}, Parallel: true},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: &DeploymentResponse{
				Deployments: []map[string]DeploymentRollout{
					{name1: DeploymentRollout{Kind: pmTypes.ReplicaSet, Active: "controller-1-1", Rolling: "controller-1-2"}},
					{name2: DeploymentRollout{Kind: pmTypes.ReplicationController, Active: "controller-2-1", Rolling: "controller-2-2"}},
				},
				PodStatusWebsocket: fmt.Sprintf("/watchapi/v2/paas-mediation/namespaces/%s/rollout-status?replicas=replication-controller:%s%%3Breplica-set:%s",
					testNamespace, "controller-2-2", "controller-1-2")},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().RolloutDeploymentsInParallel(gomock.Any(), testNamespace, []string{name1, name2}).
					Return(&entity.DeploymentResponse{Deployments: []entity.DeploymentRollout{
						{Name: name1, Kind: pmTypes.ReplicaSet, Active: "controller-1-1", Rolling: "controller-1-2"},
						{Name: name2, Kind: pmTypes.ReplicationController, Active: "controller-2-1", Rolling: "controller-2-2"},
					}}, nil)
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/rollout", testNamespace), 200},
			reqBody:  RolloutDeploymentBody{DeploymentNames: []string{name1, name2}, Parallel: false},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: &DeploymentResponse{
				Deployments: []map[string]DeploymentRollout{
					{name1: DeploymentRollout{Kind: pmTypes.ReplicaSet, Active: "controller-1-1", Rolling: "controller-1-2"}},
					{name2: DeploymentRollout{Kind: pmTypes.ReplicationController, Active: "controller-2-1", Rolling: "controller-2-2"}},
				},
				PodStatusWebsocket: fmt.Sprintf("/watchapi/v2/paas-mediation/namespaces/%s/rollout-status?replicas=replication-controller:%s%%3Breplica-set:%s",
					testNamespace, "controller-2-2", "controller-1-2")},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().RolloutDeployments(gomock.Any(), testNamespace, []string{name1, name2}).
					Return(&entity.DeploymentResponse{Deployments: []entity.DeploymentRollout{
						{Name: name1, Kind: pmTypes.ReplicaSet, Active: "controller-1-1", Rolling: "controller-1-2"},
						{Name: name2, Kind: pmTypes.ReplicationController, Active: "controller-2-1", Rolling: "controller-2-2"},
					}}, nil)
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/rollout", testNamespace), 500},
			reqBody:  RolloutDeploymentBody{DeploymentNames: []string{name1, name2}, Parallel: false},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": "an error occurred while rolling out deployments: test error"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().RolloutDeployments(gomock.Any(), testNamespace, []string{name1, name2}).
					Return(nil, errors.New("test error"))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_Swagger(t *testing.T) {
	tests := []*testCase{
		{
			rest:     r{"GET", url("/swagger-ui/swagger.json"), 200},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: docs.SwaggerInfo.ReadDoc(),
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}
