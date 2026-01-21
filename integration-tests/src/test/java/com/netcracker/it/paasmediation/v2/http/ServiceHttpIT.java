package com.netcracker.it.paasmediation.v2.http;

import com.fasterxml.jackson.core.type.TypeReference;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.domain.AnnotationResource;
import com.netcracker.it.paasmediation.v2.domain.MediationService;
import com.netcracker.it.paasmediation.v2.helpers.ServiceHelper;
import io.fabric8.kubernetes.api.model.Service;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import okhttp3.Response;
import org.hamcrest.Matchers;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.*;

@Slf4j
@SmokeTest
public class ServiceHttpIT extends ServiceHelper {

    @AfterEach
    public void cleanup() {
        super.cleanup();
    }

    @Test
    public void checkGetServiceAPI() throws IOException {
        log.info("Check get a service API");
        Service testService = createTestService(serviceName1);
        log.info("Created service {}", testService);
        paasUtils.createService(testService);
        assertNotNull(kubernetesClient.services().withName(serviceName1).get(), "Test service was not created");
        log.info("Service was created, try get it via paas-mediation API");
        MediationService mediationService = paasMediationUtils.doRequest(paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, serviceName1, namespace, "GET", null, null), 200, MediationService.class);
        assertNotNull(mediationService, "Created service was not found via paas-mediation API");
        log.info("Service received through paas-mediation {}", mediationService);
        checkService(testService, mediationService);

