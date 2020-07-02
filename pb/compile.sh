#!/usr/bin/env bash
protoc -I . log.proto --go_out=plugins=grpc:.
