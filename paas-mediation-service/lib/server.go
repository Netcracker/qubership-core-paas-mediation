package lib

import (
	"context"
	fibersec "github.com/netcracker/qubership-core-lib-go-fiber-server-utils/v2/security"
	"github.com/netcracker/qubership-core-lib-go-fiber-server-utils/v2/server"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-lib-go-rest-utils/v2/consul-propertysource"
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"github.com/netcracker/qubership-core-lib-go/v3/security"
	"github.com/netcracker/qubership-core-lib-go/v3/serviceloader"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/controller"
	apiV2 "github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/controller/v2"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/utils"
	"k8s.io/apimachinery/pkg/api/resource"
	"runtime/debug"
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

	serviceloader.Register(1, &security.DummyToken{})
	serviceloader.Register(1, &fibersec.DummyFiberServerSecurityMiddleware{})
}

func RunServer() {
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
