#!/bin/bash

protoc \
  --plugin="/home/mjnk9xw/go/bin/protoc-gen-go-grpc" \
  --plugin="/home/mjnk9xw/go/bin/protoc-gen-openapiv2" \
  --plugin="/home/mjnk9xw/go/bin/protoc-gen-grpc-gateway" \
  --proto_path=../../mezon-server/common \
  --go_out=../mezon-protobuf \
  --go-grpc_out=../mezon-protobuf \
  --grpc-gateway_out=../mezon-protobuf \
  ../../mezon-server/common/rtapi/realtime.proto \
  ../../mezon-server/common/api/api.proto


