# workspace (GOPATH) configured at /go
FROM golang:1.15 as builder


#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/macbro/mb_corporate_service
WORKDIR $GOPATH/src/gitlab.udevs.io/macbro/mb_corporate_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/mb_corporate_service /

ENTRYPOINT ["/mb_corporate_service"]