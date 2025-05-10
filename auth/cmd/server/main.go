package main

import (
	"context"
	"fmt"

	desc "github.com/ArsenyGorokhov/grpc_test/auth/pkg/user"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserServer
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get Handler been used")
	return &desc.GetResponse{}, nil
}

func main() {

}
