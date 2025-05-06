package v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

// GetRoute godoc
//
// @Summary Get Route by name and namespace
// @Description Get Route by name and namespace
// @Tags since:2.0
// @ID v2-get-route
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Route
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/routes/{name} [get]
func (ctr *HttpController) GetRoute(ctx *fiber.Ctx) error {
	return getAdapter(types.Route, ctr.Platform.GetRoute, ToRoute, ctx)
}

// GetRouteList godoc
//
// @Summary Get Route by name and namespace
// @Description Get Route by name and namespace
// @Tags since:2.0
// @ID v2-get-route-list
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	annotations		query     string     false  "resource name"
// @Param	labels		    query     string     false  "resource name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.Route
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/routes [get]
func (ctr *HttpController) GetRouteList(ctx *fiber.Ctx) error {
	return listAdapter(types.Route, ctr.Platform.GetRouteList, ToRoute, ctx)
}

// CreateRoute godoc
//
// @Summary Create Route in namespace
// @Description Create Route in namespace
// @Tags since:2.0
// @ID v2-create-route
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	request		body     v2.Route     true  "resource body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Route
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 409 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/routes [post]
func (ctr *HttpController) CreateRoute(ctx *fiber.Ctx) error {
	return createAdapter[entity.Route, Route](types.Route, ctr.Platform.CreateRoute, ToRoute, FromRoute, ctx)
}

// UpdateOrCreateRoute godoc
//
// @Summary Update or Create Route in namespace
// @Description Update or Create Route in namespace
// @Tags since:2.0
// @ID v2-update-or-create-route
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	request		body     v2.Route     true  "resource body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Route
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 409 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/routes [put]
func (ctr *HttpController) UpdateOrCreateRoute(ctx *fiber.Ctx) error {
	return updateAdapter[entity.Route, Route](types.Route, ctr.Platform.UpdateOrCreateRoute, ToRoute, FromRoute, ctx)
}

// DeleteRoute godoc
//
// @Summary Delete Route with name in namespace
// @Description Delete Route with name in namespace
// @Tags since:2.0
// @ID v2-delete-route
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Route
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/routes/{name} [delete]
func (ctr *HttpController) DeleteRoute(ctx *fiber.Ctx) error {
	return deleteAdapter(types.Route, ctr.Platform.DeleteRoute, ctx)
}
