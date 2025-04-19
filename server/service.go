package main

import (
	"context"
	"log"

	pb "github.com/spiffgreen/learn-grpc/learn"
)

type HelloApp struct {
	pb.UnimplementedHelloAppServer
}

func (h HelloApp) GetUser(_ context.Context, input *pb.UserRequest) (*pb.User, error) {
	log.Println("Received input from user", input.GetId())
	return &pb.User{
		Id:   input.Id,
		Name: "Spiff Jekey-Green",
		Age:  23,
	}, nil
}

func (h HelloApp) GetUsers(_ *pb.EmptyMessage, srv pb.HelloApp_GetUsersServer) error {
	users := []*pb.User{
		{
			Id:   1,
			Name: "Spiff Jekey-Green",
			Age:  23,
		},
		{
			Id:   2,
			Name: "John Doe",
			Age:  30,
		},
	}

	for _, user := range users {
		if err := srv.Send(user); err != nil {
			log.Println("error generating response", user)
		}
	}

	return nil
}
