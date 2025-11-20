package v2

import (
	"errors"
	"fmt"
	"testing"

	"github.com/knadh/koanf/providers/confmap"
	fibersec "github.com/netcracker/qubership-core-lib-go-fiber-server-utils/v2/security"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"github.com/netcracker/qubership-core-lib-go/v3/serviceloader"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	"go.uber.org/mock/gomock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func initTestConfigWithFeatureFlag(enabled bool) {
	configloader.Init(&configloader.PropertySource{Provider: configloader.AsPropertyProvider(confmap.Provider(
		map[string]any{
			"microservice.namespace":             "test-namespace",
			"core.paas.mediation.gw.api.enabled": enabled,
		}, "."))})

	serviceloader.Register(1, &fibersec.DummyFiberServerSecurityMiddleware{})
}

func Test_GetHttpRouteList_FeatureDisabled(t *testing.T) {
	initTestConfigWithFeatureFlag(false)
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/gateway/httproutes", testNamespace), 404},
			respBody: map[string]string{"error": "gateway routes feature is disabled"},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_GetHttpRouteList_FeatureEnabled(t *testing.T) {
	initTestConfigWithFeatureFlag(true)
	expectedRoutes := []entity.HttpRoute{
		{HTTPRoute: &gatewayv1.HTTPRoute{
			ObjectMeta: metav1.ObjectMeta{Name: "route-1", Namespace: testNamespace},
		}},
		{HTTPRoute: &gatewayv1.HTTPRoute{
			ObjectMeta: metav1.ObjectMeta{Name: "route-2", Namespace: testNamespace},
		}},
	}
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/gateway/httproutes", testNamespace), 200},
			respBody: expectedRoutes,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetHttpRouteList(gomock.Any(), testNamespace, gomock.Any()).
					Return(expectedRoutes, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/gateway/httproutes", testNamespace), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: test error", types.HttpRoute, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetHttpRouteList(gomock.Any(), testNamespace, gomock.Any()).
					Return(nil, errors.New("test error"))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_GetGrpcRouteList_FeatureDisabled(t *testing.T) {
	initTestConfigWithFeatureFlag(false)
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/gateway/grpcroutes", testNamespace), 404},
			respBody: map[string]string{"error": "gateway routes feature is disabled"},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_GetGrpcRouteList_FeatureEnabled(t *testing.T) {
	initTestConfigWithFeatureFlag(true)
	expectedRoutes := []entity.GrpcRoute{
		{GRPCRoute: &gatewayv1.GRPCRoute{
			ObjectMeta: metav1.ObjectMeta{Name: "grpc-route-1", Namespace: testNamespace},
		}},
		{GRPCRoute: &gatewayv1.GRPCRoute{
			ObjectMeta: metav1.ObjectMeta{Name: "grpc-route-2", Namespace: testNamespace},
		}},
	}
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/gateway/grpcroutes", testNamespace), 200},
			respBody: expectedRoutes,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetGrpcRouteList(gomock.Any(), testNamespace, gomock.Any()).
					Return(expectedRoutes, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/gateway/grpcroutes", testNamespace), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: test error", types.GrpcRoute, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetGrpcRouteList(gomock.Any(), testNamespace, gomock.Any()).
					Return(nil, errors.New("test error"))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}
