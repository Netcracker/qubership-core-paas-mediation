[![Go build](https://github.com/Netcracker/qubership-core-paas-mediation/actions/workflows/go-build.yml/badge.svg)](https://github.com/Netcracker/qubership-core-paas-mediation/actions/workflows/go-build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?metric=coverage&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![duplicated_lines_density](https://sonarcloud.io/api/project_badges/measure?metric=duplicated_lines_density&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![vulnerabilities](https://sonarcloud.io/api/project_badges/measure?metric=vulnerabilities&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![bugs](https://sonarcloud.io/api/project_badges/measure?metric=bugs&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![code_smells](https://sonarcloud.io/api/project_badges/measure?metric=code_smells&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)

# PaaS Mediation Service Documentation


| Section Link                            | Contents                                                                                                                            |
|-----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| [Overview](./docs/paas-mediation-overview.md) | [Basic concept description](./docs/paas-mediation-overview.md) and [Deploy parameters](./docs/paas-mediation-overview.md#deploy-parameters) |
| [REST API](./docs/rest_api.md)           | PaaS Mediation REST API description                                                                                                 |
| [Websocket API](./docs/websocket_api.md) | PaaS Mediation Websocket API description                                                                                            |
| [Gateway API HTTPRoute timeouts](#gateway-api-httproute-timeouts) | Idle timeout via `spec.streamIdleTimeout` (not an annotation) | 


# How to run locally
1. Switch kube context and run devbox with port-forwarding to your namespace
2. Add environment var: \
   for kubernetes ```PAAS_PLATFORM=KUBERNETES;IDP_CLIENT_USERNAME=paas-mediation;IDP_CLIENT_PASSWORD=<secret>;MICROSERVICE_NAMESPACE=<ns>``` \
   for openshift ```PAAS_PLATFORM=OPENSHIFT;IDP_CLIENT_USERNAME=paas-mediation;IDP_CLIENT_PASSWORD=<secret>;MICROSERVICE_NAMESPACE=<ns>```
3. Run with flag ```-local```

# Gateway API HTTPRoute timeouts

When `GATEWAY_SYSTEM_TYPE` contains `gateway-api-default`, create/update route also creates an
[Envoy Gateway BackendTrafficPolicy](https://gateway.envoyproxy.io/docs/api/extension_types/#backendtrafficpolicy)
for idle timeout (and gRPC backend protocol).

**Timeout priority (highest first):**
1. `spec.streamIdleTimeout` — per-route Gateway API duration (`"1800s"`, `"30m"`, `"1h30m"`)
2. `HTTP_ROUTE_REQUEST_IDLE_TIMEOUT` / property `http.route.request.idle.timeout` (service/platform default)
3. Legacy nginx annotations: `max(proxy-read-timeout, proxy-send-timeout)` (values in seconds)

`streamIdleTimeout` is a **spec field**, not an annotation. A metadata annotation named `streamIdleTimeout` is copied onto the HTTPRoute and **ignored**. `HTTP_ROUTE_REQUEST_IDLE_TIMEOUT` is a service-level property, not a per-request annotation.

**POST/PUT `/api/v2/namespaces/{namespace}/routes` — correct:**

```json
{
  "metadata": {
    "name": "example",
    "annotations": {
      "argocd.argoproj.io/tracking-id": "..."
    }
  },
  "spec": {
    "host": "example.com",
    "pathType": "Prefix",
    "path": "/",
    "to": { "name": "example" },
    "port": { "targetPort": 8080 },
    "streamIdleTimeout": "668s"
  }
}
```

Result: BackendTrafficPolicy with `spec.timeout.http.streamIdleTimeout: 668s`. Nginx timeout annotations may be present; they are ignored when `spec.streamIdleTimeout` is set.

**Wrong — timeout in annotations (BTP is not created from this):**

```json
{
  "metadata": {
    "name": "example",
    "annotations": {
      "streamIdleTimeout": "668s"
    }
  },
  "spec": {
    "host": "example.com",
    "path": "/",
    "to": { "name": "example" },
    "port": { "targetPort": 8080 }
  }
}
```

Result: HTTPRoute is created; `streamIdleTimeout` stays as a useless annotation. No BackendTrafficPolicy unless a platform default or nginx `proxy-read-timeout` / `proxy-send-timeout` is set.

**When BackendTrafficPolicy is created:** if a timeout is resolved from the priority list above, or the route has `nginx.ingress.kubernetes.io/backend-protocol: GRPC`. If none of these apply, HTTPRoute is still created and BTP is skipped.

Duration must include a unit (`s`, `m`, `h`, `ms`). `"668"` without a suffix is invalid. Do not use HTTPRoute `spec.timeouts.request` (kills streaming requests).


