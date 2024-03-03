FROM golang:1.21-alpine as builder
COPY . /app/
WORKDIR /app
RUN go build

FROM alpine:3.19
RUN apk add mariadb-client openssl docker-compose tzdata
COPY --from=builder /app/automanage /usr/local/bin
COPY ./config.yaml.default /etc/machhub/config.yaml
COPY ./generatesmtp /app/generatesmtp
COPY ./web_template.sql /app/web_template.sql
WORKDIR /app
CMD automanage