package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	pmWatch "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/watch"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/controller"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAppResponses(t *testing.T) {
	srv, app := newMockedApp(t, false)
	srv.EXPECT().GetConfigMapList(gomock.Any(), testNamespace, gomock.Any()).Return([]entity.ConfigMap{}, nil).AnyTimes()
	tests := []struct {
		name       string
		target     string
		wantStatus int
		wantBody   string
	}{
		{name: "api-version", target: "/api-version", wantStatus: http.StatusOK},
		{name: "resource list", target: "/api/v2/namespaces/" + testNamespace + "/configmaps", wantStatus: http.StatusOK, wantBody: `[]`},
		{name: "unknown path under /api/v2", target: "/api/v2/unknown", wantStatus: http.StatusNotFound, wantBody: `{"error":"Not Found"}`},
		{name: "dot-dot segment in the namespace path", target: "/api/v2/namespaces/other/../" + testNamespace + "/configmaps", wantStatus: http.StatusOK,
			wantBody: `[]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, contentType, body := sendTestRequest(t, app, httptest.NewRequest(http.MethodGet, tt.target, nil))

			assert.Equal(t, tt.wantStatus, status, "status, body %s", body)
			assert.Equal(t, "application/json", contentType, "Content-Type")
			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, body, "body")
			}
		})
	}
}

func TestCreateConfigMap_RequestContentType(t *testing.T) {
	created := &entity.ConfigMap{Metadata: entity.Metadata{Kind: "ConfigMap", Name: "test-1"}, Data: map[string]string{}}
	createdBody, err := json.Marshal(ToConfigMap(*created))
	require.NoError(t, err)
	const validBody = `{"metadata":{"name":"test-1"}}`
	tests := []struct {
		name        string
		contentType string
		body        string
		wantStatus  int
		wantBody    string
	}{
		{name: "application/json", contentType: "application/json", body: validBody, wantStatus: http.StatusCreated, wantBody: string(createdBody)},
		{name: "application/json with charset", contentType: "application/json; charset=utf-8", body: validBody, wantStatus: http.StatusCreated, wantBody: string(createdBody)},
		{name: "text/json", contentType: "text/json", body: validBody, wantStatus: http.StatusBadRequest, wantBody: `{"error":"failed to parse request body: Unprocessable Entity"}`},
		{name: "vendor +json type", contentType: "application/vnd.netcracker+json", body: validBody, wantStatus: http.StatusCreated, wantBody: string(createdBody)},
		{name: "no Content-Type", body: validBody, wantStatus: http.StatusBadRequest, wantBody: `{"error":"failed to parse request body: Unprocessable Entity"}`},
		{name: "text/plain", contentType: "text/plain", body: validBody, wantStatus: http.StatusBadRequest, wantBody: `{"error":"failed to parse request body: Unprocessable Entity"}`},
		{name: "malformed JSON", contentType: "application/json", body: `{"metadata":`, wantStatus: http.StatusBadRequest, wantBody: `{"error":"failed to parse request body: bind from body: unexpected end of JSON input"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, app := newMockedApp(t, false)
			if tt.wantStatus == http.StatusCreated {
				srv.EXPECT().CreateConfigMap(gomock.Any(), gomock.Any(), testNamespace).Return(created, nil)
			}
			req := httptest.NewRequest(http.MethodPost, "/api/v2/namespaces/"+testNamespace+"/configmaps", strings.NewReader(tt.body))
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}

			status, contentType, body := sendTestRequest(t, app, req)

			assert.Equal(t, tt.wantStatus, status, "status")
			assert.Equal(t, "application/json", contentType, "Content-Type")
			assert.Equal(t, tt.wantBody, body, "body")
		})
	}
}

func TestRequestLogLine(t *testing.T) {
	_, app := newMockedApp(t, false)
	controllerLogger := logging.GetLogger("controller")
	previousLevel := controllerLogger.GetLevel()
	controllerLogger.SetLevel(logging.LvlDebug)
	t.Cleanup(func() { controllerLogger.SetLevel(previousLevel) })

	output := captureOutput(t, func() {
		req := httptest.NewRequest(http.MethodPost, "/api/v2/unknown?x=1", strings.NewReader(`{"a":1}`))
		req.Header.Set("Content-Type", "application/json")
		sendTestRequest(t, app, req)
	})

	pattern := regexp.MustCompile(`Processed request: pid=\d+ source=\S+:\d+ latency=\s*\S+\n` +
		`POST /api/v2/unknown\?x=1\n` +
		`body: \{"a":1\}\n` +
		`status: 404 resBody: \{"error":"Not Found"\}`)
	assert.Regexp(t, pattern, output)
}

