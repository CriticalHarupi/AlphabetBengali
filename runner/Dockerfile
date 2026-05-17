FROM ghcr.io/actions/actions-runner:2.334.0

USER root

RUN apt-get update && apt-get install -y wget ca-certificates && \
    wget https://go.dev/dl/go1.26.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.26.2.linux-amd64.tar.gz && \
    rm go1.26.2.linux-amd64.tar.gz && \
    apt-get clean

ENV PATH=$PATH:/usr/local/go/bin

USER runner
