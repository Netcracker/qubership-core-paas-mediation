package v2

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/watch"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/context"
)

var (
	// DefaultUpgrader specifies the parameters for upgrading an HTTP
	// connection to a WebSocket connection.

	DefaultUpgrader = &websocket.FastHTTPUpgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(ctx *fasthttp.RequestCtx) bool { return true },
	}
)

//go:generate mockgen -source=wsController.go -destination=wsController_mock.go -package=v2
type WsConn interface {
	ReadMessage() (msgType int, p []byte, err error)
	WriteMessage(msgType int, data []byte) error
	WriteControl(msgType int, data []byte, deadline time.Time) error
	SetCloseHandler(h func(code int, text string) error)
	Close() error
}

type WsController struct {
	Platform service.PlatformService
	Features Features
}

func (contr *WsController) WatchServices(c *fiber.Ctx) error {
	return contr.establishWebSocket(types.Service, c, contr.Platform.WatchServices)
}

func (contr *WsController) WatchConfigMaps(c *fiber.Ctx) error {
	return contr.establishWebSocket(types.ConfigMap, c, contr.Platform.WatchConfigMaps)
}

func (contr *WsController) WatchRoutes(c *fiber.Ctx) error {
	return contr.establishWebSocket(types.Route, c, contr.Platform.WatchRoutes)
}

func (contr *WsController) WatchGatewayHTTPRoutes(c *fiber.Ctx) error {
	if !contr.Features.GatewayRoutesEnabled {
		return respondWithErrorGatewayApiRoutesDisabled(c)
	}
	return contr.establishWebSocket(types.HttpRoute, c, contr.Platform.WatchGatewayHTTPRoutes)
}

func (contr *WsController) WatchGatewayGRPCRoutes(c *fiber.Ctx) error {
	if !contr.Features.GatewayRoutesEnabled {
		return respondWithErrorGatewayApiRoutesDisabled(c)
	}
	return contr.establishWebSocket(types.GrpcRoute, c, contr.Platform.WatchGatewayGRPCRoutes)
}

func (contr *WsController) WatchNamespaces(c *fiber.Ctx) error {
	return contr.establishWebSocket(types.Namespace, c, func(ctx context.Context, namespace string, filter filter.Meta) (*watch.Handler, error) {
		return contr.Platform.WatchNamespaces(ctx, namespace)
	})
}

func (contr *WsController) WatchRollout(c *fiber.Ctx) error {
	queryArgs := &fasthttp.Args{}
	c.Context().QueryArgs().CopyTo(queryArgs)
	return contr.establishWebSocket("rollout", c, func(ctx context.Context, namespace string, filter filter.Meta) (*watch.Handler, error) {
		replicasMap, err := parseFilterParamFromRollout(queryArgs, "replicas")
		if err != nil {
			logger.ErrorC(ctx, "Error while Parsing parameters", err)
			return nil, err
		}
		return contr.Platform.WatchPodsRestarting(ctx, namespace, filter, replicasMap)
	})
}

func (contr *WsController) establishWebSocket(resourceType string, c *fiber.Ctx, watchFunc func(ctx context.Context, namespace string, filter filter.Meta) (*watch.Handler, error)) error {
	ctx := c.UserContext()
	namespace := contr.getNamespace(c)
	logger.InfoC(ctx, "Received a request to watch resource='%s' in namespace=%s", resourceType, namespace)
	metaFilter, err := buildFilterFromParams(c)
	if err != nil {
		logger.Error("An error occurred while parsing 'filter' params: %s", err.Error())
		return respondWithError(ctx, c, 400, err.Error())
	}
	logger.InfoC(ctx, "Applying filter=%+v", metaFilter)

	err = DefaultUpgrader.Upgrade(c.Context(), func(webSocketConn *websocket.Conn) {
		defer func() {
			if webSocketConn != nil {
				if closeErr := webSocketConn.Close(); closeErr != nil {
					logger.ErrorC(ctx, "Failed to close websocket connection resource='%s' in namespace=%s. Error: %s", resourceType, namespace, closeErr.Error())
				}
				logger.DebugC(ctx, "Closed websocket connection resource='%s' in namespace=%s", resourceType, namespace)
			}
		}()
		if webSocketConn == nil {
			logger.ErrorC(ctx, "Web socket connection is nil. Cannot do further processing. resource='%s' in namespace=%s", resourceType, namespace)
			return
		}
		if watchErr := contr.watch(ctx, namespace, resourceType, metaFilter, webSocketConn, watchFunc); watchErr != nil {
			logger.ErrorC(ctx, "Error occurred during watching web socket connection for resource='%s' in namespace=%s: %s", resourceType, namespace, watchErr.Error())
		}
	})
	if err != nil {
		logger.ErrorC(ctx, "Failed to upgrade to websocket connection for resource='%s' in namespace=%s: %s", resourceType, namespace, err.Error())
		return respondWithError(ctx, c, 500, err.Error())
	}
	return nil
}

