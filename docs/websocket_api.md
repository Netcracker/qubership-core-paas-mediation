This section describes the following PaaS Mediation Websocket APIs:

* [Service watch API](#service-watch-api)
* [ConfigMap watch API](#configmap-watch-api)
* [Route watch API](#route-watch-api)
* [Namespace watch API](#namespace-watch-api)
* [Pods restart status watch API](#pods-restart-status-watch-api)

If you want to get notifications about creating, deleting, modifying resources then you should use the below API. All watch APIs are websocket API, so for establishing connection you should use websocket client.

API endpoints shown below are served by internal-gateway-service

# Service watch API

Allow to get events when some actions were performed under a service. To watch for services only with particular annotations and/or labels the filter request parameters can be used.
Annotation value can be equal to * which means any value.
* **API:**  `/watchapi/v2/paas-mediation/namespaces/{namespace}/services`
* **Request params (Optional):** `annotations=<annt_name1>:<annt_value1>;<annt_name2>:<annt_value2>&labels=<label_name1>:<label_value1>;<label_name2>:<label_value2>`
* **Headers:**      
  `Authorization: Bearer <jwt>`
* **Authorization:**
  M2M
  Response body:
```json
{
  "type": "<ADDED||MODIFIED||DELETED>",
  "object": {
        "metadata": {
          "kind": "Service",
          "name": "<service_name>",
          "namespace": "<namespace>",
          "annotations": { 
              "<annotation_one>": "<annotation_one_value>"
          },
          "labels": { 
              "<label_one>": "<label_one_value>"
          }
        },
        "spec": {
            "ports": [
              {
                "name": "<port_name>",
                "protocol": "<protocol>",
                "port": <port>,
                "targetPort": <targetPort>,
                "nodePort": <nodePort>
              }
            ],
            "selector": {
                "<selector_one>": "<selector_one_value>"
            },
            "clusterIP": "<cluster_ip>",
            "type": "<type>"
        }
  }
}
```
* **Error Response:**   
  *Response body:*
```json
{
  "type": "ERROR",
  "object": {
      error message in json format
  }
}
```

# ConfigMap watch API

Allow to get events when some actions were performed under a configMap. To watch for configMaps only with particular annotations and/or labels the filter request parameters can be used.
Annotation value can be equal to * which means any value.
* **API:**  `/watchapi/v2/paas-mediation/namespaces/{namespace}/configmaps`
* **Request params (Optional):** `annotations=<annt_name1>:<annt_value1>;<annt_name2>:<annt_value2>&labels=<label_name1>:<label_value1>;<label_name2>:<label_value2>`
* **Headers:**      
  `Authorization: Bearer <jwt>`
* **Authorization:**
  M2M
  Response body:
```json
{
  "type": "<ADDED||MODIFIED||DELETED>",
  "object": {
      "metadata": {
        "kind": "ConfigMap",
        "name": "<configMap_name>",
        "namespace": "<namespace>",
        "annotations": { 
            "<annotation_one>": "<annotation_one_value>"
        },
        "labels": { 
            "<label_one>": "<label_one_value>"
        }
      },
      "data": {
          "field_one": "<value_one>"      
      }
  }
}
```
* **Error Response:**   
  *Response body:*
```json
{
  "type": "ERROR",
  "object": {
      error message in json format
  }
}
```

# Route watch API

Allow to get events when some actions were performed under a route. To watch for routes only with particular annotations and/or labels the filter request parameters can be used.
Annotation value can be equal to * which means any value.
* **API:**  `/watchapi/v2/paas-mediation/namespaces/{namespace}/routes`
* **Request params (Optional):** `annotations=<annt_name1>:<annt_value1>;<annt_name2>:<annt_value2>&labels=<label_name1>:<label_value1>;<label_name2>:<label_value2>`
* **Headers:**      
  `Authorization: Bearer <jwt>`
* **Authorization:**
  M2M
  Response body:
```json
{
  "type": "<ADDED||MODIFIED||DELETED>",
  "object": {
      "metadata": {
        "kind": "Route",
        "name": "<route_name>",
        "namespace": "<namespace_name>",
        "annotations": { 
            "<annotation_one>": "<annotation_one_value>"
        },
        "labels": { 
            "<label_one>": "<label_one_value>"
        }
      },
      "spec": {
          "host": "<host_name>",
          "path": "<path>",
          "to": {
              "name": "<service_name>"
          }
      }
  }
}
```
* **Error Response:**   
  *Response body:*
```json
{
  "type": "ERROR",
  "object": {
      error message in json format
  }
}
```

# Namespace watch API

Allow to get events when some actions were performed under a namespace.
* **API:**  `/watchapi/v2/paas-mediation/namespaces/{namespace}/namespaces`
* **Headers:**      
  `Authorization: Bearer <jwt>`
* **Authorization:**
  M2M
  Response body:
```json
{
  "type": "<ADDED||MODIFIED||DELETED>",
  "object": {
      "metadata": {
            "kind": "Namespace",
            "name": "<namespace_name>",
            "namespace": "<namespace_name>",
            "annotations": { 
                "<annotation_one>": "<annotation_one_value>"
            },
            "labels": { 
                    "<label_one>": "<label_one_value>"
            }
          }
  }
}
```
* **Error Response:**   
  *Response body:*
```json
{
  "type": "ERROR",
  "object": {
      error message in json format
  }
}
```

# Pods restart status watch API

This API can be used for receiving pod status after Restart a deployment API calling.
Since pods appear sequentially then response body is supplemented gradually and the whole body will be before close code (1000). In general you do not build this API yourself and instead of it, you should use the one which you get from response body of Restart a deployment API.
* **API:**  /watchapi/v2/paas-mediation/namespaces/{namespace}/rollout-status
* **Request params (Optional):** `replicas=replication-controller:<replication_controller_name_1>,<replication_controller_name_2>;replica-set:<replica_set_name_1>,<replica_set_name_2>`
* **Headers:**      
  `Authorization: Bearer <jwt>`
* **Authorization:**
  M2M
  Response body:
```json
{
  "type":"CLOSE_CONTROL_MESSAGE",
  "object":{
    "<deployment_name>":
      [
        {
          "name":"<pod_name>",
          "status":"Ready"
        },
        {
          "name":"<pod_name>",
          "status":"Ready"
        }
      ],
    "<deployment_name>":
      [
        {
          "name":"<pod_name>",
          "status":"Ready"
        }
      ]
  }
}
```
* **Error Response:**   
  *Response body:*
```json
{
  "type": "ERROR",
  "object": {
      error message in json format
  }
}
```

* **Sample call**  
  *URL* "/watchapi/v2/paas-mediation/namespaces/cloudbss311-platform-core-support-dev2/rollout-status?replication-controller=deployment:maas-agent-4&replica-set=name:paas-mediation|pod-template-hash:375894095"

* **Response body**
 ```json
{ 
  "type":"CLOSE_CONTROL_MESSAGE",
  "object":{
    "maas-agent":
      [
        {
          "name":"maas-agent-4-4bsbb",
          "status":"Ready"
        },
        {
          "name":"maas-agent-4-66bmr",
          "status":"Ready"
        }
      ],
    "paas-mediation":
      [
        {
          "name":"paas-mediation-7c9df84f9-mjbk8",
          "status":"Ready"
        },
        {
          "name":"paas-mediation-7c9df84f9-zmsbn",
          "status":"Ready"
        }
      ]
  }
}
```
 