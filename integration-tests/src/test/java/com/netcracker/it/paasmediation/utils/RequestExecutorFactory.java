package com.netcracker.it.paasmediation.utils;

import okhttp3.OkHttpClient;

public class RequestExecutorFactory {
    
    public enum ExecutionMode {
        PORT_FORWARD,
        EXEC_IN_POD
    }
    
    private static ExecutionMode currentMode = ExecutionMode.PORT_FORWARD;
    private static PaasUtils paasUtils;
    private static OkHttpClient okHttpClient;
    
    public static void init(OkHttpClient client, PaasUtils utils) {
        okHttpClient = client;
        paasUtils = utils;
    }
    
    public static void setMode(ExecutionMode mode) {
        currentMode = mode;
    }
    
    public static ExecutionMode getMode() {
        return currentMode;
    }
    
    public static RequestExecutor createExecutor() {
        return createExecutor(currentMode);
    }
    
    public static RequestExecutor createExecutor(ExecutionMode mode) {
        switch (mode) {
            case PORT_FORWARD:
                if (okHttpClient == null) {
                    throw new IllegalStateException("OkHttpClient not initialized. Call init() first.");
                }
                return new PortForwardRequestExecutor(okHttpClient);
            case EXEC_IN_POD:
                if (paasUtils == null) {
                    throw new IllegalStateException("PaasUtils not initialized. Call init() first.");
                }
                return new ExecInPodRequestExecutor(paasUtils, "name", "paas-mediation");
            default:
                throw new IllegalArgumentException("Unknown mode: " + mode);
        }
    }
}