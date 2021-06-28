compile-python:
	protoc --proto_path=. --python_out=. proto/*.proto

compile-go:
	protoc --proto_path=. --go_out=proto/ proto/*.proto