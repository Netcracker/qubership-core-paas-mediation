package v2

import (
	"encoding/json"

	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// HTTPRouteFilter mirrors Gateway API HTTPRouteFilter for the REST/OpenAPI surface.
// JSON is compatible with sigs.k8s.io/gateway-api/apis/v1.HTTPRouteFilter.
type HTTPRouteFilter struct {
	// Type identifies the filter. Supported values include RequestHeaderModifier,
	// ResponseHeaderModifier, RequestMirror, RequestRedirect, URLRewrite, CORS,
	// ExternalAuth, ExtensionRef.
	Type                   string                     `json:"type" example:"ResponseHeaderModifier"`
	RequestHeaderModifier  *HTTPHeaderFilter          `json:"requestHeaderModifier,omitempty"`
	ResponseHeaderModifier *HTTPHeaderFilter          `json:"responseHeaderModifier,omitempty"`
	RequestMirror          *HTTPRequestMirrorFilter   `json:"requestMirror,omitempty"`
	RequestRedirect        *HTTPRequestRedirectFilter `json:"requestRedirect,omitempty"`
	URLRewrite             *HTTPURLRewriteFilter      `json:"urlRewrite,omitempty"`
	CORS                   *HTTPCORSFilter            `json:"cors,omitempty"`
	ExternalAuth           *HTTPExternalAuthFilter    `json:"externalAuth,omitempty"`
	ExtensionRef           *LocalObjectReference      `json:"extensionRef,omitempty"`
}

// HTTPHeaderFilter modifies request or response headers.
type HTTPHeaderFilter struct {
	Set    []HTTPHeader `json:"set,omitempty"`
	Add    []HTTPHeader `json:"add,omitempty"`
	Remove []string     `json:"remove,omitempty"`
}

// HTTPHeader is a name/value HTTP header pair.
type HTTPHeader struct {
	Name  string `json:"name" example:"X-Request-Id"`
	Value string `json:"value" example:"abc"`
}

// HTTPPathModifier rewrites or redirects a request path.
type HTTPPathModifier struct {
	Type               string  `json:"type" example:"ReplacePrefixMatch"`
	ReplaceFullPath    *string `json:"replaceFullPath,omitempty"`
	ReplacePrefixMatch *string `json:"replacePrefixMatch,omitempty"`
}

// HTTPRequestRedirectFilter configures an HTTP redirect response.
type HTTPRequestRedirectFilter struct {
	Scheme     *string           `json:"scheme,omitempty" example:"https"`
	Hostname   *string           `json:"hostname,omitempty"`
	Path       *HTTPPathModifier `json:"path,omitempty"`
	Port       *int32            `json:"port,omitempty"`
	StatusCode *int              `json:"statusCode,omitempty" example:"302"`
}

// HTTPURLRewriteFilter rewrites host and/or path while forwarding.
type HTTPURLRewriteFilter struct {
	Hostname *string           `json:"hostname,omitempty"`
	Path     *HTTPPathModifier `json:"path,omitempty"`
}

// HTTPRequestMirrorFilter mirrors a fraction of requests to another backend.
type HTTPRequestMirrorFilter struct {
	BackendRef BackendObjectReference `json:"backendRef"`
	Percent    *int32                 `json:"percent,omitempty"`
	Fraction   *Fraction              `json:"fraction,omitempty"`
}

// HTTPCORSFilter configures CORS response headers.
type HTTPCORSFilter struct {
	AllowOrigins     []string `json:"allowOrigins,omitempty"`
	AllowCredentials *bool    `json:"allowCredentials,omitempty"`
	AllowMethods     []string `json:"allowMethods,omitempty"`
	AllowHeaders     []string `json:"allowHeaders,omitempty"`
	ExposeHeaders    []string `json:"exposeHeaders,omitempty"`
	MaxAge           *int64   `json:"maxAge,omitempty"`
}

// HTTPExternalAuthFilter sends the request to an external auth service.
type HTTPExternalAuthFilter struct {
	Protocol   string                 `json:"protocol,omitempty" example:"HTTP"`
	BackendRef BackendObjectReference `json:"backendRef,omitempty"`
}

// LocalObjectReference references a namespaced API object in the same namespace.
type LocalObjectReference struct {
	Group string `json:"group" example:""`
	Kind  string `json:"kind" example:"ConfigMap"`
	Name  string `json:"name" example:"my-filter"`
}

// BackendObjectReference references a backend (typically a Service).
type BackendObjectReference struct {
	Group     *string `json:"group,omitempty"`
	Kind      *string `json:"kind,omitempty" example:"Service"`
	Name      string  `json:"name" example:"backend"`
	Namespace *string `json:"namespace,omitempty"`
	Port      *int32  `json:"port,omitempty" example:"8080"`
}

// Fraction represents a numerator/denominator pair.
type Fraction struct {
	Numerator   int32  `json:"numerator"`
	Denominator *int32 `json:"denominator,omitempty"`
}

func toAPIFilters(src []gatewayv1.HTTPRouteFilter) []HTTPRouteFilter {
	if src == nil {
		return nil
	}
	raw, err := json.Marshal(src)
	if err != nil {
		return nil
	}
	var dst []HTTPRouteFilter
	if err := json.Unmarshal(raw, &dst); err != nil {
		return nil
	}
	return dst
}

func fromAPIFilters(src []HTTPRouteFilter) []gatewayv1.HTTPRouteFilter {
	if src == nil {
		return nil
	}
	raw, err := json.Marshal(src)
	if err != nil {
		return nil
	}
	var dst []gatewayv1.HTTPRouteFilter
	if err := json.Unmarshal(raw, &dst); err != nil {
		return nil
	}
	return dst
}
