FMT_TARGET = $(shell find . -type f -name "*.go" -not -path "./vendor/*")
VET_TARGET = $(shell go list ./...)
TEST_TARGET = ./...

default: build

setup:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

dep:
	dep ensure

fmt:
	goimports -w $(FMT_TARGET)

vet:
	go vet $(VET_TARGET)
	golint $(VET_TARGET)

test:
	go test $(TEST_TARGET)

build: vet test
	go build

.PHONY: default setup dep fmt vet test build
