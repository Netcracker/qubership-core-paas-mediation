package v2

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"testing"
)

func Test_Deployment(t *testing.T) {
	initTestConfigLoader()
	resourceName1 := "test-1"
	resourceName2 := "test-2"
	resourceName3 := "test-3"
	resourceName4 := "test-4"

	resource1 := &entity.Deployment{Metadata: entity.Metadata{Name: resourceName1,
		Annotations: map[string]string{"annotation-1": "value-1", "all": "all"},
		Labels:      map[string]string{"label-1": "value-1", "all": "all"},
	}}
	resource2 := &entity.Deployment{Metadata: entity.Metadata{Name: resourceName2,
		Annotations: map[string]string{"annotation-2": "value-2", "all": "all"},
		Labels:      map[string]string{"label-2": "value-2", "all": "all"},
	}}
	modelResource1 := ToDeployment(*resource1)
	modelResource2 := ToDeployment(*resource2)

	resourceType := types.Deployment
	tests := []*testCase{
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},
			respBody: []Deployment{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{}}).
					Return([]entity.Deployment{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all", testNamespace, resourceType), 200},
			respBody: []Deployment{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all"}, Labels: map[string]string{}}).
					Return([]entity.Deployment{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all", testNamespace, resourceType), 200},
			respBody: []Deployment{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all"}}).
					Return([]entity.Deployment{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all&labels=all:all", testNamespace, resourceType), 200},
			respBody: []Deployment{modelResource1, modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all"}, Labels: map[string]string{"all": "all"}}).
					Return([]entity.Deployment{*resource1, *resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?annotations=all:all;annotation-2:value-2", testNamespace, resourceType), 200},
			respBody: []Deployment{modelResource2},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{"all": "all", "annotation-2": "value-2"}, Labels: map[string]string{}}).
					Return([]entity.Deployment{*resource2}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all;label-1:value-1", testNamespace, resourceType), 200},
			respBody: []Deployment{modelResource1},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all", "label-1": "value-1"}}).
					Return([]entity.Deployment{*resource1}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss?labels=all:all;label-unknown:value-unknown", testNamespace, resourceType), 200},
			respBody: []Deployment{},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace,
					filter.Meta{Annotations: map[string]string{}, Labels: map[string]string{"all": "all", "label-unknown": "value-unknown"}}).
					Return([]entity.Deployment{}, nil)
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName1), 200},
			respBody: modelResource1,
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeployment(gomock.Any(), resourceName1, testNamespace).Return(resource1, nil)
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
				srv.EXPECT().GetDeployment(gomock.Any(), resourceName3, testNamespace).Return(nil, nil)
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName4), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s with name=%s in namespace=%s: "+
				"test error", resourceType, resourceName4, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeployment(gomock.Any(), resourceName4, testNamespace).Return(nil, errors.New("test error"))
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, resourceName4), 403},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s with name=%s in namespace=%s: %s \"%s\" is forbidden: test reason",
				resourceType, resourceName4, testNamespace, resourceType, resourceName4)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeployment(gomock.Any(), resourceName4, testNamespace).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), resourceName4, errors.New("test reason")))
			},
		},
		{
			rest:     r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 500},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: test error", resourceType, testNamespace)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace, gomock.Any()).Return(nil, errors.New("test error"))
			},
		},
		{
			rest: r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 403},
			respBody: map[string]string{"error": fmt.Sprintf("error occurred while getting %s list in namespace=%s: %s is forbidden: test reason",
				resourceType, testNamespace, resourceType)},
			mock: func(srv *service.MockPlatformService) {
				srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace, gomock.Any()).
					Return(nil, k8sErrors.NewForbidden(corev1.Resource(resourceType), "", errors.New("test reason")))
			},
		},
	}
	for _, tc := range tests {
		runTestCase(t, tc)
	}
}

func Test_DeploymentGetConcurrency(t *testing.T) {
	testGetConcurrency[entity.Deployment](t, types.Deployment, func(srv *service.MockPlatformService) *gomock.Call {
		return srv.EXPECT().GetDeployment(gomock.Any(), gomock.Any(), testNamespace)
	})
}

func Test_DeploymentListConcurrency(t *testing.T) {
	testListConcurrency[entity.Deployment](t, types.Deployment, func(srv *service.MockPlatformService) *gomock.Call {
		return srv.EXPECT().GetDeploymentList(gomock.Any(), testNamespace, gomock.Any())
	})
}
