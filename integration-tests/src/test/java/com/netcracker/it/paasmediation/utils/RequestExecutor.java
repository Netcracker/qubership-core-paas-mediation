package com.netcracker.it.paasmediation.utils;

import okhttp3.Request;
import okhttp3.Response;

import java.io.IOException;

public interface RequestExecutor {
    Response execute(Request request) throws IOException;
}