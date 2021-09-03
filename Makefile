.PHONY: build lint all

DEFAULT: all

lint:
	@go fmt ./...
	@go vet ./...
	@golangci-lint run

build:
	@go test -v ./...
	@go build ./...

run: 
	@go run ./...



all: build lint