#!/bin/bash

protoc -I ./helloworldpb \
   --go_out ./helloworldpb --go_opt paths=source_relative \
   --go-grpc_out ./helloworldpb --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./helloworldpb --grpc-gateway_opt paths=source_relative \
   ./helloworldpb/hello_world.proto