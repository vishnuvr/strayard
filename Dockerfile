FROM golang:1.4

RUN apt-get update && rm -rf /var/lib/apt/lists/*

ENV STRAYARD_DIR /go/src/github.com/strayard/strayard
ENV GOPATH $STRAYARD_DIR/Godeps/_workspace:$GOPATH
ENV PATH $GOPATH/bin:$PATH

WORKDIR $STRAYARD_DIR
COPY . $STRAYARD_DIR
RUN make clean build

EXPOSE 8080 
ENTRYPOINT ["dashboard"]
