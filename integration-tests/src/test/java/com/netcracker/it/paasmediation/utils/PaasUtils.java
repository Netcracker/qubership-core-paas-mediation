package com.netcracker.it.paasmediation.utils;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.fabric8.kubernetes.api.model.ConfigMap;
import io.fabric8.kubernetes.api.model.Pod;
import io.fabric8.kubernetes.api.model.Secret;
import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.api.model.apps.ReplicaSet;
import io.fabric8.kubernetes.api.model.networking.v1.Ingress;
import io.fabric8.kubernetes.api.model.networking.v1.IngressBuilder;
import io.fabric8.kubernetes.api.model.StatusDetails;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.dsl.ExecWatch;
import lombok.extern.slf4j.Slf4j;
import okhttp3.*;
import okio.Buffer;

import java.util.ArrayList;
import java.util.Collections;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

@Slf4j
public class PaasUtils {

    private final KubernetesClient kubernetesClient;
    private final ObjectMapper objectMapper;

    public PaasUtils(KubernetesClient kubernetesClient) {
        this.kubernetesClient = kubernetesClient;
        this.objectMapper = new ObjectMapper();
    }

    public String execInPod(String podName, String... command) throws Exception {
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        ByteArrayOutputStream err = new ByteArrayOutputStream();
        
        log.info("Executing in pod {}: {}", podName, String.join(" ", command));
        
        try (ExecWatch execWatch = kubernetesClient.pods()
                .withName(podName)
                .writingOutput(out)
                .writingError(err)
                .exec(command)) {
            
            Integer exitCode = execWatch.exitCode().get(60, TimeUnit.SECONDS);
            String output = out.toString();
            String error = err.toString();
            
            log.info("Exit code: {}", exitCode);
            
            if (!output.isEmpty()) {
                log.debug("Output length: {}", output.length());
                if (log.isTraceEnabled()) {
                    log.trace("Output: {}", output);
                }
            }
            
            if (!error.isEmpty()) {
                log.warn("Error output: {}", error);
            }
            
            if (exitCode != null && exitCode != 0) {
                throw new RuntimeException(String.format("Command failed with exit code %d. Error: %s", exitCode, error));
            }
            
            return output;
        } catch (Exception e) {
            log.error("Failed to exec in pod {}: {}", podName, e.getMessage());
            log.error("Command was: {}", String.join(" ", command));
            throw e;
        }
    }

    public String execInPodByLabel(String labelKey, String labelValue, String... command) throws Exception {
        String podName = getPodNameByLabel(labelKey, labelValue);
        if (podName == null) {
            throw new RuntimeException(String.format("No pod found with label %s=%s", labelKey, labelValue));
        }
        return execInPod(podName, command);
    }

    public Response doRequestFromInclusterPod(String podName, Request request) throws Exception {
        String curlCommand = buildCurlCommand(request);
        log.info("Executing curl from pod {}: {}", podName, curlCommand);

        String output = execInPod(podName, "sh", "-c", curlCommand);

        return parseCurlResponse(output);
    }

    public Response doRequestFromInclusterPodByLabel(String labelKey, String labelValue, Request request)
            throws Exception {
        String podName = getPodNameByLabel(labelKey, labelValue);
        if (podName == null) {
            throw new RuntimeException(String.format("No pod found with label %s=%s", labelKey, labelValue));
        }
        return doRequestFromInclusterPod(podName, request);
    }

