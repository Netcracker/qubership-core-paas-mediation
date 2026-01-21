package com.netcracker.it.paasmediation.v2.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.NonNull;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class WatchEvent<T> {
    @NonNull
    private String type;
    @NonNull
    private T object;
}
