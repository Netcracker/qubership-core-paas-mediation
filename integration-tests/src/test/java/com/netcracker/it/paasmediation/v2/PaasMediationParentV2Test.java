package com.netcracker.it.paasmediation.v2;

import com.netcracker.it.paasmediation.PaasMediationParentTest;
import com.netcracker.it.paasmediation.utils.PaasMediationUtils;
import com.netcracker.it.paasmediation.v2.helpers.DeploymentHelper;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Tag;

@Slf4j
@Tag("v2")
public class PaasMediationParentV2Test extends PaasMediationParentTest {

    protected static PaasMediationUtils paasMediationUtils;
    protected static DeploymentHelper deploymentHelper;

    @BeforeAll
    public static void initClass() {
        paasMediationUtils = new PaasMediationUtils("v2", internalGateway.toString(), okHttpClient, objectMapper);
        deploymentHelper = new DeploymentHelper(kubernetesClient, paasUtils, paasMediationUtils, namespace);
    }

}
