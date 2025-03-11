SERVICE_PATH=rkosykh/wallet

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin
LOCAL_BIN:=$(CURDIR)/bin
BUF_BIN:=$(LOCAL_BIN)/buf

.PHONY: bin-deps
bin-deps: bin-deps-test
	$(info Installing bin dependencies...)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.16.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.16.2
	GOBIN=$(LOCAL_BIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.4.0
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.26.0

bin-deps-test:
	$(info Installing bin dependencies for test...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.21.1
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest

.PHONY: deps
deps:
	$(info Install dependencies...)
	go mod download
	go mod tidy

.PHONY: generate
generate: bin-deps
	$(BUF_BIN) generate

.PHONY: build
build: deps generate
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-o ./bin/wallet$(shell go env GOEXE) ./cmd/wallet/main.go