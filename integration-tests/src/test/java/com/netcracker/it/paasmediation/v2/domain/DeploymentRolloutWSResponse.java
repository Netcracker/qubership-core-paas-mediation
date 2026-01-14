package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.HashMap;
import java.util.List;

@NoArgsConstructor
public class DeploymentRolloutWSResponse extends HashMap<String, List<DeploymentRolloutWSResponse.WSResponsePod>> {
    @Data
    public static class WSResponsePod {
        String name;
        String status;
    }
}
