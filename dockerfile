FROM ubuntu:20.04

WORKDIR /app

COPY audit-log /app/

COPY .env .

RUN \
    apt-get update && \
    apt-get install -y ca-certificates && update-ca-certificates

CMD ["./audit-log", "-n"]