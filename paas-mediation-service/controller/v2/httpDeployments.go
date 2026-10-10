package v2

import (
	"github.com/gofiber/fiber/v3"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

// GetDeployment godoc
//
// @Summary Get Deployment by name and namespace
// @Description Get Deployment by name and namespace
// @Tags since:2.0
// @ID v2-get-deployment
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Deployment
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/deployments/{name} [get]
func (ctr *HttpController) GetDeployment(ctx *fiber.Ctx) error {
	return getAdapter(types.Deployment, ctr.Platform.GetDeployment, ToDeployment, ctx)
}

// GetDeploymentList godoc
//
// @Summary Get Deployment by name and namespace
// @Description Get Deployment by name and namespace
// @Tags since:2.0
// @ID v2-get-deployment-list
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	annotations		query     string     false  "resource name"
// @Param	labels		    query     string     false  "resource name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.Deployment
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/deployments [get]
func (ctr *HttpController) GetDeploymentList(ctx *fiber.Ctx) error {
	return listAdapter(types.Deployment, ctr.Platform.GetDeploymentList, ToDeployment, ctx)
}
