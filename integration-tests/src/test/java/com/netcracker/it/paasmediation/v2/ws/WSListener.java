package com.netcracker.it.paasmediation.v2.ws;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import okhttp3.*;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

@Slf4j
public class WSListener extends WebSocketListener {
    private final ObjectMapper objectMapper;
    private CountDownLatch openWsLatch;
    private final List<Throwable> exceptions = new ArrayList<>();
    private final OkHttpClient okHttpClient;
    private final Request request;
    private WebSocket webSocket;
    private final List<Watcher<?>> watchers;
    private boolean closed;

    public WSListener(OkHttpClient okHttpClient, Request request, Watcher<?>... watchers) {
        this.okHttpClient = okHttpClient;
        this.request = request;
        this.openWsLatch = new CountDownLatch(1);
        this.watchers = Arrays.stream(watchers).collect(Collectors.toList());
        this.objectMapper = new ObjectMapper().configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        this.webSocket = connect();
    }

    private WebSocket connect() {
        return this.okHttpClient.newWebSocket(this.request, this);
    }

    public void waitConnected(int timeout, TimeUnit unit) throws InterruptedException {
        if (!openWsLatch.await(timeout, unit)) {
            throw new IllegalStateException(String.format("Websocket connection was not established in specified period: %d %s. Exception: %s",
                    timeout, unit, getException()));
        }
    }

    public Throwable getException() {
        if (!exceptions.isEmpty()) {
            if (exceptions.size() == 1) {
                return exceptions.get(0);
            } else {
                Exception exception = new Exception("Websocket listener got exceptions");
                exceptions.forEach(exception::addSuppressed);
                return exception;
            }
        }
        return null;
    }

    public boolean isClosed() {
        return closed;
    }

    @Override
    public void onOpen(WebSocket webSocket, Response response) {
        log.info("onOpen '{}'", getWebsocketUrl(webSocket));
        openWsLatch.countDown();
    }

    @Override
    public void onMessage(WebSocket webSocket, String text) {
        try {
            JsonNode jsonNode = objectMapper.readTree(text);
            String type = jsonNode.get("type").asText();
            log.info("onMessage '{}'. [{}] {}",  getWebsocketUrl(webSocket), type,
                    Optional.ofNullable(jsonNode.get("object"))
                            .map(object -> object.get("metadata"))
                            .map(object -> object.get("name")).map(Object::toString).orElse(""));
            this.watchers.stream().filter(watcher -> watcher.getTypeFilter().test(type))
                    .forEach(watcher -> watcher.receiveEvent(type, jsonNode.get("object").toString(), objectMapper));
        } catch (IOException e) {
            log.warn("Failed to parse event ", e);
            exceptions.add(e);
        }
    }

    @Override
    public void onClosed(WebSocket webSocket, int code, String reason) {
        log.info("onClosed '{}', code: {}, reason: {}",  getWebsocketUrl(webSocket), code, reason);
        closed = true;
    }


    @Override
    public void onFailure(WebSocket webSocket, Throwable t, Response response) {
        log.warn("onFailure '{}'  exception: {}, response: {}",  getWebsocketUrl(webSocket), t, response);
        exceptions.add(new Exception(String.format("Fail, exception %s, response %s", t, response)));
    }

    public void close() {
        closed = true;
        Optional.ofNullable(this.webSocket).ifPresent(ws -> ws.close(1001, ""));
    }

    private String getWebsocketUrl(WebSocket webSocket) {
        return Optional.ofNullable(webSocket).map(WebSocket::request).map(Request::url).map(HttpUrl::toString).orElse("null");
    }
}
