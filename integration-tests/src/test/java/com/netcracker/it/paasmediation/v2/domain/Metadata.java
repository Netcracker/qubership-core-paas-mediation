package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;

import java.util.Map;

@Data
public class Metadata {
    private String kind;
    private String name;
    private String namespace;
    private Map<String, String> annotations;
    private Map<String, String> labels;
}