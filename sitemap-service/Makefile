PROTO_DIR=../protos/content-service
PROTO_FILE=$(PROTO_DIR)/content.proto
GO_OUT=./src/protos

.PHONY: proto

proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)
