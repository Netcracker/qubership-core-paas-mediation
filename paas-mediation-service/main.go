package main

//go:generate go get github.com/swaggo/swag/cmd/swag@v1.16.3
//go:generate go run github.com/swaggo/swag/cmd/swag init --generalInfo /main.go --parseDependency --parseDepth 2

import (
	"context"
	"github.com/netcracker/qubership-core-lib-go-fiber-server-utils/v2/server"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-lib-go-rest-utils/v2/consul-propertysource"
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"github.com/netcracker/qubership-core-paas-mediation/controller"
	apiV2 "github.com/netcracker/qubership-core-paas-mediation/controller/v2"
	"github.com/netcracker/qubership-core-paas-mediation/utils"
	"k8s.io/apimachinery/pkg/api/resource"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"runtime/debug"
	// swagger docs
	_ "github.com/netcracker/qubership-core-paas-mediation/docs"
)

var (
	logger logging.Logger
)

func init() {
	consulPS := consul.NewLoggingPropertySource()
	propertySources := consul.AddConsulPropertySource(configloader.BasePropertySources())
	configloader.InitWithSourcesArray(append(propertySources, consulPS))
	consul.StartWatchingForPropertiesWithRetry(context.Background(), consulPS, func(event interface{}, err error) {})
	logger = logging.GetLogger("main")
}

// @title           			Paas Mediation API
// @version         			2.0
// @description     			API for Paas Mediation.
// @tag.name                    api version info
// @tag.description             Apis provides information related to versions
// @tag.name                    since:2.0
// @tag.description             Apis existed since 2.0 version
// @Produce 					json
// @securityDefinitions.apikey 	ApiKeyAuth
// @in 							header
// @name 						Authorization
func main() {
	ctx := context.Background()
	// need to set soft memory limit based on what was passed to k8s container
	memoryLimit := utils.GetMemoryLimit()
	memoryLimit = resource.NewQuantity(memoryLimit.Value()*9/10, memoryLimit.Format)
	logger.Info("Setting memoryLimit to: %s", memoryLimit.String())
	debug.SetMemoryLimit(memoryLimit.Value())

	platformClient, err := service.NewPlatformClientBuilder().WithWatchClientTimeout(utils.GetWatchClientTimeout()).
		WithAllCaches().WithCacheSettings(utils.GetCacheSettings()).Build()
	if err != nil {
		panic("Cannot create Platform Builder: " + err.Error())
	}
	errorHandler := controller.NewErrorHandler()
	app, err := controller.InitFiber(ctx, platformClient, errorHandler, true, true, true)
	if err != nil {
		panic("Cannot create Fiber app: " + err.Error())
	}
	namespace := configloader.GetKoanf().MustString("microservice.namespace")

	apiV2.SetupRoutes(app, platformClient)
	apiV2.WithRoutes(namespace)
	apiV2.ErrorHandler(errorHandler)

	server.StartServer(app, "http.server.bind")
}
