package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;
import lombok.EqualsAndHashCode;

import java.util.List;
import java.util.Map;

@Data
@EqualsAndHashCode(callSuper = true)
public class MediationService extends MediationMetadata {
    private ServiceSpec spec;

    @Data
    public static class ServiceSpec {
        private List<Port> ports;
        private Map<String, String> selector;
        private String clusterIP;
        private String type;
    }

    @Data
    public static class Port {
        private String name;
        private String protocol;
        private Integer port;
        private Integer targetPort;
        private Integer nodePort;
    }

}
