package com.netcracker.it.paasmediation.v2.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import io.fabric8.kubernetes.api.model.networking.v1.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

import java.util.Optional;

@Data
@AllArgsConstructor
@NoArgsConstructor
@EqualsAndHashCode(callSuper = true)
public class MediationRoute extends MediationMetadata {
    private RouteSpec spec;

    public MediationRoute(Ingress route) {
        metadata = new Metadata();
        metadata.setAnnotations(route.getMetadata().getAnnotations());
        metadata.setKind(route.getKind());
        metadata.setLabels(route.getMetadata().getLabels());
        metadata.setName(route.getMetadata().getName());
        metadata.setNamespace(route.getMetadata().getNamespace());
        spec = new RouteSpec();
        spec.host = Optional.of(route.getSpec().getRules())
                .map(rules-> rules.isEmpty() ? null : rules.getFirst())
                .map(IngressRule::getHost)
                .orElse(null);
        spec.path = Optional.of(route.getSpec().getRules())
                .map(rules -> rules.isEmpty() ? null : rules.getFirst())
                .map(IngressRule::getHttp).map(HTTPIngressRuleValue::getPaths)
                .map(paths -> paths.isEmpty() ? null : paths.getFirst())
                .map(HTTPIngressPath::getPath)
                .orElse(null);
        spec.service = new Target();
        spec.service.name = Optional.of(route.getSpec().getRules())
                .map(rules -> rules.isEmpty() ? null : rules.getFirst())
                .map(IngressRule::getHttp).map(HTTPIngressRuleValue::getPaths)
                .map(paths -> paths.isEmpty() ? null : paths.getFirst())
                .map(HTTPIngressPath::getBackend)
                .map(IngressBackend::getService)
                .map(IngressServiceBackend::getName)
                .orElse(null);
    }

    @Data
    public static class RouteSpec {
        private String host;
        private String path;
        @JsonProperty("to")
        private Target service;
    }

    @Data
    public static class Target {
        private String name;
    }
}
