package com.netcracker.it.paasmediation.v2.domain;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Data;
import lombok.EqualsAndHashCode;

import java.util.List;
import java.util.Map;

@Data
@EqualsAndHashCode(callSuper = true)
public class MediationFullDeploymentInfo extends MediationMetadata{
    private DeploymentSpec spec;
    private DeploymentStatus status;

    @Data
    public static class DeploymentSpec {
        private PodTemplateSpec template;
        private DeploymentStrategy strategy;
        private Long replicas;
        private Long revisionHistoryLimit;
    }

    @Data
    public static class PodTemplateSpec {
        private TemplateMetadata metadata;
        private MediationFullPodInfo.PodSpec spec;
    }

    @Data
    public static class DeploymentStrategy {
        private String type;
    }

    @Data
    public static class TemplateMetadata {
        private Map<String, String> labels;
    }

    @Data
    public static class SpecVolume {
        private String name;
        private VolumeSecret secret;
    }

    @Data
    public static class VolumeSecret {
        private String secretName;
        private int defaultMode;
    }

    @Data
    public static class SpecContainer {
        private String name;
        private String image;
        private List<ContainerPort> ports;
        private List<ContainerEnv> env;
        private ContainerResources resources;
        private List<ContainerVolumeMount> volumeMounts;
        private String ImagePullPolicy;
    }

    @Data
    public static class ContainerPort {
        private int containerPort;
        private String protocol;
    }

    @Data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class ContainerEnv {
        private String name;
        private String value;
    }

    @Data
    public static class ContainerResources {
        private CpuMemoryResource limits;
        private CpuMemoryResource requests;
    }

    @Data
    public static class CpuMemoryResource {
        private String cpu;
        private String memory;
    }

    @Data
    public static class ContainerVolumeMount {
        private String name;
        private String mountPath;
        private String readOnly;
    }

    @Data
    public static class DeploymentStatus {
        private Long availableReplicas;
        private List<DeploymentCondition> conditions;
        private Long observedGeneration;
        private Integer readyReplicas;
        private Integer replicas;
        private Long updatedReplicas;
    }

    @Data
    public static class DeploymentCondition {
        private String lastTransitionTime;
        private String lastUpdateTime;
        private String message;
        private String reason;
        private String status;
        private String type;
    }

    @Data
    public static class StatusCondition {
        private String type;
        private String status;
        private String lastProbeTime;
        private String lastTransitionTime;
    }

    @Data
    public static class ContainerStatus {
        private String name;
        private ContainerState state;
        private ContainerState lastState;
        private boolean ready;
        private int restoreCount;
        private String image;
        private String imageID;
        private String containerID;
    }

    @Data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class ContainerState {
        private ContainerStateRunning running;
        private ContainerStateTerminated terminated;
        private ContainerStateWaiting waiting;
    }

    @Data
    public static class ContainerStateRunning {
        private String startedAt;
    }

    @Data
    public static class ContainerStateTerminated {
        private String containerID;
        private int exitCode;
        private String finishedAt;
        private String Reason;
        private String startedAt;
    }

    @Data
    public static class ContainerStateWaiting {
        private String message;
        private String reason;
    }


}
