package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.domain.ApplicationData;
import com.netcracker.it.paasmediation.v2.helpers.ConfigMapHelper;
import io.fabric8.kubernetes.api.model.ConfigMap;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

import java.io.IOException;
import java.util.*;
import java.util.function.Function;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

@Slf4j
@SmokeTest
//@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public class VersionsIT extends ConfigMapHelper {

    private static final String VERSION_CONFIG_MAP_NAME = "version";

    @BeforeAll
    public void createTestVersionConfigMap() {
        ConfigMap versionConfigMap = createTestConfigMap(VERSION_CONFIG_MAP_NAME,
                Map.of(
                        "cloud-core.2025-12-31-23-59-59-249.user_a", "main-20251231.110000-62",
                        "cloud-core.2026-01-01-13-00-00-316.user_b", "main-20260101.110000-63"));
        paasUtils.createConfigMap(versionConfigMap);
        log.info("Created test version ConfigMap: {}", VERSION_CONFIG_MAP_NAME);
    }

    @AfterAll
    public void deleteTestVersionConfigMap() {
        if (paasUtils.getConfigMapByName(VERSION_CONFIG_MAP_NAME) != null) {
            paasUtils.deleteConfigMap(VERSION_CONFIG_MAP_NAME);
            log.info("Deleted test version ConfigMap: {}", VERSION_CONFIG_MAP_NAME);
        }
    }

    @Test
    public void checkGetVersionsAPI() throws IOException {
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.CONFIGMAPS, "versions", namespace, "GET", null, null);

        log.info("Request to {}", request.url());
        List<ApplicationData> versions = Arrays.asList(paasMediationUtils.doRequest(request, 200, ApplicationData[].class));
        log.info("Returned versions {}", versions);
        assertNotNull(versions);
        // compare content of 'versions' configmap with what returned by paas-mediation api
        String deploymentVersion = getLastCloudCoreVersion();
        log.info("deploymentVersion = {}", deploymentVersion);

        assertTrue(isCloudCorePresents(versions, deploymentVersion),
                String.format("Invalid api version response (missing cloud-core application): \n%s",
                        objectMapper.writeValueAsString(versions)));
    }

    private boolean isCloudCorePresents(List<ApplicationData> versions, String deploymentVersion) {
        return versions.stream().anyMatch(data -> "cloud-core".equalsIgnoreCase(data.getAppName()) &&
                Objects.equals(data.getAppVersion(), deploymentVersion));
    }

    private static String getLastCloudCoreVersion() {
        ConfigMap versionMap = kubernetesClient.configMaps().withName("version").get();
        assertNotNull(versionMap);
        return Optional.ofNullable(versionMap.getData().entrySet().stream()
                .filter(entry -> entry.getKey().toLowerCase().startsWith("cloud-core"))
                .reduce(null, (entry1, entry2) -> {
                    if (entry1 == null) {
                        return entry2;
                    }
                    Function<String, String> timeStampVerFunc = (key) -> {
                        Pattern pattern = Pattern.compile(".*\\.([\\d-]+)\\..*");
                        Matcher matcher = pattern.matcher(key);
                        if (matcher.matches()) {
                            return matcher.group(1);
                        } else {
                            throw new IllegalArgumentException(String.format("unsupported version format: %s. Expecting in pattern: %s", key, pattern));
                        }
                    };
                    String compareKey1 = timeStampVerFunc.apply(entry1.getKey());
                    String compareKey2 = timeStampVerFunc.apply(entry2.getKey());
                    return compareKey1.compareTo(compareKey2) >= 0 ? entry1 : entry2;
                })).map(Map.Entry::getValue).orElse("unknown");
    }
}
