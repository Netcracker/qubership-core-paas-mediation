package com.netcracker.it.paasmediation.v2.ws;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.netcracker.it.paasmediation.v2.domain.WatchEvent;
import lombok.Getter;
import lombok.extern.slf4j.Slf4j;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.function.Predicate;

@Getter
@Slf4j
public class Watcher<T> {
    private final Class<T> _class;
    private final int eventsToReceive;
    private CountDownLatch eventLatch;
    private final List<T> events = new ArrayList<>();
    private final Predicate<String> typeFilter;
    private final Predicate<WatchEvent<T>> eventFilter;

    public Watcher(Class<T> _class, int eventsToReceive, Predicate<String> typeFilter, Predicate<WatchEvent<T>> eventFilter) {
        this._class = _class;
        this.eventsToReceive = eventsToReceive;
        this.eventLatch = new CountDownLatch(eventsToReceive);
        this.typeFilter = typeFilter;
        this.eventFilter = eventFilter;
    }

    public void waitMessagesReceived(int timeout, TimeUnit unit) throws InterruptedException {
        if (!eventLatch.await(timeout, unit)) {
            throw new IllegalStateException(String.format("Watch events (amount=%d) were not received in specified period: %d %s",
                    this.eventsToReceive, timeout, unit));
        }
    }

    public void receiveEvent(String type, String obj, ObjectMapper objectMapper) {
        try {
            T resource = objectMapper.readValue(obj, _class);
            WatchEvent<T> event = new WatchEvent<>(type, resource);
            if (this.eventFilter.test(event)) {
                log.info("Got event that satisfies eventFilter. Event: {}", event);
                events.add(event.getObject());
                eventLatch.countDown();
            }
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }

    public List<T> getEvents() {
        return events;
    }
}
