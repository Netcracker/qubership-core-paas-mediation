package com.netcracker.it.paasmediation.v2.helpers;

import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationSecret;
import io.fabric8.kubernetes.api.model.Secret;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;

import java.util.Collections;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

public abstract class SecretHelper extends PaasMediationParentV2Test {

    protected String secretNamePrefix = "paas-mediation-it-test-secret";
    protected String secretName1 = secretNamePrefix + "-1";
    protected String secretName2 = secretNamePrefix + "-2";

    @BeforeAll
    @AfterAll
    public void cleanup() {
        Stream.of(secretName1, secretName2).forEach(name -> {
            if (paasUtils.getSecretByName(name) != null) {
                paasUtils.deleteSecret(name);
                assertNull(paasUtils.getSecretByName(name));
            }
        });
    }

    protected void checkSecret(Secret expectSecret, MediationSecret mediationSecret) {
        assertEquals(expectSecret.getKind(), mediationSecret.getMetadata().getKind());
        assertEquals(expectSecret.getMetadata().getName(), mediationSecret.getMetadata().getName());
        assertEquals(expectSecret.getMetadata().getNamespace(), mediationSecret.getMetadata().getNamespace());
        assertEquals(expectSecret.getMetadata().getLabels(), mediationSecret.getMetadata().getLabels());
        assertTrue(expectSecret.getMetadata().getAnnotations().entrySet().containsAll(mediationSecret.getMetadata().getAnnotations().entrySet()));
        assertTrue(mediationSecret.getData().containsKey("paas-mediation-secret-data"));
    }

    protected Secret createTestSecret(String name) {
        Secret secret = new Secret();
        secret.setMetadata(createTestMetadata(name));
        secret.setKind("Secret");
        secret.setData(Collections.singletonMap("paas-mediation-secret-data", "cmVhbGx5X3NlY3JldF92YWx1ZTEK"));
        return secret;
    }

}
