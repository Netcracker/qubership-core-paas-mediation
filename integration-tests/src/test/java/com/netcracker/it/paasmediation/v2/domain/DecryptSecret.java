package com.netcracker.it.paasmediation.v2.domain;

import java.util.Base64;
import java.util.Map;

public class DecryptSecret {
    String name;
    Map<String, String> data;

    public DecryptSecret(String name, Map<String, String> data) {
        this.name = name;
        this.data = data;
    }

    public String get(String name) {
        return new String(Base64.getDecoder().decode(data.get(name)));
    }
}
