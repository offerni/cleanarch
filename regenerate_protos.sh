#!/bin/bash

protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
echo "Protobuf files generated successfully!"
