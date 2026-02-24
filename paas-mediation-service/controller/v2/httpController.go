package v2

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/constants"
	pmErrors "github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/errors"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/pmservice"
	"github.com/valyala/fasthttp"
)

const (
	ParamResourceName = "resource_name"
	ParamNamespace    = "namespace"
)

type HttpController struct {
	Platform  service.PlatformService
	PmService *pmservice.PmService
	Features  Features
}

// GetDeploymentFamilyVersions godoc
//
// @Summary Get DeploymentFamily data based on Deployments labeled with 'family_name' label with value specified via 'deployment-family' path param
// @Description Get DeploymentFamily data based on Deployments labeled with 'family_name' label with value specified via 'deployment-family' path param
// @Tags since:2.0
// @ID v2-get-deploymentfamily-versions
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	family_name		path      string     true  "family name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.DeploymentFamilyVersion
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/deployment-family/{family_name} [get]
func (ctr *HttpController) GetDeploymentFamilyVersions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	namespace := c.Params(ParamNamespace)
	familyName := c.Params("family_name", "")
	if strings.TrimSpace(familyName) == "" {
		paramsErr := fmt.Errorf("family_name path parameter cannot be empty")
		return pmErrors.New(fiber.StatusBadRequest, paramsErr)
	}
	logger.InfoC(ctx, "Received a request to list deployments with family_name=%s in namespace=%s", familyName, namespace)
	resources, err := ctr.Platform.GetDeploymentFamilyVersions(ctx, familyName, namespace)
	if err != nil {
		logger.Error("An error occurred while getting deployments: %+v", err)
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, err)
	}
	return respondWithJson(ctx, c, fiber.StatusOK, ToDeploymentFamilyVersion(resources))
}

// GetNamespaces godoc
//
// @Summary Get namespaces
// @Description Get namespaces
// @Tags since:2.0
// @ID v2-get-namespaces
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array}		v2.Namespace
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces [get]
func (ctr *HttpController) GetNamespaces(c *fiber.Ctx) error {
	ctx := c.UserContext()
	logger.InfoC(ctx, "Received request on get list of available projects")
	namespaces, err := ctr.Platform.GetNamespaces(ctx, filter.Meta{})
	if err != nil {
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, err)
	}
	var result []Namespace
	for _, ns := range namespaces {
		result = append(result, ToNamespace(ns))
	}
	return respondWithJson(ctx, c, fiber.StatusOK, result)
}

// GetAnnotationResource godoc
//
// @Summary Get resources by resource type and annotation name
// @Description Get resources by resource type and annotation name in namespace
// @Tags since:2.0
// @ID v2-get-annotationresource
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	resourceType	query     string     true   "resource type"
// @Param	annotation		query     string     false  "annotation name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.AnnotationResource
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/annotations [get]
func (ctr *HttpController) GetAnnotationResource(c *fiber.Ctx) error {
	ctx := c.UserContext()
	logger.InfoC(ctx, "Received a request to get an annotation resource")

	namespace := c.Params(ParamNamespace)
	annotationKey, err := getParamsToAnnotationKey(c.Request().URI().QueryArgs())
	if err != nil {
		paramsErr := fmt.Errorf("failed to parse request parameters: %w", err)
		return pmErrors.New(fiber.StatusBadRequest, paramsErr)
	}
	logger.InfoC(ctx, "Request with resource=%s and annotation=%s", annotationKey.ResourceType, annotationKey.Annotation)
	metadata, err := ctr.PmService.GetResourceMeta(ctx, annotationKey.ResourceType, namespace)
	if err != nil {
		var unsupportedErr pmservice.UnsupportedResourceTypeErr
		if errors.As(err, &unsupportedErr) {
			return pmErrors.New(fiber.StatusBadRequest, unsupportedErr)
		}
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, err)
	}
	resources := getAnnotationResourcesFromMetadata(metadata, annotationKey.Annotation, namespace)
	return respondWithJson(ctx, c, fiber.StatusOK, resources)
}

