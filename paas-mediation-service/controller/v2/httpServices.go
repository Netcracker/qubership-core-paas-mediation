package v2

import (
	"github.com/gofiber/fiber/v2"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-paas-mediation/types"
)

// GetService godoc
//
// @Summary Get Service by name and namespace
// @Description Get Service by name and namespace
// @Tags since:2.0
// @ID v2-get-service
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Service
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/services/{name} [get]
func (ctr *HttpController) GetService(ctx *fiber.Ctx) error {
	return getAdapter(types.Service, ctr.Platform.GetService, ToService, ctx)
}

// GetServiceList godoc
//
// @Summary Get Service by name and namespace
// @Description Get Service by name and namespace
// @Tags since:2.0
// @ID v2-get-service-list
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	annotations		query     string     false  "resource name"
// @Param	labels		    query     string     false  "resource name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.Service
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/services [get]
func (ctr *HttpController) GetServiceList(ctx *fiber.Ctx) error {
	return listAdapter(types.Service, ctr.Platform.GetServiceList, ToService, ctx)
}

// CreateService godoc
//
// @Summary Create Service in namespace
// @Description Create Service in namespace
// @Tags since:2.0
// @ID v2-create-service
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	request		body     v2.Service     true  "resource body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Service
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 409 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/services [post]
func (ctr *HttpController) CreateService(ctx *fiber.Ctx) error {
	return createAdapter[entity.Service, Service](types.Service, ctr.Platform.CreateService, ToService, FromService, ctx)
}

// UpdateOrCreateService godoc
//
// @Summary Update or Create Service in namespace
// @Description Update or Create Service in namespace
// @Tags since:2.0
// @ID v2-update-or-create-service
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	request		body     v2.Service     true  "resource body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Service
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 409 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/services [put]
func (ctr *HttpController) UpdateOrCreateService(ctx *fiber.Ctx) error {
	return updateAdapter[entity.Service, Service](types.Service, ctr.Platform.UpdateOrCreateService, ToService, FromService, ctx)
}

// DeleteService godoc
//
// @Summary Delete Service with name in namespace
// @Description Delete Service with name in namespace
// @Tags since:2.0
// @ID v2-delete-service
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Service
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/services/{name} [delete]
func (ctr *HttpController) DeleteService(ctx *fiber.Ctx) error {
	return deleteAdapter(types.Service, ctr.Platform.DeleteService, ctx)
}
