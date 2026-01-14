package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.PortForward;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.Value;
import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

@Slf4j
public class HealthIT extends PaasMediationParentV2Test {

    @PortForward(serviceName = @Value("paas-mediation"))
    private static URL paasMediationAddress;

    @Test
    public void checkHealth() throws IOException {
        log.info("check health status");
        Request request = new Request.Builder()
                .url(paasMediationAddress + "health")
                .get()
                .build();
        paasMediationUtils.doRequest(request, 200, null); // status must be 200
    }
}
