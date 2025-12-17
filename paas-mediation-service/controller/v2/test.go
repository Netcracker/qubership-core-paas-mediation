package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/knadh/koanf/providers/confmap"
	fibersec "github.com/netcracker/qubership-core-lib-go-fiber-server-utils/v2/security"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"github.com/netcracker/qubership-core-lib-go/v3/serviceloader"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/controller"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

const (
	testNamespace = "test-namespace"
	metadataName1 = "metadataName1"
	metadataName2 = "metadataName2"
	metadataName3 = "metadataName3"
)

func initTestConfig() {
	configloader.Init(&configloader.PropertySource{Provider: configloader.AsPropertyProvider(confmap.Provider(
		map[string]any{
			"microservice.namespace": "test-namespace",
			"policy.update.enabled":  "false",
			"policy.file.name":       "test/test-policies.conf"}, "."))})

	serviceloader.Register(1, &fibersec.DummyFiberServerSecurityMiddleware{})
}

type fiberAndMockSrv struct {
	srv *service.MockPlatformService
	app *fiber.App
}

func runTestCase(t *testing.T, test *testCase, fiberAndMockSrvOpt ...*fiberAndMockSrv) {
	name := fmt.Sprintf("[%s]-%s-code:[%d]", test.rest.method, test.rest.url, test.rest.code)
	assertions := require.New(t)
	var fAnds *fiberAndMockSrv
	if len(fiberAndMockSrvOpt) > 0 {
		fAnds = fiberAndMockSrvOpt[0]
	}
	t.Run(name, func(t *testing.T) {
		if fAnds == nil {
			fAnds = &fiberAndMockSrv{}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // to make sure mock was executed as expected on each test case
			fAnds.srv = service.NewMockPlatformService(ctrl)
			errorHandler := controller.NewErrorHandler()
			ErrorHandler(errorHandler)
			var err error
			fAnds.app, err = controller.InitFiber(context.Background(), fAnds.srv, errorHandler, false, false, false)
			assertions.Nil(err)
			SetupRoutes(fAnds.app, fAnds.srv, Features{GatewayRoutesEnabled: utils.IsGatewayRoutesEnabled()})
		}
		if test.mock != nil {
			test.mock(fAnds.srv)
		}
		var reqBodyReader io.Reader
		if test.reqBody != nil {
			b, mErr := json.Marshal(test.reqBody)
			assertions.Nil(mErr)
			reqBodyReader = bytes.NewReader(b)
		}
		req := httptest.NewRequest(test.rest.method, test.rest.url, reqBodyReader)
		if test.alterReq != nil {
			test.alterReq(req)
		}
		resp, rErr := fAnds.app.Test(req, -1)
		assertions.Nil(rErr)
		assertions.Equal(test.rest.code, resp.StatusCode)
		if test.respBody != nil {
			bodyBytes, readErr := io.ReadAll(resp.Body)
			assertions.Nil(readErr)
			defer resp.Body.Close()
			var expectedBody string
			switch b := test.respBody.(type) {
			case string:
				expectedBody = b
			case []byte:
				expectedBody = string(b)
			default:
				body, mErr := json.Marshal(b)
				assertions.Nil(mErr)
				expectedBody = string(body)
			}
			assertions.Equal(expectedBody, string(bodyBytes))
		}
		if test.assertResponse != nil {
			test.assertResponse(assertions, resp)
		}
	})
}

type testCase struct {
	rest           r
	alterReq       func(req *http.Request)
	reqBody        any
	respBody       any
	assertResponse func(assertions *require.Assertions, resp *http.Response)
	mock           func(srv *service.MockPlatformService)
}

type r struct {
	method string // GET POST DELETE etc
	url    string // route path to test
	code   int    // expected HTTP status code
}

func url(url string, params ...any) string {
	return fmt.Sprintf(url, params...)
}

