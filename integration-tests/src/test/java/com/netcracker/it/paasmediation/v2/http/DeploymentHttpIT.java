package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationFullDeploymentInfo;
import io.fabric8.kubernetes.api.model.HasMetadata;
import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.api.model.apps.DeploymentStatus;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.TimeUnit;

import static org.junit.jupiter.api.Assertions.*;

@Slf4j
@SmokeTest
public class DeploymentHttpIT extends PaasMediationParentV2Test {

    public static final String paasMediationName = "paas-mediation";
    public static final String deploymentName = "paas-mediation-deployment-get-test";

    @BeforeEach
    @AfterEach
    public void cleanup() {
        kubernetesClient.apps().deployments().withName(deploymentName).withTimeout(1, TimeUnit.MINUTES).delete();
    }

    @Test
    public void checkGetDeploymentListAPI() throws IOException {
        log.info("Check get a list deployments");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.DEPLOYMENTS, null, namespace, "GET", null, "labels=app.kubernetes.io/part-of:Cloud-Core");
        List<MediationFullDeploymentInfo> mediationDeployments = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationFullDeploymentInfo[].class));
        assertNotNull(mediationDeployments, "Can't get a list of deployments");
        log.debug("List of deployments {}", mediationDeployments);
        checkGetDeploymentListAPI(mediationDeployments, getAllDeploymentsByNamespace(namespace));
    }

    @Test
    void checkGetDeploymentByNameAPI() throws Exception {
        log.info("Check deployment");
        deploymentHelper.createDeployment(deploymentName, deploymentHelper.getImage(paasMediationName));
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.DEPLOYMENTS, deploymentName, namespace, "GET", null, null);
        MediationFullDeploymentInfo mediationDeployment = paasMediationUtils.doRequest(request, 200, MediationFullDeploymentInfo.class);
        log.debug("Deployment {}", mediationDeployment);
        checkGetDeploymentAPI(mediationDeployment, getAllDeploymentsByNamespace(namespace));
    }

    protected void checkGetDeploymentListAPI(List<MediationFullDeploymentInfo> mediationDeployments, List<HasMetadata> expectedListOfDeployments) {
        assertNotNull(expectedListOfDeployments, "Expected deployment names was not returned");
        for (MediationFullDeploymentInfo mediationDpl : mediationDeployments) {
            assertTrue(expectedListOfDeployments.stream().anyMatch(dpl -> dpl.getMetadata().getName().equals(mediationDpl.getMetadata().getName())),
                    mediationDpl.getMetadata().getName() + " that deployment expected but not found");
            checkGetDeploymentAPI(mediationDpl, expectedListOfDeployments);
        }
    }

    protected void compareDeployment(Deployment expectedDpl, MediationFullDeploymentInfo mediationDpl) {
        assertEquals(expectedDpl.getSpec().getTemplate().getSpec().getDnsPolicy(), mediationDpl.getSpec().getTemplate().getSpec().getDnsPolicy());
        assertEquals(expectedDpl.getSpec().getTemplate().getSpec().getNodeName(), mediationDpl.getSpec().getTemplate().getSpec().getNodeName());
        assertEquals(expectedDpl.getSpec().getTemplate().getSpec().getRestartPolicy(), mediationDpl.getSpec().getTemplate().getSpec().getRestartPolicy());
        assertEquals(expectedDpl.getSpec().getTemplate().getSpec().getTerminationGracePeriodSeconds(), mediationDpl.getSpec().getTemplate().getSpec().getTerminationGracePeriodSeconds());

        assertEquals(Optional.ofNullable(expectedDpl.getStatus()).map(DeploymentStatus::getReplicas).orElse(0),
                Optional.ofNullable(mediationDpl.getStatus()).map(MediationFullDeploymentInfo.DeploymentStatus::getReplicas).orElse(0));
        assertEquals(Optional.ofNullable(expectedDpl.getStatus()).map(DeploymentStatus::getReadyReplicas).orElse(0),
                Optional.ofNullable(mediationDpl.getStatus()).map(MediationFullDeploymentInfo.DeploymentStatus::getReadyReplicas).orElse(0));
        assertEquals(Optional.ofNullable(expectedDpl.getStatus()).map(DeploymentStatus::getConditions).map(List::size).orElse(0),
                Optional.ofNullable(mediationDpl.getStatus()).map(MediationFullDeploymentInfo.DeploymentStatus::getConditions).map(List::size).orElse(0));
    }

    protected void checkGetDeploymentAPI(MediationFullDeploymentInfo mediationDeployment, List<HasMetadata> expectedListOfDeployments) {
        assertNotNull(expectedListOfDeployments, "Expected deployment names was not returned");
        assertTrue(expectedListOfDeployments.stream().anyMatch(dpl -> dpl.getMetadata().getName().equals(mediationDeployment.getMetadata().getName())),
                mediationDeployment.getMetadata().getName() + " that deployment expected but not found");

        for (HasMetadata expectedDpl : expectedListOfDeployments) {
            if (mediationDeployment.getMetadata().getName().equals(expectedDpl.getMetadata().getName())) {
                compareDeployment((Deployment) expectedDpl, mediationDeployment);
            }
        }
    }

    protected List<HasMetadata> getAllDeploymentsByNamespace(String namespace) {
        return new ArrayList<>(paasUtils.getAllDeploymentsByNamespace(namespace));
    }
}
