FROM golang:1.13-alpine as builder
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
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