func AddHeader(name, value string) func(req *http.Request) {
	return func(req *http.Request) {
		req.Header.Add(name, value)
	}
}
func testGetConcurrency[T entity.HasMetadata](t *testing.T, resourceType string, mfunc func(srv *service.MockPlatformService) *gomock.Call) {
	testGetOrListConcurrency[T](t, resourceType, true, mfunc)
}

func testListConcurrency[T entity.HasMetadata](t *testing.T, resourceType string, mfunc func(srv *service.MockPlatformService) *gomock.Call) {
	testGetOrListConcurrency[T](t, resourceType, false, mfunc)
}

func testGetOrListConcurrency[T entity.HasMetadata](t *testing.T, resourceType string, get bool, mfunc func(srv *service.MockPlatformService) *gomock.Call) {
	initTestConfig()
	var concurrencyLevel int
	if get {
		concurrencyLevel = concurrencyDefault
	} else {
		concurrencyLevel = concurrencyList
	}
	amount := concurrencyLevel * 4
	var wgList []*sync.WaitGroup
	for amount > 0 {
		amount--
		wgList = append(wgList, createWaitGroup())
	}
	parallelControlChannel := make(chan struct{}, concurrencyLevel)
	process := func() {
		select {
		case parallelControlChannel <- struct{}{}:
		default:
			panic(fmt.Sprintf("more then %d parallel request detected", concurrencyLevel))
		}
		time.Sleep(100 * time.Millisecond)
		<-parallelControlChannel
	}
	var createTest func(i int, wg *sync.WaitGroup) *testCase
	if get {
		createTest = func(i int, wg *sync.WaitGroup) *testCase {
			return &testCase{
				rest: r{"GET", url("/api/v2/namespaces/%s/%ss/%s", testNamespace, resourceType, fmt.Sprintf("test-name-%d", i)), 200},
				mock: func(srv *service.MockPlatformService) {
					mfunc(srv).DoAndReturn(func(ctx context.Context, name, namespace string) (*T, error) {
						var result T
						process()
						return &result, nil
					})
				},
				assertResponse: func(assertions *require.Assertions, resp *http.Response) {
					wg.Done()
				},
			}
		}
	} else {
		createTest = func(i int, wg *sync.WaitGroup) *testCase {
			return &testCase{
				rest: r{"GET", url("/api/v2/namespaces/%s/%ss", testNamespace, resourceType), 200},
				mock: func(srv *service.MockPlatformService) {
					mfunc(srv).DoAndReturn(func(ctx context.Context, namespace string, filter filter.Meta) (result []T, err error) {
						process()
						return
					})
				},
				assertResponse: func(assertions *require.Assertions, resp *http.Response) {
					wg.Done()
				},
			}
		}
	}
	var tests []*testCase
	for i, wg := range wgList {
		tests = append(tests, createTest(i, wg))
	}
	assertions := require.New(t)
	ctrl := gomock.NewController(t)
	srv := service.NewMockPlatformService(ctrl)
	errorHandler := controller.NewErrorHandler()
	ErrorHandler(errorHandler)
	app, err := controller.InitFiber(context.Background(), srv, errorHandler, false, false, false)
	assertions.Nil(err)
	SetupRoutes(app, srv, Features{GatewayRoutesEnabled: utils.IsGatewayRoutesEnabled()})
	fiberAndMockSrvOpt := &fiberAndMockSrv{srv: srv, app: app}

	for _, tc := range tests {
		go runTestCase(t, tc, fiberAndMockSrvOpt)
	}
	assertions.True(waitWG(5*time.Second, wgList...))
}

func createWaitGroup() *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	return wg
}

func waitWG(timeout time.Duration, groups ...*sync.WaitGroup) bool {
	finalWg := &sync.WaitGroup{}
	finalWg.Add(len(groups))
	c := make(chan struct{})
	for _, wg := range groups {
		go func(wg *sync.WaitGroup) {
			wg.Wait()
			finalWg.Done()
		}(wg)
	}
	go func() {
		finalWg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(timeout)
	select {
	case <-c:
		timer.Stop()
		return true
	case <-timer.C:
		return false
	}
}
