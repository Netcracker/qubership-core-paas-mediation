package com.netcracker.it.paasmediation.utils;


import io.fabric8.kubernetes.api.model.ConfigMap;
import io.fabric8.kubernetes.api.model.Pod;
import io.fabric8.kubernetes.api.model.Secret;
import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.api.model.apps.ReplicaSet;
import io.fabric8.kubernetes.api.model.networking.v1.Ingress;
import io.fabric8.kubernetes.client.KubernetesClient;
import lombok.extern.slf4j.Slf4j;

import java.util.Comparator;
import java.util.List;
import java.util.Set;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

@Slf4j
public class PaasUtils {

    private final KubernetesClient kubernetesClient;

    public PaasUtils(KubernetesClient kubernetesClient) {
        this.kubernetesClient = kubernetesClient;
    }

    public void deleteIngress(String name) {
        log.info("Start delete ingress {}", name);
        kubernetesClient.network().v1().ingresses().withName(name).withTimeout(1, TimeUnit.MINUTES).delete();
        log.info("Ingress {} was deleted", name);
    }

    public Ingress getIngressByName(String name) {
        return kubernetesClient.network().v1().ingresses().withName(name).get();
    }

    // Service

    public Service createService(Service service) {
        log.info("Start create service {}", service.getMetadata().getName());
        Service createdService = kubernetesClient.services().resource(service).create();
        log.info("Service was created {}", service);
        return createdService;
    }

    public Pod createPod(Pod pod) {
        log.info("Start create pod {}", pod.getMetadata().getName());
        Pod createdPod = kubernetesClient.pods().resource(pod).create();
        log.info("Pod was created {}", pod);
        return createdPod;
    }

    public void deleteService(String name) {
        log.info("Start delete service {}", name);
        kubernetesClient.services().withName(name).withTimeout(1, TimeUnit.MINUTES).delete();
        log.info("service was deleted {}", name);
    }

    public Service getServiceByName(String name) {
        return kubernetesClient.services().withName(name).get();
    }

    public Pod getPodByName(String name) {
        return kubernetesClient.pods().withName(name).get();
    }

    // Secret

    public Secret createSecret(Secret secret) {
        log.info("Start create secret {}", secret.getMetadata().getName());
        Secret createdSecret = kubernetesClient.secrets().resource(secret).create();
        log.info("Secret {} was created", secret.getMetadata().getName());
        return createdSecret;
    }

    public void deleteSecret(String name) {
        log.info("Start delete secret {}", name);
        kubernetesClient.secrets().withName(name).delete();
        log.info("Secret {} was deleted", name);

    }

    public Secret getSecretByName(String name) {
        return kubernetesClient.secrets().withName(name).get();
    }

    // ConfigMap
    public ConfigMap createConfigMap(ConfigMap configMap) {
        log.info("Start create configMap {}", configMap.getMetadata().getName());
        ConfigMap createdConfigMap = kubernetesClient.configMaps().resource(configMap).create();
        log.info("ConfigMap {} was created", createdConfigMap.getMetadata().getName());
        return createdConfigMap;
    }

    public void deleteConfigMap(String name) {
        log.info("Start delete configMap {}", name);
        kubernetesClient.configMaps().withName(name).delete();
        log.info("ConfigMap {} was deleted", name);
    }

    public ConfigMap getConfigMapByName(String name) {
        return kubernetesClient.configMaps().withName(name).get();
    }

    public String getImage(String name) {
        Deployment deployment = kubernetesClient.apps().deployments().withName(name).get();
        if (deployment != null)
            return deployment.getSpec().getTemplate().getSpec().getContainers().getFirst().getImage();
        return null;
    }

    public Set<String> getPodNamesByReplica(String replicaName) {
        return kubernetesClient.pods().list().getItems().stream()
                .filter(pod-> pod.getMetadata().getName().startsWith(replicaName) && !pod.getMetadata().getName().endsWith("deploy"))
                .map(pod-> pod.getMetadata().getName())
                .collect(Collectors.toSet());
    }

    public ReplicaSet getLatestReplicaSetByDeploymentName(String deploymentName) {
        return kubernetesClient.apps().replicaSets().list().getItems().stream()
                .filter(rs-> rs.getMetadata().getName().startsWith(deploymentName))
                .max(Comparator.comparing(rs -> Integer.parseInt(rs.getMetadata().getAnnotations().get("deployment.kubernetes.io/revision"))))
                .orElse(null);
    }

    public List<Pod> getAllPodsByNamespace(String namespace) {
        return kubernetesClient.pods().inNamespace(namespace).list().getItems();
    }

    public List<Deployment> getAllDeploymentsByNamespace(String namespace) {
        return kubernetesClient.apps().deployments().inNamespace(namespace).list().getItems();
    }
}