// RestartDeployment godoc
//
// @Summary RestartDeployment
// @Description RestartDeployment by name in namespace
// @Tags since:2.0
// @ID v2-post-restartdeployment
// @Accept  json
// @Produce  json
// @Param	namespace		path     string     true   "target namespace"
// @Param	resource-name	path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.DeploymentResponse
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/rollout/{resource-name} [post]
func (ctr *HttpController) RestartDeployment(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resourceName, namespace := getURLParams(ctx, c)
	var err error
	if strings.TrimSpace(resourceName) == "" {
		paramsErr := fmt.Errorf("resource-name path parameter cannot be empty")
		return pmErrors.New(fiber.StatusBadRequest, paramsErr)
	}
	deploymentsNames := []string{resourceName}
	logger.InfoC(ctx, "Get request to restart deployment=%s in namespace=%s", deploymentsNames, namespace)
	result, err := ctr.Platform.RolloutDeployments(ctx, namespace, deploymentsNames)
	if err != nil {
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, fmt.Errorf("an error occurred while rolling out deployment: %w", err))
	}
	if result == nil {
		logger.Error("Empty rollout result")
		return pmErrors.New(fiber.StatusInternalServerError, err)
	}
	return respondWithJson(ctx, c, fiber.StatusOK, ToDeploymentResponse(namespace, result))
}

// RestartDeploymentsBulk godoc
//
// @Summary Restart Deployments in bulk by names in namespace in parallel or sequentially
// @Description Restart Deployments in bulk by names in namespace in parallel or sequentially
// @Tags since:2.0
// @ID v2-post-restartdeployments-bulk
// @Accept  json
// @Produce  json
// @Param	namespace		path     string     				 true   "target namespace"
// @Param	request 		body     v2.RolloutDeploymentBody    true   "request body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.DeploymentResponse
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/rollout [post]
func (ctr *HttpController) RestartDeploymentsBulk(c *fiber.Ctx) error {
	ctx := c.UserContext()
	namespace := c.Params(ParamNamespace)
	body, pErr := parseRolloutDeploymentBody(c)
	if pErr != nil {
		return pmErrors.New(fiber.StatusBadRequest, fmt.Errorf("invalid body provided. Err: %v", pErr))
	}
	deploymentsNames := body.DeploymentNames
	parallel := body.Parallel
	logger.InfoC(ctx, "Get request to restart deployments=%v in namespace=%s, parallel=%t", deploymentsNames, namespace, parallel)
	var rolloutFunc func(ctx context.Context, namespace string, deploymentNames []string) (*entity.DeploymentResponse, error)
	if parallel {
		rolloutFunc = ctr.Platform.RolloutDeploymentsInParallel
	} else {
		rolloutFunc = ctr.Platform.RolloutDeployments
	}
	result, err := rolloutFunc(ctx, namespace, deploymentsNames)
	if err != nil {
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, fmt.Errorf("an error occurred while rolling out deployments: %w", err))
	}
	if result == nil {
		return pmErrors.New(fiber.StatusInternalServerError, errors.New("empty rollout result"))
	}
	return respondWithJson(ctx, c, fiber.StatusOK, ToDeploymentResponse(namespace, result))
}

// GetBgVersionMap godoc
//
// @Summary Get Blue-Green version ('bg-version') ConfigMap
// @Description Get Blue-Green version ('bg-version') ConfigMap
// @Tags since:2.0
// @ID v2-get-bg-versions
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.AppVersionData
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps/bg-version [get]
func (ctr *HttpController) GetBgVersionMap(c *fiber.Ctx) error {
	ctx := c.UserContext()
	_, namespace := getURLParams(ctx, c)
	versionsMap, err := ctr.Platform.GetConfigMap(ctx, constants.BlueGreenVersionConfigMap, namespace)
	if err != nil {
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, err)
	} else if versionsMap == nil {
		notFoundErr := fmt.Errorf("configmap '%s' not found", constants.BlueGreenVersionConfigMap)
		return pmErrors.New(fiber.StatusNotFound, notFoundErr)
	}
	return respondWithJson(ctx, c, fiber.StatusOK, versionsMap)
}

