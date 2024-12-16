#!/bin/bash


# change plugin path
# example
# PLUGIN_PATH="/home/minhnv/go/bin"
PLUGIN_PATH=""

protoc \
  --plugin="${PLUGIN_PATH}/protoc-gen-go-grpc" \
  --plugin="${PLUGIN_PATH}/protoc-gen-openapiv2" \
  --plugin="${PLUGIN_PATH}/protoc-gen-grpc-gateway" \
  --proto_path=../../mezon-server/common \
  --go_out=../mezon-protobuf \
  --go-grpc_out=../mezon-protobuf \
  --grpc-gateway_out=../mezon-protobuf \
  ../../mezon-server/common/rtapi/realtime.proto \
  ../../mezon-server/common/api/api.proto


