# Docker configuration to build Tekton Results apiserver image
FROM registry.ci.openshift.org/openshift/release:golang-1.18 AS builder

WORKDIR /opt/app-root/src

COPY go.mod go.mod
COPY go.sum go.sum
# COPY vendor/ vendor/
COPY gen/ gen/
COPY your/ your/
COPY tools/ tools/
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -a -o api main.go

FROM registry.access.redhat.com/ubi9-minimal:9.1.0
COPY --from=builder /opt/app-root/src/api /usr/local/bin/api
USER 65532:65532

ENTRYPOINT ["api"]