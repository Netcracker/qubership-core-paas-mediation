
# How to generate swagger doc manually

Prefer the version-pinned `go:generate` directives in `main.go` (single source of truth).
They must not change `go.mod` / `go.sum`:

```
cd paas-mediation-service
go generate .
```

That runs, in order:

1. OpenAPI generation (`swag`):
   ```
   go run github.com/swaggo/swag/cmd/swag@v1.16.6 init --generalInfo main.go --parseDependency --parseDepth 2
   ```
2. Markdown docs from `swagger.json` (`go-swagger`):
   ```
   go run github.com/go-swagger/go-swagger/cmd/swagger@v0.36.6 generate markdown -f ./docs/swagger.json --output ../docs/rest_api.md
   ```

`--parseDependency --parseDepth 2` is required so Gateway API types such as
`gatewayv1.HTTPRouteFilter` (including nested ExternalAuth fields) are expanded
in the OpenAPI schema instead of `items: {}`.

Do **not** run `go get` for these tools as part of generation — that can rewrite module
dependencies.
