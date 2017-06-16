#!/bin/bash

cd vendor/github.com/golang/protobuf/protoc-gen-go
go build -o ../../../../../bin/protoc-gen-go
cd ../../../../../
ls bin
protoc --proto_path=proto --go_out=plugins=grpc:template/ proto/*.proto
