package pmservice

import (
	"context"
	"fmt"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-paas-mediation/types"
	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	"strings"
)

type PmService struct {
	Platform service.PlatformService
}

type listFunc[T entity.HasMetadata] func(ctx context.Context, namespace string, filter filter.Meta) ([]T, error)

func getMetadata[T entity.HasMetadata](ctx context.Context, namespace string, GetResourceList listFunc[T]) ([]entity.Metadata, error) {
	list, err := GetResourceList(ctx, namespace, filter.Meta{})
	if err != nil {
		return nil, err
	}
	withMetadata := make([]entity.Metadata, len(list))
	for i, v := range list {
		withMetadata[i] = v.GetMetadata()
	}
	return withMetadata, nil
}

func (pm *PmService) GetResourceMeta(ctx context.Context, resourceType string, namespace string) ([]entity.Metadata, error) {
	resourceType = strings.ToLower(resourceType)
	switch resourceType {
	case types.ConfigMap:
		return getMetadata[entity.ConfigMap](ctx, namespace, pm.Platform.GetConfigMapList)
	case types.Deployment:
		return getMetadata[entity.Deployment](ctx, namespace, pm.Platform.GetDeploymentList)
	case types.Pod:
		return getMetadata[entity.Pod](ctx, namespace, pm.Platform.GetPodList)
	case types.Route:
		return getMetadata[entity.Route](ctx, namespace, pm.Platform.GetRouteList)
	case types.Service:
		return getMetadata[entity.Service](ctx, namespace, pm.Platform.GetServiceList)
	case types.ServiceAccount:
		return getMetadata[entity.ServiceAccount](ctx, namespace, pm.Platform.GetServiceAccountList)
	case types.Namespace:
		return getMetadata[entity.Namespace](ctx, namespace, func(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Namespace, error) {
			return pm.Platform.GetNamespaces(ctx, filter)
		})
	default:
		return nil, NewUnsupportedResourceTypeErr(resourceType)
	}
}

type UnsupportedResourceTypeErr struct {
	Type string
}

func NewUnsupportedResourceTypeErr(resourceType string) UnsupportedResourceTypeErr {
	return UnsupportedResourceTypeErr{Type: resourceType}
}

func (e UnsupportedResourceTypeErr) Error() string {
	return fmt.Sprintf("Unsupported resource type '%s'", e.Type)
}

func ErrorConverterToStatusCode(err error) int {
	switch {
	case errors.IsAlreadyExists(err):
		return http.StatusConflict
	case errors.IsNotFound(err):
		return http.StatusNotFound
	case errors.IsConflict(err):
		return http.StatusConflict
	case errors.IsForbidden(err):
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
