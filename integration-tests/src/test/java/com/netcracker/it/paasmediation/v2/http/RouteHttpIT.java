package com.netcracker.it.paasmediation.v2.http;

import com.fasterxml.jackson.core.type.TypeReference;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.PortForward;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.SmokeTest;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.Value;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.domain.AnnotationResource;
import com.netcracker.it.paasmediation.v2.domain.HealthProbe;
import com.netcracker.it.paasmediation.v2.domain.MediationRoute;
import com.netcracker.it.paasmediation.v2.helpers.RouteHelper;
import io.fabric8.kubernetes.api.model.networking.v1.Ingress;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import okhttp3.Response;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

import java.io.IOException;
import java.net.URL;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;


@Slf4j
@SmokeTest
public class RouteHttpIT extends RouteHelper {

    @PortForward(serviceName = @Value("paas-mediation"))
    private static URL paasMediationService;

    @Nested
    @TestInstance(TestInstance.Lifecycle.PER_CLASS)
    class CreateRoute {
        MediationRoute mediationRoute;
        MediationRoute createdRoute;
        Ingress actualRoute;

        @BeforeAll
        void createRoute() throws Exception {
            Ingress testRoute = createTestRoute(routeName1);
            mediationRoute = new MediationRoute(testRoute);
            Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "POST", mediationRoute, null);
            createdRoute = paasMediationUtils.doRequest(request, 201, MediationRoute.class);
            assertNotNull(createdRoute);
            log.info("Route created {}", createdRoute);
            actualRoute = paasUtils.getIngressByName(routeName1);
            assertNotNull(actualRoute);
        }

        @Test
        void testCreatedRoute() {
            checkServiceRoutes(mediationRoute, createdRoute);
            Ingress actualRoute = paasUtils.getIngressByName(routeName1);
            log.info("Actual route {}", actualRoute);
            checkRoute(actualRoute, mediationRoute);
        }

        @Test
        public void checkGetRouteAPI() throws IOException {
            Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, routeName1, namespace, "GET", null, null);
            MediationRoute mediationRoute = paasMediationUtils.doRequest(request, 200, MediationRoute.class);
            assertNotNull(mediationRoute);
            checkRoute(actualRoute, mediationRoute);
        }

        @Test
        public void checkGetRouteListAPI() throws IOException {
            Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "GET", null, null);
            MediationRoute mediationRoute = Arrays.stream(paasMediationUtils.doRequest(request, 200, MediationRoute[].class))
                    .filter(paasRoute -> paasRoute.getMetadata().getName().equals(routeName1)).findFirst().orElse(null);
            assertNotNull(mediationRoute, String.format("Test route '%s' not present in LIST response", routeName1));
            log.info("Route {}", mediationRoute);
            checkRoute(actualRoute, mediationRoute);
        }

        @Test
        void checkGetRouteAnnotation() throws IOException {
            Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ANNOTATIONS, "", namespace, "GET", null, "annotation=owner&resourceType=Route");
            List<AnnotationResource> annotationResource;
            try (Response response = paasMediationUtils.doRequest(request)) {
                String respBody = response.body().string();
                log.info("Response: {}, body {}", response, respBody);
                assertEquals(200, response.code());
                annotationResource = objectMapper.readValue(respBody, new TypeReference<>() {
                });
            }
            assertNotNull(annotationResource, "Created route was not found via paas-mediation API");
            log.info("Route annotation resource received through paas-mediation {}", annotationResource);
            assertEquals(actualRoute.getMetadata().getAnnotations().get("owner"), annotationResource.get(0).getAnnotationValue());
        }

        @Test
        public void checkCreateRouteAPIWithConflictResponse() throws IOException {
            Ingress testRoute = createTestRoute(routeName1);
            MediationRoute mediationRoute = new MediationRoute(testRoute);
            Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "POST", mediationRoute, null);
            paasMediationUtils.doRequest(request, 409, null);
        }

        @Nested
        @TestInstance(TestInstance.Lifecycle.PER_CLASS)
        class Update {
            Ingress updatedRoute;
            MediationRoute mediationUpdatedRoute;

            @BeforeAll
            void updateRoute() throws Exception {
                updatedRoute = actualRoute.toBuilder().editSpec()
                        .editFirstRule()
                        .editHttp()
                        .editFirstPath()
                        .editBackend()
                        .editService().withName("point-to-endless").endService()
                        .endBackend()
                        .endPath()
                        .endHttp()
                        .endRule()
                        .endSpec()
                        .build();
                log.info("Request to paas-mediation for updating created route");
                MediationRoute updatedMediationRoute = new MediationRoute(updatedRoute);
                Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "PUT", updatedMediationRoute, null);
                mediationUpdatedRoute = paasMediationUtils.doRequestWithRetry(request, 200, 10, MediationRoute.class);
            }

            @Test
            public void checkUpdatedRouteAPI() {
                checkRoute(updatedRoute, mediationUpdatedRoute);
            }

            @Nested
            @TestInstance(TestInstance.Lifecycle.PER_CLASS)
            class Delete {

                @BeforeAll
                void deleteRoute() throws Exception {
                    log.info("Request to paas-mediation for deleting created route");
                    Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, routeName1, namespace, "DELETE", null, null);
                    paasMediationUtils.doRequest(request, 200, null);
                }

                @Test
                public void checkRouteNotFoundAfterDeletion() {
                    log.info("Check that route was deleted");
                    assertNull(paasUtils.getIngressByName(routeName1), "Route was not deleted");
                }

                @Test
                public void checkDeleteRouteAPIWithNotFoundResponse() throws IOException {
                    log.info("Request to paas-mediation for deleting route");
                    Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, routeName1, namespace, "DELETE", null, null);
                    paasMediationUtils.doRequest(request, 404, null);
                }

                @Nested
                @TestInstance(TestInstance.Lifecycle.PER_CLASS)
                class CreateOrUpdate {
                    @BeforeAll
                    void createOrUpdateRoute() throws Exception {
                        log.info("Request to paas-mediation for updating/creating route");
                        Request request = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "PUT", mediationRoute, null);
                        createdRoute = paasMediationUtils.doRequestWithRetry(request, 200, 10, MediationRoute.class);
                        assertNotNull(createdRoute);
                        actualRoute = paasUtils.getIngressByName(routeName1);
                        assertNotNull(actualRoute);
                    }

                    @Test
                    public void checkCreatedOrUpdatedRoute() {
                        log.info("Actual route {}", actualRoute);
                        checkRoute(actualRoute, createdRoute);
                    }
                }
            }
        }
    }

    @Test
    public void checkBadRoutesAPI() throws IOException {
        Request request = new Request.Builder().url(paasMediationService.toString() + "health").get().build();

        HealthProbe healthProbe = paasMediationUtils.doRequest(request, 200, HealthProbe.class);
        assertEquals("UP", healthProbe.getStatus());
        assertNotNull(healthProbe.getBadResourcesHealthCheck());
        assertNotNull(healthProbe.getBadResourcesHealthCheck().getBadRoutes());
        assertEquals(0, healthProbe.getBadResourcesHealthCheck().getBadRoutes().size());
    }
}