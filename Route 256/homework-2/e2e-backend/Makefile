LOCAL_BIN:=$(CURDIR)/bin
LOCAL_PROTO:=$(CURDIR)/proto
PB_PATH=./internal/pb

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go get google.golang.org/grpc
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

grpc:
	protoc --proto_path=$(LOCAL_PROTO) wallet.proto --go-grpc_out=${PB_PATH}  --go_out=${PB_PATH}

gproto: PROTO=$PROTO
gproto:
	./gproto.sh ${PROTO}

test:
	go test ./tests/...