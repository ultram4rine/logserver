#!/usr/bin/env bash
protoc -I . --go_out . --go-grpc_out . \
    --go_opt paths=source_relative \
    --go-grpc_opt paths=source_relative log.proto
protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative log.proto
protoc -I . --openapiv2_out . \
    --openapiv2_opt logtostderr=true log.proto
