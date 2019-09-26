FROM golang:1.13-buster

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/gorilla/mux

COPY ./clients/go /app/
COPY ./config/protocol_buffers/compiled/go /app
WORKDIR /app

EXPOSE 8080
CMD ["go", "run", "main.go"]
