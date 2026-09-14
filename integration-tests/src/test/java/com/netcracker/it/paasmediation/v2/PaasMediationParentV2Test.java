package com.netcracker.it.paasmediation.v2;

import com.netcracker.it.paasmediation.PaasMediationParentTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.utils.PaasUtils;
import com.netcracker.it.paasmediation.utils.RequestExecutor;
import com.netcracker.it.paasmediation.utils.RequestExecutorFactory;
import com.netcracker.it.paasmediation.v2.helpers.DeploymentHelper;
import lombok.extern.slf4j.Slf4j;

import java.util.concurrent.TimeUnit;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Tag;
import okhttp3.OkHttpClient;

@Slf4j
@Tag("v2")
public class PaasMediationParentV2Test extends PaasMediationParentTest {

    protected static PaasMediationUtils paasMediationUtils;
    protected static DeploymentHelper deploymentHelper;

    private static String meshType() {
        String meshType = System.getProperty("SERVICE_MESH_TYPE");
        if (meshType == null || meshType.isEmpty()) {
            meshType = System.getenv("SERVICE_MESH_TYPE");
        }
        return meshType == null || meshType.isEmpty() ? "Core" : meshType;
    }

    @BeforeAll
    public static void initClass() {
        okHttpClient = new OkHttpClient.Builder()
                .readTimeout(60, TimeUnit.SECONDS)
                .connectTimeout(60, TimeUnit.SECONDS)
                .build();
        
        paasUtils = new PaasUtils(kubernetesClient);
        
        RequestExecutorFactory.init(okHttpClient, paasUtils);

        String mode = System.getProperty("executor.mode", "PORT_FORWARD");
        RequestExecutorFactory.setMode(RequestExecutorFactory.ExecutionMode.valueOf(mode));

        log.info("Executor mode set to: {}", mode);

        RequestExecutor executor = RequestExecutorFactory.createExecutor();

        // websocket requests do not go through the executor, they are sent from outside the service mesh:
        // in Istio mesh mode the internal gateway routes are applied by the mesh proxies only, so the watch
        // api is called on the service itself
        boolean istioMesh = "Istio".equalsIgnoreCase(meshType()) && paasMediationService != null;
        String watchBaseUrl = istioMesh ? paasMediationService.toString() : internalGateway.toString();
        String watchServicePrefix = istioMesh ? "" : "/paas-mediation";
        log.info("Service mesh type is {}, watch api base url is {}", meshType(), watchBaseUrl);

        paasMediationUtils = new PaasMediationUtils(
            "v2", 
            internalGateway.toString(), 
            privateGateway.toString(),
            watchBaseUrl,
            watchServicePrefix,
            executor, 
            objectMapper
        );

        deploymentHelper = new DeploymentHelper(kubernetesClient, paasUtils, paasMediationUtils, namespace);
    }

}
