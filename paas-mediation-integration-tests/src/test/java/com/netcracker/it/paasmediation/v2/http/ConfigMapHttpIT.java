package com.netcracker.it.paasmediation.v2.http;

import com.fasterxml.jackson.core.type.TypeReference;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.domain.AnnotationResource;
import com.netcracker.it.paasmediation.v2.domain.MediationConfigMap;
import com.netcracker.it.paasmediation.v2.helpers.ConfigMapHelper;
import io.fabric8.kubernetes.api.model.ConfigMap;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import okhttp3.Response;
import org.hamcrest.Matchers;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.*;

@Slf4j
@SmokeTest
public class ConfigMapHttpIT extends ConfigMapHelper {

    @AfterEach
    public void cleanup() {
        super.cleanup();
    }

    @Test
    public void checkGetConfigMapAPI() throws IOException {
        log.info("Check get a configmap API");
        ConfigMap configMap = createTestConfigMap(configMapName1);
        log.info("Created configmap {}", configMap);
        paasUtils.createConfigMap(configMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configmap was not created");
        log.info("Configmap was created, try get it via paas-mediation API");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "GET", null, null);
        MediationConfigMap mediationConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        assertNotNull(mediationConfigMap, "Created configmap was not found via paas-mediation API");
        log.info("ConfigMap received through paas-mediation {}", mediationConfigMap);
        checkConfigMap(configMap, mediationConfigMap);
    }

