#!/bin/bash

cd proto

~/Downloads/protoc/bin/protoc --go_out=. --go-grpc_out=. user.proto
