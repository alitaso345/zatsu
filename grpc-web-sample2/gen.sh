#!/bin/sh

CLIENT_OUTPUT_DIR=client/src/helloworld
SERVER_OUTPUT_DIR=server/helloworld

mkdir -p ${CLIENT_OUTPUT_DIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=./proto ./proto/helloworld.proto \
       --js_out=import_style=commonjs:${CLIENT_OUTPUT_DIR} \
       --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTPUT_DIR} \
       --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
