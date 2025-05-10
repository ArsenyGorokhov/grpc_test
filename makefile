# Variables
LOCAL_BIN := $(CURDIR)/bin
PROTOC := protoc

# Auth service paths
AUTH_PROTO_DIR = auth/api/user
AUTH_PROTO_FILE = $(AUTH_PROTO_DIR)/user.proto
AUTH_GO_OUT = auth/pkg/user

# Chat service paths
CHAT_PROTO_DIR = chat-server/api/chat
CHAT_PROTO_FILE = $(CHAT_PROTO_DIR)/chat.proto
CHAT_GO_OUT = chat-server/pkg/chat

.PHONY: install-deps get-deps generate generate-auth generate-chat-server clean

install-deps:
	@echo "Installing dependencies..."
	@GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	@GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

get-deps:
	@echo "Getting dependencies..."
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	@go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate: generate-auth generate-chat-server

generate-auth: install-deps
	@echo "Generating Auth service code..."
	@mkdir -p $(AUTH_GO_OUT)
	@$(PROTOC) --proto_path=$(AUTH_PROTO_DIR) \
	--go_out=$(AUTH_GO_OUT) --go_opt=paths=source_relative \
	--go-grpc_out=$(AUTH_GO_OUT) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	$(AUTH_PROTO_FILE)

generate-chat-server: install-deps
	@echo "Generating Chat service code..."
	@mkdir -p $(CHAT_GO_OUT)
	@$(PROTOC) --proto_path=$(CHAT_PROTO_DIR) \
	--go_out=$(CHAT_GO_OUT) --go_opt=paths=source_relative \
	--go-grpc_out=$(CHAT_GO_OUT) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	$(CHAT_PROTO_FILE)

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(AUTH_GO_OUT) $(CHAT_GO_OUT)