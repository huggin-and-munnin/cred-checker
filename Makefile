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
	go install github.com/golang/mock/mockgen@v1.6.0

mocks:
	@echo "Generating mocks..."
	mockgen -source=internal/use_cases/get_credentials/use_case.go -package mocks -destination internal/use_cases/get_credentials/mocks/mocks.go &
	mockgen -source=internal/app/cred_checker/service.go -package mocks -destination internal/app/cred_checker/mocks/mocks.go


proto-deps:
	@echo "No proto dependencies"

proto: proto-deps
	@echo "Generating protbuf..."
	mkdir -p pb
	protoc --proto_path=api --go_out=pb/ cred-checker.proto
	protoc --proto_path=api --go-grpc_out=pb/ cred-checker.proto
	protoc --proto_path=api --go_out=pb/ health.proto
	protoc --proto_path=api --go-grpc_out=pb/ health.proto

generate: proto mocks
