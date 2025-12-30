package com.netcracker.it.paasmediation.v2.helpers;

import com.netcracker.it.paasmediation.v2.PaasMediationParentV2Test;
import com.netcracker.it.paasmediation.v2.domain.MediationRoute;
import io.fabric8.kubernetes.api.model.ObjectMeta;
import io.fabric8.kubernetes.api.model.networking.v1.*;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;

import java.util.Collections;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

public abstract class RouteHelper extends PaasMediationParentV2Test {

    protected static String routeNamePrefix = "paas-mediation-it-test-route";
    protected static String routeName1 = routeNamePrefix + "-1";
    protected static String routeName2 = routeNamePrefix + "-2";

    @BeforeAll
    @AfterAll
    public static void cleanup() {
        Stream.of(routeName1, routeName2).forEach(name -> {
            if (paasUtils.getIngressByName(name) != null) {
                paasUtils.deleteIngress(name);
                assertNull(paasUtils.getIngressByName(name));
            }
        });
    }

    protected void checkServiceRoutes(MediationRoute expectRoute, MediationRoute actualRoute) {
        // todo change kind from Route to Ingress in paas-mediation
//        assertEquals(expectRoute.getMetadata().getKind(), actualRoute.getMetadata().getKind());
        assertEquals(expectRoute.getMetadata().getName(), actualRoute.getMetadata().getName());
        assertEquals(expectRoute.getMetadata().getNamespace(), actualRoute.getMetadata().getNamespace());
        assertTrue(actualRoute.getMetadata().getAnnotations().entrySet().containsAll(expectRoute.getMetadata().getAnnotations().entrySet()));
        assertEquals(expectRoute.getMetadata().getLabels(), actualRoute.getMetadata().getLabels());
        assertEquals(expectRoute.getSpec().getService().getName(), actualRoute.getSpec().getService().getName());
    }

    protected void checkRoute(Ingress expectRoute, MediationRoute actualRoute) {
        // todo change kind from Route to Ingress in paas-mediation
//        assertEquals(expectRoute.getKind(), actualRoute.getMetadata().getKind());
        assertEquals(expectRoute.getMetadata().getName(), actualRoute.getMetadata().getName());
        assertEquals(expectRoute.getMetadata().getNamespace(), actualRoute.getMetadata().getNamespace());
        assertEquals(expectRoute.getMetadata().getLabels(), actualRoute.getMetadata().getLabels());
        assertTrue(expectRoute.getMetadata().getAnnotations().entrySet().containsAll(actualRoute.getMetadata().getAnnotations().entrySet()));
        assertEquals(expectRoute.getSpec().getRules().getFirst().getHttp().getPaths().getFirst().getBackend().getService().getName(), actualRoute.getSpec().getService().getName());
    }

    protected Ingress createTestRoute(String name) {
        ObjectMeta objectMeta = new ObjectMeta();
        objectMeta.setName(name);
        objectMeta.setNamespace(namespace);
        objectMeta.setLabels(Collections.singletonMap("it", "paas-mediation"));
        objectMeta.setAnnotations(Collections.singletonMap("owner", "qubership"));
        String host = name + "-" + namespace + ".local.qubership.org";
        HTTPIngressRuleValue rule = new HTTPIngressRuleValue().edit()
                .withPaths(new HTTPIngressPath(new IngressBackend().edit()
                        .withService(new IngressServiceBackend().edit()
                                .withName("paas-mediation")
                                .withPort(new ServiceBackendPort().edit()
                                        .withName("web")
                                        .build())
                                .build())
                        .build(), "/", "Prefix"))
                .build();
        IngressSpec spec = new IngressSpec().edit()
                .withRules(new IngressRule(host, rule))
                .build();
        return new Ingress().edit()
                .withMetadata(objectMeta)
                .withSpec(spec)
                .build();
    }

}
