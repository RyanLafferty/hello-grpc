#/bin/bash

go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

echo 'Please ensure that $GOPATH/bin is added to your path'
echo 'Example export PATH=$PATH:$GOPATH/bin'

