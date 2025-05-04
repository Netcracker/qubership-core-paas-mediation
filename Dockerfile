FROM ghcr.io/netcracker/qubership/core-base:1.0.0

COPY --chown=10001:0 paas-mediation-service/bin/paas-mediation-service /app/paas-mediation
COPY --chown=10001:0 ["paas-mediation-service/application.yaml", "paas-mediation-service/policies.conf", "paas-mediation-service/docs", "/app/"]