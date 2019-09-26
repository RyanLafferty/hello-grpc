#!/bin/bash

if [[ $(command -v protoc) == "" ]]; then
    echo "Please Install protobuf compiler"
fi

if [[ $(command -v grpc_tools_ruby_protoc) == "" ]]; then
    echo "Please Install ruby protobuf compiler"
fi

# output dir
BASE_COMPILED_DIR='./config/protocol_buffers/compiled'
GO_COMPILED_DIR="${BASE_COMPILED_DIR}/go"
RUBY_COMPILED_DIR="${BASE_COMPILED_DIR}/ruby"

# golang
protoc -I ./config/protocol_buffers/helloworld --go_out=plugins=grpc:${GO_COMPILED_DIR}/helloworld ./config/protocol_buffers/helloworld/helloworld.proto
protoc -I ./config/protocol_buffers/store --go_out=plugins=grpc:${GO_COMPILED_DIR}/store ./config/protocol_buffers/store/*.proto

# ruby
grpc_tools_ruby_protoc -I ./config/protocol_buffers/helloworld --ruby_out=${RUBY_COMPILED_DIR}/helloworld --grpc_out=${RUBY_COMPILED_DIR}/helloworld ./config/protocol_buffers/helloworld/helloworld.proto
grpc_tools_ruby_protoc -I ./config/protocol_buffers/store --ruby_out=${RUBY_COMPILED_DIR}/store --grpc_out=${RUBY_COMPILED_DIR}/store  ./config/protocol_buffers/store/*.proto
