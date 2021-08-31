.PHONY: build lint all

DEFAULT: all

lint:
	@go fmt
	@go vet
	@golangci-lint run

build:
	@go build main.go

run: lint build
	@./main



all: build lint