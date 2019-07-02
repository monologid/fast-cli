
.PHONY: test

default: dep build install

dep:
	go mod tidy && \
		go get -v

build:
	go build 

install:
	go install

test:
	go test ./... -v -cover