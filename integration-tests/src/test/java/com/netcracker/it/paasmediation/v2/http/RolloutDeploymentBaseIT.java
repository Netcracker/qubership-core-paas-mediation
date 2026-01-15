package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.v2.domain.DeploymentRequestBody;
import com.netcracker.it.paasmediation.v2.domain.DeploymentRolloutResponse;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.TimeUnit;
import java.util.function.BiFunction;

import static org.junit.jupiter.api.Assertions.*;

@Slf4j
@SmokeTest
public class RolloutDeploymentBaseIT extends AbstractRolloutDeployment {


    @Test
    public void checkRestartDeployment() throws Exception {
        String createdDeploymentName = null;
        try {
            log.debug("Start create deployment {}", deploymentName1);
            createdDeploymentName = deploymentHelper.createDeployment(deploymentName1, deploymentHelper.getImage("paas-mediation"));

            String firstReplicaSetName = deploymentHelper.getLatestReplicaSet(deploymentName1).getMetadata().getName();
            log.debug("Last replicaSet={} from deployment={}", firstReplicaSetName, deploymentName1);

            List<String> deploymentsList = Collections.singletonList(deploymentName1);
            DeploymentRolloutResponse deploymentRolloutResponse = deploymentHelper.rolloutDeployment(new DeploymentRequestBody(deploymentsList));
            DeploymentRolloutResponse.Deployment responseReplicas = deploymentRolloutResponse.getDeployments().get(0).entrySet().iterator().next().getValue();

            String secondReplicaSetName = deploymentHelper.getLatestReplicaSet(deploymentName1).getMetadata().getName();
            log.debug("Last replicaSet={} from deployment={} after restart", secondReplicaSetName, deploymentName1);

            assertEquals(firstReplicaSetName, responseReplicas.getActive(), "ReplicaSet differ from the response");
            assertEquals(secondReplicaSetName, responseReplicas.getRolling(), "ReplicaSet differ from the response");
        } finally {
            if (createdDeploymentName != null) {
                kubernetesClient.apps().deployments().withName(createdDeploymentName).delete();
                assertNull(kubernetesClient.apps().deployments().withName(deploymentName1).waitUntilCondition(Objects::isNull, 1, TimeUnit.MINUTES));
            }
        }
    }

    @Test
    public void checkBulkRestartDeployments() throws Exception {
        log.debug("Start check that platform is openshift");
        String createdDeploymentName1 = null;
        String createdDeploymentName2 = null;
        try {
            String deploymentImage = deploymentHelper.getImage("paas-mediation");
            log.debug("Start create deployment {}", deploymentName1);
            createdDeploymentName1 = deploymentHelper.createDeployment(deploymentName1, deploymentImage);

            String firstReplicaSetName1 = deploymentHelper.getLatestReplicaSet(deploymentName1).getMetadata().getName();
            log.info("Last replicaSet={} from deployment={}", firstReplicaSetName1, deploymentName1);

            log.debug("Start create deployment {}", deploymentName1);
            createdDeploymentName2 = deploymentHelper.createDeployment(deploymentName2, deploymentImage);

            String firstReplicaSetName2 = deploymentHelper.getLatestReplicaSet(deploymentName2).getMetadata().getName();
            log.info("Last replicaSet={} from deployment={}", firstReplicaSetName2, deploymentName2);

            List<String> deploymentsList = Arrays.asList(deploymentName1, deploymentName2);
            DeploymentRolloutResponse deploymentRolloutResponse = deploymentHelper.rolloutDeployment(new DeploymentRequestBody(deploymentsList));

            String secondReplicaSetName1 = deploymentHelper.getLatestReplicaSet(deploymentName1).getMetadata().getName();
            log.info("Last replicaSet={} from deployment={} after restart", secondReplicaSetName1, deploymentName1);

            String secondReplicaSetName2 = deploymentHelper.getLatestReplicaSet(deploymentName2).getMetadata().getName();
            log.info("Last replicaSet={} from deployment={} after restart", secondReplicaSetName2, deploymentName2);

            BiFunction<String, String, Boolean> validateReplicasFunc = (String name1, String name2) ->
                    deploymentRolloutResponse.getDeployments().stream().anyMatch(deploymentMap -> {
                        DeploymentRolloutResponse.Deployment responseReplicas = deploymentMap.entrySet().iterator().next().getValue();
                        return Objects.equals(name1, responseReplicas.getActive()) && Objects.equals(name2, responseReplicas.getRolling());
                    });
            assertTrue(validateReplicasFunc.apply(firstReplicaSetName1, secondReplicaSetName1), "ReplicaSet differ from the response");
            assertTrue(validateReplicasFunc.apply(firstReplicaSetName2, secondReplicaSetName2), "ReplicaSet differ from the response");
        } finally {
            if (createdDeploymentName1 != null) {
                kubernetesClient.apps().deployments().withName(createdDeploymentName1).delete();
                assertNull(kubernetesClient.apps().deployments().withName(deploymentName1).waitUntilCondition(Objects::isNull, 1, TimeUnit.MINUTES));
            }
            if (createdDeploymentName2 != null) {
                kubernetesClient.apps().deployments().withName(createdDeploymentName2).delete();
                assertNull(kubernetesClient.apps().deployments().withName(deploymentName2).waitUntilCondition(Objects::isNull, 1, TimeUnit.MINUTES));
            }
        }
    }
}
