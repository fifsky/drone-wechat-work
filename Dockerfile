FROM golang:1.23-alpine as builder
WORKDIR /build
COPY . .
RUN go build -ldflags "-s -w" -o drone-wechat .

FROM alpine:latest
RUN apk update && \
  apk add \
  ca-certificates && \
  rm -rf /var/cache/apk/*

COPY --from=builder /build/drone-wechat /bin/
ENTRYPOINT ["/bin/drone-wechat"]