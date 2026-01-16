package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@EqualsAndHashCode(callSuper = true)
public class MediationPod extends MediationMetadata {
    private PodSpec spec;

    @Data
    public static class PodSpec {
        private List<Container> containers;
        private String clusterIP;
        private String type;
    }

    @Data
    public static class Container {
        private String name;
        private String image;
        private List<Port> ports;

        @Data public static class Port {
            private String name;
            private int containerPort;
            private String protocol;
        }
    }
}
