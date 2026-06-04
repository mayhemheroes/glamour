FROM golang:1.21 AS builder

ENV GOTOOLCHAIN=auto

ARG project='glamour'
ARG fuzzTarget='FuzzRenderBytes'
ARG fuzzTargetPkg='./testdata/fuzz'

ADD . /${project}
WORKDIR /${project}

RUN go test -c -fuzz=${fuzzTarget} ${fuzzTargetPkg} -o /${fuzzTarget}

FROM ubuntu:22.04
COPY --from=builder /FuzzRenderBytes /