    private String buildCurlCommand(Request request) {
        StringBuilder curl = new StringBuilder("curl -s -i --max-time 5 --no-keepalive");
        
        // Add method
        String method = request.method();
        if (!"GET".equals(method)) {
            curl.append(" -X ").append(method);
        }
        
        // Add headers
        Headers headers = request.headers();
        boolean hasContentType = false;
        for (String name : headers.names()) {
            String value = headers.get(name);
            String escapedValue = value.replace("'", "'\\''");
            curl.append(" -H '").append(name).append(": ").append(escapedValue).append("'");
            if ("Content-Type".equalsIgnoreCase(name) && value.contains("json")) {
                hasContentType = true;
            }
        }
        RequestBody body = request.body();
        if (body != null && !"GET".equals(method) && !"DELETE".equals(method)) {
            if (!hasContentType) {
                curl.append(" -H 'Content-Type: application/json'");
            }
            
            try {
                okio.Buffer buffer = new okio.Buffer();
                body.writeTo(buffer);
                String bodyString = buffer.readUtf8();
                String escapedBody = bodyString.replace("'", "'\\''");
                curl.append(" -d '").append(escapedBody).append("'");
                
                log.debug("Request body length: {}, preview: {}", bodyString.length(), 
                    bodyString.length() > 100 ? bodyString.substring(0, 100) + "..." : bodyString);
            } catch (IOException e) {
                log.error("Failed to read request body", e);
            }
        }
        
        String url = request.url().toString();
        curl.append(" '").append(url.replace("'", "'\\''")).append("'");
        
        String curlCommand = curl.toString();
        log.info("Generated curl command (length: {}): {}", curlCommand.length(), 
            curlCommand.length() > 500 ? curlCommand.substring(0, 500) + "..." : curlCommand);
        
        return curlCommand;
    }

    private String escapeForShell(String input) {
        if (input == null) {
            return "";
        }
        return input.replace("'", "'\\''");
    }

    private Response parseCurlResponse(String output) throws IOException {
        if (output == null || output.isEmpty()) {
            throw new IOException("Empty response from pod");
        }

        String[] parts = output.split("\r\n\r\n|\n\n", 2);
        String headersPart = parts[0];
        String bodyPart = parts.length > 1 ? parts[1] : "";

        String[] headerLines = headersPart.split("\r\n|\n");
        String statusLine = headerLines[0];

        int statusCode = extractStatusCode(statusLine);
        String reasonPhrase = extractReasonPhrase(statusLine);

        Response.Builder responseBuilder = new Response.Builder()
                .request(new Request.Builder().url("http://localhost").build())
                .protocol(Protocol.HTTP_1_1)
                .code(statusCode)
                .message(reasonPhrase);

        for (int i = 1; i < headerLines.length; i++) {
            String headerLine = headerLines[i];
            int colonIndex = headerLine.indexOf(':');
            if (colonIndex > 0) {
                String headerName = headerLine.substring(0, colonIndex).trim();
                String headerValue = headerLine.substring(colonIndex + 1).trim();
                responseBuilder.addHeader(headerName, headerValue);
            }
        }

        MediaType mediaType = MediaType.parse(responseBuilder.build().header("Content-Type", "text/plain"));
        ResponseBody responseBody = ResponseBody.create(bodyPart, mediaType);
        responseBuilder.body(responseBody);

        return responseBuilder.build();
    }

    private int extractStatusCode(String statusLine) {
        String[] parts = statusLine.split(" ");
        if (parts.length >= 2) {
            try {
                return Integer.parseInt(parts[1]);
            } catch (NumberFormatException e) {
                log.warn("Failed to parse status code from: {}", statusLine);
            }
        }
        return 500;
    }

    private String extractReasonPhrase(String statusLine) {
        String[] parts = statusLine.split(" ");
        if (parts.length >= 3) {
            StringBuilder reason = new StringBuilder();
            for (int i = 2; i < parts.length; i++) {
                if (reason.length() > 0)
                    reason.append(" ");
                reason.append(parts[i]);
            }
            return reason.toString();
        }
        return "";
    }

    public <T> T doRequestFromInclusterPod(String podName, Request request, int expectStatus, Class<T> clazz)
            throws Exception {
        try (Response response = doRequestFromInclusterPod(podName, request)) {
            String respBody = response.body() != null ? response.body().string() : "";
            log.info("Response from incluster pod: status={}, body={}", response.code(), respBody);

            if (response.code() != expectStatus) {
                throw new AssertionError(String.format("Expected status %d but got %d. Body: %s",
                        expectStatus, response.code(), respBody));
            }

            if (clazz == null) {
                return null;
            }
            return objectMapper.readValue(respBody, clazz);
        }
    }

    public void deleteIngress(String name) {
        log.info("Start delete ingress {}", name);
        kubernetesClient.network().v1().ingresses().withName(name).withTimeout(1, TimeUnit.MINUTES).delete();
        log.info("Ingress {} was deleted", name);
    }

