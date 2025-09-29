package v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

func (ctr *HttpController) GetHttpRouteList(ctx *fiber.Ctx) error {
	return listAdapter(types.HttpRoute, ctr.Platform.GetHttpRouteList, Same, ctx)
}

func (ctr *HttpController) GetGrpcRouteList(ctx *fiber.Ctx) error {
	return listAdapter(types.GrpcRoute, ctr.Platform.GetGrpcRouteList, Same, ctx)
}
