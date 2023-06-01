# FROM ubuntu:20.04
# RUN \
#   apt-get update && \
#   DEBIAN_FRONTEND=noninteractive apt install -y mysql-server && \
#   service mysql start
# WORKDIR /app
# COPY audit-log .
# EXPOSE 3306
# CMD ["service mysql start && ./audit-log"]

FROM ubuntu:20.04

WORKDIR /app

COPY audit-log /app/

COPY .env .

# Install necessary packages
RUN \
    apt-get update && \
    # apt-get install -y supervisor && \
    apt-get install -y ca-certificates && update-ca-certificates

CMD ["./audit-log", "-n"]