# https://github.com/protocolbuffers/protobuf
PROTOC_VERSION ?= 3.19.1

# https://github.com/golang/protobuf
PROTOC_GEN_GO_VERSION ?= 1.5.2

# https://github.com/grpc/grpc-go
PROTOC_GEN_GO_GRPC_VERSION ?= 1.0.1

build/validate:
	@protoc -I protos --go_out=. --go_opt=module=github.com/picatz/protoc-gen-go-validate protos/validate/validate.proto
build/example:
	@protoc -I protos --go_out=. --go-validate_out=. --go-validate_opt logtostderr=true example.proto
build: build/validate install build/example test
test:
	@go test -timeout 30s -run ^TestValidate github.com/picatz/protoc-gen-go-validate/pkg/example -v
install:
	@go build -o protoc-gen-go-validate .
	@sudo mv protoc-gen-go-validate /usr/local/bin
install/protoc:
	@echo "• Installing protoc v${PROTOC_VERSION}"
	@curl -sLO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
	@unzip -q -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d ~/.local
	@rm protoc-${PROTOC_VERSION}-linux-x86_64.zip
	@echo "• Installing protoc-gen-go v${PROTOC_GEN_GO_VERSION}"
	@go install github.com/golang/protobuf/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
	@echo "• Installing protoc-gen-go-grpc v${PROTOC_GEN_GO_GRPC_VERSION}"
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}
	@echo "✓ Finnished install protoc dependencies"