func (contr *WsController) watch(ctx context.Context, namespace, resourceType string, filter filter.Meta, wsConn WsConn,
	watchFunc func(ctx context.Context, namespace string, filter filter.Meta) (*watch.Handler, error)) error {

	ctx, abortWatchF := context.WithCancel(ctx)
	handler, err := watchFunc(ctx, namespace, filter)
	if err != nil {
		logger.ErrorC(ctx, "Failed to establish watch connection: %s", err.Error())
		// close message cannot be larger than 125, so send constant err
		_ = sendCloseMessage(ctx, wsConn, websocket.CloseInternalServerErr, "failed to establish watch connection")
		return err
	}
	setupPingPong(ctx, wsConn, time.Minute)
	// read from client websocket connection. This allows to monitor connection tear down from the client's side
	go func(wsConn WsConn, cancelWatch context.CancelFunc) {
		defer func() {
			// stop watching resources, because the client has closed the connection
			logger.InfoC(ctx, "Client terminated connection. Cancelling watch context")
			cancelWatch()
		}()
		for {
			if _, _, readErr := wsConn.ReadMessage(); readErr != nil {
				logger.DebugC(ctx, "Error while reading from the client: %+v", readErr)
				// default wsConn's CloseHandler will send close message back to the client
				return
			}
		}
	}(wsConn, abortWatchF)
	// read events from the watch handler
	for {
		select {
		case <-ctx.Done():
			return nil
		case out, ok := <-handler.Channel:
			if ok {
				if err := sendEventViaLimiter(ctx, resourceType, out, wsConn); err != nil {
					return err
				}
			} else {
				logger.InfoC(ctx, "Remote connections was closed. Sending CloseMessage to the client")
				return sendCloseMessage(ctx, wsConn, websocket.CloseNormalClosure, "")
			}
		}
	}
}

func sendEventViaLimiter(ctx context.Context, resourceType string, event watch.ApiEvent, wsConn WsConn) error {
	limitChanWebSocket[resourceType] <- struct{}{}
	defer func() {
		<-limitChanWebSocket[resourceType]
	}()
	body, er := convertToByte(event)
	if er != nil {
		logger.ErrorC(ctx, "Failed to convert event to json: %s", er.Error())
		_ = sendCloseMessage(ctx, wsConn, websocket.CloseInternalServerErr, "")
		return er
	}
	if handleAsCloseControlMessage(ctx, event, wsConn) {
		return nil
	} else {
		logger.DebugC(ctx, "Sending message to the client: %s", body)
		if err := wsConn.WriteMessage(websocket.TextMessage, body); err != nil {
			logger.ErrorC(ctx, "Error during sending a message to the client: %+v", err)
			_ = sendCloseMessage(ctx, wsConn, websocket.CloseInternalServerErr, "")
			return err
		}
	}
	return nil
}

func sendCloseMessage(ctx context.Context, wsConn WsConn, code int, msg string) error {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("recovered from panic which occurred during sendCloseMessage: %+v", err)
		}
	}()
	logger.DebugC(ctx, "Sending close message code=%d, msg=%s", code, msg)
	message := websocket.FormatCloseMessage(code, msg)
	return wsConn.WriteControl(websocket.CloseMessage, message, time.Now().Add(5*time.Second))
}

func setupPingPong(ctx context.Context, conn WsConn, interval time.Duration) {
	go func(ctx context.Context, conn WsConn) {
		pingTicker := time.NewTicker(interval)
		defer pingTicker.Stop()
		for {
			select {
			case <-pingTicker.C:
				if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(5*time.Second)); err != nil {
					logger.WarnC(ctx, "Failed to send ping message: %v", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx, conn)
}

func handleAsCloseControlMessage(ctx context.Context, out watch.ApiEvent, webSocketConn WsConn) bool {
	if out.Type == "CLOSE_CONTROL_MESSAGE" {
		body, err := convertToByte(out)
		logger.DebugC(ctx, "Message to client before closing conn: %s", body)
		if err = webSocketConn.WriteMessage(websocket.TextMessage, body); err != nil {
			logger.ErrorC(ctx, "Error during send a TextMessage to client: %+v", err)
		}

		logger.DebugC(ctx, "Remote connections was closed. Sending CloseMessage to local client: %s", body)
		message := websocket.FormatCloseMessage(out.GetControlMessageDetails().CloseCode, out.GetControlMessageDetails().CloseMessage)
		webSocketConn.WriteControl(out.GetControlMessageDetails().MessageType, message, time.Now().Add(5*time.Second))
		return true
	} else {
		return false
	}
}

func (*WsController) getNamespace(c *fiber.Ctx) string {
	return c.Params(ParamNamespace)
}

func convertToByte(body interface{}) ([]byte, error) {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(body)
	if err != nil {
		return nil, err
	}
	return reqBodyBytes.Bytes(), nil
}

// parse a parameter from URI by its name. The parameter's name and value must be of the following format
// <param_name>=<key_1>:<key_1_value_1>,<key_1_value_2>;<key_2>:<key_2_value_1>,<key_2_value_2>...<key_N>:<key_N_value_1>,..,<key_N_value_N>
// returns the map of parsed keys and list values
// returns error in case param's value has invalid format
func parseFilterParamFromRollout(queryArgs *fasthttp.Args, paramName string) (map[string][]string, error) {
	result := make(map[string][]string)
	paramBytes := queryArgs.Peek(paramName)
	if paramBytes != nil && len(paramBytes) != 0 {
		param := string(paramBytes)
		// parse 'filter' param
		entries := strings.Split(param, ";")
		for _, entry := range entries {
			keyValue := strings.Split(entry, ":")
			if len(keyValue) != 2 {
				return result, errors.New("Invalid parameter '" + paramName + "' provided. Failed to parse entry: '" + entry + "'. " +
					"Valid param structure is <param_name>=<key_1>:<key_1_value>;<key_2>:<key_2_value>")
			} else {
				valuesList := strings.Split(keyValue[1], ",")
				result[keyValue[0]] = valuesList
			}
		}
	}
	return result, nil
}
