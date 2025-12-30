package com.netcracker.it.paasmediation.v2.domain;

import lombok.Data;

@Data
public class ApplicationData {
    private String appName;
    private String appVersion;
    private String deployTime;
}
