build/validate:
	@protoc -I protos --go_out=. --go_opt=module=github.com/picatz/protoc-gen-go-validate protos/validate/validate.proto

build/example:
	@protoc -I protos --go_out=. --go-validate_out=. --go-validate_opt logtostderr=true example.proto

build/readme:
	@cd cmd/protoc-gen-go-validate-readme && go build -o protoc-gen-go-validate-readme . && mv protoc-gen-go-validate-readme /usr/local/bin 
	@protoc -I protos --go-validate-readme_out=. --go-validate-readme_opt logtostderr=true protos/validate/validate.proto

build/protos: install build/validate build/example

build/docker/protoc:
	@docker build -t go-validate-protoc -f Dockerfile.protoc .

build/docker/protos:
	@docker run --rm -v $(CURDIR):/workdir go-validate-protoc:latest make build/protos

build/docker/readme:
	@docker run --rm -v $(CURDIR):/workdir go-validate-protoc:latest make build/readme

install:
	@go build -o protoc-gen-go-validate .
	@mv protoc-gen-go-validate /usr/local/bin

test:
	@go test -timeout 30s -run ^TestValidate github.com/picatz/protoc-gen-go-validate/pkg/example -v