package v2

import (
	"context"
	"errors"
	"github.com/fasthttp/websocket"
	"github.com/golang/mock/gomock"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service"
	pmWatch "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/watch"
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"k8s.io/apimachinery/pkg/watch"
	"sync"
	"testing"
	"time"
)

func TestParseFilterParamFromRollout(t *testing.T) {
	queryArgs := fasthttp.Args{}
	queryArgs.Parse("replicas=replication-controller:dbaas-agent-1,tenant-manager-1;replica-set:paas-mediation-2")
	result, _ := parseFilterParamFromRollout(&queryArgs, "replicas")
	assert.Equal(t, "dbaas-agent-1", result["replication-controller"][0])
	assert.Equal(t, "tenant-manager-1", result["replication-controller"][1])
	assert.Equal(t, "paas-mediation-2", result["replica-set"][0])
}

func TestParseFilterParamFromRolloutEmptyVal(t *testing.T) {
	queryArgs := fasthttp.Args{}
	result, err := parseFilterParamFromRollout(&queryArgs, "replicas")
	assert.Empty(t, result)
	assert.Nil(t, err)
}

func TestParseFilterParamFromRolloutEmptyArgs(t *testing.T) {
	queryArgs := fasthttp.Args{}
	queryArgs.Parse("")
	result, err := parseFilterParamFromRollout(&queryArgs, "replicas")
	assert.Empty(t, result)
	assert.Nil(t, err)
}

func TestParseFilterParamFromRolloutError(t *testing.T) {
	queryArgs := fasthttp.Args{}
	queryArgs.Parse("replicas=;;;")
	result, err := parseFilterParamFromRollout(&queryArgs, "replicas")
	assert.Empty(t, result)
	assert.NotNil(t, err)
}

func TestWsController_SetupPingPongCancel(t *testing.T) {
	assertions := require.New(t)
	ctx, cancelFunc := context.WithCancel(context.Background())
	wsConn := NewMockWsConn(gomock.NewController(t))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wsConn.EXPECT().WriteControl(websocket.PingMessage, nil, gomock.Any()).
		DoAndReturn(func(msgType int, data []byte, deadline time.Time) error {
			cancelFunc()
			wg.Done()
			return errors.New("test error")
		})
	setupPingPong(ctx, wsConn, 50*time.Millisecond)
	assertions.True(waitTimeout(wg, 3*time.Second), "timed out")
}

func TestWsController_WatchError(t *testing.T) {
	assertions := require.New(t)
	wsConn := NewMockWsConn(gomock.NewController(t))
	formatCloseMessage := websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "failed to establish watch connection")
	wsConn.EXPECT().WriteControl(websocket.CloseMessage, formatCloseMessage, gomock.Any()).
		DoAndReturn(func(msgType int, data []byte, deadline time.Time) error {
			return errors.New("test error")
		})
	controller := WsController{}
	err := controller.watch(context.Background(), testNamespace, types.ConfigMap, filter.Meta{}, wsConn,
		func(ctx context.Context, namespace string, filter filter.Meta) (*pmWatch.Handler, error) {
			return nil, errors.New("test error")
		})
	assertions.NotNil(err)
	assertions.Equal("test error", err.Error())
}

func TestWsControllerWatch_ClosureAfterWrite(t *testing.T) {
	assertions := require.New(t)
	ctrl := gomock.NewController(t)
	platformService := service.NewMockPlatformService(ctrl)
	eventsChannel := make(chan pmWatch.ApiEvent, 1)
	watchWasAbortedViaContext := &sync.WaitGroup{}
	watchWasAbortedViaContext.Add(1)
	watchHandler := &pmWatch.Handler{Channel: eventsChannel, StopWatching: func() {}}
	platformService.EXPECT().WatchConfigMaps(gomock.Any(), testNamespace, gomock.Any()).
		DoAndReturn(func(ctx context.Context, namespace string, filter filter.Meta) (*pmWatch.Handler, error) {
			go func() {
				<-ctx.Done()
				watchWasAbortedViaContext.Done()
			}()
			return watchHandler, nil
		})
	eventsChannel <- pmWatch.ApiEvent{
		Type:   string(watch.Added),
		Object: nil,
	}
	wsConn := NewMockWsConn(ctrl)

	writeHappened := &sync.WaitGroup{}
	writeHappened.Add(1)
	wsConn.EXPECT().
		ReadMessage().
		DoAndReturn(func() (int, []byte, error) {
			writeHappened.Wait()
			return 0, nil, &websocket.CloseError{Code: websocket.CloseInternalServerErr}
		})

	wsConn.EXPECT().
		WriteMessage(gomock.Any(), gomock.Any()).
		DoAndReturn(func(msgType int, data []byte) error {
			writeHappened.Done()
			return errors.New("network err")
		})

	wsConn.EXPECT().WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, ""), gomock.Any())

	controller := WsController{platformService}
	watchErr := controller.watch(context.Background(), testNamespace, types.ConfigMap, filter.Meta{}, wsConn, platformService.WatchConfigMaps)

	assertions.NotNil(watchErr)
	assertions.Equal("network err", watchErr.Error())

	AssertWithDeadline(t, func() {
		assertions.Empty(watchHandler.Channel)
		watchWasAbortedViaContext.Wait()
	}, time.Now().Add(time.Second*10))
}

