# PaaS Mediation Service Documentation


| Section Link                            | Contents                                                                                                                            |
|-----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| [Overview](/docs/paas-mediation-overview.md) | [Basic concept description](/docs/paas-mediation-overview.md) and [Deploy parameters](/docs/paas-mediation-overview.md#deploy-parameters) |
| [REST API](/docs/rest_api.md)           | PaaS Mediation REST API description                                                                                                 |
| [Websocket API](/docs/websocket_api.md) | PaaS Mediation Websocket API description                                                                                            | 


# How to run locally
1. Switch kube context and run devbox with port-forwarding to your namespace
2. Add environment var: \
   for kubernetes ```PAAS_PLATFORM=KUBERNETES;IDP_CLIENT_USERNAME=paas-mediation;IDP_CLIENT_PASSWORD=<secret>;MICROSERVICE_NAMESPACE=<ns>``` \
   for openshift ```PAAS_PLATFORM=OPENSHIFT;IDP_CLIENT_USERNAME=paas-mediation;IDP_CLIENT_PASSWORD=<secret>;MICROSERVICE_NAMESPACE=<ns>```
3. Run with flag ```-local```
