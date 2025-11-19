package v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

func (ctr *HttpController) GetHttpRouteList(ctx *fiber.Ctx) error {
	if !ctr.Features.GatewayRoutesEnabled {
		return respondWithError(ctx.UserContext(), ctx, 404, "gateway routes feature is disabled")
	}
	return listAdapter(types.HttpRoute, ctr.Platform.GetHttpRouteList, Same, ctx)
}

func (ctr *HttpController) GetGrpcRouteList(ctx *fiber.Ctx) error {
	if !ctr.Features.GatewayRoutesEnabled {
		return respondWithError(ctx.UserContext(), ctx, 404, "gateway routes feature is disabled")
	}
	return listAdapter(types.GrpcRoute, ctr.Platform.GetGrpcRouteList, Same, ctx)
}
