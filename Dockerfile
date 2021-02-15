FROM restreamio/gstreamer:latest-dev

RUN apt-get update && \
    apt-get dist-upgrade -y && \
    apt-get install -y --no-install-recommends \ 
    golang-go \ 
    nodejs \ 
    curl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV GOPATH /root/go
ENV PATH $GOPATH/bin:$PATH
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.36.0

WORKDIR /opt