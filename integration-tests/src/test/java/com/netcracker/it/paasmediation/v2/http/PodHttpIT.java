package com.netcracker.it.paasmediation.v2.http;

import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.domain.MediationFullPodInfo;
import com.netcracker.it.paasmediation.v2.domain.MediationPod;
import com.netcracker.it.paasmediation.v2.helpers.PodHelper;
import io.fabric8.kubernetes.api.model.Pod;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.hamcrest.Matchers;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.*;

@Slf4j
@SmokeTest
public class PodHttpIT extends PodHelper {


        private final String ANNOTATIONS = "annotations";
        private final String LABELS = "labels";

    @Test
    public void checkGetPodListAPI() throws IOException {
        log.info("Check get a list pods");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.PODS, null, namespace, "GET", null, null);
        List<MediationFullPodInfo> mediationPods = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationFullPodInfo[].class));
        assertNotNull(mediationPods, "Can't get a list of pods");
        log.debug("List of pods {}", mediationPods);
        List<Pod> expectedListOfPods = paasUtils.getAllPodsByNamespace(namespace);
        assertNotNull(expectedListOfPods, "Expected pod names was not returned");
        for(MediationFullPodInfo mediationPod : mediationPods){
            assertTrue(expectedListOfPods.stream().anyMatch(pod -> pod.getMetadata().getName().equals(mediationPod.getMetadata().getName())),
                    mediationPod.getMetadata().getName() + " that pod expected but not found");

        }
        for(MediationFullPodInfo mediationPod : mediationPods){
            for(Pod expectedPod : expectedListOfPods){
                if(mediationPod.getMetadata().getName().equals(expectedPod.getMetadata().getName())){
                    assertEquals(expectedPod.getSpec().getDnsPolicy(), mediationPod.getSpec().getDnsPolicy());
                    assertEquals(expectedPod.getSpec().getNodeName(), mediationPod.getSpec().getNodeName());
                    assertEquals(expectedPod.getSpec().getRestartPolicy(), mediationPod.getSpec().getRestartPolicy());
                    assertEquals(expectedPod.getSpec().getTerminationGracePeriodSeconds(), mediationPod.getSpec().getTerminationGracePeriodSeconds());

                    assertEquals(expectedPod.getStatus().getHostIP(), mediationPod.getStatus().getHostIP());
                    assertEquals(expectedPod.getStatus().getPodIP(), mediationPod.getStatus().getPodIP());
                    assertEquals(expectedPod.getStatus().getPhase(), mediationPod.getStatus().getPhase());
                }
            }
        }
    }

    @Test
    public void checkGetPodAPI() throws IOException {
        log.info("Check get a pod");
        Pod testPod = createTestPod(podName);
        log.info("Created pod {}", testPod);
        paasUtils.createPod(testPod);
        assertNotNull(paasUtils.getPodByName(podName), "Test pod was not created");
        log.info("Pod was created, try get it via paas-mediation API");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.PODS, null, namespace, "GET", null, null);
        List<MediationPod> mediationPods = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationPod[].class));
        log.info("Got list of pods via paas-mediation API: {}", mediationPods);
        assertThat(mediationPods.size(), Matchers.greaterThanOrEqualTo(1));
        MediationPod actualPod = mediationPods.parallelStream()
                .filter(mediationPod -> mediationPod.getMetadata().getName().equals(podName))
                .findAny()
                .orElse(null);
        assertNotNull(actualPod, "Expected pod was not returned through paas-mediation API");
        checkPod(testPod, actualPod);
    }

    @Test
    public void checkGetPodListAPIFilteredByLabel() throws IOException {
        checkGetPodListAPIFiltered(LABELS, "it", "paas-mediation");
    }

    @Test
    public void checkGetPodListAPIFilteredByAnnotation() throws IOException {
        checkGetPodListAPIFiltered(ANNOTATIONS, "owner", "Nectracker-company");
    }

    private void checkGetPodListAPIFiltered(String paramName, String paramKey, String paramValue) throws IOException {
        log.info("Check get a filtered list of pods");
        Pod testPod = createTestPod(podName);
        log.info("Created pod {}", testPod);
        paasUtils.createPod(testPod);
        assertNotNull(paasUtils.getPodByName(podName), "Test pod was not created");
        log.info("Pod was created, try get it via paas-mediation API");
        String fullParameter = paramName + "=" + paramKey + ":" +paramValue;
        log.info("Get filtered list of pods by {}", fullParameter);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.PODS, null, namespace, "GET", null, fullParameter);
        List<MediationPod> mediationPods = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationPod[].class));
        log.info("Got list of pods via paas-mediation API: {}", mediationPods);
        assertThat(mediationPods.size(), Matchers.greaterThanOrEqualTo(1));

        MediationPod actualPod = null;

        switch (paramName) {
            case LABELS: actualPod = mediationPods.parallelStream()
                    .filter(mediationPod -> mediationPod.getMetadata().getLabels().get(paramKey).equals(paramValue))
                    .findAny()
                    .orElse(null);
            break;
            case ANNOTATIONS: actualPod = mediationPods.parallelStream()
                    .filter(mediationPod -> mediationPod.getMetadata().getAnnotations().get(paramKey).equals(paramValue))
                    .findAny()
                    .orElse(null);
            break;
        }
        assertNotNull(actualPod, "Expected pod was not returned through paas-mediation API");
        checkPod(testPod, actualPod);
    }
}
