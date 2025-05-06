package pmservice

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	metadataName1 = "metadataName1"
	metadataName2 = "metadataName2"
	testNamespace = "test-namespace"
)

var (
	testMetadata = entity.Metadata{
		Name:      "test-name",
		Namespace: testNamespace,
		Annotations: map[string]string{
			"test": "test",
		},
	}
	testMetadataSlice = []entity.Metadata{testMetadata}
)

func getTestMetadata() (entity.Metadata, entity.Metadata) {
	return entity.Metadata{Name: metadataName1, Annotations: map[string]string{"owner": "someOwner"}},
		entity.Metadata{Name: metadataName2, Annotations: map[string]string{"notOwner": "nothing"}}
}

func TestCastSecretToMetadata(t *testing.T) {
	testMetadata1, testMetadata2 := getTestMetadata()
	testGetMetadata(t, []entity.Secret{{Metadata: testMetadata1}, {Metadata: testMetadata2}})
}

func TestCastConfigMapToMetadata(t *testing.T) {
	testMetadata1, testMetadata2 := getTestMetadata()
	testGetMetadata(t, []entity.ConfigMap{{Metadata: testMetadata1}, {Metadata: testMetadata2}})
}

func TestCastRouteToMetadata(t *testing.T) {
	testMetadata1, testMetadata2 := getTestMetadata()
	testGetMetadata(t, []entity.Route{{Metadata: testMetadata1}, {Metadata: testMetadata2}})
}

func TestCastServiceAccountsToMetadata(t *testing.T) {
	testMetadata1, testMetadata2 := getTestMetadata()
	testGetMetadata(t, []entity.ServiceAccount{{Metadata: testMetadata1}, {Metadata: testMetadata2}})
}

func TestCastPodsToMetadata(t *testing.T) {
	testMetadata1, testMetadata2 := getTestMetadata()
	testGetMetadata(t, []entity.Pod{{Metadata: testMetadata1}, {Metadata: testMetadata2}})
}

func TestCastNamespacesToMetadata(t *testing.T) {
	testMetadata1, testMetadata2 := getTestMetadata()
	testGetMetadata(t, []entity.Namespace{{Metadata: testMetadata1}, {Metadata: testMetadata2}})
}

func testGetMetadata[T entity.HasMetadata](t *testing.T, list []T) {
	assertions := require.New(t)
	testMetadata1, testMetadata2 := getTestMetadata()
	metadataSlice, _ := getMetadata[T](context.Background(), testMetadata.Namespace,
		func(ctx context.Context, namespace string, filter filter.Meta) ([]T, error) {
			return list, nil
		})
	assertions.Equal(testMetadata1, metadataSlice[0])
	assertions.Equal(testMetadata2, metadataSlice[1])
}

func TestGetListOfUnknownResource(t *testing.T) {
	assertions := require.New(t)
	resourceType := "unknown"
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	srv := PmService{srvMock}
	meta, err := srv.GetResourceMeta(context.Background(), resourceType, testMetadata.Namespace)
	assertions.Nil(meta)
	assertions.Equal(UnsupportedResourceTypeErr{Type: resourceType}, err)
}

func TestGetListOfServiceResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	srvMock.EXPECT().GetServiceList(gomock.Any(), testMetadata.Namespace, gomock.Any()).
		Return([]entity.Service{{Metadata: testMetadata}}, nil)
	testListOfResource(t, types.Service, testNamespace, srvMock, testMetadataSlice, nil)
}

func TestGetListOfServiceAccountResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	srvMock.EXPECT().GetServiceAccountList(gomock.Any(), testMetadata.Namespace, gomock.Any()).
		Return([]entity.ServiceAccount{{Metadata: testMetadata}}, nil)
	testListOfResource(t, types.ServiceAccount, testNamespace, srvMock, testMetadataSlice, nil)
}

func TestGetListOfSecretResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	testListOfResource(t, types.Secret, testNamespace, srvMock, nil, NewUnsupportedResourceTypeErr(types.Secret))
}

func TestGetListOfConfigMapResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	srvMock.EXPECT().GetConfigMapList(gomock.Any(), testMetadata.Namespace, gomock.Any()).
		Return([]entity.ConfigMap{{Metadata: testMetadata}}, nil)
	testListOfResource(t, types.ConfigMap, testNamespace, srvMock, testMetadataSlice, nil)
}

func TestGetListOfPodResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	srvMock.EXPECT().GetPodList(gomock.Any(), testMetadata.Namespace, gomock.Any()).
		Return([]entity.Pod{{Metadata: testMetadata}}, nil)
	testListOfResource(t, types.Pod, testNamespace, srvMock, testMetadataSlice, nil)
}

func TestGetListOfRouteResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	srvMock := service.NewMockPlatformService(ctrl)
	srvMock.EXPECT().GetRouteList(gomock.Any(), testMetadata.Namespace, gomock.Any()).
		Return([]entity.Route{{Metadata: testMetadata}}, nil)
	testListOfResource(t, types.Route, testNamespace, srvMock, testMetadataSlice, nil)
}

func testListOfResource(t *testing.T, resourceName, namespace string, mock *service.MockPlatformService, expectedMeta []entity.Metadata, expectedErr error) {
	assertions := require.New(t)
	service := PmService{mock}
	meta, err := service.GetResourceMeta(context.Background(), resourceName, namespace)
	assertions.Equal(expectedErr, err)
	assertions.Equal(expectedMeta, meta)
}
