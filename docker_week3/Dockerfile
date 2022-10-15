FROM ubuntu:latest
ENV VERSION=123456
RUN mkdir /app
COPY http_server /app/http_server
RUN chmod +x /app/http_server
WORKDIR /app
EXPOSE 80
CMD [ "/app/http_server" ]