package com.netcracker.it.paasmediation.v2.ws;

import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Request;
import org.junit.jupiter.api.*;

import java.util.concurrent.TimeUnit;


@Slf4j
@Tag("watch")
public class NamespaceWebSocketIT extends PaasMediationParentV2Test {

    private WSListener wsListener;

    @BeforeAll
    void createConfigMap() {
        Request request = paasMediationUtils.createWsRequest(PaasMediationUtils.Resources.NAMESPACES, null);
        wsListener = new WSListener(okHttpClient, request);
    }

    @AfterAll
    public void afterTest() {
        if (wsListener != null) {
            wsListener.close();
        }
    }

    @Test
    public void checkWebsocketNamespaceEvent() throws InterruptedException {
        wsListener.waitConnected(30, TimeUnit.SECONDS);
        Assertions.assertNull(wsListener.getException());
    }
}
