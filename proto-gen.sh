#!/bin/bash

protoc -I=. --go-grpc_out=require_unimplemented_servers=false:./internal --go_out=:./internal protos/genericsapi.proto
