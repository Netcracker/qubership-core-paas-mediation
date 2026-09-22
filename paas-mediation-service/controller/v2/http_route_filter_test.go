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
