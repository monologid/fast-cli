
.PHONY: test

default: dep test build install

dep:
	go mod tidy && \
		go get -v

build:
	go build 

install:
	go install

test:
	go test ./... -cover -coverprofile=coverage.txt -covermode=atomic