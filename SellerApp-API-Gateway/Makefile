SHELL := /bin/sh

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go

swag: ## Generate swagger docs
	swag init -g cmd/main.go -o ./cmd/docs

swagger: ##Insatall swagger
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest


proto1:
	wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip

proto2:
	unzip protoc-3.19.4-linux-x86_64.zip -d ~/protobuf
proto3:
	export PATH="$PATH:$HOME/protobuf/bin"
proto4:
	source ~/.bashrc
proto5:
	protoc --version
proto6:
	go clean -i github.com/golang/protobuf/protoc-gen-go
proto7:
	go install github.com/golang/protobuf/protoc-gen-go