// GetVersions godoc
//
// @Summary Get versions from 'version' ConfigMap
// @Description Get versions from 'version' ConfigMap
// @Tags since:2.0
// @ID v2-get-versions
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.AppVersionData
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps/versions [get]
func (ctr *HttpController) GetVersions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	_, namespace := getURLParams(ctx, c)
	result, err := ctr.getVersions(ctx, namespace)
	if err != nil {
		logger.Error("An error occurred while getting versions: %+v", err)
		statusCode := pmservice.ErrorConverterToStatusCode(err)
		return pmErrors.New(statusCode, err)
	} else if result == nil {
		notFoundErr := fmt.Errorf("configmap '%s' not found", constants.VersionConfigMap)
		return pmErrors.New(fiber.StatusNotFound, notFoundErr)
	}
	return respondWithJson(ctx, c, fiber.StatusOK, result)
}

func (ctr *HttpController) getVersions(ctx context.Context, namespace string) ([]AppVersionData, error) {
	appData := make([]AppVersionData, 0)
	versionsMap, err := ctr.Platform.GetConfigMap(ctx, constants.VersionConfigMap, namespace)
	if err != nil {
		return nil, err
	}
	if versionsMap == nil {
		return nil, nil
	}
	versions := extractVersionsFromConfigMap(ctx, versionsMap)
	for _, app := range versions {
		if app != nil {
			appData = append(appData, *app)
		}
	}
	return appData, err
}

// app-name.2023-01-26-11-44-51-309 or app.name.2023-01-26-11-44-51-309.Ivan_Ivanov
var versionKeyRegex = regexp.MustCompile(`^(.*)\.([\d-]+)(\.[^:]+)?$`)

