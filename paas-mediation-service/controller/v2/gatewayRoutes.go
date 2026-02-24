package v2

import (
	"github.com/gofiber/fiber/v3"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

// GetHttpRouteList godoc
//
// @Summary Get Gateway API HTTP Routes in namespace
// @Description Get Gateway API HTTP Routes in namespace. This endpoint requires the CORE_PAAS_MEDIATION_GW_API_ENABLED feature flag to be enabled. If the feature flag is disabled, the endpoint will return a 404 error.
// @Tags since:2.0
// @ID v2-get-gateway-httproutes
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Security ApiKeyAuth
// @Success 200 {array}	object
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}	v2.ErrorResponse "Not Found - Gateway routes feature is disabled"
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/gateway/httproutes [get]
func (ctr *HttpController) GetHttpRouteList(ctx *fiber.Ctx) error {
	if !ctr.Features.GatewayRoutesEnabled {
		return respondWithErrorGatewayApiRoutesDisabled(ctx)
	}
	return listAdapter(types.HttpRoute, ctr.Platform.GetHttpRouteList, Same, ctx)
}

// GetGrpcRouteList godoc
//
// @Summary Get Gateway API GRPC Routes in namespace
// @Description Get Gateway API GRPC Routes in namespace. This endpoint requires the CORE_PAAS_MEDIATION_GW_API_ENABLED feature flag to be enabled. If the feature flag is disabled, the endpoint will return a 404 error.
// @Tags since:2.0
// @ID v2-get-gateway-grpcroutes
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Security ApiKeyAuth
// @Success 200 {array}	object
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}	v2.ErrorResponse "Not Found - Gateway routes feature is disabled"
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/gateway/grpcroutes [get]
func (ctr *HttpController) GetGrpcRouteList(ctx *fiber.Ctx) error {
	if !ctr.Features.GatewayRoutesEnabled {
		return respondWithErrorGatewayApiRoutesDisabled(ctx)
	}
	return listAdapter(types.GrpcRoute, ctr.Platform.GetGrpcRouteList, Same, ctx)
}
