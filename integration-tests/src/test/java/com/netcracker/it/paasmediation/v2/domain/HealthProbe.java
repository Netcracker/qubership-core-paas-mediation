package com.netcracker.it.paasmediation.v2.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Map;
import java.util.Set;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class HealthProbe {
    private String status;
    private HealthCheck badResourcesHealthCheck;

    @Data
    public class HealthCheck {
        private String status;
        private Map<String, Set<String>> badRoutes;
    }
}
