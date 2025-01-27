FROM golang:1.21.0-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /

COPY . .

WORKDIR /app/checkout

RUN go mod tidy && go build -o exe

EXPOSE 8888

CMD ["./exe"]