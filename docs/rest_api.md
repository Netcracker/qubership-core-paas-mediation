


# Paas Mediation API
API for Paas Mediation.
  

## Informations

### Version

2.0

### Contact

  

## Tags

  ### <span id="tag-api-version-info"></span>api version info

Apis provides information related to versions

  ### <span id="tag-since-2-0"></span>since:2.0

Apis existed since 2.0 version

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### ApiKeyAuth (header: Authorization)



> **Type**: apikey

## All endpoints

###  api_version_info

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api-version | [api version](#api-version) | Get Api Version information |
  


###  since_2_0

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/namespaces/{namespace}/configmaps | [v2 create configmap](#v2-create-configmap) | Create ConfigMap in namespace |
| POST | /api/v2/namespaces/{namespace}/routes | [v2 create route](#v2-create-route) | Create Route in namespace |
| POST | /api/v2/namespaces/{namespace}/services | [v2 create service](#v2-create-service) | Create Service in namespace |
| DELETE | /api/v2/namespaces/{namespace}/configmaps/{name} | [v2 delete configmap](#v2-delete-configmap) | Delete ConfigMap with name in namespace |
| DELETE | /api/v2/namespaces/{namespace}/routes/{name} | [v2 delete route](#v2-delete-route) | Delete Route with name in namespace |
| DELETE | /api/v2/namespaces/{namespace}/services/{name} | [v2 delete service](#v2-delete-service) | Delete Service with name in namespace |
| GET | /api/v2/namespaces/{namespace}/annotations | [v2 get annotationresource](#v2-get-annotationresource) | Get resources by resource type and annotation name |
| GET | /api/v2/namespaces/{namespace}/configmaps/bg-version | [v2 get bg versions](#v2-get-bg-versions) | Get Blue-Green version ('bg-version') ConfigMap |
| GET | /api/v2/namespaces/{namespace}/configmaps/{name} | [v2 get configmap](#v2-get-configmap) | Get ConfigMap by name and namespace |
| GET | /api/v2/namespaces/{namespace}/configmaps | [v2 get configmap list](#v2-get-configmap-list) | Get ConfigMap by name and namespace |
| GET | /api/v2/namespaces/{namespace}/deployments/{name} | [v2 get deployment](#v2-get-deployment) | Get Deployment by name and namespace |
| GET | /api/v2/namespaces/{namespace}/deployments | [v2 get deployment list](#v2-get-deployment-list) | Get Deployment by name and namespace |
| GET | /api/v2/namespaces/{namespace}/deployment-family/{family_name} | [v2 get deploymentfamily versions](#v2-get-deploymentfamily-versions) | Get DeploymentFamily data based on Deployments labeled with 'family_name' label with value specified via 'deployment-family' path param |
| GET | /api/v2/namespaces/{namespace}/gateway/grpcroutes | [v2 get gateway grpcroutes](#v2-get-gateway-grpcroutes) | Get Gateway API GRPC Routes in namespace |
| GET | /api/v2/namespaces/{namespace}/gateway/httproutes | [v2 get gateway httproutes](#v2-get-gateway-httproutes) | Get Gateway API HTTP Routes in namespace |
| GET | /api/v2/namespaces | [v2 get namespaces](#v2-get-namespaces) | Get namespaces |
| GET | /api/v2/namespaces/{namespace}/pods/{name} | [v2 get pod](#v2-get-pod) | Get Pod by name and namespace |
| GET | /api/v2/namespaces/{namespace}/pods | [v2 get pod list](#v2-get-pod-list) | Get Pod by name and namespace |
| GET | /api/v2/namespaces/{namespace}/routes/{name} | [v2 get route](#v2-get-route) | Get Route by name and namespace |
| GET | /api/v2/namespaces/{namespace}/routes | [v2 get route list](#v2-get-route-list) | Get Route by name and namespace |
| GET | /api/v2/namespaces/{namespace}/services/{name} | [v2 get service](#v2-get-service) | Get Service by name and namespace |
| GET | /api/v2/namespaces/{namespace}/services | [v2 get service list](#v2-get-service-list) | Get Service by name and namespace |
| GET | /api/v2/namespaces/{namespace}/configmaps/versions | [v2 get versions](#v2-get-versions) | Get versions from 'version' ConfigMap |
| POST | /api/v2/namespaces/{namespace}/rollout/{resource-name} | [v2 post restartdeployment](#v2-post-restartdeployment) | RestartDeployment |
| POST | /api/v2/namespaces/{namespace}/rollout | [v2 post restartdeployments bulk](#v2-post-restartdeployments-bulk) | Restart Deployments in bulk by names in namespace in parallel or sequentially |
| PUT | /api/v2/namespaces/{namespace}/configmaps | [v2 update or create configmap](#v2-update-or-create-configmap) | Update or Create ConfigMap in namespace |
| PUT | /api/v2/namespaces/{namespace}/routes | [v2 update or create route](#v2-update-or-create-route) | Update or Create Route in namespace |
| PUT | /api/v2/namespaces/{namespace}/services | [v2 update or create service](#v2-update-or-create-service) | Update or Create Service in namespace |
  


## Paths

### <span id="api-version"></span> Get Api Version information (*api-version*)

```
GET /api-version
```

Get Major, Minor and Supported Major versions

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#api-version-200) | OK | OK |  | [schema](#api-version-200-schema) |

#### Responses


##### <span id="api-version-200"></span> 200 - OK
Status: OK

###### <span id="api-version-200-schema"></span> Schema
   
  

[ControllerAPIVersionResponse](#controller-api-version-response)

### <span id="v2-create-configmap"></span> Create ConfigMap in namespace (*v2-create-configmap*)

```
POST /api/v2/namespaces/{namespace}/configmaps
```

Create ConfigMap in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2ConfigMap](#v2-config-map) | `models.V2ConfigMap` | | ✓ | | resource body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-create-configmap-200) | OK | OK |  | [schema](#v2-create-configmap-200-schema) |
| [400](#v2-create-configmap-400) | Bad Request | Bad Request |  | [schema](#v2-create-configmap-400-schema) |
| [403](#v2-create-configmap-403) | Forbidden | Forbidden |  | [schema](#v2-create-configmap-403-schema) |
| [409](#v2-create-configmap-409) | Conflict | Conflict |  | [schema](#v2-create-configmap-409-schema) |
| [500](#v2-create-configmap-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-create-configmap-500-schema) |

#### Responses


##### <span id="v2-create-configmap-200"></span> 200 - OK
Status: OK

###### <span id="v2-create-configmap-200-schema"></span> Schema
   
  

[V2ConfigMap](#v2-config-map)

##### <span id="v2-create-configmap-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-create-configmap-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-configmap-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-create-configmap-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-configmap-409"></span> 409 - Conflict
Status: Conflict

###### <span id="v2-create-configmap-409-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-configmap-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-create-configmap-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-create-route"></span> Create Route in namespace (*v2-create-route*)

```
POST /api/v2/namespaces/{namespace}/routes
```

Create Route in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2Route](#v2-route) | `models.V2Route` | | ✓ | | resource body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-create-route-200) | OK | OK |  | [schema](#v2-create-route-200-schema) |
| [400](#v2-create-route-400) | Bad Request | Bad Request |  | [schema](#v2-create-route-400-schema) |
| [403](#v2-create-route-403) | Forbidden | Forbidden |  | [schema](#v2-create-route-403-schema) |
| [409](#v2-create-route-409) | Conflict | Conflict |  | [schema](#v2-create-route-409-schema) |
| [500](#v2-create-route-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-create-route-500-schema) |

#### Responses


##### <span id="v2-create-route-200"></span> 200 - OK
Status: OK

###### <span id="v2-create-route-200-schema"></span> Schema
   
  

[V2Route](#v2-route)

##### <span id="v2-create-route-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-create-route-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-route-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-create-route-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-route-409"></span> 409 - Conflict
Status: Conflict

###### <span id="v2-create-route-409-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-route-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-create-route-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-create-service"></span> Create Service in namespace (*v2-create-service*)

```
POST /api/v2/namespaces/{namespace}/services
```

Create Service in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2Service](#v2-service) | `models.V2Service` | | ✓ | | resource body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-create-service-200) | OK | OK |  | [schema](#v2-create-service-200-schema) |
| [400](#v2-create-service-400) | Bad Request | Bad Request |  | [schema](#v2-create-service-400-schema) |
| [403](#v2-create-service-403) | Forbidden | Forbidden |  | [schema](#v2-create-service-403-schema) |
| [409](#v2-create-service-409) | Conflict | Conflict |  | [schema](#v2-create-service-409-schema) |
| [500](#v2-create-service-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-create-service-500-schema) |

#### Responses


##### <span id="v2-create-service-200"></span> 200 - OK
Status: OK

###### <span id="v2-create-service-200-schema"></span> Schema
   
  

[V2Service](#v2-service)

##### <span id="v2-create-service-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-create-service-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-service-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-create-service-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-service-409"></span> 409 - Conflict
Status: Conflict

###### <span id="v2-create-service-409-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-create-service-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-create-service-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-delete-configmap"></span> Delete ConfigMap with name in namespace (*v2-delete-configmap*)

```
DELETE /api/v2/namespaces/{namespace}/configmaps/{name}
```

Delete ConfigMap with name in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-delete-configmap-200) | OK | OK |  | [schema](#v2-delete-configmap-200-schema) |
| [400](#v2-delete-configmap-400) | Bad Request | Bad Request |  | [schema](#v2-delete-configmap-400-schema) |
| [403](#v2-delete-configmap-403) | Forbidden | Forbidden |  | [schema](#v2-delete-configmap-403-schema) |
| [500](#v2-delete-configmap-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-delete-configmap-500-schema) |

#### Responses


##### <span id="v2-delete-configmap-200"></span> 200 - OK
Status: OK

###### <span id="v2-delete-configmap-200-schema"></span> Schema
   
  

[V2ConfigMap](#v2-config-map)

##### <span id="v2-delete-configmap-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-delete-configmap-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-delete-configmap-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-delete-configmap-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-delete-configmap-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-delete-configmap-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-delete-route"></span> Delete Route with name in namespace (*v2-delete-route*)

```
DELETE /api/v2/namespaces/{namespace}/routes/{name}
```

Delete Route with name in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-delete-route-200) | OK | OK |  | [schema](#v2-delete-route-200-schema) |
| [400](#v2-delete-route-400) | Bad Request | Bad Request |  | [schema](#v2-delete-route-400-schema) |
| [403](#v2-delete-route-403) | Forbidden | Forbidden |  | [schema](#v2-delete-route-403-schema) |
| [500](#v2-delete-route-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-delete-route-500-schema) |

#### Responses


##### <span id="v2-delete-route-200"></span> 200 - OK
Status: OK

###### <span id="v2-delete-route-200-schema"></span> Schema
   
  

[V2Route](#v2-route)

##### <span id="v2-delete-route-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-delete-route-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-delete-route-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-delete-route-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-delete-route-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-delete-route-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-delete-service"></span> Delete Service with name in namespace (*v2-delete-service*)

```
DELETE /api/v2/namespaces/{namespace}/services/{name}
```

Delete Service with name in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-delete-service-200) | OK | OK |  | [schema](#v2-delete-service-200-schema) |
| [400](#v2-delete-service-400) | Bad Request | Bad Request |  | [schema](#v2-delete-service-400-schema) |
| [403](#v2-delete-service-403) | Forbidden | Forbidden |  | [schema](#v2-delete-service-403-schema) |
| [500](#v2-delete-service-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-delete-service-500-schema) |

#### Responses


##### <span id="v2-delete-service-200"></span> 200 - OK
Status: OK

###### <span id="v2-delete-service-200-schema"></span> Schema
   
  

[V2Service](#v2-service)

##### <span id="v2-delete-service-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-delete-service-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-delete-service-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-delete-service-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-delete-service-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-delete-service-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-annotationresource"></span> Get resources by resource type and annotation name (*v2-get-annotationresource*)

```
GET /api/v2/namespaces/{namespace}/annotations
```

Get resources by resource type and annotation name in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| annotation | `query` | string | `string` |  |  |  | annotation name |
| resourceType | `query` | string | `string` |  | ✓ |  | resource type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-annotationresource-200) | OK | OK |  | [schema](#v2-get-annotationresource-200-schema) |
| [400](#v2-get-annotationresource-400) | Bad Request | Bad Request |  | [schema](#v2-get-annotationresource-400-schema) |
| [403](#v2-get-annotationresource-403) | Forbidden | Forbidden |  | [schema](#v2-get-annotationresource-403-schema) |
| [500](#v2-get-annotationresource-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-annotationresource-500-schema) |

#### Responses


##### <span id="v2-get-annotationresource-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-annotationresource-200-schema"></span> Schema
   
  

[][V2AnnotationResource](#v2-annotation-resource)

##### <span id="v2-get-annotationresource-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-annotationresource-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-annotationresource-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-annotationresource-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-annotationresource-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-annotationresource-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-bg-versions"></span> Get Blue-Green version ('bg-version') ConfigMap (*v2-get-bg-versions*)

```
GET /api/v2/namespaces/{namespace}/configmaps/bg-version
```

Get Blue-Green version ('bg-version') ConfigMap

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-bg-versions-200) | OK | OK |  | [schema](#v2-get-bg-versions-200-schema) |
| [403](#v2-get-bg-versions-403) | Forbidden | Forbidden |  | [schema](#v2-get-bg-versions-403-schema) |
| [404](#v2-get-bg-versions-404) | Not Found | Not Found |  | [schema](#v2-get-bg-versions-404-schema) |
| [500](#v2-get-bg-versions-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-bg-versions-500-schema) |

#### Responses


##### <span id="v2-get-bg-versions-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-bg-versions-200-schema"></span> Schema
   
  

[][V2AppVersionData](#v2-app-version-data)

##### <span id="v2-get-bg-versions-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-bg-versions-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-bg-versions-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-bg-versions-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-bg-versions-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-bg-versions-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-configmap"></span> Get ConfigMap by name and namespace (*v2-get-configmap*)

```
GET /api/v2/namespaces/{namespace}/configmaps/{name}
```

Get ConfigMap by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-configmap-200) | OK | OK |  | [schema](#v2-get-configmap-200-schema) |
| [400](#v2-get-configmap-400) | Bad Request | Bad Request |  | [schema](#v2-get-configmap-400-schema) |
| [403](#v2-get-configmap-403) | Forbidden | Forbidden |  | [schema](#v2-get-configmap-403-schema) |
| [404](#v2-get-configmap-404) | Not Found | Not Found |  | [schema](#v2-get-configmap-404-schema) |
| [500](#v2-get-configmap-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-configmap-500-schema) |

#### Responses


##### <span id="v2-get-configmap-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-configmap-200-schema"></span> Schema
   
  

[V2ConfigMap](#v2-config-map)

##### <span id="v2-get-configmap-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-configmap-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-configmap-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-configmap-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-configmap-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-configmap-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-configmap-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-configmap-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-configmap-list"></span> Get ConfigMap by name and namespace (*v2-get-configmap-list*)

```
GET /api/v2/namespaces/{namespace}/configmaps
```

Get ConfigMap by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| annotations | `query` | string | `string` |  |  |  | resource name |
| labels | `query` | string | `string` |  |  |  | resource name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-configmap-list-200) | OK | OK |  | [schema](#v2-get-configmap-list-200-schema) |
| [400](#v2-get-configmap-list-400) | Bad Request | Bad Request |  | [schema](#v2-get-configmap-list-400-schema) |
| [403](#v2-get-configmap-list-403) | Forbidden | Forbidden |  | [schema](#v2-get-configmap-list-403-schema) |
| [500](#v2-get-configmap-list-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-configmap-list-500-schema) |

#### Responses


##### <span id="v2-get-configmap-list-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-configmap-list-200-schema"></span> Schema
   
  

[][V2ConfigMap](#v2-config-map)

##### <span id="v2-get-configmap-list-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-configmap-list-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-configmap-list-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-configmap-list-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-configmap-list-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-configmap-list-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-deployment"></span> Get Deployment by name and namespace (*v2-get-deployment*)

```
GET /api/v2/namespaces/{namespace}/deployments/{name}
```

Get Deployment by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-deployment-200) | OK | OK |  | [schema](#v2-get-deployment-200-schema) |
| [400](#v2-get-deployment-400) | Bad Request | Bad Request |  | [schema](#v2-get-deployment-400-schema) |
| [403](#v2-get-deployment-403) | Forbidden | Forbidden |  | [schema](#v2-get-deployment-403-schema) |
| [404](#v2-get-deployment-404) | Not Found | Not Found |  | [schema](#v2-get-deployment-404-schema) |
| [500](#v2-get-deployment-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-deployment-500-schema) |

#### Responses


##### <span id="v2-get-deployment-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-deployment-200-schema"></span> Schema
   
  

[V2Deployment](#v2-deployment)

##### <span id="v2-get-deployment-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-deployment-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deployment-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-deployment-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deployment-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-deployment-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deployment-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-deployment-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-deployment-list"></span> Get Deployment by name and namespace (*v2-get-deployment-list*)

```
GET /api/v2/namespaces/{namespace}/deployments
```

Get Deployment by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| annotations | `query` | string | `string` |  |  |  | resource name |
| labels | `query` | string | `string` |  |  |  | resource name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-deployment-list-200) | OK | OK |  | [schema](#v2-get-deployment-list-200-schema) |
| [400](#v2-get-deployment-list-400) | Bad Request | Bad Request |  | [schema](#v2-get-deployment-list-400-schema) |
| [403](#v2-get-deployment-list-403) | Forbidden | Forbidden |  | [schema](#v2-get-deployment-list-403-schema) |
| [500](#v2-get-deployment-list-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-deployment-list-500-schema) |

#### Responses


##### <span id="v2-get-deployment-list-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-deployment-list-200-schema"></span> Schema
   
  

[][V2Deployment](#v2-deployment)

##### <span id="v2-get-deployment-list-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-deployment-list-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deployment-list-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-deployment-list-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deployment-list-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-deployment-list-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-deploymentfamily-versions"></span> Get DeploymentFamily data based on Deployments labeled with 'family_name' label with value specified via 'deployment-family' path param (*v2-get-deploymentfamily-versions*)

```
GET /api/v2/namespaces/{namespace}/deployment-family/{family_name}
```

Get DeploymentFamily data based on Deployments labeled with 'family_name' label with value specified via 'deployment-family' path param

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| family_name | `path` | string | `string` |  | ✓ |  | family name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-deploymentfamily-versions-200) | OK | OK |  | [schema](#v2-get-deploymentfamily-versions-200-schema) |
| [400](#v2-get-deploymentfamily-versions-400) | Bad Request | Bad Request |  | [schema](#v2-get-deploymentfamily-versions-400-schema) |
| [403](#v2-get-deploymentfamily-versions-403) | Forbidden | Forbidden |  | [schema](#v2-get-deploymentfamily-versions-403-schema) |
| [500](#v2-get-deploymentfamily-versions-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-deploymentfamily-versions-500-schema) |

#### Responses


##### <span id="v2-get-deploymentfamily-versions-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-deploymentfamily-versions-200-schema"></span> Schema
   
  

[][V2DeploymentFamilyVersion](#v2-deployment-family-version)

##### <span id="v2-get-deploymentfamily-versions-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-deploymentfamily-versions-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deploymentfamily-versions-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-deploymentfamily-versions-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-deploymentfamily-versions-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-deploymentfamily-versions-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-gateway-grpcroutes"></span> Get Gateway API GRPC Routes in namespace (*v2-get-gateway-grpcroutes*)

```
GET /api/v2/namespaces/{namespace}/gateway/grpcroutes
```

Get Gateway API GRPC Routes in namespace. This endpoint requires the GATEWAY_SYSTEM_TYPE feature flag to contain gateway-api-default. If the feature flag is absent, the endpoint will return a 404 error.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-gateway-grpcroutes-200) | OK | OK |  | [schema](#v2-get-gateway-grpcroutes-200-schema) |
| [400](#v2-get-gateway-grpcroutes-400) | Bad Request | Bad Request |  | [schema](#v2-get-gateway-grpcroutes-400-schema) |
| [403](#v2-get-gateway-grpcroutes-403) | Forbidden | Forbidden |  | [schema](#v2-get-gateway-grpcroutes-403-schema) |
| [404](#v2-get-gateway-grpcroutes-404) | Not Found | Not Found - Gateway routes feature is disabled |  | [schema](#v2-get-gateway-grpcroutes-404-schema) |
| [500](#v2-get-gateway-grpcroutes-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-gateway-grpcroutes-500-schema) |

#### Responses


##### <span id="v2-get-gateway-grpcroutes-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-gateway-grpcroutes-200-schema"></span> Schema
   
  

[][any](#any)

##### <span id="v2-get-gateway-grpcroutes-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-gateway-grpcroutes-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-gateway-grpcroutes-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-gateway-grpcroutes-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-gateway-grpcroutes-404"></span> 404 - Not Found - Gateway routes feature is disabled
Status: Not Found

###### <span id="v2-get-gateway-grpcroutes-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-gateway-grpcroutes-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-gateway-grpcroutes-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-gateway-httproutes"></span> Get Gateway API HTTP Routes in namespace (*v2-get-gateway-httproutes*)

```
GET /api/v2/namespaces/{namespace}/gateway/httproutes
```

Get Gateway API HTTP Routes in namespace. This endpoint requires the GATEWAY_SYSTEM_TYPE feature flag to contain gateway-api-default. If the feature flag is absent, the endpoint will return a 404 error.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-gateway-httproutes-200) | OK | OK |  | [schema](#v2-get-gateway-httproutes-200-schema) |
| [400](#v2-get-gateway-httproutes-400) | Bad Request | Bad Request |  | [schema](#v2-get-gateway-httproutes-400-schema) |
| [403](#v2-get-gateway-httproutes-403) | Forbidden | Forbidden |  | [schema](#v2-get-gateway-httproutes-403-schema) |
| [404](#v2-get-gateway-httproutes-404) | Not Found | Not Found - Gateway routes feature is disabled |  | [schema](#v2-get-gateway-httproutes-404-schema) |
| [500](#v2-get-gateway-httproutes-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-gateway-httproutes-500-schema) |

#### Responses


##### <span id="v2-get-gateway-httproutes-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-gateway-httproutes-200-schema"></span> Schema
   
  

[][any](#any)

##### <span id="v2-get-gateway-httproutes-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-gateway-httproutes-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-gateway-httproutes-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-gateway-httproutes-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-gateway-httproutes-404"></span> 404 - Not Found - Gateway routes feature is disabled
Status: Not Found

###### <span id="v2-get-gateway-httproutes-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-gateway-httproutes-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-gateway-httproutes-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-namespaces"></span> Get namespaces (*v2-get-namespaces*)

```
GET /api/v2/namespaces
```

Get namespaces

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-namespaces-200) | OK | OK |  | [schema](#v2-get-namespaces-200-schema) |
| [400](#v2-get-namespaces-400) | Bad Request | Bad Request |  | [schema](#v2-get-namespaces-400-schema) |
| [403](#v2-get-namespaces-403) | Forbidden | Forbidden |  | [schema](#v2-get-namespaces-403-schema) |
| [500](#v2-get-namespaces-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-namespaces-500-schema) |

#### Responses


##### <span id="v2-get-namespaces-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-namespaces-200-schema"></span> Schema
   
  

[][V2Namespace](#v2-namespace)

##### <span id="v2-get-namespaces-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-namespaces-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-namespaces-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-namespaces-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-namespaces-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-namespaces-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-pod"></span> Get Pod by name and namespace (*v2-get-pod*)

```
GET /api/v2/namespaces/{namespace}/pods/{name}
```

Get Pod by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-pod-200) | OK | OK |  | [schema](#v2-get-pod-200-schema) |
| [400](#v2-get-pod-400) | Bad Request | Bad Request |  | [schema](#v2-get-pod-400-schema) |
| [403](#v2-get-pod-403) | Forbidden | Forbidden |  | [schema](#v2-get-pod-403-schema) |
| [404](#v2-get-pod-404) | Not Found | Not Found |  | [schema](#v2-get-pod-404-schema) |
| [500](#v2-get-pod-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-pod-500-schema) |

#### Responses


##### <span id="v2-get-pod-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-pod-200-schema"></span> Schema
   
  

[V2Pod](#v2-pod)

##### <span id="v2-get-pod-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-pod-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-pod-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-pod-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-pod-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-pod-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-pod-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-pod-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-pod-list"></span> Get Pod by name and namespace (*v2-get-pod-list*)

```
GET /api/v2/namespaces/{namespace}/pods
```

Get Pod by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| annotations | `query` | string | `string` |  |  |  | resource name |
| labels | `query` | string | `string` |  |  |  | resource name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-pod-list-200) | OK | OK |  | [schema](#v2-get-pod-list-200-schema) |
| [400](#v2-get-pod-list-400) | Bad Request | Bad Request |  | [schema](#v2-get-pod-list-400-schema) |
| [403](#v2-get-pod-list-403) | Forbidden | Forbidden |  | [schema](#v2-get-pod-list-403-schema) |
| [500](#v2-get-pod-list-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-pod-list-500-schema) |

#### Responses


##### <span id="v2-get-pod-list-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-pod-list-200-schema"></span> Schema
   
  

[][V2Pod](#v2-pod)

##### <span id="v2-get-pod-list-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-pod-list-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-pod-list-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-pod-list-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-pod-list-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-pod-list-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-route"></span> Get Route by name and namespace (*v2-get-route*)

```
GET /api/v2/namespaces/{namespace}/routes/{name}
```

Get Route by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-route-200) | OK | OK |  | [schema](#v2-get-route-200-schema) |
| [400](#v2-get-route-400) | Bad Request | Bad Request |  | [schema](#v2-get-route-400-schema) |
| [403](#v2-get-route-403) | Forbidden | Forbidden |  | [schema](#v2-get-route-403-schema) |
| [404](#v2-get-route-404) | Not Found | Not Found |  | [schema](#v2-get-route-404-schema) |
| [500](#v2-get-route-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-route-500-schema) |

#### Responses


##### <span id="v2-get-route-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-route-200-schema"></span> Schema
   
  

[V2Route](#v2-route)

##### <span id="v2-get-route-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-route-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-route-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-route-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-route-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-route-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-route-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-route-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-route-list"></span> Get Route by name and namespace (*v2-get-route-list*)

```
GET /api/v2/namespaces/{namespace}/routes
```

Get Route by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| annotations | `query` | string | `string` |  |  |  | resource name |
| labels | `query` | string | `string` |  |  |  | resource name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-route-list-200) | OK | OK |  | [schema](#v2-get-route-list-200-schema) |
| [400](#v2-get-route-list-400) | Bad Request | Bad Request |  | [schema](#v2-get-route-list-400-schema) |
| [403](#v2-get-route-list-403) | Forbidden | Forbidden |  | [schema](#v2-get-route-list-403-schema) |
| [500](#v2-get-route-list-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-route-list-500-schema) |

#### Responses


##### <span id="v2-get-route-list-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-route-list-200-schema"></span> Schema
   
  

[][V2Route](#v2-route)

##### <span id="v2-get-route-list-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-route-list-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-route-list-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-route-list-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-route-list-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-route-list-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-service"></span> Get Service by name and namespace (*v2-get-service*)

```
GET /api/v2/namespaces/{namespace}/services/{name}
```

Get Service by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  | resource name |
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-service-200) | OK | OK |  | [schema](#v2-get-service-200-schema) |
| [400](#v2-get-service-400) | Bad Request | Bad Request |  | [schema](#v2-get-service-400-schema) |
| [403](#v2-get-service-403) | Forbidden | Forbidden |  | [schema](#v2-get-service-403-schema) |
| [404](#v2-get-service-404) | Not Found | Not Found |  | [schema](#v2-get-service-404-schema) |
| [500](#v2-get-service-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-service-500-schema) |

#### Responses


##### <span id="v2-get-service-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-service-200-schema"></span> Schema
   
  

[V2Service](#v2-service)

##### <span id="v2-get-service-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-service-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-service-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-service-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-service-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-service-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-service-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-service-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-service-list"></span> Get Service by name and namespace (*v2-get-service-list*)

```
GET /api/v2/namespaces/{namespace}/services
```

Get Service by name and namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| annotations | `query` | string | `string` |  |  |  | resource name |
| labels | `query` | string | `string` |  |  |  | resource name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-service-list-200) | OK | OK |  | [schema](#v2-get-service-list-200-schema) |
| [400](#v2-get-service-list-400) | Bad Request | Bad Request |  | [schema](#v2-get-service-list-400-schema) |
| [403](#v2-get-service-list-403) | Forbidden | Forbidden |  | [schema](#v2-get-service-list-403-schema) |
| [500](#v2-get-service-list-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-service-list-500-schema) |

#### Responses


##### <span id="v2-get-service-list-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-service-list-200-schema"></span> Schema
   
  

[][V2Service](#v2-service)

##### <span id="v2-get-service-list-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-get-service-list-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-service-list-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-service-list-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-service-list-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-service-list-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-get-versions"></span> Get versions from 'version' ConfigMap (*v2-get-versions*)

```
GET /api/v2/namespaces/{namespace}/configmaps/versions
```

Get versions from 'version' ConfigMap

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-get-versions-200) | OK | OK |  | [schema](#v2-get-versions-200-schema) |
| [403](#v2-get-versions-403) | Forbidden | Forbidden |  | [schema](#v2-get-versions-403-schema) |
| [404](#v2-get-versions-404) | Not Found | Not Found |  | [schema](#v2-get-versions-404-schema) |
| [500](#v2-get-versions-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-get-versions-500-schema) |

#### Responses


##### <span id="v2-get-versions-200"></span> 200 - OK
Status: OK

###### <span id="v2-get-versions-200-schema"></span> Schema
   
  

[][V2AppVersionData](#v2-app-version-data)

##### <span id="v2-get-versions-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-get-versions-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-versions-404"></span> 404 - Not Found
Status: Not Found

###### <span id="v2-get-versions-404-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-get-versions-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-get-versions-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-post-restartdeployment"></span> RestartDeployment (*v2-post-restartdeployment*)

```
POST /api/v2/namespaces/{namespace}/rollout/{resource-name}
```

RestartDeployment by name in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| resource-name | `path` | string | `string` |  | ✓ |  | resource name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-post-restartdeployment-200) | OK | OK |  | [schema](#v2-post-restartdeployment-200-schema) |
| [400](#v2-post-restartdeployment-400) | Bad Request | Bad Request |  | [schema](#v2-post-restartdeployment-400-schema) |
| [403](#v2-post-restartdeployment-403) | Forbidden | Forbidden |  | [schema](#v2-post-restartdeployment-403-schema) |
| [500](#v2-post-restartdeployment-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-post-restartdeployment-500-schema) |

#### Responses


##### <span id="v2-post-restartdeployment-200"></span> 200 - OK
Status: OK

###### <span id="v2-post-restartdeployment-200-schema"></span> Schema
   
  

[V2DeploymentResponse](#v2-deployment-response)

##### <span id="v2-post-restartdeployment-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-post-restartdeployment-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-post-restartdeployment-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-post-restartdeployment-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-post-restartdeployment-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-post-restartdeployment-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-post-restartdeployments-bulk"></span> Restart Deployments in bulk by names in namespace in parallel or sequentially (*v2-post-restartdeployments-bulk*)

```
POST /api/v2/namespaces/{namespace}/rollout
```

Restart Deployments in bulk by names in namespace in parallel or sequentially

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2RolloutDeploymentBody](#v2-rollout-deployment-body) | `models.V2RolloutDeploymentBody` | | ✓ | | request body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-post-restartdeployments-bulk-200) | OK | OK |  | [schema](#v2-post-restartdeployments-bulk-200-schema) |
| [400](#v2-post-restartdeployments-bulk-400) | Bad Request | Bad Request |  | [schema](#v2-post-restartdeployments-bulk-400-schema) |
| [403](#v2-post-restartdeployments-bulk-403) | Forbidden | Forbidden |  | [schema](#v2-post-restartdeployments-bulk-403-schema) |
| [500](#v2-post-restartdeployments-bulk-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-post-restartdeployments-bulk-500-schema) |

#### Responses


##### <span id="v2-post-restartdeployments-bulk-200"></span> 200 - OK
Status: OK

###### <span id="v2-post-restartdeployments-bulk-200-schema"></span> Schema
   
  

[V2DeploymentResponse](#v2-deployment-response)

##### <span id="v2-post-restartdeployments-bulk-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-post-restartdeployments-bulk-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-post-restartdeployments-bulk-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-post-restartdeployments-bulk-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-post-restartdeployments-bulk-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-post-restartdeployments-bulk-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-update-or-create-configmap"></span> Update or Create ConfigMap in namespace (*v2-update-or-create-configmap*)

```
PUT /api/v2/namespaces/{namespace}/configmaps
```

Update or Create ConfigMap in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2ConfigMap](#v2-config-map) | `models.V2ConfigMap` | | ✓ | | resource body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-update-or-create-configmap-200) | OK | OK |  | [schema](#v2-update-or-create-configmap-200-schema) |
| [400](#v2-update-or-create-configmap-400) | Bad Request | Bad Request |  | [schema](#v2-update-or-create-configmap-400-schema) |
| [403](#v2-update-or-create-configmap-403) | Forbidden | Forbidden |  | [schema](#v2-update-or-create-configmap-403-schema) |
| [409](#v2-update-or-create-configmap-409) | Conflict | Conflict |  | [schema](#v2-update-or-create-configmap-409-schema) |
| [500](#v2-update-or-create-configmap-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-update-or-create-configmap-500-schema) |

#### Responses


##### <span id="v2-update-or-create-configmap-200"></span> 200 - OK
Status: OK

###### <span id="v2-update-or-create-configmap-200-schema"></span> Schema
   
  

[V2ConfigMap](#v2-config-map)

##### <span id="v2-update-or-create-configmap-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-update-or-create-configmap-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-configmap-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-update-or-create-configmap-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-configmap-409"></span> 409 - Conflict
Status: Conflict

###### <span id="v2-update-or-create-configmap-409-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-configmap-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-update-or-create-configmap-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-update-or-create-route"></span> Update or Create Route in namespace (*v2-update-or-create-route*)

```
PUT /api/v2/namespaces/{namespace}/routes
```

Update or Create Route in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2Route](#v2-route) | `models.V2Route` | | ✓ | | resource body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-update-or-create-route-200) | OK | OK |  | [schema](#v2-update-or-create-route-200-schema) |
| [400](#v2-update-or-create-route-400) | Bad Request | Bad Request |  | [schema](#v2-update-or-create-route-400-schema) |
| [403](#v2-update-or-create-route-403) | Forbidden | Forbidden |  | [schema](#v2-update-or-create-route-403-schema) |
| [409](#v2-update-or-create-route-409) | Conflict | Conflict |  | [schema](#v2-update-or-create-route-409-schema) |
| [500](#v2-update-or-create-route-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-update-or-create-route-500-schema) |

#### Responses


##### <span id="v2-update-or-create-route-200"></span> 200 - OK
Status: OK

###### <span id="v2-update-or-create-route-200-schema"></span> Schema
   
  

[V2Route](#v2-route)

##### <span id="v2-update-or-create-route-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-update-or-create-route-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-route-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-update-or-create-route-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-route-409"></span> 409 - Conflict
Status: Conflict

###### <span id="v2-update-or-create-route-409-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-route-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-update-or-create-route-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

### <span id="v2-update-or-create-service"></span> Update or Create Service in namespace (*v2-update-or-create-service*)

```
PUT /api/v2/namespaces/{namespace}/services
```

Update or Create Service in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * ApiKeyAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| namespace | `path` | string | `string` |  | ✓ |  | target namespace |
| request | `body` | [V2Service](#v2-service) | `models.V2Service` | | ✓ | | resource body |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#v2-update-or-create-service-200) | OK | OK |  | [schema](#v2-update-or-create-service-200-schema) |
| [400](#v2-update-or-create-service-400) | Bad Request | Bad Request |  | [schema](#v2-update-or-create-service-400-schema) |
| [403](#v2-update-or-create-service-403) | Forbidden | Forbidden |  | [schema](#v2-update-or-create-service-403-schema) |
| [409](#v2-update-or-create-service-409) | Conflict | Conflict |  | [schema](#v2-update-or-create-service-409-schema) |
| [500](#v2-update-or-create-service-500) | Internal Server Error | Internal Server Error |  | [schema](#v2-update-or-create-service-500-schema) |

#### Responses


##### <span id="v2-update-or-create-service-200"></span> 200 - OK
Status: OK

###### <span id="v2-update-or-create-service-200-schema"></span> Schema
   
  

[V2Service](#v2-service)

##### <span id="v2-update-or-create-service-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="v2-update-or-create-service-400-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-service-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="v2-update-or-create-service-403-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-service-409"></span> 409 - Conflict
Status: Conflict

###### <span id="v2-update-or-create-service-409-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

##### <span id="v2-update-or-create-service-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="v2-update-or-create-service-500-schema"></span> Schema
   
  

[V2ErrorResponse](#v2-error-response)

## Models

### <span id="controller-api-version-response"></span> controller.ApiVersionResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| major | integer| `int64` |  | |  |  |
| minor | integer| `int64` |  | |  |  |
| specRootUrl | string| `string` |  | |  |  |
| specs | [][ControllerInfo](#controller-info)| `[]*ControllerInfo` |  | |  |  |
| supportedMajors | []integer| `[]int64` |  | |  |  |



### <span id="controller-info"></span> controller.Info


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| major | integer| `int64` |  | |  |  |
| minor | integer| `int64` |  | |  |  |
| specRootUrl | string| `string` |  | |  |  |
| supportedMajors | []integer| `[]int64` |  | |  |  |



### <span id="sigs-k8s-io-gateway-api-apis-v1-http-header"></span> sigs_k8s_io_gateway-api_apis_v1.HTTPHeader


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | Name is the name of the HTTP Header to be matched. Name matching MUST be</br>case-insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).</br></br>If multiple entries specify equivalent header names, the first entry with</br>an equivalent name MUST be considered for a match. Subsequent entries</br>with an equivalent header name MUST be ignored. Due to the</br>case-insensitivity of header names, "foo" and "Foo" are considered</br>equivalent.</br>+required |  |
| value | string| `string` |  | | Value is the value of HTTP Header to be matched.</br><gateway:experimental:description></br>Must consist of printable US-ASCII characters, optionally separated</br>by single tabs or spaces. See: https://tools.ietf.org/html/rfc7230#section-3.2</br></gateway:experimental:description></br></br>+kubebuilder:validation:MinLength=1</br>+kubebuilder:validation:MaxLength=4096</br>+required</br><gateway:experimental:validation:Pattern=`^[!-~]+([\t ]?[!-~]+)*$`> |  |



### <span id="sigs-k8s-io-gateway-api-apis-v1-local-object-reference"></span> sigs_k8s_io_gateway-api_apis_v1.LocalObjectReference


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| group | string| `string` |  | | Group is the group of the referent. For example, "gateway.networking.k8s.io".</br>When unspecified or empty string, core API group is inferred.</br>+required |  |
| kind | string| `string` |  | | Kind is kind of the referent. For example "HTTPRoute" or "Service".</br>+required |  |
| name | string| `string` |  | | Name is the name of the referent.</br>+required |  |



### <span id="v1-backend-object-reference"></span> v1.BackendObjectReference


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| group | string| `string` |  | | Group is the group of the referent. For example, "gateway.networking.k8s.io".</br>When unspecified or empty string, core API group is inferred.</br></br>+optional</br>+kubebuilder:default="" |  |
| kind | string| `string` |  | | Kind is the Kubernetes resource kind of the referent. For example</br>"Service".</br></br>Defaults to "Service" when not specified.</br></br>ExternalName services can refer to CNAME DNS records that may live</br>outside of the cluster and as such are difficult to reason about in</br>terms of conformance. They also may not be safe to forward to (see</br>CVE-2021-25740 for more information). Implementations SHOULD NOT</br>support ExternalName Services.</br></br>Support: Core (Services with a type other than ExternalName)</br></br>Support: Implementation-specific (Services with type ExternalName)</br></br>+optional</br>+kubebuilder:default=Service |  |
| name | string| `string` |  | | Name is the name of the referent.</br>+required |  |
| namespace | string| `string` |  | | Namespace is the namespace of the backend. When unspecified, the local</br>namespace is inferred.</br></br>Note that when a namespace different than the local namespace is specified,</br>a ReferenceGrant object is required in the referent namespace to allow that</br>namespace's owner to accept the reference. See the ReferenceGrant</br>documentation for details.</br></br>Support: Core</br></br>+optional |  |
| port | integer| `int64` |  | | Port specifies the destination port number to use for this resource.</br>Port is required when the referent is a Kubernetes Service. In this</br>case, the port number is the service port number, not the target port.</br>For other resources, destination port might be derived from the referent</br>resource or this field.</br></br>+optional</br>+kubebuilder:validation:Minimum=1</br>+kubebuilder:validation:Maximum=65535 |  |



### <span id="v1-forward-body-config"></span> v1.ForwardBodyConfig


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| maxSize | integer| `int64` |  | | MaxSize specifies how large in bytes the largest body that will be buffered</br>and sent to the authorization server. If the body size is larger than</br>`maxSize`, then the body sent to the authorization server must be</br>truncated to `maxSize` bytes.</br></br>Experimental note: This behavior needs to be checked against</br>various dataplanes; it may need to be changed.</br>See https://github.com/kubernetes-sigs/gateway-api/pull/4001#discussion_r2291405746</br>for more.</br></br>If 0, the body will not be sent to the authorization server.</br>+optional |  |



### <span id="v1-fraction"></span> v1.Fraction


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| denominator | integer| `int64` |  | | +optional</br>+kubebuilder:default=100</br>+kubebuilder:validation:Minimum=1 |  |
| numerator | integer| `int64` |  | | +kubebuilder:validation:Minimum=0</br>+required |  |



### <span id="v1-g-rpc-auth-config"></span> v1.GRPCAuthConfig


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allowedHeaders | []string| `[]string` |  | | AllowedRequestHeaders specifies what headers from the client request</br>will be sent to the authorization server.</br></br>If this list is empty, then all headers must be sent.</br></br>If the list has entries, only those entries must be sent.</br></br>+optional</br>+listType=set</br>+kubebuilder:validation:MaxItems=64 |  |



### <span id="v1-http-auth-config"></span> v1.HTTPAuthConfig


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allowedHeaders | []string| `[]string` |  | | AllowedRequestHeaders specifies what additional headers from the client request</br>will be sent to the authorization server.</br></br>The following headers must always be sent to the authorization server,</br>regardless of this setting:</br></br>* `Host`</br>* `Method`</br>* `Path`</br>* `Content-Length`</br>* `Authorization`</br></br>If this list is empty, then only those headers must be sent.</br></br>Note that `Content-Length` has a special behavior, in that the length</br>sent must be correct for the actual request to the external authorization</br>server - that is, it must reflect the actual number of bytes sent in the</br>body of the request to the authorization server.</br></br>So if the `forwardBody` stanza is unset, or `forwardBody.maxSize` is set</br>to `0`, then `Content-Length` must be `0`. If `forwardBody.maxSize` is set</br>to anything other than `0`, then the `Content-Length` of the authorization</br>request must be set to the actual number of bytes forwarded.</br></br>+optional</br>+listType=set</br>+kubebuilder:validation:MaxItems=64 |  |
| allowedResponseHeaders | []string| `[]string` |  | | AllowedResponseHeaders specifies what headers from the authorization response</br>will be copied into the request to the backend.</br></br>If this list is empty, then all headers from the authorization server</br>except Authority or Host must be copied.</br></br>+optional</br>+listType=set</br>+kubebuilder:validation:MaxItems=64 |  |
| path | string| `string` |  | | Path sets the prefix that paths from the client request will have added</br>when forwarded to the authorization server.</br></br>When empty or unspecified, no prefix is added.</br></br>Valid values are the same as the "value" regex for path values in the `match`</br>stanza, and the validation regex will screen out invalid paths in the same way.</br>Even with the validation, implementations MUST sanitize this input before using it</br>directly.</br></br>+optional</br>+kubebuilder:validation:MaxLength=1024</br>+kubebuilder:validation:Pattern="^(?:[-A-Za-z0-9/._~!$&'()*+,;=:@]|[%][0-9a-fA-F]{2})+$" |  |



### <span id="v1-http-c-o-r-s-filter"></span> v1.HTTPCORSFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allowCredentials | boolean| `bool` |  | | AllowCredentials indicates whether the actual cross-origin request allows</br>to include credentials.</br></br>When set to true, the gateway will include the `Access-Control-Allow-Credentials`</br>response header with value true (case-sensitive).</br></br>When set to false or omitted the gateway will omit the header</br>`Access-Control-Allow-Credentials` entirely (this is the standard CORS</br>behavior).</br></br>Support: Extended</br></br>+optional |  |
| allowHeaders | []string| `[]string` |  | | AllowHeaders indicates which HTTP request headers are supported for</br>accessing the requested resource.</br></br>Header names are not case-sensitive.</br></br>Multiple header names in the value of the `Access-Control-Allow-Headers`</br>response header are separated by a comma (",").</br></br>When the `AllowHeaders` field is configured with one or more headers, the</br>gateway must return the `Access-Control-Allow-Headers` response header</br>which value is present in the `AllowHeaders` field.</br></br>If any header name in the `Access-Control-Request-Headers` request header</br>is not included in the list of header names specified by the response</br>header `Access-Control-Allow-Headers`, it will present an error on the</br>client side.</br></br>If any header name in the `Access-Control-Allow-Headers` response header</br>does not recognize by the client, it will also occur an error on the</br>client side.</br></br>A wildcard indicates that the requests with all HTTP headers are allowed.</br>If config contains the wildcard "*" in allowHeaders and the request is</br>not credentialed, the `Access-Control-Allow-Headers` response header</br>can either use the `*` wildcard or the value of</br>Access-Control-Request-Headers from the request.</br></br>When the request is credentialed, the gateway must not specify the `*`</br>wildcard in the `Access-Control-Allow-Headers` response header. When</br>also the `AllowCredentials` field is true and `AllowHeaders` field</br>is specified with the `*` wildcard, the gateway must specify one or more</br>HTTP headers in the value of the `Access-Control-Allow-Headers` response</br>header. The value of the header `Access-Control-Allow-Headers` is same as</br>the `Access-Control-Request-Headers` header provided by the client. If</br>the header `Access-Control-Request-Headers` is not included in the</br>request, the gateway will omit the `Access-Control-Allow-Headers`</br>response header, instead of specifying the `*` wildcard.</br></br>Support: Extended</br></br>+listType=set</br>+kubebuilder:validation:MaxItems=64</br>+kubebuilder:validation:XValidation:message="AllowHeaders cannot contain '*' alongside other methods",rule="!('*' in self && self.size() > 1)"</br>+optional |  |
| allowMethods | []string| `[]string` |  | | AllowMethods indicates which HTTP methods are supported for accessing the</br>requested resource.</br></br>Valid values are any method defined by RFC9110, along with the special</br>value `*`, which represents all HTTP methods are allowed.</br></br>Method names are case-sensitive, so these values are also case-sensitive.</br>(See https://www.rfc-editor.org/rfc/rfc2616#section-5.1.1)</br></br>Multiple method names in the value of the `Access-Control-Allow-Methods`</br>response header are separated by a comma (",").</br></br>A CORS-safelisted method is a method that is `GET`, `HEAD`, or `POST`.</br>(See https://fetch.spec.whatwg.org/#cors-safelisted-method) The</br>CORS-safelisted methods are always allowed, regardless of whether they</br>are specified in the `AllowMethods` field.</br></br>When the `AllowMethods` field is configured with one or more methods, the</br>gateway must return the `Access-Control-Allow-Methods` response header</br>which value is present in the `AllowMethods` field.</br></br>If the HTTP method of the `Access-Control-Request-Method` request header</br>is not included in the list of methods specified by the response header</br>`Access-Control-Allow-Methods`, it will present an error on the client</br>side.</br></br>If config contains the wildcard "*" in allowMethods and the request is</br>not credentialed, the `Access-Control-Allow-Methods` response header</br>can either use the `*` wildcard or the value of</br>Access-Control-Request-Method from the request.</br></br>When the request is credentialed, the gateway must not specify the `*`</br>wildcard in the `Access-Control-Allow-Methods` response header. When</br>also the `AllowCredentials` field is true and `AllowMethods` field</br>specified with the `*` wildcard, the gateway must specify one HTTP method</br>in the value of the Access-Control-Allow-Methods response header. The</br>value of the header `Access-Control-Allow-Methods` is same as the</br>`Access-Control-Request-Method` header provided by the client. If the</br>header `Access-Control-Request-Method` is not included in the request,</br>the gateway will omit the `Access-Control-Allow-Methods` response header,</br>instead of specifying the `*` wildcard.</br></br>Support: Extended</br></br>+listType=set</br>+kubebuilder:validation:MaxItems=9</br>+kubebuilder:validation:XValidation:message="AllowMethods cannot contain '*' alongside other methods",rule="!('*' in self && self.size() > 1)"</br>+optional |  |
| allowOrigins | []string| `[]string` |  | | AllowOrigins indicates whether the response can be shared with requested</br>resource from the given `Origin`.</br></br>The `Origin` consists of a scheme and a host, with an optional port, and</br>takes the form `<scheme>://<host>(:<port>)`.</br></br>Valid values for scheme are: `http` and `https`.</br></br>Valid values for port are any integer between 1 and 65535 (the list of</br>available TCP/UDP ports). Note that, if not included, port `80` is</br>assumed for `http` scheme origins, and port `443` is assumed for `https`</br>origins. This may affect origin matching.</br></br>The host part of the origin may contain the wildcard character `*`. These</br>wildcard characters behave as follows:</br></br>* `*` is a greedy match to the _left_, including any number of</br>  DNS labels to the left of its position. This also means that</br>  `*` will include any number of period `.` characters to the</br>  left of its position.</br>* A wildcard by itself matches all hosts.</br></br>An origin value that includes _only_ the `*` character indicates requests</br>from all `Origin`s are allowed.</br></br>When the `AllowOrigins` field is configured with multiple origins, it</br>means the server supports clients from multiple origins. If the request</br>`Origin` matches the configured allowed origins, the gateway must return</br>the given `Origin` and sets value of the header</br>`Access-Control-Allow-Origin` same as the `Origin` header provided by the</br>client.</br></br>The status code of a successful response to a "preflight" request is</br>always an OK status (i.e., 204 or 200).</br></br>If the request `Origin` does not match the configured allowed origins,</br>the gateway returns 204/200 response but doesn't set the relevant</br>cross-origin response headers. Alternatively, the gateway responds with</br>403 status to the "preflight" request is denied, coupled with omitting</br>the CORS headers. The cross-origin request fails on the client side.</br>Therefore, the client doesn't attempt the actual cross-origin request.</br></br>Conversely, if the request `Origin` matches one of the configured</br>allowed origins, the gateway sets the response header</br>`Access-Control-Allow-Origin` to the same value as the `Origin`</br>header provided by the client.</br></br>When config has the wildcard ("*") in allowOrigins, and the request</br>is not credentialed (e.g., it is a preflight request), the</br>`Access-Control-Allow-Origin` response header either contains the</br>wildcard as well or the Origin from the request.</br></br>When the request is credentialed, the gateway must not specify the `*`</br>wildcard in the `Access-Control-Allow-Origin` response header. When</br>also the `AllowCredentials` field is true and `AllowOrigins` field</br>specified with the `*` wildcard, the gateway must return a single origin</br>in the value of the `Access-Control-Allow-Origin` response header,</br>instead of specifying the `*` wildcard. The value of the header</br>`Access-Control-Allow-Origin` is same as the `Origin` header provided by</br>the client.</br></br>Support: Extended</br>+listType=set</br>+kubebuilder:validation:MaxItems=64</br>+kubebuilder:validation:XValidation:message="AllowOrigins cannot contain '*' alongside other origins",rule="!('*' in self && self.size() > 1)"</br>+optional |  |
| exposeHeaders | []string| `[]string` |  | | ExposeHeaders indicates which HTTP response headers can be exposed</br>to client-side scripts in response to a cross-origin request.</br></br>A CORS-safelisted response header is an HTTP header in a CORS response</br>that it is considered safe to expose to the client scripts.</br>The CORS-safelisted response headers include the following headers:</br>`Cache-Control`</br>`Content-Language`</br>`Content-Length`</br>`Content-Type`</br>`Expires`</br>`Last-Modified`</br>`Pragma`</br>(See https://fetch.spec.whatwg.org/#cors-safelisted-response-header-name)</br>The CORS-safelisted response headers are exposed to client by default.</br></br>When an HTTP header name is specified using the `ExposeHeaders` field,</br>this additional header will be exposed as part of the response to the</br>client.</br></br>Header names are not case-sensitive.</br></br>Multiple header names in the value of the `Access-Control-Expose-Headers`</br>response header are separated by a comma (",").</br></br>A wildcard indicates that the responses with all HTTP headers are exposed</br>to clients. The `Access-Control-Expose-Headers` response header can only</br>use `*` wildcard as value when the request is not credentialed.</br></br>When the `exposeHeaders` config field contains the "*" wildcard and</br>the request is credentialed, the gateway cannot use the `*` wildcard in</br>the `Access-Control-Expose-Headers` response header.</br></br>Support: Extended</br></br>+optional</br>+listType=set</br>+kubebuilder:validation:MaxItems=64 |  |
| maxAge | integer| `int64` |  | | MaxAge indicates the duration (in seconds) for the client to cache the</br>results of a "preflight" request.</br></br>The information provided by the `Access-Control-Allow-Methods` and</br>`Access-Control-Allow-Headers` response headers can be cached by the</br>client until the time specified by `Access-Control-Max-Age` elapses.</br></br>The default value of `Access-Control-Max-Age` response header is 5</br>(seconds).</br></br>When the `MaxAge` field is unspecified, the gateway sets the response</br>header "Access-Control-Max-Age: 5" by default.</br></br>+optional</br>+kubebuilder:default=5</br>+kubebuilder:validation:Minimum=1 |  |



### <span id="v1-http-external-auth-filter"></span> v1.HTTPExternalAuthFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| backendRef | [V1HTTPExternalAuthFilter](#v1-http-external-auth-filter)| `V1HTTPExternalAuthFilter` |  | | BackendRef is a reference to a backend to send authorization</br>requests to.</br></br>The backend must speak the selected protocol (GRPC or HTTP) on the</br>referenced port.</br></br>If the backend service requires TLS, use BackendTLSPolicy to tell the</br>implementation to supply the TLS details to be used to connect to that</br>backend.</br></br>+required |  |
| forwardBody | [V1HTTPExternalAuthFilter](#v1-http-external-auth-filter)| `V1HTTPExternalAuthFilter` |  | | ForwardBody controls if requests to the authorization server should include</br>the body of the client request; and if so, how big that body is allowed</br>to be.</br></br>It is expected that implementations will buffer the request body up to</br>`forwardBody.maxSize` bytes. Bodies over that size must be rejected with a</br>4xx series error (413 or 403 are common examples), and fail processing</br>of the filter.</br></br>If unset, or `forwardBody.maxSize` is set to `0`, then the body will not</br>be forwarded.</br></br>Feature Name: HTTPRouteExternalAuthForwardBody</br></br>+optional |  |
| grpc | [V1HTTPExternalAuthFilter](#v1-http-external-auth-filter)| `V1HTTPExternalAuthFilter` |  | | GRPCAuthConfig contains configuration for communication with ext_authz</br>protocol-speaking backends.</br></br>If unset, implementations must assume the default behavior for each</br>included field is intended.</br></br>+optional |  |
| http | [V1HTTPExternalAuthFilter](#v1-http-external-auth-filter)| `V1HTTPExternalAuthFilter` |  | | HTTPAuthConfig contains configuration for communication with HTTP-speaking</br>backends.</br></br>If unset, implementations must assume the default behavior for each</br>included field is intended.</br></br>+optional |  |
| protocol | [V1HTTPExternalAuthFilter](#v1-http-external-auth-filter)| `V1HTTPExternalAuthFilter` |  | | ExternalAuthProtocol describes which protocol to use when communicating with an</br>ext_authz authorization server.</br></br>When this is set to GRPC, each backend must use the Envoy ext_authz protocol</br>on the port specified in `backendRefs`. Requests and responses are defined</br>in the protobufs explained at:</br>https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/auth/v3/external_auth.proto</br></br>When this is set to HTTP, each backend must respond with a `200` status</br>code in on a successful authorization. Any other code is considered</br>an authorization failure.</br></br>Feature Names:</br>GRPC Support - HTTPRouteExternalAuthGRPC</br>HTTP Support - HTTPRouteExternalAuthHTTP</br></br>+unionDiscriminator</br>+required</br>+kubebuilder:validation:Enum=HTTP;GRPC |  |



### <span id="v1-http-header-filter"></span> v1.HTTPHeaderFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| add | [][SigsK8sIoGatewayAPIApisV1HTTPHeader](#sigs-k8s-io-gateway-api-apis-v1-http-header)| `[]*SigsK8sIoGatewayAPIApisV1HTTPHeader` |  | | Add adds the given header(s) (name, value) to the request</br>before the action. It appends to any existing values associated</br>with the header name.</br></br>Input:</br>  GET /foo HTTP/1.1</br>  my-header: foo</br></br>Config:</br>  add:</br>  - name: "my-header"</br>    value: "bar,baz"</br></br>Output:</br>  GET /foo HTTP/1.1</br>  my-header: foo,bar,baz</br></br>+optional</br>+listType=map</br>+listMapKey=name</br>+kubebuilder:validation:MaxItems=16 |  |
| remove | []string| `[]string` |  | | Remove the given header(s) from the HTTP request before the action. The</br>value of Remove is a list of HTTP header names. Note that the header</br>names are case-insensitive (see</br>https://datatracker.ietf.org/doc/html/rfc2616#section-4.2).</br></br>Input:</br>  GET /foo HTTP/1.1</br>  my-header1: foo</br>  my-header2: bar</br>  my-header3: baz</br></br>Config:</br>  remove: ["my-header1", "my-header3"]</br></br>Output:</br>  GET /foo HTTP/1.1</br>  my-header2: bar</br></br>+optional</br>+listType=set</br>+kubebuilder:validation:MaxItems=16 |  |
| set | [][SigsK8sIoGatewayAPIApisV1HTTPHeader](#sigs-k8s-io-gateway-api-apis-v1-http-header)| `[]*SigsK8sIoGatewayAPIApisV1HTTPHeader` |  | | Set overwrites the request with the given header (name, value)</br>before the action.</br></br>Input:</br>  GET /foo HTTP/1.1</br>  my-header: foo</br></br>Config:</br>  set:</br>  - name: "my-header"</br>    value: "bar"</br></br>Output:</br>  GET /foo HTTP/1.1</br>  my-header: bar</br></br>+optional</br>+listType=map</br>+listMapKey=name</br>+kubebuilder:validation:MaxItems=16 |  |



### <span id="v1-http-path-modifier"></span> v1.HTTPPathModifier


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| replaceFullPath | string| `string` |  | | ReplaceFullPath specifies the value with which to replace the full path</br>of a request during a rewrite or redirect.</br></br>+kubebuilder:validation:MaxLength=1024</br>+optional |  |
| replacePrefixMatch | string| `string` |  | | ReplacePrefixMatch specifies the value with which to replace the prefix</br>match of a request during a rewrite or redirect. For example, a request</br>to "/foo/bar" with a prefix match of "/foo" and a ReplacePrefixMatch</br>of "/xyz" would be modified to "/xyz/bar".</br></br>Note that this matches the behavior of the PathPrefix match type. This</br>matches full path elements. A path element refers to the list of labels</br>in the path split by the `/` separator. When specified, a trailing `/` is</br>ignored. For example, the paths `/abc`, `/abc/`, and `/abc/def` would all</br>match the prefix `/abc`, but the path `/abcd` would not.</br></br>ReplacePrefixMatch is only compatible with a `PathPrefix` HTTPRouteMatch.</br>Using any other HTTPRouteMatch type on the same HTTPRouteRule will result in</br>the implementation setting the Accepted Condition for the Route to `status: False`.</br></br>Request Path | Prefix Match | Replace Prefix | Modified Path</br>-------------|--------------|----------------|----------</br>/foo/bar     | /foo         | /xyz           | /xyz/bar</br>/foo/bar     | /foo         | /xyz/          | /xyz/bar</br>/foo/bar     | /foo/        | /xyz           | /xyz/bar</br>/foo/bar     | /foo/        | /xyz/          | /xyz/bar</br>/foo         | /foo         | /xyz           | /xyz</br>/foo/        | /foo         | /xyz           | /xyz/</br>/foo/bar     | /foo         | <empty string> | /bar</br>/foo/        | /foo         | <empty string> | /</br>/foo         | /foo         | <empty string> | /</br>/foo/        | /foo         | /              | /</br>/foo         | /foo         | /              | /</br></br>+kubebuilder:validation:MaxLength=1024</br>+optional |  |
| type | [V1HTTPPathModifier](#v1-http-path-modifier)| `V1HTTPPathModifier` |  | | Type defines the type of path modifier. Additional types may be</br>added in a future release of the API.</br></br>Note that values may be added to this enum, implementations</br>must ensure that unknown values will not cause a crash.</br></br>Unknown values here must result in the implementation setting the</br>Accepted Condition for the Route to `status: False`, with a</br>Reason of `UnsupportedValue`.</br></br>+kubebuilder:validation:Enum=ReplaceFullPath;ReplacePrefixMatch</br>+required |  |



### <span id="v1-http-path-modifier-type"></span> v1.HTTPPathModifierType


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| v1.HTTPPathModifierType | string| string | |  |  |



### <span id="v1-http-request-mirror-filter"></span> v1.HTTPRequestMirrorFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| backendRef | [V1HTTPRequestMirrorFilter](#v1-http-request-mirror-filter)| `V1HTTPRequestMirrorFilter` |  | | BackendRef references a resource where mirrored requests are sent.</br></br>Mirrored requests must be sent only to a single destination endpoint</br>within this BackendRef, irrespective of how many endpoints are present</br>within this BackendRef.</br></br>If the referent cannot be found, this BackendRef is invalid and must be</br>dropped from the Gateway. The controller must ensure the "ResolvedRefs"</br>condition on the Route status is set to `status: False` and not configure</br>this backend in the underlying implementation.</br></br>If there is a cross-namespace reference to an *existing* object</br>that is not allowed by a ReferenceGrant, the controller must ensure the</br>"ResolvedRefs"  condition on the Route is set to `status: False`,</br>with the "RefNotPermitted" reason and not configure this backend in the</br>underlying implementation.</br></br>In either error case, the Message of the `ResolvedRefs` Condition</br>should be used to provide more detail about the problem.</br></br>Support: Extended for Kubernetes Service</br></br>Support: Implementation-specific for any other resource</br>+required |  |
| fraction | [V1HTTPRequestMirrorFilter](#v1-http-request-mirror-filter)| `V1HTTPRequestMirrorFilter` |  | | Fraction represents the fraction of requests that should be</br>mirrored to BackendRef.</br></br>Only one of Fraction or Percent may be specified. If neither field</br>is specified, 100% of requests will be mirrored.</br></br>+optional |  |
| percent | integer| `int64` |  | | Percent represents the percentage of requests that should be</br>mirrored to BackendRef. Its minimum value is 0 (indicating 0% of</br>requests) and its maximum value is 100 (indicating 100% of requests).</br></br>Only one of Fraction or Percent may be specified. If neither field</br>is specified, 100% of requests will be mirrored.</br></br>+optional</br>+kubebuilder:validation:Minimum=0</br>+kubebuilder:validation:Maximum=100 |  |



### <span id="v1-http-request-redirect-filter"></span> v1.HTTPRequestRedirectFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| hostname | string| `string` |  | | Hostname is the hostname to be used in the value of the `Location`</br>header in the response.</br>When empty, the hostname in the `Host` header of the request is used.</br></br>Support: Core</br></br>+optional |  |
| path | [V1HTTPRequestRedirectFilter](#v1-http-request-redirect-filter)| `V1HTTPRequestRedirectFilter` |  | | Path defines parameters used to modify the path of the incoming request.</br>The modified path is then used to construct the `Location` header. When</br>empty, the request path is used as-is.</br></br>Support: Extended</br></br>+optional |  |
| port | integer| `int64` |  | | Port is the port to be used in the value of the `Location`</br>header in the response.</br></br>If no port is specified, the redirect port MUST be derived using the</br>following rules:</br></br>* If redirect scheme is not-empty, the redirect port MUST be the well-known</br>  port associated with the redirect scheme. Specifically "http" to port 80</br>  and "https" to port 443. If the redirect scheme does not have a</br>  well-known port, the listener port of the Gateway SHOULD be used.</br>* If redirect scheme is empty, the redirect port MUST be the Gateway</br>  Listener port.</br></br>Implementations SHOULD NOT add the port number in the 'Location'</br>header in the following cases:</br></br>* A Location header that will use HTTP (whether that is determined via</br>  the Listener protocol or the Scheme field) _and_ use port 80.</br>* A Location header that will use HTTPS (whether that is determined via</br>  the Listener protocol or the Scheme field) _and_ use port 443.</br></br>Support: Extended</br></br>+optional</br></br>+kubebuilder:validation:Minimum=1</br>+kubebuilder:validation:Maximum=65535 |  |
| scheme | string| `string` |  | | Scheme is the scheme to be used in the value of the `Location` header in</br>the response. When empty, the scheme of the request is used.</br></br>Scheme redirects can affect the port of the redirect, for more information,</br>refer to the documentation for the port field of this filter.</br></br>Note that values may be added to this enum, implementations</br>must ensure that unknown values will not cause a crash.</br></br>Unknown values here must result in the implementation setting the</br>Accepted Condition for the Route to `status: False`, with a</br>Reason of `UnsupportedValue`.</br></br>Support: Extended</br></br>+optional</br>+kubebuilder:validation:Enum=http;https |  |
| statusCode | integer| `int64` |  | | StatusCode is the HTTP status code to be used in response.</br></br>Note that values may be added to this enum, implementations</br>must ensure that unknown values will not cause a crash.</br></br>Unknown values here must result in the implementation setting the</br>Accepted Condition for the Route to `status: False`, with a</br>Reason of `UnsupportedValue`.</br></br>Support: Core</br></br>+optional</br>+kubebuilder:default=302</br>+kubebuilder:validation:Enum=301;302;303;307;308 |  |



### <span id="v1-http-route-external-auth-protocol"></span> v1.HTTPRouteExternalAuthProtocol


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| v1.HTTPRouteExternalAuthProtocol | string| string | |  |  |



### <span id="v1-http-route-filter"></span> v1.HTTPRouteFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cors | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | CORS defines a schema for a filter that responds to the</br>cross-origin request based on HTTP response header.</br></br>Support: Extended</br></br>+optional |  |
| extensionRef | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | ExtensionRef is an optional, implementation-specific extension to the</br>"filter" behavior.  For example, resource "myroutefilter" in group</br>"networking.example.net"). ExtensionRef MUST NOT be used for core and</br>extended filters.</br></br>This filter can be used multiple times within the same rule.</br></br>Support: Implementation-specific</br></br>+optional |  |
| externalAuth | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | ExternalAuth configures settings related to sending request details</br>to an external auth service. The external service MUST authenticate</br>the request, and MAY authorize the request as well.</br></br>If there is any problem communicating with the external service,</br>this filter MUST fail closed.</br></br>Support: Extended</br></br>+optional</br><gateway:experimental> |  |
| requestHeaderModifier | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | RequestHeaderModifier defines a schema for a filter that modifies request</br>headers.</br></br>Support: Core</br></br>+optional |  |
| requestMirror | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | RequestMirror defines a schema for a filter that mirrors requests.</br>Requests are sent to the specified destination, but responses from</br>that destination are ignored.</br></br>This filter can be used multiple times within the same rule. Note that</br>not all implementations will be able to support mirroring to multiple</br>backends.</br></br>Support: Extended</br></br>+optional</br></br>+kubebuilder:validation:XValidation:message="Only one of percent or fraction may be specified in HTTPRequestMirrorFilter",rule="!(has(self.percent) && has(self.fraction))" |  |
| requestRedirect | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | RequestRedirect defines a schema for a filter that responds to the</br>request with an HTTP redirection.</br></br>Support: Core</br></br>+optional |  |
| responseHeaderModifier | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | ResponseHeaderModifier defines a schema for a filter that modifies response</br>headers.</br></br>Support: Extended</br></br>+optional |  |
| type | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | Type identifies the type of filter to apply. As with other API fields,</br>types are classified into three conformance levels:</br></br>- Core: Filter types and their corresponding configuration defined by</br>  "Support: Core" in this package, e.g. "RequestHeaderModifier". All</br>  implementations must support core filters.</br></br>- Extended: Filter types and their corresponding configuration defined by</br>  "Support: Extended" in this package, e.g. "RequestMirror". Implementers</br>  are encouraged to support extended filters.</br></br>- Implementation-specific: Filters that are defined and supported by</br>  specific vendors.</br>  In the future, filters showing convergence in behavior across multiple</br>  implementations will be considered for inclusion in extended or core</br>  conformance levels. Filter-specific configuration for such filters</br>  is specified using the ExtensionRef field. `Type` should be set to</br>  "ExtensionRef" for custom filters.</br></br>Implementers are encouraged to define custom implementation types to</br>extend the core API with implementation-specific behavior.</br></br>If a reference to a custom filter type cannot be resolved, the filter</br>MUST NOT be skipped. Instead, requests that would have been processed by</br>that filter MUST receive a HTTP error response.</br></br>Note that values may be added to this enum, implementations</br>must ensure that unknown values will not cause a crash.</br></br>Unknown values here must result in the implementation setting the</br>Accepted Condition for the Route to `status: False`, with a</br>Reason of `UnsupportedValue`.</br></br>+unionDiscriminator</br>+kubebuilder:validation:Enum=RequestHeaderModifier;ResponseHeaderModifier;RequestMirror;RequestRedirect;URLRewrite;ExtensionRef;CORS</br><gateway:experimental:validation:Enum=RequestHeaderModifier;ResponseHeaderModifier;RequestMirror;RequestRedirect;URLRewrite;ExtensionRef;CORS;ExternalAuth></br>+required |  |
| urlRewrite | [V1HTTPRouteFilter](#v1-http-route-filter)| `V1HTTPRouteFilter` |  | | URLRewrite defines a schema for a filter that modifies a request during forwarding.</br></br>Support: Extended</br></br>+optional |  |



### <span id="v1-http-route-filter-type"></span> v1.HTTPRouteFilterType


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| v1.HTTPRouteFilterType | string| string | |  |  |



### <span id="v1-http-url-rewrite-filter"></span> v1.HTTPURLRewriteFilter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| hostname | string| `string` |  | | Hostname is the value to be used to replace the Host header value during</br>forwarding.</br></br>Support: Extended</br></br>+optional |  |
| path | [V1HTTPURLRewriteFilter](#v1-http-url-rewrite-filter)| `V1HTTPURLRewriteFilter` |  | | Path defines a path rewrite.</br></br>Support: Extended</br></br>+optional |  |



### <span id="v2-annotation-resource"></span> v2.AnnotationResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| annotationValue | string| `string` |  | |  |  |
| namespace | string| `string` |  | |  |  |
| resourceName | string| `string` |  | |  |  |



### <span id="v2-app-version-data"></span> v2.AppVersionData


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| appName | string| `string` |  | |  |  |
| appVersion | string| `string` |  | |  |  |
| deployTime | string| `string` |  | |  |  |



### <span id="v2-config-map"></span> v2.ConfigMap


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | map of string| `map[string]string` |  | |  |  |
| metadata | [V2Metadata](#v2-metadata)| `V2Metadata` |  | |  |  |



### <span id="v2-container-env"></span> v2.ContainerEnv


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  |  |
| value | string| `string` |  | |  |  |
| valueFrom | [V2ValueFrom](#v2-value-from)| `V2ValueFrom` |  | |  |  |



### <span id="v2-container-port"></span> v2.ContainerPort


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| containerPort | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| protocol | string| `string` |  | |  |  |



### <span id="v2-container-resources"></span> v2.ContainerResources


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| limits | [V2CPUMemoryResource](#v2-cpu-memory-resource)| `V2CPUMemoryResource` |  | |  |  |
| requests | [V2CPUMemoryResource](#v2-cpu-memory-resource)| `V2CPUMemoryResource` |  | |  |  |



### <span id="v2-container-state"></span> v2.ContainerState


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| running | [V2ContainerStateRunning](#v2-container-state-running)| `V2ContainerStateRunning` |  | |  |  |
| terminated | [V2ContainerStateTerminated](#v2-container-state-terminated)| `V2ContainerStateTerminated` |  | |  |  |
| waiting | [V2ContainerStateWaiting](#v2-container-state-waiting)| `V2ContainerStateWaiting` |  | |  |  |



### <span id="v2-container-state-running"></span> v2.ContainerStateRunning


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| startedAt | string| `string` |  | |  |  |



### <span id="v2-container-state-terminated"></span> v2.ContainerStateTerminated


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| containerID | string| `string` |  | |  |  |
| exitCode | integer| `int64` |  | |  |  |
| finishedAt | string| `string` |  | |  |  |
| reason | string| `string` |  | |  |  |
| startedAt | string| `string` |  | |  |  |



### <span id="v2-container-state-waiting"></span> v2.ContainerStateWaiting


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| message | string| `string` |  | |  |  |
| reason | string| `string` |  | |  |  |



### <span id="v2-container-status"></span> v2.ContainerStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| containerID | string| `string` |  | |  |  |
| image | string| `string` |  | |  |  |
| imageID | string| `string` |  | |  |  |
| lastState | [V2ContainerState](#v2-container-state)| `V2ContainerState` |  | |  |  |
| name | string| `string` |  | |  |  |
| ready | boolean| `bool` |  | |  |  |
| restartCount | integer| `int64` |  | |  |  |
| state | [V2ContainerState](#v2-container-state)| `V2ContainerState` |  | |  |  |



### <span id="v2-container-volume-mount"></span> v2.ContainerVolumeMount


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| mountPath | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| readOnly | boolean| `bool` |  | |  |  |



### <span id="v2-cpu-memory-resource"></span> v2.CpuMemoryResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cpu | string| `string` |  | |  |  |
| memory | string| `string` |  | |  |  |



### <span id="v2-deployment"></span> v2.Deployment


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [V2Metadata](#v2-metadata)| `V2Metadata` |  | |  |  |
| spec | [V2DeploymentSpec](#v2-deployment-spec)| `V2DeploymentSpec` |  | |  |  |
| status | [V2DeploymentStatus](#v2-deployment-status)| `V2DeploymentStatus` |  | |  |  |



### <span id="v2-deployment-condition"></span> v2.DeploymentCondition


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| lastTransitionTime | string| `string` |  | |  |  |
| lastUpdateTime | string| `string` |  | |  |  |
| message | string| `string` |  | |  |  |
| reason | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="v2-deployment-family-version"></span> v2.DeploymentFamilyVersion


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| app_name | string| `string` |  | |  |  |
| app_version | string| `string` |  | |  |  |
| bluegreen_version | string| `string` |  | |  |  |
| family_name | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| state | string| `string` |  | |  |  |
| version | string| `string` |  | |  |  |



### <span id="v2-deployment-response"></span> v2.DeploymentResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| deployments | [][map[string]V2DeploymentRollout](#map-string-v2-deployment-rollout)| `[]map[string]V2DeploymentRollout` |  | |  |  |
| pod_status_websocket | string| `string` |  | |  |  |



### <span id="v2-deployment-rollout"></span> v2.DeploymentRollout


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| active | string| `string` |  | |  |  |
| kind | string| `string` |  | |  |  |
| rolling | string| `string` |  | |  |  |



### <span id="v2-deployment-spec"></span> v2.DeploymentSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| replicas | integer| `int64` |  | |  |  |
| revisionHistoryLimit | integer| `int64` |  | |  |  |
| strategy | [V2DeploymentStrategy](#v2-deployment-strategy)| `V2DeploymentStrategy` |  | |  |  |
| template | [V2PodTemplateSpec](#v2-pod-template-spec)| `V2PodTemplateSpec` |  | |  |  |



### <span id="v2-deployment-status"></span> v2.DeploymentStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| availableReplicas | integer| `int64` |  | |  |  |
| conditions | [][V2DeploymentCondition](#v2-deployment-condition)| `[]*V2DeploymentCondition` |  | |  |  |
| observedGeneration | integer| `int64` |  | |  |  |
| readyReplicas | integer| `int64` |  | |  |  |
| replicas | integer| `int64` |  | |  |  |
| updatedReplicas | integer| `int64` |  | |  |  |



### <span id="v2-deployment-strategy"></span> v2.DeploymentStrategy


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| type | string| `string` |  | |  |  |



### <span id="v2-error-response"></span> v2.ErrorResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="v2-field-ref"></span> v2.FieldRef


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| apiVersion | string| `string` |  | |  |  |
| fieldPath | string| `string` |  | |  |  |



### <span id="v2-metadata"></span> v2.Metadata


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| annotations | map of string| `map[string]string` |  | |  |  |
| generation | integer| `int64` |  | |  |  |
| kind | string| `string` |  | |  |  |
| labels | map of string| `map[string]string` |  | |  |  |
| name | string| `string` |  | |  |  |
| namespace | string| `string` |  | |  |  |
| resourceVersion | string| `string` |  | |  |  |
| uid | string| `string` |  | |  |  |



### <span id="v2-namespace"></span> v2.Namespace


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [V2Metadata](#v2-metadata)| `V2Metadata` |  | |  |  |



### <span id="v2-pod"></span> v2.Pod


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [V2Metadata](#v2-metadata)| `V2Metadata` |  | |  |  |
| spec | [V2PodSpec](#v2-pod-spec)| `V2PodSpec` |  | |  |  |
| status | [V2PodStatus](#v2-pod-status)| `V2PodStatus` |  | |  |  |



### <span id="v2-pod-spec"></span> v2.PodSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| containers | [][V2SpecContainer](#v2-spec-container)| `[]*V2SpecContainer` |  | |  |  |
| dnsPolicy | string| `string` |  | |  |  |
| nodeName | string| `string` |  | |  |  |
| restartPolicy | string| `string` |  | |  |  |
| terminationGracePeriodSeconds | integer| `int64` |  | |  |  |
| volumes | [][V2SpecVolume](#v2-spec-volume)| `[]*V2SpecVolume` |  | |  |  |



### <span id="v2-pod-status"></span> v2.PodStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| conditions | [][V2StatusCondition](#v2-status-condition)| `[]*V2StatusCondition` |  | |  |  |
| containerStatuses | [][V2ContainerStatus](#v2-container-status)| `[]*V2ContainerStatus` |  | |  |  |
| hostIP | string| `string` |  | |  |  |
| phase | string| `string` |  | |  |  |
| podIP | string| `string` |  | |  |  |
| startTime | string| `string` |  | |  |  |



### <span id="v2-pod-template-spec"></span> v2.PodTemplateSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [V2TemplateMetadata](#v2-template-metadata)| `V2TemplateMetadata` |  | |  |  |
| spec | [V2PodSpec](#v2-pod-spec)| `V2PodSpec` |  | |  |  |



### <span id="v2-rollout-deployment-body"></span> v2.RolloutDeploymentBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| deployment_names | []string| `[]string` |  | |  |  |
| parallel | boolean| `bool` |  | |  |  |



### <span id="v2-route"></span> v2.Route


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [V2Metadata](#v2-metadata)| `V2Metadata` |  | |  |  |
| spec | [V2RouteSpec](#v2-route-spec)| `V2RouteSpec` |  | |  |  |



### <span id="v2-route-port"></span> v2.RoutePort


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| targetPort | integer| `int64` |  | |  |  |



### <span id="v2-route-spec"></span> v2.RouteSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| filters | [][V1HTTPRouteFilter](#v1-http-route-filter)| `[]*V1HTTPRouteFilter` |  | |  |  |
| host | string| `string` |  | |  |  |
| ingressClassName | string| `string` |  | |  |  |
| path | string| `string` |  | |  |  |
| pathType | string| `string` |  | |  |  |
| port | [V2RoutePort](#v2-route-port)| `V2RoutePort` |  | |  |  |
| streamIdleTimeout | string| `string` |  | |  |  |
| to | [V2Target](#v2-target)| `V2Target` |  | |  |  |



### <span id="v2-secret-key-ref"></span> v2.SecretKeyRef


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| key | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="v2-service"></span> v2.Service


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [V2Metadata](#v2-metadata)| `V2Metadata` |  | |  |  |
| spec | [V2ServiceSpec](#v2-service-spec)| `V2ServiceSpec` |  | |  |  |



### <span id="v2-service-port"></span> v2.ServicePort


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  |  |
| nodePort | integer| `int64` |  | |  |  |
| port | integer| `int64` |  | |  |  |
| protocol | string| `string` |  | |  |  |
| targetPort | integer| `int64` |  | |  |  |



### <span id="v2-service-spec"></span> v2.ServiceSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| clusterIP | string| `string` |  | |  |  |
| ports | [][V2ServicePort](#v2-service-port)| `[]*V2ServicePort` |  | |  |  |
| selector | map of string| `map[string]string` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="v2-spec-container"></span> v2.SpecContainer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| args | []string| `[]string` |  | |  |  |
| env | [][V2ContainerEnv](#v2-container-env)| `[]*V2ContainerEnv` |  | |  |  |
| image | string| `string` |  | |  |  |
| imagePullPolicy | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| ports | [][V2ContainerPort](#v2-container-port)| `[]*V2ContainerPort` |  | |  |  |
| resources | [V2ContainerResources](#v2-container-resources)| `V2ContainerResources` |  | |  |  |
| volumeMounts | [][V2ContainerVolumeMount](#v2-container-volume-mount)| `[]*V2ContainerVolumeMount` |  | |  |  |



### <span id="v2-spec-volume"></span> v2.SpecVolume


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  |  |
| secret | [V2VolumesSecret](#v2-volumes-secret)| `V2VolumesSecret` |  | |  |  |



### <span id="v2-status-condition"></span> v2.StatusCondition


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| lastProbeTime | string| `string` |  | |  |  |
| lastTransitionTime | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="v2-target"></span> v2.Target


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | |  |  |



### <span id="v2-template-metadata"></span> v2.TemplateMetadata


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| labels | map of string| `map[string]string` |  | |  |  |



### <span id="v2-value-from"></span> v2.ValueFrom


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| fieldRef | [V2FieldRef](#v2-field-ref)| `V2FieldRef` |  | |  |  |
| secretKeyRef | [V2SecretKeyRef](#v2-secret-key-ref)| `V2SecretKeyRef` |  | |  |  |



### <span id="v2-volumes-secret"></span> v2.VolumesSecret


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| defaultMode | integer| `int64` |  | |  |  |
| secretName | string| `string` |  | |  |  |


