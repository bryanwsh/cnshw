FROM golang:1.19.3-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /build
COPY . /build/
RUN go build -o http_server http_server_week2/main.go
WORKDIR /dist
RUN cp /build/http_server /dist/http_server
FROM ubuntu:latest
RUN mkdir /app
COPY --from=builder /dist/http_server /app/http_server
RUN chmod +x /app/http_server
WORKDIR /app
EXPOSE 80
CMD [ "/app/http_server" ]