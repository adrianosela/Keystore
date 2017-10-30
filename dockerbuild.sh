#!/bin/bash +x

GOOS=linux GOARCH=amd64 go build -a -o Keystore

docker build -t keystore .
