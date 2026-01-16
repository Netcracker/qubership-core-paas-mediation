package com.netcracker.it.paasmediation.v2.helpers;

import com.fasterxml.jackson.dataformat.yaml.YAMLMapper;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.utils.PaasUtils;
import com.netcracker.it.paasmediation.v2.domain.DeploymentRequestBody;
import com.netcracker.it.paasmediation.v2.domain.DeploymentRolloutResponse;
import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.api.model.apps.ReplicaSet;
import io.fabric8.kubernetes.client.KubernetesClient;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;

import java.io.File;
import java.io.IOException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.text.MessageFormat;
import java.util.Set;
import java.util.concurrent.TimeUnit;

@Slf4j
public class DeploymentHelper {

    public static final int WAIT_FOR_DEPLOYMENT_CREATION_MIN = 2;
    private final PaasUtils paasUtils;
    private final KubernetesClient kubernetesClient;
    private final PaasMediationUtils paasMediationUtils;
    private final String namespace;

    public DeploymentHelper(KubernetesClient kubernetesClient, PaasUtils paasUtils, PaasMediationUtils paasMediationUtils, String namespace) {
        this.kubernetesClient = kubernetesClient;
        this.paasUtils = paasUtils;
        this.paasMediationUtils = paasMediationUtils;
        this.namespace = namespace;
    }

    public Set<String> getPodNamesByReplicaName(String replicaName) {
        return paasUtils.getPodNamesByReplica(replicaName);
    }

    public ReplicaSet getLatestReplicaSet(String deploymentName) {
        return paasUtils.getLatestReplicaSetByDeploymentName(deploymentName);
    }

    public String getImage(String deploymentName) {
        return paasUtils.getImage(deploymentName);
    }

    public String createDeployment(String name, String image) throws Exception {
        Deployment deployment = getDeploymentContentFromTemplate(name, image);
        String createdDeployment = kubernetesClient.apps().deployments().resource(deployment).create().getMetadata().getName();
        kubernetesClient.apps().deployments().withName(createdDeployment).waitUntilReady(WAIT_FOR_DEPLOYMENT_CREATION_MIN, TimeUnit.MINUTES);
        log.info("Successful created deployment={}", createdDeployment);
        return createdDeployment;
    }

    public DeploymentRolloutResponse rolloutDeployment(DeploymentRequestBody deploymentRequestBody) throws IOException {
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROLLOUT, null, namespace, "POST", deploymentRequestBody, null);
        DeploymentRolloutResponse deploymentRolloutResponse = paasMediationUtils.doRequest(request, 200, DeploymentRolloutResponse.class);
        log.info("Rollout response {}", deploymentRolloutResponse);
        return deploymentRolloutResponse;
    }

    public Deployment getDeploymentContentFromTemplate(String name, String image) throws Exception {
        String fileContent;
        URL fileUrl = getClass().getResource("/deployments/tests/deployment-template.yaml");
        try {
            fileContent = new String(Files.readAllBytes(Paths.get(new File(fileUrl.getFile()).toURI())));
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        String kubernetesVersion = kubernetesClient.getKubernetesVersion().getGitVersion();
        YAMLMapper mapper = new YAMLMapper();
        fileContent = MessageFormat.format(fileContent, "KUBERNETES", kubernetesVersion, namespace, image, name);
        return mapper.readValue(fileContent, Deployment.class);
    }
}