func TestPrometheusEndpoint(t *testing.T) {
	registry := prometheus.NewRegistry()
	previousRegisterer, previousGatherer := prometheus.DefaultRegisterer, prometheus.DefaultGatherer
	prometheus.DefaultRegisterer, prometheus.DefaultGatherer = registry, registry
	t.Cleanup(func() { prometheus.DefaultRegisterer, prometheus.DefaultGatherer = previousRegisterer, previousGatherer })
	_, app := newMockedApp(t, true)

	status, _, body := sendTestRequest(t, app, httptest.NewRequest(http.MethodGet, "/prometheus", nil))

	assert.Equal(t, http.StatusOK, status, "status, body %s", body)
}

func TestWatchConfigMapsOverWebSocket(t *testing.T) {
	srv, app := newMockedApp(t, false)
	events := make(chan pmWatch.ApiEvent, 1)
	events <- pmWatch.ApiEvent{Type: "ADDED", Object: map[string]string{"name": "test-1"}}
	srv.EXPECT().WatchConfigMaps(gomock.Any(), testNamespace, gomock.Any()).
		DoAndReturn(func(ctx context.Context, namespace string, filter filter.Meta) (*pmWatch.Handler, error) {
			go func() {
				time.Sleep(200 * time.Millisecond)
				close(events)
			}()
			return &pmWatch.Handler{Channel: events, StopWatching: func() {}}, nil
		})
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	require.NoError(t, err)
	go func() { _ = app.Listener(ln) }()
	t.Cleanup(func() { _ = app.Shutdown() })
	url := fmt.Sprintf("ws://%s/watchapi/v2/namespaces/%s/configmaps", ln.Addr().String(), testNamespace)

	conn, resp, err := websocket.DefaultDialer.Dial(url, nil)

	require.NoError(t, err, "dial %s", url)
	defer conn.Close()
	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode, "upgrade status")
	require.NoError(t, conn.SetReadDeadline(time.Now().Add(5*time.Second)))
	messageType, message, err := conn.ReadMessage()
	require.NoError(t, err, "read the event")
	assert.Equal(t, websocket.TextMessage, messageType, "message type")
	assert.Equal(t, `{"type":"ADDED","object":{"name":"test-1"}}`+"\n", string(message), "event")
	_, _, err = conn.ReadMessage()
	var closeErr *websocket.CloseError
	require.True(t, errors.As(err, &closeErr), "read after the watch ends, error %v", err)
	assert.Equal(t, websocket.CloseNormalClosure, closeErr.Code, "close code")
}

func newMockedApp(t *testing.T, withPrometheus bool) (*service.MockPlatformService, *fiber.App) {
	t.Helper()
	initTestConfig()
	srv := service.NewMockPlatformService(gomock.NewController(t))
	errorHandler := controller.NewErrorHandler()
	ErrorHandler(errorHandler)
	app, err := controller.InitFiber(context.Background(), srv, errorHandler, false, withPrometheus, false)
	require.NoError(t, err)
	SetupRoutes(app, srv, Features{})
	return srv, app
}

func sendTestRequest(t *testing.T, app *fiber.App, req *http.Request) (int, string, string) {
	t.Helper()
	resp, err := app.Test(req, fiber.TestConfig{Timeout: 0})
	require.NoError(t, err, "%s %s", req.Method, req.RequestURI)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "read response body")
	return resp.StatusCode, resp.Header.Get("Content-Type"), string(body)
}

func captureOutput(t *testing.T, action func()) string {
	t.Helper()
	original := os.Stdout
	reader, writer, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = writer
	captured := make(chan string)
	go func() {
		var buffer bytes.Buffer
		_, _ = io.Copy(&buffer, reader)
		captured <- buffer.String()
	}()
	defer func() {
		os.Stdout = original
	}()
	action()
	os.Stdout = original
	require.NoError(t, writer.Close())
	return <-captured
}
