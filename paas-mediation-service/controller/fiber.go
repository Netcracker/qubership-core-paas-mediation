package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/netcracker/qubership-core-lib-go-actuator-common/v2/health"
	"github.com/netcracker/qubership-core-lib-go-actuator-common/v2/tracing"
	fiberserver "github.com/netcracker/qubership-core-lib-go-fiber-server-utils/v2"
	paasMediation "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"github.com/netcracker/qubership-core-paas-mediation/docs"
	pmErrors "github.com/netcracker/qubership-core-paas-mediation/errors"
	"io"
	"net/http"
	"strings"
	"time"
)

var logger = logging.GetLogger("controller")

func InitFiber(ctx context.Context, platformService paasMediation.PlatformService, errorHandler *ErrorHandler,
	withPprof, withPrometheus, withTracer bool) (*fiber.App, error) {
	healthCheck := func() (status health.Status) {
		badRoutes, err := platformService.GetBadRouteLists(ctx)
		if err != nil {
			logger.Error("An error occurred while to GET bad routes: %+v", err)
		}
		return health.Status{Name: health.StatusUp,
			Details: map[string]interface{}{
				"badRoutes": badRoutes,
			},
		}
	}
	healthService, err := health.NewHealthService()
	if err != nil {
		return nil, fmt.Errorf("couldn't create healthService: %w", err)
	}
	healthService.AddCheck("badResourcesHealthCheck", healthCheck)
	builder := fiberserver.New(fiber.Config{
		Network:      fiber.NetworkTCP,
		IdleTimeout:  30 * time.Second,
		ErrorHandler: errorHandler.Handler()})
	if withPprof {
		builder.WithPprof("3030")
	}
	if withPrometheus {
		builder.WithPrometheus("prometheus")
	}
	if withTracer {
		builder.WithTracer(tracing.NewZipkinTracer())
	}
	app, err := builder.
		WithHealth("/health", healthService).
		WithDeprecatedApiSwitchedOff().
		WithLogLevelsInfo().
		ProcessWithContext(ctx)
	if err != nil {
		return nil, err
	}
	// logging
	app.Use(fiberLogger.New(fiberLogger.Config{
		Done: func(c *fiber.Ctx, logString []byte) {
			logger.DebugC(c.UserContext(), string(logString))
		},
		Format: "Processed request: pid=${pid} source=${ip}:${port} latency=${latency}\n" +
			"${method} ${url}\n" +
			"body: ${body}\n" +
			"status: ${status} resBody: ${resBody}",
		Output: io.Discard,
	}))
	// swagger
	app.Get("/swagger-ui/swagger.json", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")
		return ctx.Status(http.StatusOK).SendString(docs.SwaggerInfo.ReadDoc())
	})
	// api-version
	app.Get("/api-version", ApiVersion)
	return app, nil
}

// ApiVersion godoc
//
// @Summary Get Api Version information
// @Description Get Major, Minor and Supported Major versions
// @Tags api version info
// @ID api-version
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object}	controller.ApiVersionResponse
// @Router /api-version [get]
func ApiVersion(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(
		ApiVersionResponse{
			Info: Info{
				Major:           &docs.MajorVersion,
				Minor:           &docs.MinorVersion,
				SupportedMajors: docs.SupportedMajors,
			},
			Specs: []Info{
				{
					SpecRootUrl:     "/api",
					Major:           &docs.MajorVersion,
					Minor:           &docs.MinorVersion,
					SupportedMajors: docs.SupportedMajors,
				},
			},
		})
}

type Info struct {
	SpecRootUrl     string `json:"specRootUrl"`
	Major           *int   `json:"major"`
	Minor           *int   `json:"minor"`
	SupportedMajors []int  `json:"supportedMajors"`
}

type ApiVersionResponse struct {
	Info
	Specs []Info `json:"specs"`
}

type ErrorHandlerFunc func(err error) any

type ErrorHandler struct {
	handlers map[string]ErrorHandlerFunc
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{handlers: make(map[string]ErrorHandlerFunc, 0)}
}

func (h *ErrorHandler) Handler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var httpErr pmErrors.HttpError
		var fiberErr *fiber.Error
		var targetErr error
		if errors.As(err, &httpErr) {
			code = httpErr.Code
			targetErr = httpErr.Err
		} else if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			targetErr = fiberErr
		} else {
			code = http.StatusInternalServerError
			targetErr = err
		}
		currentPath := c.Path()
		var handler ErrorHandlerFunc
		for pathPrefix, handle := range h.handlers {
			if strings.HasPrefix(currentPath, pathPrefix) {
				handler = handle
				break
			}
		}
		var errorBody any
		if handler != nil {
			errorBody = handler(targetErr)
		} else {
			errorBody = map[string]string{
				"error": targetErr.Error(),
			}
		}
		var logFunc func(ctx context.Context, format string, args ...any)
		if code >= 400 && code <= 499 {
			logFunc = logger.WarnC
		} else if code >= 500 {
			logFunc = logger.ErrorC
		} else {
			// ignore
			logFunc = func(ctx context.Context, format string, args ...any) {}
		}
		logFunc(c.UserContext(), fmt.Sprintf("%s : response code: %d %s", targetErr.Error(), code, http.StatusText(code)))
		return c.Status(code).JSON(errorBody)
	}
}

func (h *ErrorHandler) WithErrorHandler(pathPrefix string, handler func(error) any) {
	h.handlers[pathPrefix] = handler
}