        MediationService[] array = paasMediationUtils.doRequest(paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "GET", null, "annotations=owner:Nectracker-company"), 200, MediationService[].class);
        assertEquals(1, array.length, "Created service was not found via paas-mediation API by annotation key:value");

        array = paasMediationUtils.doRequest(paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "GET", null, "annotations=owner:*"), 200, MediationService[].class);
        assertEquals(1, array.length, "Created service was not found via paas-mediation API by annotation name");

        array = paasMediationUtils.doRequest(paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "GET", null, "annotations=owner:Some-company"), 200, MediationService[].class);
        assertEquals(0, array.length, "Created service shouldn't be found via paas-mediation API by incorrect annotation value");
    }

    @Test
    public void checkGetServiceListAPI() throws IOException {
        log.info("Check get a service API");
        Service testService = createTestService(serviceName1);
        log.info("Created service {}", testService);
        paasUtils.createService(testService);
        assertNotNull(paasUtils.getServiceByName(serviceName1), "Test service was not created");
        log.info("Service was created, try get it via paas-mediation API");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "GET", null, null);
        List<MediationService> mediationServices = Arrays.asList(paasMediationUtils.doRequest(request, 200, MediationService[].class));
        log.info("Got list of services via paas-mediation API: {}", mediationServices);
        assertThat(mediationServices.size(), Matchers.greaterThanOrEqualTo(1));
        MediationService actualService = mediationServices.parallelStream()
                .filter(mediationService -> mediationService.getMetadata().getName().equals(serviceName1))
                .findAny()
                .orElse(null);
        assertNotNull(actualService, "Expected service was not returned through paas-mediation API");
        checkService(testService, actualService);
    }

    @Test
    public void checkCreateServiceAPI() throws IOException {
        log.info("Check create service API");
        Service expectedService = createTestService(serviceName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "POST", expectedService, null);
        log.info("Created service {}", expectedService);
        MediationService actualService = paasMediationUtils.doRequest(request, 201, MediationService.class);
        assertNotNull(actualService);
        assertNotNull(actualService, "Created service was not found via paas-mediation API");
        log.info("Service received through paas-mediation {}", actualService);
        checkService(expectedService, actualService);
    }

    @Test
    public void checkCreateServiceAPIWithConflictResponse() throws IOException {
        log.info("Check create service API");
        Service expectedService = createTestService(serviceName1);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "POST", expectedService, null);
        log.info("Created service {}", expectedService);
        MediationService actualService = paasMediationUtils.doRequest(request, 201, MediationService.class);
        assertNotNull(actualService);
        log.info("Service created {}", actualService);
        paasMediationUtils.doRequest(request, 409, null);
    }

    @Test
    public void checkDeleteServiceAPI() throws IOException {
        log.info("Check delete service API");
        Service testService = createTestService(serviceName1);
        paasUtils.createService(testService);
        assertNotNull(paasUtils.getServiceByName(serviceName1));
        log.info("Service was created {}", testService);

        log.info("Request to paas-mediation for deleting created service");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, serviceName1, namespace, "DELETE", testService, null);
        paasMediationUtils.doRequest(request, 200, null);
        log.info("Check that service was deleted");
        assertNull(paasUtils.getServiceByName(serviceName1), "Service was not deleted");
    }

    @Test
    public void checkDeleteServiceAPIWithNotFoundResponse() throws IOException {
        log.info("Check delete service API");
        Service testService = createTestService(serviceName1);

        log.info("Request to paas-mediation for deleting created service");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, serviceName1, namespace, "DELETE", testService, null);
        paasMediationUtils.doRequest(request, 404, null);
    }

    @Test
    public void checkUpdateServiceAPI() throws IOException {
        log.info("Check update service API");
        Service expectedService = createTestService(serviceName1);
        paasUtils.createService(expectedService);
        assertNotNull(paasUtils.getServiceByName(serviceName1));
        log.info("Service was created {}", expectedService);

        log.info("Request to paas-mediation for updating created service");
        expectedService.getSpec().getPorts().get(0).setPort(9090);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "PUT", expectedService, null);
        MediationService actualService = paasMediationUtils.doRequest(request, 200, MediationService.class);
        log.info("Check that service was updated");
        checkService(expectedService, actualService);
    }

    @Test
    public void checkUpdateServiceAPIToCreateService() throws IOException {
        log.info("Check update service API");
        Service expectedService = createTestService(serviceName1);

        log.info("Request to paas-mediation for creating service");
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.SERVICES, null, namespace, "PUT", expectedService, null);
        MediationService actualService = paasMediationUtils.doRequest(request, 200, MediationService.class);
        log.info("Check that service was updated");
        checkService(expectedService, actualService);
    }

    @Test
    void checkGetServiceAnnotation() throws IOException {
        log.info("Check get a service annotation resource");
        Service testService = createTestService(serviceName1);
        paasUtils.createService(testService);
        assertNotNull(paasUtils.getServiceByName(serviceName1), "Test service was not created");
        log.info("Serice was created {}", testService);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ANNOTATIONS, "", namespace, "GET", null, "annotation=owner&resourceType=Service");
        List<AnnotationResource> annotationResource;
        try (Response response = paasMediationUtils.doRequest(request)) {
            String respBody = response.body().string();
            log.info("Response: {}, body {}", response, respBody);
            assertEquals(200, response.code());
            annotationResource = objectMapper.readValue(respBody, new TypeReference<List<AnnotationResource>>() {
            });
        }

        assertNotNull(annotationResource, "Created service was not found via paas-mediation API");
        log.info("Route annotation resource received through paas-mediation {}", annotationResource);
        assertEquals(testService.getMetadata().getAnnotations().get("owner"), annotationResource.get(0).getAnnotationValue());
    }

    @Test
    void checkResourceTypeCase() throws IOException {
        log.info("Check get a service annotation resource");
        Service testService = createTestService(serviceName1);
        paasUtils.createService(testService);
        assertNotNull(paasUtils.getServiceByName(serviceName1), "Test service was not created");
        log.info("Service was created {}", testService);
        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ANNOTATIONS, "", namespace, "GET", null, "annotation=owner&resourceType=SERVICE");
        List<AnnotationResource> annotationResource;
        try (Response response = paasMediationUtils.doRequest(request)) {
            String respBody = response.body().string();
            log.info("Response: {}, body {}", response, respBody);
            assertEquals(200, response.code());
            annotationResource = objectMapper.readValue(respBody, new TypeReference<List<AnnotationResource>>() {
            });
        }

        assertNotNull(annotationResource, "Created service was not found via paas-mediation API");
        log.info("Route annotation resource received through paas-mediation {}", annotationResource);
        assertEquals(testService.getMetadata().getAnnotations().get("owner"), annotationResource.get(0).getAnnotationValue());
    }
}
