package com.netcracker.it.paasmediation.v2.ws;

import com.netcracker.it.paasmediation.v2.domain.DeploymentRequestBody;
import com.netcracker.it.paasmediation.v2.domain.DeploymentRolloutResponse;
import com.netcracker.it.paasmediation.v2.domain.DeploymentRolloutWSResponse;
import com.netcracker.it.paasmediation.v2.http.AbstractRolloutDeployment;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;


@Slf4j
@Tag("watch")
public class RolloutDeploymentWebSocketIT extends AbstractRolloutDeployment {

    protected WSListener webSocketListener;

    @AfterEach
    public void afterTest() {
        if (webSocketListener != null) {
            webSocketListener.close();
        }
    }

    @Test
    public void checkWebSocketRestartPodsEvent() throws Exception {
        String createdDeployment1 = null;
        String createdDeployment2 = null;
        try {
            String deploymentImage = deploymentHelper.getImage("paas-mediation");
            log.info("Start creating Deployment={}", deploymentName1);
            createdDeployment1 = deploymentHelper.createDeployment(deploymentName1, deploymentImage);

            log.info("Start creating Deployment={}", deploymentName2);
            createdDeployment2 = deploymentHelper.createDeployment(deploymentName2, deploymentImage);

            Watcher<DeploymentRolloutWSResponse> closeWatcher = new Watcher<>(DeploymentRolloutWSResponse.class, 1,
                    "CLOSE_CONTROL_MESSAGE"::equals, event -> {
                DeploymentRolloutWSResponse deploymentMap = event.getObject();
                return checkCorrectPods(deploymentHelper.getLatestReplicaSet(deploymentName1).getMetadata().getName(),
                        deploymentMap.getOrDefault(deploymentName1, List.of())) &&
                       checkCorrectPods(deploymentHelper.getLatestReplicaSet(deploymentName2).getMetadata().getName(),
                               deploymentMap.getOrDefault(deploymentName2, List.of()));
            });


            List<String> deploymentsList = Arrays.asList(deploymentName1, deploymentName2);
            DeploymentRequestBody deploymentRequestBody = new DeploymentRequestBody(deploymentsList);
            DeploymentRolloutResponse deploymentRolloutResponse = deploymentHelper.rolloutDeployment(deploymentRequestBody);

            log.info("Create webSocket");
            Request request = paasMediationUtils.createWsRequest(deploymentRolloutResponse.getPodStatusWebSocket().substring(1));
            webSocketListener = new WSListener(okHttpClient, request, closeWatcher);

            // we must be sure that connect is established
            Assertions.assertDoesNotThrow(() -> webSocketListener.waitConnected(10, TimeUnit.SECONDS));
            Assertions.assertDoesNotThrow(() -> closeWatcher.waitMessagesReceived(60, TimeUnit.SECONDS));
        } finally {
            if (createdDeployment1 != null) {
                kubernetesClient.apps().deployments().withName(createdDeployment1).delete();
                kubernetesClient.apps().deployments().withName(createdDeployment1).waitUntilCondition(Objects::isNull, 1, TimeUnit.MINUTES);
            }
            if (createdDeployment2 != null) {
                kubernetesClient.apps().deployments().withName(createdDeployment2).delete();
                kubernetesClient.apps().deployments().withName(createdDeployment2).waitUntilCondition(Objects::isNull, 1, TimeUnit.MINUTES);
            }
        }
    }

    private boolean checkCorrectPods(String replicaName, List<DeploymentRolloutWSResponse.WSResponsePod> mediationPodList) {
        Set<String> podsNamesExpected = deploymentHelper.getPodNamesByReplicaName(replicaName);
        Set<String> podsNamesActual = mediationPodList.stream().map(DeploymentRolloutWSResponse.WSResponsePod::getName).collect(Collectors.toSet());
        return Objects.equals(podsNamesExpected, podsNamesActual);
    }
}
