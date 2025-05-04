package v2

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	pmErrors "github.com/netcracker/qubership-core-paas-mediation/errors"
	"github.com/netcracker/qubership-core-paas-mediation/pmservice"
	"github.com/netcracker/qubership-core-paas-mediation/types"
)

type GetResource[T any] func(ctx context.Context, resourceName string, namespace string) (T, error)
type ListResource[T any] func(ctx context.Context, namespace string, filter filter.Meta) ([]T, error)
type DeleteResource func(ctx context.Context, name string, namespace string) error
type CreateResource[T any] func(ctx context.Context, resource *T, namespace string) (*T, error)

const (
	concurrencyDefault   = 4
	concurrencyList      = 1
	concurrencyWebSocket = 4
)

var (
	limitTypes         = []string{types.ConfigMap, types.Deployment, types.Pod, types.Route, types.Service}
	limitChanGet       = createLimitChanMap(concurrencyDefault, limitTypes...)
	limitChanList      = createLimitChanMap(concurrencyList, limitTypes...)
	limitChanCreate    = createLimitChanMap(concurrencyDefault, []string{types.ConfigMap, types.Route, types.Service}...)
	limitChanDelete    = createLimitChanMap(concurrencyDefault, []string{types.ConfigMap, types.Route, types.Service}...)
	watchLimitTypes    = []string{types.ConfigMap, types.Namespace, "rollout", types.Route, types.Service}
	limitChanWebSocket = createLimitChanMap(concurrencyWebSocket, watchLimitTypes...)
)

func createLimitChanMap(limit int, types ...string) map[string]chan struct{} {
	result := make(map[string]chan struct{}, len(types))
	for _, t := range types {
		result[t] = make(chan struct{}, limit)
	}
	return result
}

// limit parallel processing of get requests to max 4 concurrent requests per resource type
func getAdapter[T any, R any](resourceType string, get GetResource[*T], convert func(resource T) R, c *fiber.Ctx) error {
	ctx := c.UserContext()
	resourceName, namespace := getURLParams(ctx, c)
	logger.InfoC(ctx, "Received a request to get %s with name=%s in namespace=%s",
		resourceType, resourceName, namespace)
	limitChanGet[resourceType] <- struct{}{}
	defer func() {
		<-limitChanGet[resourceType]
	}()
	result, err := get(ctx, resourceName, namespace)
	if err != nil {
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		dErr := fmt.Errorf("error occurred while getting %s with name=%s in namespace=%s: %w",
			resourceType, resourceName, namespace, err)
		return pmErrors.New(statusCode, dErr)
	} else if result == nil {
		return pmErrors.New(fiber.StatusNotFound, fmt.Errorf("%s '%s' not found", resourceType, resourceName))
	}
	return respondWithJson(ctx, c, fiber.StatusOK, convert(*result))
}

// for now list requests must not be executed in parallel because they may consume a lot of memory for heavy resources which can lead to OutOfMemory fails
// so limit list request processing to max 1 concurrent processing per resource type
func listAdapter[T any, R any](resourceType string, list ListResource[T], convert func(resource T) R, c *fiber.Ctx) error {
	ctx := c.UserContext()
	_, namespace := getURLParams(ctx, c)
	logger.InfoC(ctx, "Received a request to get %s list in namespace=%s", resourceType, namespace)
	limitChanList[resourceType] <- struct{}{}
	defer func() {
		<-limitChanList[resourceType]
	}()
	metaFilter, fErr := buildFilterFromParams(c)
	if fErr != nil {
		return pmErrors.New(fiber.StatusBadRequest, fErr)
	}
	result, err := list(ctx, namespace, metaFilter)
	if err != nil {
		dErr := fmt.Errorf("error occurred while getting %s list in namespace=%s: %w", resourceType, namespace, err)
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, dErr)
	}
	modelResources := make([]R, 0)
	for _, resource := range result {
		modelResources = append(modelResources, convert(resource))
	}
	return respondWithJson(ctx, c, fiber.StatusOK, modelResources)
}

func createAdapter[T any, R any](resourceType string, createFunc CreateResource[T], convertTo func(T) R, convertFrom func(R) T, c *fiber.Ctx) error {
	return updateOrCreateAdapter("create", resourceType, fiber.StatusCreated, createFunc, convertTo, convertFrom, c)
}

func updateAdapter[T any, R any](resourceType string, createFunc CreateResource[T], convertTo func(T) R, convertFrom func(R) T, c *fiber.Ctx) error {
	return updateOrCreateAdapter("update or create", resourceType, fiber.StatusOK, createFunc, convertTo, convertFrom, c)
}

func updateOrCreateAdapter[T any, R any](action, resourceType string, successCode int,
	createFunc CreateResource[T], convertTo func(resource T) R, convertFrom func(resource R) T, c *fiber.Ctx) error {
	ctx := c.UserContext()
	_, namespace := getURLParams(ctx, c)
	logger.InfoC(ctx, "Received request to %s %s in namespace %s", action, resourceType, namespace)
	limitChanCreate[resourceType] <- struct{}{}
	defer func() {
		<-limitChanCreate[resourceType]
	}()
	var modelResource R
	if err := c.BodyParser(&modelResource); err != nil {
		pErr := fmt.Errorf("failed to parse request body: %w", err)
		return pmErrors.New(fiber.StatusBadRequest, pErr)
	}
	entityResource := convertFrom(modelResource)
	result, err := createFunc(ctx, &entityResource, namespace)
	if err != nil {
		cErr := fmt.Errorf("failed to %s %s: %w", action, resourceType, err)
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, cErr)
	}
	return respondWithJson(ctx, c, successCode, convertTo(*result))
}

func deleteAdapter(resourceType string, deleteFunc DeleteResource, c *fiber.Ctx) error {
	ctx := c.UserContext()
	resourceName, namespace := getURLParams(ctx, c)
	logger.InfoC(ctx, "Received request to delete %s with name=%s", resourceType, resourceName)
	limitChanDelete[resourceType] <- struct{}{}
	defer func() {
		<-limitChanDelete[resourceType]
	}()
	err := deleteFunc(ctx, resourceName, namespace)
	if err != nil {
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		dErr := fmt.Errorf("failed to delete %s '%s': %w", resourceType, resourceName, err)
		return pmErrors.New(statusCode, dErr)
	}
	logger.DebugC(ctx, "%s '%s' was successfully deleted", resourceType, resourceName)
	c.Status(fiber.StatusOK)
	return nil
}
