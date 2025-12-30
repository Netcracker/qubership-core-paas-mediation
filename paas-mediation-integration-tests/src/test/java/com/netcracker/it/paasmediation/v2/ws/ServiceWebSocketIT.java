package com.netcracker.it.paasmediation.v2.ws;

import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.utils.PaasRequestFilter;
import com.netcracker.it.paasmediation.v2.domain.MediationService;
import com.netcracker.it.paasmediation.v2.helpers.ServiceHelper;
import io.fabric8.kubernetes.api.model.Service;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.*;

import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Stream;

@Slf4j
@Tag("watch")
public class ServiceWebSocketIT extends ServiceHelper {

    private WSListener wsListener;
    private WSListener wsListenerWithFilter;

    private final Watcher<MediationService> onAddedWatcher = new Watcher<>(MediationService.class, 1,
            "ADDED"::equals, event -> Objects.equals(event.getObject().getMetadata().getName(), serviceName1));
    private final Watcher<MediationService> onDeletedWatcher = new Watcher<>(MediationService.class, 1,
            "DELETED"::equals, event -> Objects.equals(event.getObject().getMetadata().getName(), serviceName1));
    private final Watcher<MediationService> onAddedWatcherWithFilter = new Watcher<>(MediationService.class, 1,
            "ADDED"::equals, event -> event.getObject().getMetadata().getName().startsWith(serviceNamePrefix));
    private final Watcher<MediationService> onDeletedWatcherWithFilter = new Watcher<>(MediationService.class, 1,
            "DELETED"::equals, event -> event.getObject().getMetadata().getName().startsWith(serviceNamePrefix));
    String label_1 = "1";
    String label_1_val = "1";
    String label_2 = "2";
    String label_2_val = "2";

    @BeforeAll
    void connect() throws Exception {
        Request request1 = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.SERVICES, namespace);
        wsListener = new WSListener(okHttpClient, request1, onAddedWatcher, onDeletedWatcher);
        wsListener.waitConnected(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);

        PaasRequestFilter filter = new PaasRequestFilter().withLabel(label_2, label_2_val);
        Request request2 = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.SERVICES, namespace, filter);
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
        Service testService1;
        Service testService2;

        @BeforeAll
        void createService() {
            testService1 = createTestService(serviceName1);
            Map<String, String> labels_1 = new HashMap<>();
            labels_1.put(label_1, label_1_val);
            testService1.getMetadata().setLabels(labels_1);
            paasUtils.createService(testService1);

            testService2 = createTestService(serviceName2);
            Map<String, String> labels_2 = new HashMap<>();
            labels_2.put(label_2, label_2_val);
            testService2.getMetadata().setLabels(labels_2);
            paasUtils.createService(testService2);
        }

        @Test
        void testOnCreated() throws Exception {
            onAddedWatcher.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
            MediationService actualMediationService = onAddedWatcher.getEvents().get(0);
            checkService(testService1, actualMediationService);
        }

        @Test
        void testOnCreatedWithFilter() throws Exception {
            onAddedWatcherWithFilter.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
            MediationService actualMediationService = onAddedWatcherWithFilter.getEvents().get(0);
            checkService(testService2, actualMediationService);
        }

        @Nested
        @TestInstance(TestInstance.Lifecycle.PER_CLASS)
        class AfterDeleted {

            @BeforeAll
            void deleteService() {
                paasUtils.deleteService(serviceName1);
                paasUtils.deleteService(serviceName2);
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
