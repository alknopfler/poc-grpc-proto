#!/bin/bash

cd vendor/github.com/golang/protobuf/protoc-gen-go
go build -o ../../../../../bin/protoc-gen-go
cd ../../../../../
ls bin
#protoc --proto_path=proto --python_out=template proto/*.proto
python3.6 -m grpc_tools.protoc -Iproto --python_out=template --grpc_python_out=template proto/*.proto