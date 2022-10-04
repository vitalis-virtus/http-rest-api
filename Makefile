.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: lint
lint:
	golangci-lint run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast:
	golangci-lint run ./... --fast --config=./.golangci.yml

.DEFAULT_GOAL := build
