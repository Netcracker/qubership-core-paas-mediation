package com.netcracker.it.paasmediation.v2.helpers;

import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationPod;
import io.fabric8.kubernetes.api.model.Container;
import io.fabric8.kubernetes.api.model.ContainerPort;
import io.fabric8.kubernetes.api.model.Pod;
import io.fabric8.kubernetes.api.model.PodSpec;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;

import static org.junit.jupiter.api.Assertions.*;

public abstract class PodHelper extends PaasMediationParentV2Test {

    protected String podName = "paas-mediation-pod-it-test";
    protected String containerName = "paas-mediation-container-it-test";
    protected String imageName = "paas-mediation-image-it-test";

    @AfterEach
    public void tearDown() {
        kubernetesClient.pods().withName(podName).withTimeout(1, TimeUnit.MINUTES).delete();
    }

    @BeforeEach
    public void tearUp() {
        // we must be sure that pod with checking name does not exist
        kubernetesClient.pods().withName(podName).withTimeout(1, TimeUnit.MINUTES).delete();
        assertNull(paasUtils.getPodByName(podName));
    }

    protected void checkPod(Pod expectPod, MediationPod actualPod) {
        assertEquals(expectPod.getKind(), actualPod.getMetadata().getKind());
        assertEquals(expectPod.getMetadata().getName(), actualPod.getMetadata().getName());
        assertEquals(expectPod.getMetadata().getNamespace(), actualPod.getMetadata().getNamespace());
        assertEquals(expectPod.getMetadata().getLabels(), actualPod.getMetadata().getLabels());
        assertTrue(actualPod.getMetadata().getAnnotations().entrySet().containsAll(expectPod.getMetadata().getAnnotations().entrySet()));
        MediationPod.Container actualContainer = actualPod.getSpec().getContainers().get(0);
        Container expectedContainer = expectPod.getSpec().getContainers().get(0);
        assertEquals(expectedContainer.getName(), actualContainer.getName());
        assertEquals(expectedContainer.getImage(), actualContainer.getImage());
        MediationPod.Container.Port actualPort = actualContainer.getPorts().get(0);
        ContainerPort expectPort = expectedContainer.getPorts().get(0);
        assertEquals(expectPort.getName(), actualPort.getName());
        assertEquals(expectPort.getProtocol(), actualPort.getProtocol());
        assertEquals(expectPort.getContainerPort(), actualPort.getContainerPort());
    }

    protected Pod createTestPod(String name) {
        PodSpec spec = new PodSpec();
        Map<String, String> selector = new HashMap<>();
        selector.put("name", "paas-mediation-pod-it");
        spec.setNodeSelector(selector);

        List<ContainerPort> ports = new ArrayList<>();
        ContainerPort containerPort = new ContainerPort();
        containerPort.setName("web");
        containerPort.setProtocol("TCP");
        containerPort.setContainerPort(8881);
        ports.add(containerPort);
        List<Container> containers = new ArrayList<>();
        Container container = new Container();
        container.setPorts(ports);
        container.setName(containerName);
        container.setImage(imageName);
        containers.add(container);
        spec.setContainers(containers);

        Pod pod = new Pod();
        pod.setKind("Pod");
        pod.setSpec(spec);
        pod.setMetadata(createTestMetadata(name));
        return pod;
    }

}
