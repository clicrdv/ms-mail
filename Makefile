BINARY=ms-mail
VERSION:=0.1.3

.PHONY: all
.DEFAULT_GOAL := all

test:
	go test  -v ./...

get:
	go get

docker:
	docker build -t clicrdv/${BINARY}:${VERSION} .

tag-latest:
	docker tag clicrdv/${BINARY}:${VERSION} clicrdv/${BINARY}:latest

binary:
	go build -o ${BINARY}-osx main.go
	env GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux main.go

all: get test binary docker tag-latest
