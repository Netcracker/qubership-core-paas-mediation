package com.netcracker.it.paasmediation.v2.ws;

import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.utils.PaasRequestFilter;
import com.netcracker.it.paasmediation.v2.domain.MediationConfigMap;
import com.netcracker.it.paasmediation.v2.helpers.ConfigMapHelper;
import io.fabric8.kubernetes.api.model.ConfigMap;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.*;

import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Stream;

@Slf4j
@Tag("watch")
public class ConfigMapWebSocketIT extends ConfigMapHelper {

    private WSListener wsListener;
    private WSListener wsListenerWithFilter;

    private final Watcher<MediationConfigMap> onAddedWatcher = new Watcher<>(MediationConfigMap.class, 1,
            "ADDED"::equals, event -> Objects.equals(event.getObject().getMetadata().getName(), configMapName1));
    private final Watcher<MediationConfigMap> onDeletedWatcher = new Watcher<>(MediationConfigMap.class, 1,
            "DELETED"::equals, event -> Objects.equals(event.getObject().getMetadata().getName(), configMapName1));
    private final Watcher<MediationConfigMap> onAddedWatcherWithFilter = new Watcher<>(MediationConfigMap.class, 1,
            "ADDED"::equals, event -> event.getObject().getMetadata().getName().startsWith(configMapNamePrefix));
    private final Watcher<MediationConfigMap> onDeletedWatcherWithFilter = new Watcher<>(MediationConfigMap.class, 1,
            "DELETED"::equals, event -> event.getObject().getMetadata().getName().startsWith(configMapNamePrefix));

    String label_1 = "1";
    String label_1_val = "1";
    String label_2 = "2";
    String label_2_val = "2";

    @BeforeAll
    void connect() throws Exception {
        Request request1 = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.CONFIGMAPS, namespace);
        wsListener = new WSListener(okHttpClient, request1, onAddedWatcher, onDeletedWatcher);
        wsListener.waitConnected(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);

        PaasRequestFilter filter = new PaasRequestFilter().withLabel(label_2, label_2_val);
        Request request2 = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.CONFIGMAPS, namespace, filter);
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
        ConfigMap testConfigMap1;
        ConfigMap testConfigMap2;

        @BeforeAll
        void createConfigMap() {
            testConfigMap1 = createTestConfigMap(configMapName1);
            Map<String, String> labels_1 = new HashMap<>();
            labels_1.put(label_1, label_1_val);
            testConfigMap1.getMetadata().setLabels(labels_1);
            paasUtils.createConfigMap(testConfigMap1);

            testConfigMap2 = createTestConfigMap(configMapName2);
            Map<String, String> labels_2 = new HashMap<>();
            labels_2.put(label_2, label_2_val);
            testConfigMap2.getMetadata().setLabels(labels_2);
            paasUtils.createConfigMap(testConfigMap2);
        }

        @Test
        void testOnCreated() throws Exception {
            onAddedWatcher.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
            MediationConfigMap actualMediationConfigMap = onAddedWatcher.getEvents().get(0);
            checkConfigMap(testConfigMap1, actualMediationConfigMap);
        }

        @Test
        void testOnCreatedWithFilter() throws Exception {
            onAddedWatcherWithFilter.waitMessagesReceived(WAIT_WS_TIMEOUT, WAIT_WS_TIMEUNIT);
            MediationConfigMap actualMediationConfigMap = onAddedWatcherWithFilter.getEvents().get(0);
            checkConfigMap(testConfigMap2, actualMediationConfigMap);
        }

        @Nested
        @TestInstance(TestInstance.Lifecycle.PER_CLASS)
        class AfterDeleted {

            @BeforeAll
            void deleteConfigMap() {
                paasUtils.deleteConfigMap(configMapName1);
                paasUtils.deleteConfigMap(configMapName2);
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
