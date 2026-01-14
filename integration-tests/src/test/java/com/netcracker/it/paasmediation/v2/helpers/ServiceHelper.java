package com.netcracker.it.paasmediation.v2.helpers;

import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationService;
import io.fabric8.kubernetes.api.model.IntOrString;
import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.ServicePort;
import io.fabric8.kubernetes.api.model.ServiceSpec;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

public abstract class ServiceHelper extends PaasMediationParentV2Test {

    protected String serviceNamePrefix = "paas-mediation-it-test-service";
    protected String serviceName1 = serviceNamePrefix + "-1";
    protected String serviceName2 = serviceNamePrefix + "-2";


    @BeforeAll
    @AfterAll
    public void cleanup() {
        Stream.of(serviceName1, serviceName2).forEach(name -> {
            if (kubernetesClient.services().withName(name).get() != null) {
                kubernetesClient.services().withName(name).withTimeout(1, TimeUnit.MINUTES).delete();
                assertNull(kubernetesClient.services().withName(name).get());
            }
        });
    }

    protected void checkService(Service expectService, MediationService actualService) {
        assertEquals(expectService.getKind(), actualService.getMetadata().getKind());
        assertEquals(expectService.getMetadata().getName(), actualService.getMetadata().getName());
        assertEquals(expectService.getMetadata().getNamespace(), actualService.getMetadata().getNamespace());
        assertEquals(expectService.getMetadata().getLabels(), actualService.getMetadata().getLabels());
        assertTrue(actualService.getMetadata().getAnnotations().entrySet().containsAll(expectService.getMetadata().getAnnotations().entrySet()));
        assertTrue(actualService.getSpec().getSelector().entrySet().containsAll(expectService.getSpec().getSelector().entrySet()));
        MediationService.Port actualPort = actualService.getSpec().getPorts().get(0);
        ServicePort expectPort = expectService.getSpec().getPorts().get(0);
        assertEquals(expectPort.getName(), actualPort.getName());
        assertEquals(expectPort.getProtocol(), actualPort.getProtocol());
        assertEquals(expectPort.getPort(), actualPort.getPort());
        assertEquals(expectPort.getTargetPort().getIntVal(), actualPort.getTargetPort());
    }

    protected Service createTestService(String name) {
        ServiceSpec spec = new ServiceSpec();
        Map<String, String> selector = new HashMap<>();
        selector.put("name", "paas-mediation-service-it");
        spec.setSelector(selector);

        List<ServicePort> ports = new ArrayList<>();
        ServicePort servicePort = new ServicePort();
        servicePort.setName("web");
        servicePort.setProtocol("TCP");
        servicePort.setPort(8080);
        servicePort.setTargetPort(new IntOrString(8080));
        ports.add(servicePort);
        spec.setPorts(ports);

        Service service = new Service();
        service.setKind("Service");
        service.setSpec(spec);
        service.setMetadata(createTestMetadata(name));
        return service;
    }

}
