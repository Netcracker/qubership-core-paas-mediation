package com.netcracker.it.paasmediation.v2.domain;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Data;
import lombok.EqualsAndHashCode;

import java.util.List;

@Data
@EqualsAndHashCode(callSuper = true)
public class MediationFullPodInfo extends MediationMetadata {
    private PodSpec spec;
    private PodStatus status;

    @Data
    public static class PodSpec{
        private List<SpecVolume> volumes;
        private List<SpecContainer> containers;
        private String restartPolicy;
        private Long terminationGracePeriodSeconds;
        private String dnsPolicy;
        private String nodeName;
    }

    @Data
    public static class SpecVolume{
        private String name;
        private VolumeSecret secret;
    }

    @Data
    public static class VolumeSecret{
        private String secretName;
        private int defaultMode;
    }

    @Data
    public static class SpecContainer{
        private String name;
        private String image;
        private List<ContainerPort> ports;
        private List<ContainerEnv> env;
        private ContainerResources resources;
        private List<ContainerVolumeMount> volumeMounts;
        private String ImagePullPolicy;
    }

    @Data
    public static class ContainerPort{
        private int containerPort;
        private String protocol;
    }

    @Data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    public static class ContainerEnv{
        private String name;
        private String value;
    }

    @Data
    public static class ContainerResources{
        private CpuMemoryResource limits;
        private CpuMemoryResource requests;
    }

    @Data
    public static class CpuMemoryResource{
        private String cpu;
        private String memory;
    }

    @Data
    public static class ContainerVolumeMount{
        private String name;
        private String mountPath;
        private String readOnly;
    }

    @Data
    public static class PodStatus{
        private String phase;
        private List<StatusCondition> conditions;
        private String hostIP;
        private String podIP;
        private String startTime;
        private List<ContainerStatus> containerStatuses;
    }

    @Data
    public static class StatusCondition{
        private String type;
        private String status;
        private String lastProbeTime;
        private String lastTransitionTime;
    }

    @Data
    public static class ContainerStatus{
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
    public static class ContainerState{
        private ContainerStateRunning running;
        private ContainerStateTerminated terminated;
        private ContainerStateWaiting waiting;
    }

    @Data
    public static class ContainerStateRunning{
        private String startedAt;
    }

    @Data
    public static class ContainerStateTerminated{
        private String containerID;
        private int exitCode;
        private String finishedAt;
        private String Reason;
        private String startedAt;
    }

    @Data
    public static class ContainerStateWaiting{
        private String message;
        private String reason;
    }


}