func TestWsControllerWatch_ReadTerminatedBeforeWrite(t *testing.T) {
	assertions := require.New(t)
	ctrl := gomock.NewController(t)
	platformService := service.NewMockPlatformService(ctrl)
	eventsChannel := make(chan pmWatch.ApiEvent)
	watchWasAbortedViaContext := &sync.WaitGroup{}
	watchWasAbortedViaContext.Add(1)
	watchHandler := &pmWatch.Handler{Channel: eventsChannel, StopWatching: func() {}}
	platformService.EXPECT().WatchConfigMaps(gomock.Any(), testNamespace, gomock.Any()).
		DoAndReturn(func(ctx context.Context, namespace string, filter filter.Meta) (*pmWatch.Handler, error) {
			go func() {
				<-ctx.Done()
				watchWasAbortedViaContext.Done()
			}()
			return watchHandler, nil
		})
	wsConn := NewMockWsConn(ctrl)

	wsConn.EXPECT().ReadMessage().
		DoAndReturn(func() (int, []byte, error) {
			return 0, nil, &websocket.CloseError{Code: websocket.CloseNormalClosure}
		})

	controller := WsController{platformService}
	watchErr := controller.watch(context.Background(), testNamespace, types.ConfigMap, filter.Meta{}, wsConn, platformService.WatchConfigMaps)
	assertions.Nil(watchErr)

	AssertWithDeadline(t, func() {
		assertions.Empty(watchHandler.Channel)
		watchWasAbortedViaContext.Wait()
	}, time.Now().Add(time.Second*10))
}

func TestWsControllerWatch_WatchHandlerClosed(t *testing.T) {
	assertions := require.New(t)
	ctrl := gomock.NewController(t)
	platformService := service.NewMockPlatformService(ctrl)
	eventsChannel := make(chan pmWatch.ApiEvent)
	watchWasAbortedViaContext := &sync.WaitGroup{}
	watchWasAbortedViaContext.Add(1)
	watchHandler := &pmWatch.Handler{Channel: eventsChannel, StopWatching: func() {}}
	platformService.EXPECT().WatchConfigMaps(gomock.Any(), testNamespace, gomock.Any()).
		DoAndReturn(func(ctx context.Context, namespace string, filter filter.Meta) (*pmWatch.Handler, error) {
			go func() {
				<-ctx.Done()
				watchWasAbortedViaContext.Done()
			}()
			return watchHandler, nil
		})
	wsConn := NewMockWsConn(ctrl)
	closeMsgWasSent := &sync.WaitGroup{}
	closeMsgWasSent.Add(1)

	wsConn.EXPECT().ReadMessage().
		DoAndReturn(func() (int, []byte, error) {
			closeMsgWasSent.Wait() // block on read until close sent to the client
			return 0, nil, &websocket.CloseError{Code: websocket.CloseNormalClosure, Text: "test"}
		})
	wsConn.EXPECT().WriteControl(websocket.CloseMessage, gomock.Any(), gomock.Any()).
		DoAndReturn(func(msgType int, data []byte, deadline time.Time) error {
			closeMsgWasSent.Done()
			return nil
		})

	controller := WsController{platformService}
	close(eventsChannel)
	watchErr := controller.watch(context.Background(), testNamespace, types.ConfigMap, filter.Meta{}, wsConn, platformService.WatchConfigMaps)
	assertions.Nil(watchErr)

	AssertWithDeadline(t, func() {
		_, opened := <-watchHandler.Channel
		assertions.False(opened)
		assertions.Empty(watchHandler.Channel)
		watchWasAbortedViaContext.Wait()
	}, time.Now().Add(time.Second*10))
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return true // completed normally
	case <-time.After(timeout):
		return false // timed out
	}
}

func AssertWithDeadline(t *testing.T, asserter func(), deadline time.Time) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	assertCompleted := make(chan struct{})
	backgroundAsserter := func() {
		asserter()
		close(assertCompleted)
	}
	go backgroundAsserter()

loop:
	for range ticker.C {
		t.Logf("Try assertion")
		if time.Now().After(deadline) {
			t.Fatalf("assertion failed, time exceeded")
			return
		}
		select {
		case <-assertCompleted:
			t.Logf("Assert happened")
			break loop
		default:
		}
	}
}
