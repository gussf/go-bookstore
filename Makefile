.PHONY: build lint all

DEFAULT: all

lint:
	@go fmt ./...
	@go vet ./...
	@golangci-lint run

build:
	@go test -v ./...
	@go build main.go

run: 
	@go run main.go



all: build lint