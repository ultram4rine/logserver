#!/usr/bin/env bash
openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out logcertificate.crt -keyout logkey.key
