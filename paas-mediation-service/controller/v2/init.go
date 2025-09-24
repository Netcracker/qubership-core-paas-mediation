package v2

import (
	"github.com/gofiber/fiber/v2"
	paasMediation "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	routeregistration "github.com/netcracker/qubership-core-lib-go-rest-utils/v2/route-registration"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/pmservice"
)

var logger logging.Logger

func init() {
	logger = logging.GetLogger("controller/v2")
}

func SetupRoutes(app *fiber.App,
	platformService paasMediation.PlatformService) {

	pmService := pmservice.PmService{Platform: platformService}
	httpContr := HttpController{Platform: platformService, PmService: &pmService}
	wsContr := WsController{Platform: platformService}

	app.Route("/api/v2", func(api fiber.Router) {
		api.Get("/namespaces", httpContr.GetNamespaces)
		api.Route("/namespaces/:namespace", func(api fiber.Router) {
			api.Route("/configmaps", func(configmaps fiber.Router) {
				configmaps.Get("", httpContr.GetConfigMapList)
				configmaps.Get("/", httpContr.GetConfigMapList)
				configmaps.Get("/versions", httpContr.GetVersions)
				configmaps.Get("/bg-version", httpContr.GetBgVersionMap)
				configmaps.Get("/:resource_name", httpContr.GetConfigMap)
				configmaps.Post("", httpContr.CreateConfigMap)
				configmaps.Put("", httpContr.UpdateOrCreateConfigMap)
				configmaps.Delete("/:resource_name", httpContr.DeleteConfigMap)
			})
			api.Route("/routes", func(routes fiber.Router) {
				routes.Get("", httpContr.GetRouteList)
				routes.Get("/", httpContr.GetRouteList)
				routes.Get("/:resource_name", httpContr.GetRoute)
				routes.Post("", httpContr.CreateRoute)
				routes.Put("", httpContr.UpdateOrCreateRoute)
				routes.Delete("/:resource_name", httpContr.DeleteRoute)
			})
			api.Route("/gateway", func(gateways fiber.Router) {
				gateways.Get("/httproutes", httpContr.GetHttpRouteList)
				gateways.Get("/grpcroutes", httpContr.GetGrpcRouteList)

			})
			api.Route("/services", func(services fiber.Router) {
				services.Get("", httpContr.GetServiceList)
				services.Get("/", httpContr.GetServiceList)
				services.Get("/:resource_name", httpContr.GetService)
				services.Post("", httpContr.CreateService)
				services.Put("", httpContr.UpdateOrCreateService)
				services.Delete("/:resource_name", httpContr.DeleteService)
			})
			api.Route("/deployments", func(deployments fiber.Router) {
				deployments.Get("", httpContr.GetDeploymentList)
				deployments.Get("/", httpContr.GetDeploymentList)
				deployments.Get("/:resource_name", httpContr.GetDeployment)
			})
			api.Route("/pods", func(pods fiber.Router) {
				pods.Get("", httpContr.GetPodList)
				pods.Get("/", httpContr.GetPodList)
				pods.Get("/:resource_name", httpContr.GetPod)
			})
			api.Route("/deployment-family", func(deploymentFamily fiber.Router) {
				deploymentFamily.Get("/:family_name", httpContr.GetDeploymentFamilyVersions)
			})
			api.Route("/rollout", func(rollout fiber.Router) {
				rollout.Post("", httpContr.RestartDeploymentsBulk)
				rollout.Post("/", httpContr.RestartDeploymentsBulk)
				rollout.Post("/:resource_name", httpContr.RestartDeployment)
			})
			api.Get("/annotations", httpContr.GetAnnotationResource)
		})
	})
	app.Route("/watchapi/v2", func(watchApi fiber.Router) {
		watchApi.Get("/namespaces", wsContr.WatchNamespaces)
		watchApi.Route("/namespaces/:namespace", func(watchApi fiber.Router) {
			watchApi.Get("/services", wsContr.WatchServices)
			watchApi.Get("/configmaps", wsContr.WatchConfigMaps)
			watchApi.Get("/routes", wsContr.WatchRoutes)
			watchApi.Get("/rollout-status", wsContr.WatchRollout)
			watchApi.Get("/gateway/httproutes", wsContr.WatchGatewayHTTPRoutes)
			watchApi.Get("/gateway/grpcroutes", wsContr.WatchGatewayGRPCRoutes)
		})
	})
}

func WithRoutes(currentNamespace string) {
	routeregistration.NewRegistrar().WithRoutes(
		routeregistration.Route{From: "/watchapi/v2/paas-mediation", To: "/watchapi/v2", RouteType: routeregistration.Internal},
		routeregistration.Route{From: "/api/v2/paas-mediation", To: "/api/v2", RouteType: routeregistration.Internal},
		routeregistration.Route{From: "/api/v2/paas-mediation/namespaces/" + currentNamespace + "/configmaps/bg-version",
			To: "/api/v2/namespaces/" + currentNamespace + "/configmaps/bg-version", RouteType: routeregistration.Private},
		routeregistration.Route{From: "/api/v2/paas-mediation/versions",
			To: "/api/v2/namespaces/" + currentNamespace + "/configmaps/versions", RouteType: routeregistration.Private},
	).Register()
}

func ErrorHandler(holder interface {
	WithErrorHandler(pathPrefix string, handler func(error) any)
}) {
	holder.WithErrorHandler("/api/v2", func(err error) any { return ErrorResponse{Error: err.Error()} })
	holder.WithErrorHandler("/watchapi/v2", func(err error) any { return ErrorResponse{Error: err.Error()} })
}
