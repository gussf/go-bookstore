.PHONY: build lint all

DEFAULT: all

lint:
	golangci-lint run

build:
	go build main.go

run: lint build
	./main



all: build lint