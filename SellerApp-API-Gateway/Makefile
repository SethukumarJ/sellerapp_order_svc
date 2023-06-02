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