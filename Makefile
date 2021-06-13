GOPATH=$(shell go env GOPATH)
BINNAME="pulumi-do-polkadot"

.PHONY: deps clean build
deps:
	@go mod vendor
clean:
	@rm -v $(BINNAME)
	@go clean -i .
build:
	@go build -o $(BINNAME)
