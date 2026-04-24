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

    @BeforeAll
    public static void initClass() {
        okHttpClient = new OkHttpClient.Builder()
                .readTimeout(60, TimeUnit.SECONDS)
                .connectTimeout(60, TimeUnit.SECONDS)
                .build();
        
        paasUtils = new PaasUtils(kubernetesClient);
        
        RequestExecutorFactory.init(okHttpClient, paasUtils);
        
        //RequestExecutorFactory.setMode(RequestExecutorFactory.ExecutionMode.PORT_FORWARD);
        //RequestExecutorFactory.setMode(RequestExecutorFactory.ExecutionMode.EXEC_IN_POD);

        String mode = System.getProperty("executor.mode", "PORT_FORWARD");
        RequestExecutorFactory.setMode(RequestExecutorFactory.ExecutionMode.valueOf(mode));

        log.info("Executor mode set to: {}", mode);

        RequestExecutor executor = RequestExecutorFactory.createExecutor();
        
        paasMediationUtils = new PaasMediationUtils(
            "v2", 
            internalGateway.toString(), 
            executor, 
            objectMapper
        );

        deploymentHelper = new DeploymentHelper(kubernetesClient, paasUtils, paasMediationUtils, namespace);
    }

}
