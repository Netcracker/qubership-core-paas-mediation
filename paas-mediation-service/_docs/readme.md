
# How to generate swagger doc manually

Prefer the version-pinned `go:generate` directives in `main.go` (single source of truth):

```
cd paas-mediation-service
go generate
```

That runs `swag` with `--parseDependency --parseDepth 2`, which is required so that
Gateway API types such as `gatewayv1.HTTPRouteFilter` (including nested ExternalAuth
fields) are expanded in the OpenAPI schema instead of `items: {}`.

Manual equivalent:

1. install swag/cmd (version must match `go:generate` in `main.go`, currently v1.16.3; v1.8.12+ is required for dependency parsing)
   ```
   go install github.com/swaggo/swag/cmd/swag@v1.16.3
   ```
2. from within `paas-mediation-service` folder execute:
   ```
   swag init --generalInfo main.go --parseDependency --parseDepth 2
   ```
3. install https://github.com/go-swagger/go-swagger
4. generate MD doc from swagger.json file (optional; regenerates the whole file):
   ```
   swagger generate markdown -f ./docs/swagger.json --output ./../docs/rest_api.md
   ```
