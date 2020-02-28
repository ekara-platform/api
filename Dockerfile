FROM ekaraplatform/base:latest

RUN mkdir -p /opt/ekara/bin
COPY rest /opt/ekara/bin/api

ENTRYPOINT ["/opt/ekara/bin/api"]
