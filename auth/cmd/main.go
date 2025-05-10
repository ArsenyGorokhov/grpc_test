package main

import (
	desc "github.com/ArsenyGorokhov/grpc_test/auth/pkg/user"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserServer
}

func main() {

}
