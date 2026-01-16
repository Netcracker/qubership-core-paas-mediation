package com.netcracker.it.paasmediation;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.Cloud;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.EnableExtension;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.PortForward;
import com.netcracker.cloud.junit.cloudcore.extension.annotations.Value;
import com.netcracker.cloud.security.core.utils.tls.TlsUtils;
import com.netcracker.it.paasmediation.utils.PaasUtils;
import io.fabric8.kubernetes.api.model.ConfigMap;
import io.fabric8.kubernetes.api.model.ObjectMeta;
import io.fabric8.kubernetes.client.KubernetesClient;
import lombok.extern.slf4j.Slf4j;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import org.junit.jupiter.api.BeforeAll;

import java.net.URL;
import java.util.Collections;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.TimeUnit;
import java.util.function.Function;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import static org.junit.jupiter.api.Assertions.assertNotNull;

@Slf4j
@EnableExtension
public class PaasMediationParentTest {

    @PortForward(serviceName = @Value("internal-gateway-service"))
    protected static URL internalGateway;

    @Cloud
    protected static KubernetesClient kubernetesClient;

    protected static PaasUtils paasUtils;
    protected static String namespace;
    protected static final ObjectMapper objectMapper = initObjectMapper();
    protected static OkHttpClient okHttpClient;
    protected final int WAIT_WS_TIMEOUT = 1;
    protected final TimeUnit WAIT_WS_TIMEUNIT = TimeUnit.MINUTES;

    @BeforeAll
    public static void initParentClass() throws Exception {
        okHttpClient = new OkHttpClient.Builder()
                .addInterceptor(chain -> {
                    Request newRequest = chain.request().newBuilder().build();
                    return chain.proceed(newRequest);
                })
                .readTimeout(30, TimeUnit.SECONDS)
                .sslSocketFactory(TlsUtils.getSslContext().getSocketFactory(), TlsUtils.getTrustManager())
                .build();
        paasUtils = new PaasUtils(kubernetesClient);
        namespace = kubernetesClient.getNamespace();
    }

    protected ObjectMeta createTestMetadata(String resourceName) {
        ObjectMeta objectMeta = new ObjectMeta();
        objectMeta.setName(resourceName);
        objectMeta.setNamespace(namespace);
        objectMeta.setLabels(Collections.singletonMap("it", "paas-mediation"));
        objectMeta.setAnnotations(Collections.singletonMap("owner", "Nectracker-company"));
        return objectMeta;
    }

    private static ObjectMapper initObjectMapper() {
        ObjectMapper objectMapper = new ObjectMapper();
        objectMapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        objectMapper.configure(SerializationFeature.INDENT_OUTPUT, true);
        return objectMapper;
    }
}
