package v2

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"testing"
)

func Test_Route(t *testing.T) {
	initTestConfig()
	resourceName1 := "test-1"
	resourceName2 := "test-2"
	resourceName3 := "test-3"
	resourceName4 := "test-4"
	resourceName5 := "test-5"

	resource1 := &entity.Route{Metadata: entity.Metadata{Name: resourceName1,
		Annotations: map[string]string{"annotation-1": "value-1", "all": "all"},
		Labels:      map[string]string{"label-1": "value-1", "all": "all"},
	}}
	resource2 := &entity.Route{Metadata: entity.Metadata{Name: resourceName2,
		Annotations: map[string]string{"annotation-2": "value-2", "all": "all"},
		Labels:      map[string]string{"label-2": "value-2", "all": "all"},
	}}
	testIngressClassName := "test-ingress-class-name"
	resourceWithIngressClassName := &entity.Route{Metadata: entity.Metadata{Name: resourceName5,
		Annotations: map[string]string{"annotation-1": "value-1", "all": "all"},
		Labels:      map[string]string{"label-1": "value-1", "all": "all"},
	},
		Spec: entity.RouteSpec{IngressClassName: &testIngressClassName}}

	modelResource1 := ToRoute(*resource1)
	modelResource2 := ToRoute(*resource2)
	modelResourceWithIngressClassName := ToRoute(*resourceWithIngressClassName)

	resourceType := types.Route
	tests := []*testCase{
		{
			rest:    r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 201},
			reqBody: resource1, respBody: modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateRoute(gomock.Any(), resource1, testNamespace).Return(resource1, nil)
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 400},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  []Route{},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewAlreadyExists(corev1.Resource(resourceType), resource1.Name))
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewConflict(corev1.Resource("configmap"), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			reqBody:  modelResource1,
			respBody: map[string]string{"error": fmt.Sprintf("failed to create %s: %s \"test-1\" is forbidden: test reason", resourceType, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"POST", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			reqBody:  modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to create %s: Internal error occurred: test reason", resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().CreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewInternalError(errors.New("test reason")))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},
			reqBody:  modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateRoute(gomock.Any(), resource1, testNamespace).Return(resource1, nil)
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},
			reqBody:  modelResourceWithIngressClassName,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: modelResourceWithIngressClassName,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateRoute(gomock.Any(), resourceWithIngressClassName, testNamespace).
					DoAndReturn(func(ctx context.Context, route *entity.Route, namespace string) (*entity.Route, error) {
						if route.Spec.IngressClassName == nil || testIngressClassName != *route.Spec.IngressClassName {
							panic("invalid IngressClassName")
						}
						return route, nil
					})
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 400},
			reqBody:  []Route{},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			reqBody:  modelResource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewAlreadyExists(corev1.Resource(resourceType), resource1.Name))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 409},
			reqBody:  resource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewConflict(corev1.Resource(resourceType), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			reqBody:  resource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to update or create %s: %s \"test-1\" is forbidden: test reason", resourceType, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resource1.Name, errors.New("test reason")))
			},
		},
		{
			rest:     r{"PUT", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			reqBody:  resource1,
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to update or create %s: Internal error occurred: test reason", resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().UpdateOrCreateRoute(gomock.Any(), resource1, testNamespace).
					Return(nil, k8sErrors.NewInternalError(errors.New("test reason")))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},
			respBody: []Route{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{}}).
					Return([]entity.Route{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all", testNamespace, resourceType), 200},
			respBody: []Route{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all"}, Labels: map[string]string{}}).
					Return([]entity.Route{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all", testNamespace, resourceType), 200},
			respBody: []Route{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all"}}).
					Return([]entity.Route{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all&labels=all:all", testNamespace, resourceType), 200},
			respBody: []Route{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all"}, Labels: map[string]string{"all": "all"}}).
					Return([]entity.Route{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all;annotation-2:value-2", testNamespace, resourceType), 200},
			respBody: []Route{modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all", "annotation-2": "value-2"}, Labels: map[string]string{}}).
					Return([]entity.Route{*resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all;label-1:value-1", testNamespace, resourceType), 200},
			respBody: []Route{modelResource1},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all", "label-1": "value-1"}}).
					Return([]entity.Route{*resource1}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all;label-unknown:value-unknown", testNamespace, resourceType), 200},
			respBody: []Route{},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all", "label-unknown": "value-unknown"}}).
					Return([]entity.Route{}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName1), 200},
			respBody: modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRoute(gomock.Any(), resourceName1, testNamespace).Return(resource1, nil)
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
				srv.EXPECT().GetRoute(gomock.Any(), resourceName3, testNamespace).Return(nil, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName5), 200},
			respBody: modelResourceWithIngressClassName,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRoute(gomock.Any(), resourceName5, testNamespace).Return(resourceWithIngressClassName, nil)
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName4), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s with name=%s in namespace=%s: "+
				"test error", resourceType, resourceName4, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRoute(gomock.Any(), resourceName4, testNamespace).Return(nil, errors.New("test error"))
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName4), 403},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s with name=%s in namespace=%s: %s \"%s\" is forbidden: test reason",
				resourceType, resourceName4, testNamespace, resourceType, resourceName4)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRoute(gomock.Any(), resourceName4, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName4, errors.New("test reason")))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: test error", resourceType, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace, gomock.Any()).Return(nil, errors.New("test error"))
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: %s is forbidden: test reason",
				resourceType, testNamespace, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetRouteList(gomock.Any(), testNamespace, gomock.Any()).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), "", errors.New("test reason")))
			},
		},
		{
			rest: r{"DELETE", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 200},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().DeleteRoute(gomock.Any(), resourceName3, testNamespace).Return(nil)
			},
		},
		{
			rest:     r{"DELETE", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 403},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to delete %s '%s': %s", resourceType, resourceName3,
				k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName3, errors.New("test reason")).Error())},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().DeleteRoute(gomock.Any(), resourceName3, testNamespace).
					Return(k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName3, errors.New("test reason")))
			},
		},
		{
			rest:     r{"DELETE", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName3), 500},
			alterReq: AddHeader(fiber.HeaderContentType, fiber.MIMEApplicationJSON),
			respBody: map[string]string{"error": fmt.Sprintf("failed to delete %s '%s': %s", resourceType, resourceName3,
				k8sErrors.NewInternalError(errors.New("test reason")).Error())},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().DeleteRoute(gomock.Any(), resourceName3, testNamespace).
					Return(k8sErrors.NewInternalError(errors.New("test reason")))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_RouteGetConcurrency(t *testing.T) {
	testGetConcurrency[entity.Route](t, types.Route, func(srv *service.MockPlatformService) *gomock.Call {
		return srv.EXPECT().GetRoute(gomock.Any(), gomock.Any(), testNamespace)
	})
}

func Test_RouteListConcurrency(t *testing.T) {
	testListConcurrency[entity.Pod](t, types.Route, func(srv *service.MockPlatformService) *gomock.Call {
		return srv.EXPECT().GetRouteList(gomock.Any(), testNamespace, gomock.Any())
	})
}
