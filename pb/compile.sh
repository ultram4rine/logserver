#!/usr/bin/env bash
protoc -I . log.proto --go_out=plugins=grpc:.
protoc -I . log.proto --js_out=import_style=commonjs:../web/client/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../web/client/src/
