build: install
	@protoc -I ../../protos --go-validate-readme_out=. --go-validate-readme_opt logtostderr=true ../../protos/validate/validate.proto
	@mv README.md ../../
install:
	@go build -o protoc-gen-go-validate-readme .
	@sudo mv protoc-gen-go-validate-readme /usr/local/bin
