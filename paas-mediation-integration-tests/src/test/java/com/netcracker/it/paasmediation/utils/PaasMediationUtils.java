package com.netcracker.it.paasmediation.utils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import okhttp3.*;
import org.apache.commons.lang3.StringUtils;

import java.io.IOException;
import java.util.Map;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;


@Slf4j
public class PaasMediationUtils {
    private static final MediaType JSON = MediaType.parse("application/json");
    private final ObjectMapper objectMapper;
    private String internalGateway;
    private OkHttpClient okHttpClient;
    private String apiVersion;

    public PaasMediationUtils(String apiVersion, String internalGateway, OkHttpClient okHttpClient, ObjectMapper objectMapper) {
        this.apiVersion = apiVersion;
        this.internalGateway = internalGateway;
        this.okHttpClient = okHttpClient;
        this.objectMapper = objectMapper;
    }

    public <T> T doRequest(Request request, int expectStatus, Class<T> clazz) throws IOException {
        try (Response response = doRequest(request)) {
            String respBody = response.body().string();
            log.info("Response: {}, body {}", response, respBody);
            assertEquals(expectStatus, response.code());
            if (clazz == null) {
                return null;
            }
            return objectMapper.readValue(respBody, clazz);
        }
    }

    public Response doRequest(Request request) throws IOException {
        return okHttpClient.newCall(request).execute();
    }

    private Response doRequestWithRetry(Request request, int expectStatus, int retryAmount) {
        int callAmount = 0;
        boolean tryAgain = true;
        Response response = null;
        log.info("Start request with retry policy. Request: {}", request);

        while (tryAgain) {
            try {
                callAmount++;
                response = doRequest(request);

                if (response == null || response.code() != expectStatus) {
                    tryAgain = callAmount < retryAmount;
                    log.info("Failed call, {} call remains", retryAmount - callAmount);
                } else {
                    tryAgain = false;
                }

            } catch (IOException e) {
                log.info("Exception during request to {}. Exception: {}", request.url(), e);
                tryAgain = callAmount < retryAmount;
            }
        }

        return response;
    }

    public <T> T doRequestWithRetry(Request request, int expectStatus, int retryAmount, Class<T> clazz) throws IOException {
        try (Response response = doRequestWithRetry(request, expectStatus, retryAmount)) {
            if (clazz == null) {
                return null;
            }
            assertEquals(expectStatus, response.code());
            String respBody = response.body().string();
            return objectMapper.readValue(respBody, clazz);
        }
    }

    public Request createRequest(Resources resource, String name, String namespace, String httpMethod, Object requestBody, String paramString) throws JsonProcessingException {
        RequestBody body = null;
        if (requestBody != null) {
            String toJson = objectMapper.writeValueAsString(requestBody);
            body = RequestBody.create(toJson, JSON);
        }
        return new Request.Builder()
                .url(internalGateway
                        + buildEndpoint(resource.name().toLowerCase(), name, namespace)
                        + (!StringUtils.isEmpty(paramString) ? "?" + paramString : ""))
                .method(httpMethod, body)
                .header("x-request-id", String.format("%s-%s", resource.getValue(), UUID.randomUUID().toString().substring(0, 8)))
                .build();
    }

    public Request createRequest(String path, String httpMethod, Object requestBody, String paramString) throws JsonProcessingException {
        RequestBody body = null;
        if (requestBody != null) {
            String toJson = objectMapper.writeValueAsString(requestBody);
            body = RequestBody.create(toJson, JSON);
        }
        return new Request.Builder()
                .url(internalGateway
                        + path
                        + (!StringUtils.isEmpty(paramString) ? "?" + paramString : ""))
                .method(httpMethod, body)
                .build();
    }


    private String buildEndpoint(String resource, String name, String namespace) {
        String url = !StringUtils.isEmpty(namespace) ?
                String.format("api/%s/paas-mediation/namespaces/%s/%s", apiVersion, namespace, resource) :
                String.format("api/%s/paas-mediation/%s", apiVersion, resource);
        url += !StringUtils.isEmpty(name) ? "/" + name : "";
        log.info("http url was built {}", url);
        return url;
    }

    public enum Resources {
        ROUTES("routes"), CONFIGMAPS("configmaps"), SERVICES("services"), SECRETS("secrets"),
        NAMESPACES("namespaces"), PODS("pods"), DEPLOYMENTS("deployments"),
        ROLLOUT("rollout"), ROLLOUT_STATUS("rollout-status"), ANNOTATIONS("annotations");

        private String value;

        public String getValue() {
            return value;
        }

        Resources(String s) {
            this.value = s;
        }
    }

    public WebSocket websocketConnect(WebSocketListener webSocketListener, Request request) {
        return okHttpClient.newWebSocket(request, webSocketListener);
    }

    public Request createWsRequest(Resources resource, String namespace) {
        return createWsRequest(resource, namespace, PaasRequestFilter.EVERYTHING);
    }

    public Request createWsRequest(Resources resource, String namespace, PaasRequestFilter filter) {
        String wsUrl = buildWsEndpoint(resource.toString().toLowerCase(), namespace, filter);
        String id = "paas-mediation-it-test-" + UUID.randomUUID().toString().substring(24);
        log.debug("wsUrl={} X-Request-Id={}", wsUrl, id);
        return new Request.Builder().url(wsUrl)
                .header("X-Request-Id", id)
                .build();
    }

    public Request createWsRequest(String wsUrl) {
        return new Request.Builder().url(internalGateway.replace("http", "ws") + wsUrl).build();
    }

    private String buildWsEndpoint(String resource, String namespace, PaasRequestFilter filter) {
        String url = internalGateway.replace("http", "ws") + String.format("watchapi/%s/paas-mediation", apiVersion);
        url += StringUtils.isEmpty(namespace) ? String.format("/%s", resource) : String.format("/namespaces/%s/%s", namespace, resource);
        String query = "";
        String annotations = getFilterParam("annotations", filter.getAnnotations());
        query = addParamToQuery(query, annotations);
        String labels = getFilterParam("labels", filter.getLabels());
        query = addParamToQuery(query, labels);
        if (query.length() > 0) {
            url += "?" + query;
        }
        log.info("watch url was built {}", url);
        return url;
    }

    private String addParamToQuery(String query, String param) {
        if (param.length() > 0) {
            if (query.length() > 0) {
                query += "&";
            }
            query += param;
        }
        return query;
    }

    private String getFilterParam(String paramName, Map<String, String> paramAsMap) {
        StringBuilder sb = new StringBuilder();
        paramAsMap.forEach((key, value) -> {
            if (sb.length() > 0) {
                sb.append("%3B");
            }
            sb.append(key).append(':').append(value);
        });
        if (sb.length() > 0) {
            return paramName + "=" + sb.toString();
        } else {
            return "";
        }
    }

}
