package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		User: &pb.User{
			Email:     "test@test.com",
			FirstName: "first",
			LastName:  "last",
		},
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterSampleServiceServer(s, &server{})
	log.Println("Server running at :50051")
	s.Serve(lis)
}
