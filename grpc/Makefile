mod:
	go mod init github.com/chepkoyallan/protos

get:
	go get

gen:
	protoc --proto_path=. --go-grpc_out=/Users/allanchepkoy/Documents/projects/rflow/protos/grpc protos/greet.proto

gengrpc:
	protoc --go_out=plugins=grpc:. protos/*.proto

server:
	go run cmd/person_server/main.go

book:
	go run cmd/book/main.go

channel:
	go run cmd/chans/main.go

client:
	go run cmd/person_client/main.go

greet-server:
	go run cmd/greet_server/main.go

greet-client:
	go run cmd/greet_client/main.go

person:
	go run cmd/person/main.go

check-version:
	go version

help:
	go doc {}