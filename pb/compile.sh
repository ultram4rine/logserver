#!/usr/bin/env bash
protoc -I. --go_out=plugins=grpc,paths=source_relative:. log.proto
protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:. log.proto