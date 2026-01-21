package com.netcracker.it.paasmediation.utils;

import lombok.Data;

import java.util.HashMap;
import java.util.Map;

@Data
public class PaasRequestFilter {
    public static PaasRequestFilter EVERYTHING = new PaasRequestFilter();

    private Map<String, String> labels = new HashMap<>();
    private Map<String, String> annotations = new HashMap<>();

    public PaasRequestFilter withLabel(String name, String value) {
        labels.put(name, value);
        return this;
    }

    public PaasRequestFilter withAnnotation(String name, String value) {
        annotations.put(name, value);
        return this;
    }
}
