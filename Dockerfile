FROM strayard/base

RUN yum install -y golang && yum clean all

WORKDIR /go/src/github.com/strayard/strayard
ADD .   /go/src/github.com/strayard/strayard
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN go get github.com/strayard/strayard && \
    hack/build-go.sh && \
    cp _output/local/go/bin/* /usr/bin/ && \
    mkdir -p /var/lib/strayard

EXPOSE 8080 8443
WORKDIR /var/lib/strayard
ENTRYPOINT ["/usr/bin/strayard"]
