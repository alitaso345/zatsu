SHELL := /bin/bash

.PHONY: generate-go
generate-go:
	protoc -I./proto/ ./proto/user.proto --go_out=plugins=grpc:./proto

.PHONY: generate-ts
generate-ts:
	protoc -I./proto/ ./proto/user.proto --js_out=import_style=commonjs:./proto \
		   --grpc-web_out=import_style=typescript,mode=grpcwebtext:./proto
	cp ./proto/*.js ./client/proto/
	cp ./proto/*.ts ./client/proto/
