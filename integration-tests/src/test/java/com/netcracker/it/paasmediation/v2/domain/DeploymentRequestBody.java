package com.netcracker.it.paasmediation.v2.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.List;

@Data
@AllArgsConstructor
public class DeploymentRequestBody {
    @JsonProperty("deployment_names")
    List<String> deploymentNames;
}
