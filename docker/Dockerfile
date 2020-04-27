FROM golang:1.14

RUN go get github.com/cespare/reflex
RUN go get github.com/go-swagger/go-swagger/cmd/swagger

# make the go module cache folder writable by all users otherwise linux users will not be able to write to it
RUN mkdir -p /go/pkg/mod/cache && chmod a+rw -R /go/pkg/mod/

# gRPC dependencies
RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

RUN apt-get update -y && \
    apt-get install unzip

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip && \
    unzip protoc-3.11.4-linux-x86_64.zip -d /protoc/
ENV PATH="$PATH:/protoc/bin"

COPY reflex.conf /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
