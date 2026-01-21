package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;

@Data
public class AnnotationResource {
    private String resourceName;
    private String namespace;
    private String annotationValue;
}
