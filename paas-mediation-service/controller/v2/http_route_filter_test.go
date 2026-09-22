package v2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestAPIFilters_NilAndRoundTrip(t *testing.T) {
	assert.Nil(t, toAPIFilters(nil))
	assert.Nil(t, fromAPIFilters(nil))

	src := []gatewayv1.HTTPRouteFilter{{
		Type: gatewayv1.HTTPRouteFilterResponseHeaderModifier,
		ResponseHeaderModifier: &gatewayv1.HTTPHeaderFilter{
			Remove: []string{"authorization"},
		},
	}}
	api := toAPIFilters(src)
	require.Len(t, api, 1)
	assert.Equal(t, "ResponseHeaderModifier", api[0].Type)

	back := fromAPIFilters(api)
	require.Len(t, back, 1)
	assert.Equal(t, gatewayv1.HTTPRouteFilterResponseHeaderModifier, back[0].Type)
}

func TestAPIFilters_ExternalAuthRoundTrip(t *testing.T) {
	kind := "Service"
	port := int32(8080)
	maxSize := uint16(1024)

	src := []gatewayv1.HTTPRouteFilter{{
		Type: gatewayv1.HTTPRouteFilterExternalAuth,
		ExternalAuth: &gatewayv1.HTTPExternalAuthFilter{
			ExternalAuthProtocol: gatewayv1.HTTPRouteExternalAuthHTTPProtocol,
			BackendRef: gatewayv1.BackendObjectReference{
				Name: "auth-service",
				Kind: (*gatewayv1.Kind)(&kind),
				Port: (*gatewayv1.PortNumber)(&port),
			},
			HTTPAuthConfig: &gatewayv1.HTTPAuthConfig{
				Path:                   "/auth",
				AllowedRequestHeaders:  []string{"Authorization"},
				AllowedResponseHeaders: []string{"X-User-Id"},
			},
			ForwardBody: &gatewayv1.ForwardBodyConfig{MaxSize: maxSize},
		},
	}}

	api := toAPIFilters(src)
	require.Len(t, api, 1)
	assert.Equal(t, "ExternalAuth", api[0].Type)
	require.NotNil(t, api[0].ExternalAuth)
	assert.Equal(t, "HTTP", api[0].ExternalAuth.Protocol)
	assert.Equal(t, "auth-service", api[0].ExternalAuth.BackendRef.Name)
	require.NotNil(t, api[0].ExternalAuth.HTTP)
	assert.Equal(t, "/auth", api[0].ExternalAuth.HTTP.Path)
	assert.Equal(t, []string{"Authorization"}, api[0].ExternalAuth.HTTP.AllowedRequestHeaders)
	assert.Equal(t, []string{"X-User-Id"}, api[0].ExternalAuth.HTTP.AllowedResponseHeaders)
	require.NotNil(t, api[0].ExternalAuth.ForwardBody)
	assert.Equal(t, maxSize, api[0].ExternalAuth.ForwardBody.MaxSize)

	back := fromAPIFilters(api)
	require.Len(t, back, 1)
	require.NotNil(t, back[0].ExternalAuth)
	assert.Equal(t, gatewayv1.HTTPRouteExternalAuthHTTPProtocol, back[0].ExternalAuth.ExternalAuthProtocol)
	require.NotNil(t, back[0].ExternalAuth.HTTPAuthConfig)
	assert.Equal(t, "/auth", back[0].ExternalAuth.HTTPAuthConfig.Path)
	require.NotNil(t, back[0].ExternalAuth.ForwardBody)
	assert.Equal(t, maxSize, back[0].ExternalAuth.ForwardBody.MaxSize)

	grpcSrc := []HTTPRouteFilter{{
		Type: "ExternalAuth",
		ExternalAuth: &HTTPExternalAuthFilter{
			Protocol:   "GRPC",
			BackendRef: BackendObjectReference{Name: "auth-grpc", Kind: &kind, Port: &port},
			GRPC: &GRPCAuthConfig{
				AllowedRequestHeaders: []string{"x-request-id"},
			},
		},
	}}
	grpcBack := fromAPIFilters(grpcSrc)
	require.Len(t, grpcBack, 1)
	require.NotNil(t, grpcBack[0].ExternalAuth)
	assert.Equal(t, gatewayv1.HTTPRouteExternalAuthGRPCProtocol, grpcBack[0].ExternalAuth.ExternalAuthProtocol)
	require.NotNil(t, grpcBack[0].ExternalAuth.GRPCAuthConfig)
	assert.Equal(t, []string{"x-request-id"}, grpcBack[0].ExternalAuth.GRPCAuthConfig.AllowedRequestHeaders)
}
