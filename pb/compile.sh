#!/usr/bin/env bash
protoc -I. --go_out=plugins=grpc,paths=source_relative:. log.proto