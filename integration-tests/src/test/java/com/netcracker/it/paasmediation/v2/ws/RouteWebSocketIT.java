package com.netcracker.it.paasmediation.v2.ws;

import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.utils.PaasRequestFilter;
import com.netcracker.it.paasmediation.v2.domain.MediationRoute;
import com.netcracker.it.paasmediation.v2.helpers.RouteHelper;
import io.fabric8.kubernetes.api.model.networking.v1.Ingress;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.*;

import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Stream;

@Slf4j
@Tag("watch")
public class RouteWebSocketIT extends RouteHelper {

    private WSListener wsListener;
    private WSListener wsListenerWithFilter;

    private final Watcher<MediationRoute> onAddedWatcher = new Watcher<>(MediationRoute.class, 1,
            "ADDED"::equals, event -> Objects.equals(event.getObject().getMetadata().getName(), routeName1));
    private final Watcher<MediationRoute> onDeletedWatcher = new Watcher<>(MediationRoute.class, 1,
            "DELETED"::equals, event -> Objects.equals(event.getObject().getMetadata().getName(), routeName1));
    private final Watcher<MediationRoute> onAddedWatcherWithFilter = new Watcher<>(MediationRoute.class, 1,
            "ADDED"::equals, event -> event.getObject().getMetadata().getName().startsWith(routeNamePrefix));
    private final Watcher<MediationRoute> onDeletedWatcherWithFilter = new Watcher<>(MediationRoute.class, 1,
            "DELETED"::equals, event -> event.getObject().getMetadata().getName().startsWith(routeNamePrefix));
    String label_1 = "1";
    String label_1_val = "1";
    String label_2 = "2";
    String label_2_val = "2";

    @BeforeAll
    void connect() throws Exception {
        Request request1 = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.ROUTES, namespace);
        wsListener = new WSListener(okHttpClient, request1, onAddedWatcher, onDeletedWatcher);
        wsListener.waitConnected(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);

        PaasRequestFilter filter = new PaasRequestFilter().withLabel(label_2, label_2_val);
        Request request2 = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.ROUTES, namespace, filter);
        wsListenerWithFilter = new WSListener(okHttpClient, request2, onAddedWatcherWithFilter, onDeletedWatcherWithFilter);
        wsListenerWithFilter.waitConnected(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
    }

    @AfterAll
    public void afterTest() {
        Stream.of(wsListener, wsListenerWithFilter).filter(Objects::nonNull).forEach(WSListener::close);
    }

    @Nested
    @TestInstance(TestInstance.Lifecycle.PER_CLASS)
    class AfterCreated {
        Ingress testRoute1;
        Ingress testRoute2;

        @BeforeAll
        void createRoute() throws Exception {
            testRoute1 = createTestRoute(routeName1);
            Map<String, String> labels_1 = new HashMap<>();
            labels_1.put(label_1, label_1_val);
            testRoute1.getMetadata().setLabels(labels_1);

            MediationRoute mediationRoute1 = new MediationRoute(testRoute1);
            Request request1 = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "POST", mediationRoute1, null);
            MediationRoute createdRoute1 = paasMediationUtils.doRequest(request1, 201, MediationRoute.class);
            Assertions.assertNotNull(createdRoute1);

            testRoute2 = createTestRoute(routeName2);
            Map<String, String> labels_2 = new HashMap<>();
            labels_2.put(label_2, label_2_val);
            testRoute2.getMetadata().setLabels(labels_2);

            MediationRoute mediationRoute2 = new MediationRoute(testRoute2);
            Request request2 = paasMediationUtils.createRequest(PaasMediationUtils.Resources.ROUTES, null, namespace, "POST", mediationRoute2, null);
            MediationRoute createdRoute2 = paasMediationUtils.doRequest(request2, 201, MediationRoute.class);
            Assertions.assertNotNull(createdRoute2);
        }

        @Test
        void testOnCreated() throws Exception {
            onAddedWatcher.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
            MediationRoute actualMediationRoute = onAddedWatcher.getEvents().get(0);
            checkRoute(testRoute1, actualMediationRoute);
        }

        @Test
        void testOnCreatedWithFilter() throws Exception {
            onAddedWatcherWithFilter.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
            MediationRoute actualMediationRoute = onAddedWatcherWithFilter.getEvents().get(0);
            checkRoute(testRoute2, actualMediationRoute);
        }

        @Nested
        @TestInstance(TestInstance.Lifecycle.PER_CLASS)
        class AfterDeleted {

            @BeforeAll
            void deleteRoute() {
                paasUtils.deleteIngress(routeName1);
                paasUtils.deleteIngress(routeName2);
            }

            @Test
            void testOnDeleted() {
                Assertions.assertDoesNotThrow(() -> onDeletedWatcher.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT));
            }

            @Test
            void testOnDeletedWithFilter() {
                Assertions.assertDoesNotThrow(() -> onDeletedWatcherWithFilter.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT));
            }
        }
    }
}
