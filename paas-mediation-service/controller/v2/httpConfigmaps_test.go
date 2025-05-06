package v2

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/constants"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"testing"
	"time"
)

func Test_ConfigMap(t *testing.T) {
	initTestConfig()
	resourceName1 := "test-1"
	resourceName2 := "test-2"
	resourceName3 := "test-3"
	resourceName4 := "test-4"

	resource1 := &entity.ConfigMap{Metadata: entity.Metadata{Kind: types.ConfigMap, Name: resourceName1,
		Annotations: map[string]string{"annotation-1": "value-1", "all": "all"},
		Labels:      map[string]string{"label-1": "value-1", "all": "all"},
	}, Data: map[string]string{}}
	resource2 := &entity.ConfigMap{Metadata: entity.Metadata{Kind: types.ConfigMap, Name: resourceName2,
		Annotations: map[string]string{"annotation-2": "value-2", "all": "all"},
		Labels:      map[string]string{"label-2": "value-2", "all": "all"},
	}, Data: map[string]string{}}
	versionsMap := &entity.ConfigMap{Metadata: entity.Metadata{Kind: types.ConfigMap, Name: constants.VersionConfigMap},
		Data: map[string]string{
			"cloud-core.2023-01-26-11-05-00-883.Ivan_Ivanov": "release-2023.1-7.6.0-20221222.090546-2-RELEASE",
		}}
	blueGreenVersionsMap := &entity.ConfigMap{Metadata: entity.Metadata{Kind: types.ConfigMap, Name: constants.BlueGreenVersionConfigMap},
		Data: map[string]string{
			"version_3": "{" +
				"\"state\": \"active\"," +
				"\"apps\": {" +
				"\"cloud-core\": \"master-20230205.235456-6242\"," +
				"\"trace-services\": \"master-20221223.120908-158-RELEASE\"" +
				"}" +
				"}",
		}}

	modelResource1 := ToConfigMap(*resource1)
	modelResource2 := ToConfigMap(*resource2)

	resourceType := types.ConfigMap
	tests := []*testCase{
		{
			rest:    r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 201},
			reqBody: resource1, respBody: modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateConfigMap(gomock.Any(), resource1, testNamespace).Return(resource1, nil)
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 400},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  []ConfigMap{},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewAlreadyExists(corev1.Resource(resourceType), resource1.Name))
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewConflict(corev1.Resource("configmap"), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  modelResource1,
			respBody: map[string]string{"error": fmt.Sprintf("failed to create %s: %s \"test-1\" is forbidden: test reason", resourceType, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			reqBody:  modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to create %s: Internal error occurred: test reason", resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewInternalError(errors.New("test reason")))
			},
		},
		{
			rest: r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},

			reqBody:  modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateConfigMap(gomock.Any(), resource1, testNamespace).Return(resource1, nil)
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 400},
			reqBody:  []ConfigMap{},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			reqBody:  modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewAlreadyExists(corev1.Resource(resourceType), resource1.Name))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			reqBody:  resource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewConflict(corev1.Resource(resourceType), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			reqBody:  resource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to update or create %s: %s \"test-1\" is forbidden: test reason", resourceType, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			reqBody:  resource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to update or create %s: Internal error occurred: test reason", resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateConfigMap(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewInternalError(errors.New("test reason")))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},
			respBody: []ConfigMap{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{}}).
					Return([]entity.ConfigMap{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all", testNamespace, resourceType), 200},
			respBody: []ConfigMap{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all"}, Labels: map[string]string{}}).
					Return([]entity.ConfigMap{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all", testNamespace, resourceType), 200},
			respBody: []ConfigMap{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all"}}).
					Return([]entity.ConfigMap{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all&labels=all:all", testNamespace, resourceType), 200},
			respBody: []ConfigMap{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all"}, Labels: map[string]string{"all": "all"}}).
					Return([]entity.ConfigMap{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all;annotation-2:value-2", testNamespace, resourceType), 200},
			respBody: []ConfigMap{modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all", "annotation-2": "value-2"}, Labels: map[string]string{}}).
					Return([]entity.ConfigMap{*resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all;label-1:value-1", testNamespace, resourceType), 200},
			respBody: []ConfigMap{modelResource1},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all", "label-1": "value-1"}}).
					Return([]entity.ConfigMap{*resource1}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all;label-unknown:value-unknown", testNamespace, resourceType), 200},
			respBody: []ConfigMap{},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all", "label-unknown": "value-unknown"}}).
					Return([]entity.ConfigMap{}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName1), 200},
			respBody: resource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMap(gomock.Any(), resourceName1, testNamespace).Return(resource1, nil)
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, "versions"), 200},
			respBody: []*AppVersionData{
				{AppName: "cloud-core",
					AppVersion: "release-2023.1-7.6.0-20221222.090546-2-RELEASE",
					DeployTime: time.Date(2023, time.January, 26, 11, 05, 00, 883000000, time.UTC),
				},
			},
			mock: func(platformService *service.MockPlatformService) {
				platformService.EXPECT().GetConfigMap(gomock.Any(), constants.VersionConfigMap, testNamespace).Return(versionsMap, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, "versions"), 404},
			respBody: map[string]string{"error": "configmap 'version' not found"},
			mock: func(platformService *service.MockPlatformService) {
				platformService.EXPECT().GetConfigMap(gomock.Any(), constants.VersionConfigMap, testNamespace).Return(nil, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, constants.BlueGreenVersionConfigMap), 200},
			respBody: blueGreenVersionsMap,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMap(gomock.Any(), constants.BlueGreenVersionConfigMap, testNamespace).Return(blueGreenVersionsMap, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, constants.BlueGreenVersionConfigMap), 404},
			respBody: map[string]string{"error": "configmap 'bg-version' not found"},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMap(gomock.Any(), constants.BlueGreenVersionConfigMap, testNamespace).Return(nil, nil)
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all=all", testNamespace, resourceType), 400},
			respBody: map[string]string{"error": "Invalid parameter 'labels' provided. Failed to parse entry: 'all=all'. " +
				"Valid param structure is param_name=key_1:key_1_value;key_2:key_2_value"},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all=all", testNamespace, resourceType), 400},
			respBody: map[string]string{"error": "Invalid parameter 'annotations' provided. Failed to parse entry: 'all=all'. " +
				"Valid param structure is param_name=key_1:key_1_value;key_2:key_2_value"},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 404},
			respBody: map[string]string{"error": fmt.Sprintf("%s '%s' not found", resourceType, resourceName3)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMap(gomock.Any(), resourceName3, testNamespace).Return(nil, nil)
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName4), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s with name=%s in namespace=%s: "+
				"test error", resourceType, resourceName4, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMap(gomock.Any(), resourceName4, testNamespace).Return(nil, errors.New("test error"))
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName4), 403},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s with name=%s in namespace=%s: %s \"%s\" is forbidden: test reason",
				resourceType, resourceName4, testNamespace, resourceType, resourceName4)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMap(gomock.Any(), resourceName4, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName4, errors.New("test reason")))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: test error", resourceType, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace, gomock.Any()).Return(nil, errors.New("test error"))
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: %s is forbidden: test reason",
				resourceType, testNamespace, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace, gomock.Any()).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), "", errors.New("test reason")))
			},
		},
		{
			rest: r{"DELETE", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 200},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().DeleteConfigMap(gomock.Any(), resourceName3, testNamespace).Return(nil)
			},
		},
		{
			rest:     r{"DELETE", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 403},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to delete %s '%s': %s", resourceType, resourceName3,
				k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName3, errors.New("test reason")).Error())},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().DeleteConfigMap(gomock.Any(), resourceName3, testNamespace).
					Return(k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName3, errors.New("test reason")))
			},
		},
		{
			rest:     r{"DELETE", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 500},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to delete %s '%s': %s", resourceType, resourceName3,
				k8sErrors.NewInternalError(errors.New("test reason")).Error())},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().DeleteConfigMap(gomock.Any(), resourceName3, testNamespace).
					Return(k8sErrors.NewInternalError(errors.New("test reason")))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_ConfigMapGetConcurrency(t *testing.T) {
	testGetConcurrency[entity.ConfigMap](t, types.ConfigMap, func(srv *service.MockPlatformService) *gomock.Call {
		return srv.EXPECT().GetConfigMap(gomock.Any(), gomock.Any(), testNamespace)
	})
}

func Test_ConfigMapListConcurrency(t *testing.T) {
	testListConcurrency[entity.ConfigMap](t, types.ConfigMap, func(srv *service.MockPlatformService) *gomock.Call {
		return srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace, gomock.Any())
	})
}
