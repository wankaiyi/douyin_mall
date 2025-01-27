FROM golang:1.21.0-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /

COPY . .

WORKDIR /app/user

RUN go mod tidy && go build -o exe

EXPOSE 8001

CMD ["./exe"]