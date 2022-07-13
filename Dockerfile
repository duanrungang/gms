FROM golang:alpine as builder

WORKDIR /go/src/gms
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="duanrungang@163.com"

WORKDIR /go/src/gms

COPY --from=0 /go/src/gms ./
COPY --from=0 /go/src/gms/resource ./resource/
COPY --from=0 /go/src/gms/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
