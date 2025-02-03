ARG NACOS_ADDR
ARG NACOS_PASSWORD
ARG env

FROM alpine:latest

WORKDIR /

COPY . .

WORKDIR /app/order

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo "$TZ" > /etc/timezone \

ENV NACOS_ADDR=$NACOS_ADDR
ENV NACOS_PASSWORD=$NACOS_PASSWORD

EXPOSE 8888

CMD ["./main"]