// 2006-01-02-15-04-05-000
var deployTimeRegex = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})-(\d{2})-(\d{2})-(\d{2})-(\d{3})$`)

func extractVersionsFromConfigMap(ctx context.Context, versionsMap *entity.ConfigMap) map[string]*AppVersionData {
	versions := map[string]*AppVersionData{}

	for appWithDate, version := range versionsMap.Data {
		//appWithDate format: ApplicationName.DeploymentDate.Username
		//example: app.name.2023-01-26-11-44-51-309.Ivan_Ivanov
		if versionKeyRegex.MatchString(appWithDate) {
			appName := versionKeyRegex.FindStringSubmatch(appWithDate)[1]
			deployTime := versionKeyRegex.FindStringSubmatch(appWithDate)[2]

			if deployTimeRegex.MatchString(deployTime) {
				//Change '-' to '.' Golang can't parse milliseconds in format 2006-01-02-15-04-05-000
				dtFormatted := strings.ReplaceAll(deployTime, constants.Dash, constants.Dot)
				deployTimeParsed, err := time.Parse(constants.DateFormat, dtFormatted)
				if err != nil {
					logger.WarnC(ctx, "failed to parse time '%s', err: %w, skipping", deployTime, err)
					continue
				}
				if appData, ok := versions[appName]; !ok {
					versions[appName] = &AppVersionData{
						AppName:    appName,
						AppVersion: version,
						DeployTime: deployTimeParsed,
					}
				} else {
					if appData.DeployTime.Before(deployTimeParsed) {
						appData.DeployTime = deployTimeParsed
						appData.AppVersion = version
					}
				}
			} else {
				logger.WarnC(ctx, "failed to parse time '%s', invalid format, skipping", deployTime)
			}
		} else {
			logger.WarnC(ctx, "failed to parse configmap 'versions' key: '%s', invalid format. Valid format: '%s' skipping",
				appWithDate, versionKeyRegex.String())
		}
	}
	return versions
}

func parseRolloutDeploymentBody(c *fiber.Ctx) (*RolloutDeploymentBody, error) {
	var body RolloutDeploymentBody
	err := c.BodyParser(&body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing json, %v", err)
	}
	return &body, nil
}

func getURLParams(ctx context.Context, c *fiber.Ctx) (resourceName, namespace string) {
	namespace = c.Params(ParamNamespace)
	resourceName = c.Params(ParamResourceName, "")
	logger.DebugC(ctx, "resources_name=%s", resourceName)

	return resourceName, namespace
}

func getAnnotationResourcesFromMetadata(metadata []entity.Metadata, annotation string, namespace string) []AnnotationResource {
	var resources []AnnotationResource
	var annotationValue string
	for _, foundedMeta := range metadata {
		annotationValue = foundedMeta.Annotations[annotation]
		if annotationValue != "" {
			resources = append(resources, AnnotationResource{ResourceName: foundedMeta.Name, AnnotationValue: annotationValue, Namespace: namespace})
		}
	}
	return resources
}

func getParamsToAnnotationKey(args *fasthttp.Args) (*AnnotationKey, error) {
	var annotation string
	var resourceType string
	resourceTypeBytes := args.Peek("resourceType")
	if resourceTypeBytes != nil && string(resourceTypeBytes) != "" {
		resourceType = string(resourceTypeBytes)
	} else {
		return nil, errors.New("'resourceType' param must be a non empty string")
	}
	annotationBytes := args.Peek("annotation")
	if annotationBytes != nil && string(annotationBytes) != "" {
		annotation = string(annotationBytes)
	}
	return &AnnotationKey{Annotation: annotation, ResourceType: resourceType}, nil
}

func buildFilterFromParams(c *fiber.Ctx) (filter.Meta, error) {
	filter := filter.Meta{}
	annotationsFilter, err := parseFilterParam(c, "annotations")
	if err != nil {
		return filter, err
	}
	labelsFilter, err := parseFilterParam(c, "labels")
	if err != nil {
		return filter, err
	}
	filter.Annotations = annotationsFilter
	filter.Labels = labelsFilter
	return filter, nil
}

// parse a parameter from URI by its name. The parameter's name and value must be of the following format
// <param_name>=<key_1>:<key_1_value>;<key_2>:<key_2_value>...<key_N>:<key_N_value>
// returns the map of parsed keys and values
// returns error in case param's value has invalid format
func parseFilterParam(c *fiber.Ctx, paramName string) (map[string]string, error) {
	result := make(map[string]string)
	if paramByte := c.Request().URI().QueryArgs().Peek(paramName); paramByte != nil {
		param := string(paramByte)
		if param != "" {
			// parse 'filter' param
			entries := strings.Split(param, ";")
			for _, entry := range entries {
				keyValue := strings.Split(entry, ":")
				if len(keyValue) != 2 {
					return result, errors.New(fmt.Sprintf("Invalid parameter '%s' provided. Failed to parse entry: '%s'. "+
						"Valid param structure is param_name=key_1:key_1_value;key_2:key_2_value", paramName, entry))
				} else {
					result[keyValue[0]] = keyValue[1]
				}
			}
		}
	}
	return result, nil
}

func respondWithErrorGatewayApiRoutesDisabled(c *fiber.Ctx) error {
	ctx := c.UserContext()
	return respondWithError(ctx, c, fiber.StatusNotFound, "Gateway API routes observing is disabled")
}

func respondWithError(ctx context.Context, c *fiber.Ctx, code int, msg string) error {
	return respondWithJson(ctx, c, code, map[string]string{"error": msg})
}

func respondWithJson(ctx context.Context, c *fiber.Ctx, code int, payload interface{}) error {
	logger.DebugC(ctx, "Send response code: %v, body: %+v", code, payload)
	return c.Status(code).JSON(payload)
}
