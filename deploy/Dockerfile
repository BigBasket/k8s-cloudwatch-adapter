FROM alpine:3.9 as alpine

RUN apk add -U --no-cache ca-certificates

FROM busybox
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY adapter /
USER 1001:1001
ENTRYPOINT ["/adapter", "--logtostderr=true"]

# FROM golang:1.14

# RUN apt-get update \
#     && apt install -y ca-certificates

# WORKDIR /build

# COPY . .

# RUN GOPRIVATE=github.com/bigbasket CGO_ENABLED=0 go build -o adapter cmd/adapter/adapter.go

# FROM busybox
# COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=0 adapter /
# USER 1001:1001
# ENTRYPOINT ["/adapter", "--logtostderr=true"]