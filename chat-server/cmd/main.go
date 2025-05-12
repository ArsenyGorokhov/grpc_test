package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	desc "github.com/ArsenyGorokhov/grpc_test/chat-server/pkg/chat"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServer
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Users: %v", req.Usernames)
	resp := &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}
	return resp, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	log.Printf("User %v has been deleted", req.Id)
	return nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	log.Printf("[%v] %s: %s", req.Timestamp.AsTime().Format(time.DateTime), req.Username, req.Text)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
}
