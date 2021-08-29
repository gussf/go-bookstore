.PHONY: build lint all

DEFAULT: all

lint:
	golangci-lint run .

build:
	go build main.go

run: lint build
	. ./.env
	./main



all: build lint