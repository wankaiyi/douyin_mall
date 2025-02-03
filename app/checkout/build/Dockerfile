ARG NACOS_ADDR
ARG NACOS_PASSWORD
ARG env

FROM alpine:latest

WORKDIR /

COPY . .

WORKDIR /app/checkout

RUN echo 'https://mirrors.aliyun.com/alpine/v3.16/main' > /etc/apk/repositories && \
    echo 'https://mirrors.aliyun.com/alpine/v3.16/community' >> /etc/apk/repositories && \
    apk update \

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo "$TZ" > /etc/timezone \

ENV NACOS_ADDR=$NACOS_ADDR
ENV NACOS_PASSWORD=$NACOS_PASSWORD

EXPOSE 8888

CMD ["./main"]