    @Test
    public void checkGetConfigMapListAPI() throws IOException {
        log.info("Check get a list of configMaps API");
        ConfigMap testConfigMap = createTestConfigMap(configMapName1);
        paasUtils.createConfigMap(testConfigMap);
        log.info("Created configMap {}", testConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        log.info("ConfigMap was created, try get it via paas-mediation API");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "GET", null, null);
        List<MediationConfigMap> mediationConfigMaps = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationConfigMap[].class));
        log.info("Got list of configMaps via paas-mediation API: {}", mediationConfigMaps);
        assertThat(mediationConfigMaps.size(), Matchers.greaterThanOrEqualTo(1));
        MediationConfigMap actualConfigMap = mediationConfigMaps.parallelStream()
                .filter(mediationService -> mediationService.getMetadata().getName().equals(configMapName1))
                .findAny()
                .orElse(null);
        assertNotNull(actualConfigMap, "Expected configMap was not returned through paas-mediation API");
        checkConfigMap(testConfigMap, actualConfigMap);
    }

    @Test
    public void CreateConfigMap() throws IOException {
        ConfigMap expectedConfigMap = createTestConfigMap(configMapName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "POST", expectedConfigMap, null);
        MediationConfigMap mediationConfigMap = paasMediationUtils.doRequest(request, 201, MediationConfigMap.class);
        log.info("Created configMap {}", mediationConfigMap);
        assertNotNull(mediationConfigMap, "Test configMap was not created using paas mediation API");
        checkConfigMap(expectedConfigMap, mediationConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "GET", null, null);
        MediationConfigMap actualConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        checkConfigMap(expectedConfigMap, actualConfigMap);
    }

    @Test
    public void CreateAndDeleteConfigMap() throws IOException {
        ConfigMap expectedConfigMap = createTestConfigMap(configMapName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "POST", expectedConfigMap, null);
        MediationConfigMap mediationConfigMap = paasMediationUtils.doRequest(request, 201, MediationConfigMap.class);
        log.info("Created configMap {}", mediationConfigMap);
        assertNotNull(mediationConfigMap, "Test configMap was not created using paas mediation API");
        checkConfigMap(expectedConfigMap, mediationConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "DELETE", null, null);
        paasMediationUtils.doRequest(request, 200, null);
        assertNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
    }

    @Test
    public void CreateTwiceConfigMap() throws IOException {
        ConfigMap expectedConfigMap = createTestConfigMap(configMapName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "POST", expectedConfigMap, null);
        paasMediationUtils.doRequest(request, 201, MediationConfigMap.class);
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "POST", expectedConfigMap, null);
        paasMediationUtils.doRequest(request, 409, MediationConfigMap.class);
    }

    @Test
    public void DeleteConfigMapNotFound() throws IOException {
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "DELETE", null, null);
        paasMediationUtils.doRequest(request, 404, MediationConfigMap.class);
    }

    @Test
    public void CreateOrUpdateConfigMap_WithCreateConfigMap() throws IOException {
        ConfigMap expectedConfigMap = createTestConfigMap(configMapName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "PUT", expectedConfigMap, null);
        MediationConfigMap mediationConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        log.info("Created configMap {}", mediationConfigMap);
        assertNotNull(mediationConfigMap, "Test configMap was not created using paas mediation API");
        checkConfigMap(expectedConfigMap, mediationConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "GET", null, null);
        MediationConfigMap actualConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        checkConfigMap(expectedConfigMap, actualConfigMap);
    }

    @Test
    public void CreateOrUpdateConfigMap_WithUpdateConfigMap() throws IOException {
        ConfigMap expectedConfigMap = createTestConfigMap(configMapName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "PUT", expectedConfigMap, null);
        paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        expectedConfigMap.getMetadata().setAnnotations(Collections.singletonMap("test-annotation", "test-annotations-value"));
        expectedConfigMap.getMetadata().setLabels(Collections.singletonMap("test-label", "test-label-value"));
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "PUT", expectedConfigMap, null);
        MediationConfigMap mediationConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        log.info("Created configMap {}", mediationConfigMap);
        assertNotNull(mediationConfigMap, "Test configMap was not created using paas mediation API");
        checkConfigMap(expectedConfigMap, mediationConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "GET", null, null);
        MediationConfigMap actualConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        checkConfigMap(expectedConfigMap, actualConfigMap);
    }

    @Test
    public void CreateOrUpdateConfigMapTwice() throws IOException {
        ConfigMap expectedConfigMap = createTestConfigMap(configMapName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "POST", expectedConfigMap, null);
        paasMediationUtils.doRequest(request, 201, MediationConfigMap.class);
        expectedConfigMap.getMetadata().setAnnotations(Collections.singletonMap("test-annotation", "test-annotations-value"));
        expectedConfigMap.getMetadata().setLabels(Collections.singletonMap("test-label", "test-label-value"));
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, null, namespace, "PUT", expectedConfigMap, null);
        MediationConfigMap mediationConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        log.info("Created configMap {}", mediationConfigMap);
        assertNotNull(mediationConfigMap, "Test configMap was not created using paas mediation API");
        checkConfigMap(expectedConfigMap, mediationConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, configMapName1, namespace, "GET", null, null);
        MediationConfigMap actualConfigMap = paasMediationUtils.doRequest(request, 200, MediationConfigMap.class);
        checkConfigMap(expectedConfigMap, actualConfigMap);
    }

    @Test
    void checkGetConfigMapAnnotation() throws IOException {
        log.info("Check get a configMap annotation resource");
        ConfigMap testConfigMap = createTestConfigMap(configMapName1);
        log.info("Created configMap {}", testConfigMap);
        paasUtils.createConfigMap(testConfigMap);
        assertNotNull(paasUtils.getConfigMapByName(configMapName1), "Test configMap was not created");
        log.info("ConfigMap was created, try get it via paas-mediation API");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ANNOTATIONS, "", namespace, "GET", null, "annotation=owner&resourceType=ConfigMap");
        List<AnnotationResource> annotationResource;
        try (Response response = paasMediationUtils.doRequest(request)) {
            String respBody = response.body().string();
            log.info("Response: {}, body {}", response, respBody);
            assertEquals(200, response.code());
            annotationResource = objectMapper.readValue(respBody, new TypeReference<List<AnnotationResource>>(){});
        }

        assertNotNull(annotationResource, "Created configMap was not found via paas-mediation API");
        log.info("ConfigMap annotation resource received through paas-mediation {}", annotationResource);
        assertEquals(testConfigMap.getMetadata().getAnnotations().get("owner"), annotationResource.get(0).getAnnotationValue());
    }
}
