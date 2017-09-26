BINARY=ms-mail
VERSION:=0.1.0

.PHONY: all
.DEFAULT_GOAL := all

test:
	go test  -v ./...

go-stub:
	 protoc -I mailservice/ mailservice/msmail.proto --go_out=plugins=grpc:mailservice

ruby-stub:
	grpc_tools_ruby_protoc -I mailservice --ruby_out=ruby-client --grpc_out=ruby-client mailservice/msmail.proto

get:
	go get

docker:
	docker build -t clicrdv/${BINARY}:${VERSION} .

tag-latest:
	docker tag clicrdv/${BINARY}:${VERSION} clicrdv/${BINARY}:latest

binary:
	go build -o ${BINARY}-osx main.go
	env GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux main.go

all: get test go-stub ruby-stub binary docker tag-latest
