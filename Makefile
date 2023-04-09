build/validate:
	@protoc -I protos --go_out=. --go_opt=module=github.com/picatz/protoc-gen-go-validate protos/validate/validate.proto

build/example:
	@protoc -I protos --go_out=. --go-validate_out=. --go-validate_opt logtostderr=true example.proto

build/protos: install build/validate build/example

build/docker/protoc:
	@docker build -t go-validate-protoc -f Dockerfile.protoc .

build/docker/protos:
	@docker run --rm -v $(CURDIR):/workdir go-validate-protoc:latest make build/protos

install:
	@go build -o protoc-gen-go-validate .
	@mv protoc-gen-go-validate /usr/local/bin

test:
	@go test -timeout 30s -run ^TestValidate github.com/picatz/protoc-gen-go-validate/pkg/example -v