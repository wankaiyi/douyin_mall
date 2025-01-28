ARG NACOS_ADDR
ARG NACOS_PASSWORD
ARG env

FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /

COPY . .

WORKDIR /app/auth

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOMAXPROCS=1

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth/ .

ENV NACOS_ADDR=$NACOS_ADDR
ENV NACOS_PASSWORD=$NACOS_PASSWORD

EXPOSE 8888

CMD ["./main"]