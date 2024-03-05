default: generate build test

run:
	go run cmd/main.go

build:
	go build -o app cmd/main.go

test:
	go test ./...

bin-deps:
	mkdir bin/
	go install google.golang.org/protobuf/cmd/protoc-gen-go

mocks:
	@echo "Generating mocks..."
	@echo "Nothin right now"

proto-deps:
	@echo "No proto dependencies"

proto: proto-deps
	@echo "Generating protbuf..."
	mkdir -p pb
	protoc --proto_path=api --go_out=pb/ cred-checker.proto
	protoc --proto_path=api --go-grpc_out=pb/ cred-checker.proto

generate: proto mocks
