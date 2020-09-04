#!/usr/bin/env bash
protoc -I. --go_out=plugins=grpc:. log.proto
protoc -I. --grpc-gateway_out=logtostderr=true:. log.proto
protoc -I. --swagger_out=logtostderr=true:. log.proto