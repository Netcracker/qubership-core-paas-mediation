package com.netcracker.it.paasmediation.utils;

import okhttp3.Request;
import okhttp3.Response;
import lombok.extern.slf4j.Slf4j;

import java.io.IOException;

@Slf4j
public class ExecInPodRequestExecutor implements RequestExecutor {
    
    private final PaasUtils paasUtils;
    private final String podLabelKey;
    private final String podLabelValue;
    
    public ExecInPodRequestExecutor(PaasUtils paasUtils, String podLabelKey, String podLabelValue) {
        this.paasUtils = paasUtils;
        this.podLabelKey = podLabelKey;
        this.podLabelValue = podLabelValue;
    }
    
    @Override
    public Response execute(Request request) throws IOException {
        try {
            log.debug("Executing request via exec in pod {}/{}: {}", podLabelKey, podLabelValue, request.url());
            return paasUtils.doRequestFromInclusterPodByLabel(podLabelKey, podLabelValue, request);
        } catch (Exception e) {
            throw new IOException("Failed to execute request in pod", e);
        }
    }
}