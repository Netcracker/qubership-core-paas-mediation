package com.netcracker.it.paasmediation.utils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import okhttp3.*;
import org.apache.commons.lang3.StringUtils;

import java.io.IOException;
import java.util.Base64;
import java.util.Map;
import java.util.UUID;


import static org.junit.jupiter.api.Assertions.assertEquals;

@Slf4j
public class PaasMediationUtils {
    private static final MediaType JSON = MediaType.parse("application/json");
    private final ObjectMapper objectMapper;
    private final String internalGateway;
    private final String apiVersion;
    private final RequestExecutor requestExecutor;

    public PaasMediationUtils(String apiVersion, String internalGateway, RequestExecutor requestExecutor, ObjectMapper objectMapper) {
        this.apiVersion = apiVersion;
        this.internalGateway = internalGateway;
        this.requestExecutor = requestExecutor;
        this.objectMapper = objectMapper;
    }

    public <T> T doRequest(Request request, int expectStatus, Class<T> clazz) throws IOException {
        try (Response response = requestExecutor.execute(request)) {
            String respBody = response.body() != null ? response.body().string() : "";
            log.info("Response: {}, body {}", response, respBody);
            assertEquals(expectStatus, response.code());
            if (clazz == null) {
                return null;
            }
            return objectMapper.readValue(respBody, clazz);
        }
    }

    public Response doRequest(Request request) throws IOException {
        return requestExecutor.execute(request);
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
            String respBody = response.body() != null ? response.body().string() : "";
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
                .url(internalGateway + path + (!StringUtils.isEmpty(paramString) ? "?" + paramString : ""))
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

        private final String value;

        Resources(String s) {
            this.value = s;
        }

        public String getValue() {
            return value;
        }
    }

    public Request createWsRequest(Resources resource, String namespace) {
        return createWsRequest(resource, namespace, PaasRequestFilter.EVERYTHING);
    }

    public Request createWsRequest(Resources resource, String namespace, PaasRequestFilter filter) {
        String wsUrl = buildWsEndpoint(resource.toString().toLowerCase(), namespace, filter);
        String id = "paas-mediation-it-test-" + UUID.randomUUID().toString().substring(24);
        log.debug("WebSocket URL built: {}", wsUrl);
        log.debug("X-Request-Id: {}", id);
        
        String secWebSocketKey = Base64.getEncoder().encodeToString(UUID.randomUUID().toString().getBytes());
        
        Request request = new Request.Builder()
                .url(wsUrl)
                .header("X-Request-Id", id)
                .header("Connection", "Upgrade")
                .header("Upgrade", "websocket")
                .header("Sec-WebSocket-Key", secWebSocketKey)
                .header("Sec-WebSocket-Version", "13")
                .build();
        log.debug("After build wsUrl={}", request.url()); 
        return request;
    }

    public Request createWsRequest(String wsUrl) {
        return new Request.Builder().url(internalGateway.replace("http", "ws") + wsUrl).build();
    }

    private String buildWsEndpoint(String resource, String namespace, PaasRequestFilter filter) {
        String base = internalGateway;
        if (base.endsWith("/")) {
            base = base.substring(0, base.length() - 1);
        }
        
        String wsBase = base.replace("http://", "ws://").replace("https://", "wss://");
        
        String url = String.format("%s/watchapi/%s/paas-mediation", wsBase, apiVersion);
        url += StringUtils.isEmpty(namespace) ? 
                String.format("/%s", resource) : 
                String.format("/namespaces/%s/%s", namespace, resource);
        
        String query = "";
        String annotations = getFilterParam("annotations", filter.getAnnotations());
        query = addParamToQuery(query, annotations);
        String labels = getFilterParam("labels", filter.getLabels());
        query = addParamToQuery(query, labels);
        
        if (query.length() > 0) {
            url += "?" + query;
        }
        
        log.info("WebSocket URL built: {}", url);
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