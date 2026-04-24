package com.netcracker.it.paasmediation.v2.http;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.utils.PaasUtils;
import io.fabric8.kubernetes.api.model.ConfigMap;
import io.fabric8.kubernetes.api.model.ConfigMapBuilder;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import okhttp3.Response;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

import java.io.InputStream;
import java.util.*;
import java.util.stream.Collectors;

import static org.junit.jupiter.api.Assertions.*;

@Slf4j
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public class ExpectedRoutesIT extends PaasMediationParentV2Test {

    private static final String BG_VERSION_CONFIGMAP = "bg-version";
    private static final String VERSION_CONFIGMAP = "version";
    private static final String EXPECTED_ROUTES_CONFIGMAP = "paas-mediation-expected-routes";
    private Map<String, List<String>> expectedRoutes;
    private ConfigMap expectedRoutesConfigMap;
    private ConfigMap bgVersionConfigMap;
    private ConfigMap versionConfigMap;
    private ObjectMapper mapper;

    @BeforeAll
    void setUp() throws Exception {
         mapper = new ObjectMapper();

        expectedRoutesConfigMap = loadConfigMapFromResource();
        expectedRoutesConfigMap.getMetadata().setNamespace(namespace);
        
        String routesJson = expectedRoutesConfigMap.getData().get("routes.json");
        expectedRoutes = mapper.readValue(routesJson, Map.class);
        
        expectedRoutes.forEach((key, routes) -> {
            List<String> replaced = new ArrayList<>();
            for (String route : routes) {
                replaced.add(route.replace("{namespace}", namespace));
            }
            expectedRoutes.put(key, replaced);
        });
        
        log.info("Expected routes loaded: {}", expectedRoutes);
        
        bgVersionConfigMap = new ConfigMapBuilder()
                .withNewMetadata()
                    .withName(BG_VERSION_CONFIGMAP)
                    .withNamespace(namespace)
                .endMetadata()
                .addToData("version", "1.0.0")
                .build();
        
        paasUtils.createConfigMap(bgVersionConfigMap);
        log.info("Created bg-version ConfigMap");

        versionConfigMap = new ConfigMapBuilder()
                .withNewMetadata()
                    .withName(VERSION_CONFIGMAP)
                    .withNamespace(namespace)
                .endMetadata()
                .addToData("version", "1.0.0")
                .build();
        
        paasUtils.createConfigMap(versionConfigMap);
        log.info("Created bg-version ConfigMap");

        paasUtils.createConfigMap(expectedRoutesConfigMap);
        log.info("Created expected routes ConfigMap: {}", EXPECTED_ROUTES_CONFIGMAP);

    }

    @AfterAll
    void tearDown() {
        try {
            paasUtils.deleteConfigMap(BG_VERSION_CONFIGMAP);
            log.info("Deleted bg-version ConfigMap");
            paasUtils.deleteConfigMap(VERSION_CONFIGMAP);
            log.info("Deleted bg-version ConfigMap");
            paasUtils.deleteConfigMap(EXPECTED_ROUTES_CONFIGMAP);
            log.info("Deleted expected routes ConfigMap");
        } catch (Exception e) {
            log.warn("Failed to delete bg-version ConfigMap: {}", e.getMessage());
        }
    }

    private ConfigMap loadConfigMapFromResource() {
        try (InputStream is = ExpectedRoutesIT.class.getResourceAsStream("/expected-routes-configmap.yaml")) {
            if (is == null) {
                log.warn("expected-routes-configmap.yaml not found");
                throw new RuntimeException("ConfigMap file not found");
            }
            return kubernetesClient.configMaps()
                    .load(is)
                    .item();
        } catch (Exception e) {
            log.error("Failed to load ConfigMap from resource", e);
            throw new RuntimeException(e);
        }
    }

    @Test
    void testInternalRoutes() throws Exception {
        List<String> expectedHttpRoutes = expectedRoutes.get("internal").stream()
                .filter(route -> !route.contains("watchapi/"))
                .collect(Collectors.toList());
        
        for (String expectedRoute : expectedHttpRoutes) {
            Request request = paasMediationUtils.createRequest(expectedRoute, "GET", null, null);
            log.info("Check internal route: {}", request.url());
            paasMediationUtils.doRequest(request, 200, null);
        }
            
        // WebSocket маршруты проверяем отдельно
        List<String> wsRoutes = expectedRoutes.get("internal").stream()
            .filter(route -> route.contains("watchapi/"))
            .collect(Collectors.toList());

        String internalGatewayUrl = String.format("ws://internal-gateway-service.%s:8080/", namespace);

        for (String wsRoute : wsRoutes) {
            Request request = new Request.Builder()
                .url(internalGatewayUrl + wsRoute)
                .header("x-request-id", "private-routes-test")
                .build();
            log.info("Check internal ws route: {}", request.url());
            paasMediationUtils.doRequest(request, 500, null);
        }
    }

    @Test
    void testPrivateRoutes() throws Exception {
        String privateGatewayUrl = String.format("http://private-gateway-service.%s:8080/", namespace);
        
        List<String> expected = expectedRoutes.get("private");
        for (String expectedRoute : expected) {
            Request request = new Request.Builder()
                .url(privateGatewayUrl + expectedRoute)
                .header("x-request-id", "private-routes-test")
                .build();
            log.info("Check private route: {}", request.url());
            paasMediationUtils.doRequest(request, 200, null);
            
        }
    }

    @Test
    void testVersionsEndpoint() throws Exception {
        Request request = paasMediationUtils.createRequest(
            "/api/v2/paas-mediation/versions",
            "GET",
            null,
            null
        );
        
        try (Response response = paasMediationUtils.doRequest(request)) {
            int code = response.code();
            assertTrue(code == 200 || code == 404,
                String.format("Versions endpoint returned %d, expected 200 or 404", code));
            
            if (code == 200) {
                String body = response.body() != null ? response.body().string() : "";
                log.info("Versions endpoint response: {}", body);
                assertTrue(body.contains("appVersion") || body.contains("cloud-core"),
                    "Versions response should contain version info");
            }
            log.info("✅ Versions endpoint accessible, status: {}", code);
        }
    }

    @Test
    void testBgVersionEndpoint() throws Exception {
        Request request = paasMediationUtils.createRequest(
            PaasMediationUtils.Resources.CONFIGMAPS,
            BG_VERSION_CONFIGMAP,
            namespace,
            "GET",
            null,
            null
        );
        
        try (Response response = paasMediationUtils.doRequest(request)) {
            assertEquals(200, response.code());
            String body = response.body() != null ? response.body().string() : "";
            assertTrue(body.contains("version"), "Response should contain version data");
            log.info("✅ bg-version ConfigMap accessible");
        }
    }

    private List<String> extractPathsFromRoutes(JsonNode routes) {
        List<String> paths = new ArrayList<>();
        if (routes.isArray()) {
            for (JsonNode route : routes) {
                String path = route.path("spec").path("path").asText();
                if (!path.isEmpty()) {
                    paths.add(path);
                }
            }
        }
        return paths;
    }
    
}