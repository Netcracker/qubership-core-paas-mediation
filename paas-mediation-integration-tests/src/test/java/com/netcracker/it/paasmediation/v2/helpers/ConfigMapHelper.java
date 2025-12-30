package com.netcracker.it.paasmediation.v2.helpers;

import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationConfigMap;
import io.fabric8.kubernetes.api.model.ConfigMap;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;

import java.util.Collections;
import java.util.Map;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

public abstract class ConfigMapHelper extends PaasMediationParentV2Test {

    protected String configMapNamePrefix = "paas-mediation-it-test-configmap";
    protected String configMapName1 = configMapNamePrefix + "-1";
    protected String configMapName2 = configMapNamePrefix + "-2";

    @BeforeAll
    @AfterAll
    public void cleanup() {
        Stream.of(configMapName1, configMapName2).forEach(name -> {
            if (paasUtils.getConfigMapByName(name) != null) {
                paasUtils.deleteConfigMap(name);
                assertNull(paasUtils.getConfigMapByName(name));
            }
        });
    }

    protected void checkConfigMap(ConfigMap expectedConfigMap, MediationConfigMap mediationConfigMap) {
        assertEquals(expectedConfigMap.getKind(), mediationConfigMap.getMetadata().getKind());
        assertEquals(expectedConfigMap.getMetadata().getName(), mediationConfigMap.getMetadata().getName());
        assertEquals(expectedConfigMap.getMetadata().getNamespace(), mediationConfigMap.getMetadata().getNamespace());
        assertEquals(expectedConfigMap.getMetadata().getLabels(), mediationConfigMap.getMetadata().getLabels());
        assertTrue(expectedConfigMap.getMetadata().getAnnotations().entrySet().containsAll(mediationConfigMap.getMetadata().getAnnotations().entrySet()));
        assertEquals("very important data", mediationConfigMap.getData().get("paas-mediation-configMap-data"));
    }

    protected ConfigMap createTestConfigMap(String name) {
        ConfigMap configMap = new ConfigMap();
        configMap.setMetadata(createTestMetadata(name));
        configMap.setKind("ConfigMap");
        configMap.setData(Collections.singletonMap("paas-mediation-configMap-data", "very important data"));
        return configMap;
    }

    protected ConfigMap createTestConfigMap(String name, Map<String, String> data) {
        ConfigMap configMap = new ConfigMap();
        configMap.setMetadata(createTestMetadata(name));
        configMap.setKind("ConfigMap");
        configMap.setData(data);
        return configMap;
    }
}
