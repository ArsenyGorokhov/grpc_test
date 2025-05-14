package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/ArsenyGorokhov/grpc_test/auth/pkg/user"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserServer
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get is here")
	password := gofakeit.Password(true, true, true, true, false, 12)
	return &desc.GetResponse{
		User: &desc.UserInfo{
			Id: gofakeit.Int64(), // Генерируем случайный ID
			Info: &desc.UserData{
				Name:            gofakeit.Name(),
				Email:           gofakeit.Email(),
				Password:        password,
				PasswordConfirm: password, // Используем тот же пароль для подтверждения
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Update is here")
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	log.Printf("Information updated: id=%d, email=%s, name=%s, role=%s",
		req.GetId(),
		req.GetEmail(),
		req.GetName(),
		req.GetRole())
	return &emptypb.Empty{}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println("Create is here")
	log.Printf("Created %v", req.Info)
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println("Delete is here")
	log.Printf("User #%d to be deleleted", req.Id)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
