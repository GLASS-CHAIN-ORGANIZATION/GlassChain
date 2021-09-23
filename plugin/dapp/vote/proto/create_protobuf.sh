#!/bin/bash
# proto    ， pb.go     types/   , chain33_path    chain33   proto  
chain33_path=$(go list -f '{{.Dir}}' "github.com/33cn/chain33")
protoc --go_out=plugins=grpc:../types ./*.proto --proto_path=. --proto_path="${chain33_path}/types/proto/"
