.PHONY: build generate tools

# Setting GOBIN forces 'go install' commands to build binaries into ./bin for 
# ease of management.
export GOBIN := $(PWD)/bin/
# Setting PATH allows us to use the tools we have installed in 'go 
export PATH := $(PATH):$(GOBIN)
SHELL := env PATH=$(PATH) /bin/bash

clean:
	rm -rf ./bin

generate: tools
	buf lint
	buf format -w
	buf generate

# All tools should be pinned in 'cmd/tools'. This will force an entry to be
# added to go.mod, which 'go install' will respect, allowing us to version our
# dependencies correctly.
tools:
	go install github.com/bufbuild/buf/cmd/buf
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

fmt: generate
	go mod tidy
	go fmt ./...

build:
	go build ./...
