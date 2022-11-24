FROM golang:1.19.3-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /build
COPY . /build/
RUN go build -o http_server http_server_week2/main.go
RUN mkdir /app
RUN cp /build/http_server /app/http_server
WORKDIR /app
CMD [ "/app/http_server" ]