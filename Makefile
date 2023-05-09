GO_INSTALLED := $(shell which go)
PROTOC_INSTALLED := $(shell which protoc)
PROTOC_GEN_GO_INSTALLED := $(shell which protoc-gen-go)

GO_FILES = $(shell go list ./... | grep -v /vendor/ | grep -v /api/ | grep -v /cmd/)

tidy-vendor tidy:
	@go mod vendor
	@go mod tidy

install-tools install:
ifndef GO_INSTALLED
	$(error "go is not installed, please run 'brew install go'")
endif
ifndef PROTOC_INSTALLED
	$(error "protoc is not installed, please run 'brew install protobuf'")
endif
ifndef PROTOC_GEN_GO_INSTALLED
	go install github.com/golang/protobuf/protoc-gen-go
endif

generate gen: install-tools
	@rm -rf pkg/proto/kbtg/*
	@protoc \
		-I. \
		-I./third_party/googleapis/ \
		--go_out=plugins=grpc:pkg/proto/ \
		api/server/*.proto
	@echo "Success! Generated server stub"

build: generate
	@rm -rf bin
	@mkdir -p bin
	@go build -o bin/server cmd/server/*.go
	@echo "Success! Binaries can be found in 'bin' dir"

vet:
	@go vet $(GO_FILES)

fmt:
	@go fmt $(GO_FILES)

test:
	@go test $(GO_FILES) -cover --race

