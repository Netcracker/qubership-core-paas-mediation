package com.netcracker.it.paasmediation.v2.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.Map;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DeploymentRolloutResponse {
    private List<Map<String, Deployment>> deployments;
    @JsonProperty("pod_status_websocket")
    private String podStatusWebSocket;

    @Data
    @AllArgsConstructor
    @NoArgsConstructor
    public  static class Deployment {
        private String kind;
        private String active;
        private String rolling;
    }
}
