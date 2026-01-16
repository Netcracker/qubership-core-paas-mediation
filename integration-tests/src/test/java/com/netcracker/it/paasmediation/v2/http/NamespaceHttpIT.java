package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationNamespace;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@Slf4j
@SmokeTest
public class NamespaceHttpIT extends PaasMediationParentV2Test {

    @Test
    public void checkGetNamespacesAPI() throws IOException {
        log.info("Check get a list namespaces");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.NAMESPACES, null, null, "GET", null, null);
        List<MediationNamespace> mediationNamespaces = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationNamespace[].class));
        assertNotNull(mediationNamespaces, "Can't get a list of namespaces");
        log.info("List of namespaces {}", mediationNamespaces);
        MediationNamespace expectNamespace = mediationNamespaces.parallelStream().filter(mediationNamespace -> mediationNamespace.getMetadata().getName().equals(namespace)).findAny().orElse(null);
        assertNotNull(expectNamespace, "Expected namespace was not returned");
        assertEquals(namespace, expectNamespace.getMetadata().getName());
    }

}
