package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;
import lombok.EqualsAndHashCode;

import java.util.Map;

@Data
@EqualsAndHashCode(callSuper = true)
public class MediationSecret extends MediationMetadata {
    private Map<String, byte[]> data;
    private String type;
}
