# This Dockerfile builds a webapp that listens on port 8080 and responds with "Hello, World"
FROM ubuntu:20.04

EXPOSE 8080

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -y busybox

RUN echo 'Hello, World!' > index.html

CMD ["busybox", "httpd", "-f", "-p", "8080"]
