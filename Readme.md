# ClicRDV Mail Microservice

This project is a PoC of a golang microsevice exposing gRPC interface.

it also include an example ruby client


## Development requirements

protoc : go get go get -u github.com/golang/protobuf/protoc-gen-go
grpc_tool_ruby : gem install grpc && gem install grpc-tools

## Running test

make test

## Generating stubs

If you modify the proto description in mailservice/msmail.proto you should regenerate the stubs.

make go-stub

make ruby-stub

## Build binaries

make binary

This will make a linux and a mac os (ms-mail-osx) binary.

## Make everything

make

will run all the command in this order :
get
test
go-stub
ruby-stub
binary

