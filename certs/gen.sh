#!/usr/bin/env bash
set -e

openssl genrsa -out ca.key 4096
openssl req -new -x509 -key ca.key -sha256 -days 365 -out ca.crt -config certificate.conf 

openssl genrsa -out logserver.key 4096
openssl req -new -key logserver.key -out logserver.csr -config certificate.conf
openssl x509 -req -in logserver.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out logserver.pem -days 365 -sha256 -extfile certificate.conf -extensions req_ext
