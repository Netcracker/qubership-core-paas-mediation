package v2

import (
	"github.com/gofiber/fiber/v3"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

// GetConfigMap godoc
//
// @Summary Get ConfigMap by name and namespace
// @Description Get ConfigMap by name and namespace
// @Tags since:2.0
// @ID v2-get-configmap
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.ConfigMap
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps/{name} [get]
func (ctr *HttpController) GetConfigMap(ctx *fiber.Ctx) error {
	return getAdapter(types.ConfigMap, ctr.Platform.GetConfigMap, ToConfigMap, ctx)
}

// GetConfigMapList godoc
//
// @Summary Get ConfigMap by name and namespace
// @Description Get ConfigMap by name and namespace
// @Tags since:2.0
// @ID v2-get-configmap-list
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	annotations		query     string     false  "resource name"
// @Param	labels		    query     string     false  "resource name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.ConfigMap
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps [get]
func (ctr *HttpController) GetConfigMapList(ctx *fiber.Ctx) error {
	return listAdapter(types.ConfigMap, ctr.Platform.GetConfigMapList, ToConfigMap, ctx)
}

// CreateConfigMap godoc
//
// @Summary Create ConfigMap in namespace
// @Description Create ConfigMap in namespace
// @Tags since:2.0
// @ID v2-create-configmap
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     	  true  "target namespace"
// @Param	request		body     v2.ConfigMap     true  "resource body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.ConfigMap
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 409 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps [post]
func (ctr *HttpController) CreateConfigMap(ctx *fiber.Ctx) error {
	return createAdapter[entity.ConfigMap, ConfigMap](types.ConfigMap, ctr.Platform.CreateConfigMap, ToConfigMap, FromConfigMap, ctx)
}

// UpdateOrCreateConfigMap godoc
//
// @Summary Update or Create ConfigMap in namespace
// @Description Update or Create ConfigMap in namespace
// @Tags since:2.0
// @ID v2-update-or-create-configmap
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	request		body     v2.ConfigMap     true  "resource body"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.ConfigMap
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 409 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps [put]
func (ctr *HttpController) UpdateOrCreateConfigMap(ctx *fiber.Ctx) error {
	return updateAdapter[entity.ConfigMap, ConfigMap](types.ConfigMap, ctr.Platform.UpdateOrCreateConfigMap, ToConfigMap, FromConfigMap, ctx)
}

// DeleteConfigMap godoc
//
// @Summary Delete ConfigMap with name in namespace
// @Description Delete ConfigMap with name in namespace
// @Tags since:2.0
// @ID v2-delete-configmap
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.ConfigMap
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/configmaps/{name} [delete]
func (ctr *HttpController) DeleteConfigMap(ctx *fiber.Ctx) error {
	return deleteAdapter(types.ConfigMap, ctr.Platform.DeleteConfigMap, ctx)
}
