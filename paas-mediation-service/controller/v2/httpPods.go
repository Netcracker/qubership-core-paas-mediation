package v2

import (
	"github.com/gofiber/fiber/v3"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
)

// GetPod godoc
//
// @Summary Get Pod by name and namespace
// @Description Get Pod by name and namespace
// @Tags since:2.0
// @ID v2-get-pod
// @Accept  json
// @Produce  json
// @Param	namespace	path     string     true  "target namespace"
// @Param	name		path     string     true  "resource name"
// @Security ApiKeyAuth
// @Success 200 {object}	v2.Pod
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 404 {object}    v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/pods/{name} [get]
func (ctr *HttpController) GetPod(ctx *fiber.Ctx) error {
	return getAdapter(types.Pod, ctr.Platform.GetPod, ToPod, ctx)
}

// GetPodList godoc
//
// @Summary Get Pod by name and namespace
// @Description Get Pod by name and namespace
// @Tags since:2.0
// @ID v2-get-pod-list
// @Accept  json
// @Produce  json
// @Param	namespace		path      string     true   "target namespace"
// @Param	annotations		query     string     false  "resource name"
// @Param	labels		    query     string     false  "resource name"
// @Security ApiKeyAuth
// @Success 200 {array}		v2.Pod
// @Failure 400 {object}	v2.ErrorResponse
// @Failure 403 {object}	v2.ErrorResponse
// @Failure 500 {object}	v2.ErrorResponse
// @Router /api/v2/namespaces/{namespace}/pods [get]
func (ctr *HttpController) GetPodList(ctx *fiber.Ctx) error {
	return listAdapter(types.Pod, ctr.Platform.GetPodList, ToPod, ctx)
}
