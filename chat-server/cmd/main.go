package main

import (
	desc "github.com/ArsenyGorokhov/grpc_test/chat-server/pkg/chat"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServer
}

func main() {

}
