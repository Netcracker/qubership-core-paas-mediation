package com.netcracker.it.paasmediation.utils;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import lombok.extern.slf4j.Slf4j;

import java.io.IOException;

@Slf4j
public class PortForwardRequestExecutor implements RequestExecutor {
    
    private final OkHttpClient okHttpClient;
    
    public PortForwardRequestExecutor(OkHttpClient okHttpClient) {
        this.okHttpClient = okHttpClient;
    }
    
    @Override
    public Response execute(Request request) throws IOException {
        log.debug("Executing request via port-forward: {}", request.url());
        return okHttpClient.newCall(request).execute();
    }
}