    public Ingress getIngressByName(String name) {
        return kubernetesClient.network().v1().ingresses().withName(name).get();
    }

    public Service createService(Service service) {
        log.info("Start create service {}", service.getMetadata().getName());
        Service createdService = kubernetesClient.services().resource(service).create();
        log.info("Service was created {}", service);
        return createdService;
    }

    public Pod createPod(Pod pod) {
        log.info("Start create pod {}", pod.getMetadata().getName());
        Pod createdPod = kubernetesClient.pods().resource(pod).create();
        log.info("Pod was created {}", pod);
        return createdPod;
    }

    public void deleteService(String name) {
        log.info("Start delete service {}", name);
        kubernetesClient.services().withName(name).withTimeout(1, TimeUnit.MINUTES).delete();
        log.info("service was deleted {}", name);
    }

    public Service getServiceByName(String name) {
        return kubernetesClient.services().withName(name).get();
    }

    public Pod getPodByName(String name) {
        return kubernetesClient.pods().withName(name).get();
    }

    public String getPodNameByLabel(String label, String value) {
        var pods = kubernetesClient.pods()
                .withLabel(label, value)
                .list();

        if (!pods.getItems().isEmpty()) {
            String podName = pods.getItems().get(0).getMetadata().getName();
            log.info("Found pod: {}", podName);
            return podName;
        }
        return null;
    }

    public Secret createSecret(Secret secret) {
        log.info("Start create secret {}", secret.getMetadata().getName());
        Secret createdSecret = kubernetesClient.secrets().resource(secret).create();
        log.info("Secret {} was created", secret.getMetadata().getName());
        return createdSecret;
    }

    public void deleteSecret(String name) {
        log.info("Start delete secret {}", name);
        kubernetesClient.secrets().withName(name).delete();
        log.info("Secret {} was deleted", name);
    }

    public Secret getSecretByName(String name) {
        return kubernetesClient.secrets().withName(name).get();
    }

    public ConfigMap createConfigMap(ConfigMap configMap) {
        log.info("Start create configMap {}", configMap.getMetadata().getName());
        ConfigMap createdConfigMap = kubernetesClient.configMaps().resource(configMap).create();
        log.info("ConfigMap {} was created", createdConfigMap.getMetadata().getName());
        return createdConfigMap;
    }

    public void deleteConfigMap(String name) {
        log.info("Start delete configMap {}", name);
        kubernetesClient.configMaps().withName(name).delete();
        log.info("ConfigMap {} was deleted", name);
    }

    public ConfigMap getConfigMapByName(String name) {
        return kubernetesClient.configMaps().withName(name).get();
    }

    public String getImage(String name) {
        Deployment deployment = kubernetesClient.apps().deployments().withName(name).get();
        if (deployment != null)
            return deployment.getSpec().getTemplate().getSpec().getContainers().getFirst().getImage();
        return null;
    }

    public Set<String> getPodNamesByReplica(String replicaName) {
        return kubernetesClient.pods().list().getItems().stream()
                .filter(pod -> pod.getMetadata().getName().startsWith(replicaName)
                        && !pod.getMetadata().getName().endsWith("deploy"))
                .map(pod -> pod.getMetadata().getName())
                .collect(Collectors.toSet());
    }

    public ReplicaSet getLatestReplicaSetByDeploymentName(String deploymentName) {
        return kubernetesClient.apps().replicaSets().list().getItems().stream()
                .filter(rs -> rs.getMetadata().getName().startsWith(deploymentName))
                .max(Comparator.comparing(rs -> Integer
                        .parseInt(rs.getMetadata().getAnnotations().get("deployment.kubernetes.io/revision"))))
                .orElse(null);
    }

    public List<Pod> getAllPodsByNamespace(String namespace) {
        return kubernetesClient.pods().inNamespace(namespace).list().getItems();
    }

    public List<Deployment> getAllDeploymentsByNamespace(String namespace) {
        return kubernetesClient.apps().deployments().inNamespace(namespace).list().getItems();
    }
}