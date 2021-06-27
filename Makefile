compile-python:
	protoc --proto_path=. --python_out=. proto/*.proto

