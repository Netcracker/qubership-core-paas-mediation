FROM golang:1.24 AS build

WORKDIR /app

COPY paas-mediation-service/ .

RUN go mod download
RUN go build -o paas-mediation-service .

FROM ghcr.io/netcracker/qubership/core-base:1.1.0 AS run

COPY --chown=10001:0 --chmod=555 --from=build app/paas-mediation-service /app/paas-mediation
COPY --chown=10001:0 --chmod=444 --from=build app/application.yaml /app/
COPY --chown=10001:0 --chmod=444 --from=build app/docs/swagger.json /app/
COPY --chown=10001:0 --chmod=444 --from=build app/docs/swagger.yaml /app/

WORKDIR /app

USER 10001:10001

CMD ["/app/paas-mediation"]