.PHONY: build generate buf

fmt: generate
	go mod tidy
	go fmt ./...

build:
	go build ./...

export GOBIN := $(PWD)/bin/
export PATH := $(PATH):$(GOBIN)
SHELL := env PATH=$(PATH) /bin/bash

generate: buf ./cloudapi
	buf lint
	buf format -w
	buf generate

buf:
	go install github.com/bufbuild/buf/cmd/buf
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
