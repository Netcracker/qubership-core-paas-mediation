package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.Test;

import java.io.IOException;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@Slf4j
@SmokeTest
public class DeploymentFamilyIT extends PaasMediationParentV2Test {

    @Test
    public void checkEmptyList() throws IOException {
        String path = String.format("api/v2/paas-mediation/namespaces/%s/deployment-family/paas-mediation", namespace);
        Request request = paasMediationUtils.createRequest(path, "GET", null, null);
        Object[] result = paasMediationUtils.doRequest(request, 200, Object[].class);
        assertNotNull(result);
        assertEquals(0, result.length);
    }
}
