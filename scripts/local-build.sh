#!/bin/bash

cd ../src
GOARCH=amd64 GOOS=linux go build -o ../bin/main . && cp -r resources ../bin/ && cd ../bin

zip -r main.zip .
