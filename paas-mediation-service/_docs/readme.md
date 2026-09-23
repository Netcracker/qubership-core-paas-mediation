
# How to generate swagger doc manually

Prefer the version-pinned `go:generate` directive in `main.go` (single source of truth).
It must not change `go.mod` / `go.sum`:

```
cd paas-mediation-service
go generate .
```

That runs:

```
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init --generalInfo main.go --parseDependency --parseDepth 2
```

`--parseDependency --parseDepth 2` is required so Gateway API types such as
`gatewayv1.HTTPRouteFilter` (including nested ExternalAuth fields) are expanded
in the OpenAPI schema instead of `items: {}`.

Do **not** run `go get` for swag as part of generation — that can rewrite module
dependencies.

Optional: regenerate Markdown from swagger.json (rewrites the whole `docs/rest_api.md`):

1. install https://github.com/go-swagger/go-swagger
2. from `paas-mediation-service`:
   ```
   swagger generate markdown -f ./docs/swagger.json --output ./../docs/rest_api.md
   ```
