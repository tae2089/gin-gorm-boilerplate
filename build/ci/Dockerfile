FROM golang:1.21.3-alpine as base
# ARG GOPATH
RUN apk --no-cache add tzdata ca-certificates
WORKDIR /src/
RUN --mount=type=bind,source=go.mod,target=go.mod \
        --mount=type=bind,source=go.sum,target=go.sum \
        --mount=type=cache,target=$GOPATH/pkg/mod/ \
        go mod download -x

FROM base as stage
WORKDIR /src/
RUN --mount=type=cache,target=/go/pkg/mod/ \
        --mount=type=bind,target=. \
        CGO_ENABLED=0  go build -o /bin/main run/main.go
    
FROM scratch
COPY --from=stage /bin/main /bin/main
COPY ./templates /templates
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV TZ=Asia/Seoul
ENTRYPOINT ["/bin/main"]
