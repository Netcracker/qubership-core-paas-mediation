package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.domain.AnnotationResource;
import com.netcracker.it.paasmediation.v2.domain.MediationSecret;
import com.netcracker.it.paasmediation.v2.helpers.SecretHelper;
import io.fabric8.kubernetes.api.model.Secret;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import okhttp3.Response;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@Slf4j
@SmokeTest
public class SecretHttpIT extends SecretHelper {

    @AfterEach
    public void cleanup() {
        super.cleanup();
    }

    @Test
    void checkGetSecretAnnotation() throws IOException {
        log.info("Check get a secret annotation resource");
        Secret testSecret = createTestSecret(secretName1);
        log.info("Created secret {}", testSecret);
        paasUtils.createSecret(testSecret);
        assertNotNull(paasUtils.getSecretByName(secretName1), "Test secret was not created");
        log.info("Secret was created, try get it via paas-mediation API");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ANNOTATIONS, "", namespace, "GET", null, "annotation=owner&resourceType=Secret");
        List<AnnotationResource> annotationResource;
        try (Response response = paasMediationUtils.doRequest(request)) {
            String respBody = response.body().string();
            log.info("Response: {}, body {}", response, respBody);
            assertEquals(400, response.code());
        }
    }
}
