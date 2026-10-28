FROM --platform=$BUILDPLATFORM golang:1.27.3@sha256:898838c455eecb68ed7b0b6f0635630ebdd657e732f07a23d0ab475007a52c36 AS build

WORKDIR /app

COPY paas-mediation-service/ .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -o paas-mediation-service .

FROM ghcr.io/netcracker/qubership-core-base:alpine-2.5.9@sha256:ab860790e4ecce56f8df673495cd5745ce587fffa7d7eb210d6d098563ccb26d AS run

COPY --chown=10001:0 --chmod=555 --from=build app/paas-mediation-service /app/paas-mediation
COPY --chown=10001:0 --chmod=444 --from=build app/application.yaml /app/
COPY --chown=10001:0 --chmod=444 --from=build app/docs/swagger.json /app/
COPY --chown=10001:0 --chmod=444 --from=build app/docs/swagger.yaml /app/

WORKDIR /app

CMD ["/app/paas-mediation"]