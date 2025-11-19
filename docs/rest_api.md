


# Paas-Mediation API
Paas-Mediation Service
  

## Informations

### Version

2.0

### Contact

  

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

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

Get Gateway API GRPC Routes in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

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
| [404](#v2-get-gateway-grpcroutes-404) | Not Found | Not Found |  | [schema](#v2-get-gateway-grpcroutes-404-schema) |
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

##### <span id="v2-get-gateway-grpcroutes-404"></span> 404 - Not Found
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

Get Gateway API HTTP Routes in namespace

#### Consumes
  * application/json

#### Produces
  * application/json

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
| [404](#v2-get-gateway-httproutes-404) | Not Found | Not Found |  | [schema](#v2-get-gateway-httproutes-404-schema) |
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

##### <span id="v2-get-gateway-httproutes-404"></span> 404 - Not Found
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
| host | string| `string` |  | |  |  |
| ingressClassName | string| `string` |  | |  |  |
| path | string| `string` |  | |  |  |
| pathType | string| `string` |  | |  |  |
| port | [V2RoutePort](#v2-route-port)| `V2RoutePort` |  | |  |  |
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


