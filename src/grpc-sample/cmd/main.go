package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/lee212400/myProject/rpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Test Say Hello",
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterSampleServiceServer(s, &server{})
	log.Println("Server running at :50051")

	reflection.Register(s)

	s.Serve(lis)
}
