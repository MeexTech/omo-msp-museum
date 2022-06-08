.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/museum/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/museum/area.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/museum/booth.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/museum/exhibit.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/museum/anchor.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/museum/sandtable.proto