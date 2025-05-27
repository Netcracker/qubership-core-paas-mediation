[![Go build](https://github.com/Netcracker/qubership-core-paas-mediation/actions/workflows/go-build.yml/badge.svg)](https://github.com/Netcracker/qubership-core-paas-mediation/actions/workflows/go-build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?metric=coverage&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![duplicated_lines_density](https://sonarcloud.io/api/project_badges/measure?metric=duplicated_lines_density&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![vulnerabilities](https://sonarcloud.io/api/project_badges/measure?metric=vulnerabilities&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![bugs](https://sonarcloud.io/api/project_badges/measure?metric=bugs&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)
[![code_smells](https://sonarcloud.io/api/project_badges/measure?metric=code_smells&project=Netcracker_qubership-core-paas-mediation)](https://sonarcloud.io/summary/overall?id=Netcracker_qubership-core-paas-mediation)

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
