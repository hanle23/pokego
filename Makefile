.DEFAULT_GOAL := build

test:
	go test ./... -cover -v
.PHONY: test

fmt: test
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go build ./...
.PHONY: build
