.PHONY: all
all: build test examples

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: examples
examples:
	go build -o bin/helloworld ./example/helloworld
