FROM golang:1.13-buster

RUN go get -u google.golang.org/grpc
RUN go get -u cloud.google.com/go/firestore

COPY ./servers/go /app/
COPY ./config/protocol_buffers/compiled/go /app
WORKDIR /app

CMD ["go", "run", "main.go"]
