package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/lee212400/myProject/rpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) GetData(req *pb.StreamRequest, stream pb.SampleService_GetDataServer) error {
	for i := 0; i <= 100; i++ {
		res := &pb.StreamResponse{
			Name:  fmt.Sprintf("name%d", i),
			Email: fmt.Sprintf("sample%d@test.com", i),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterSampleServiceServer(s, &server{})
	log.Println("Server running at :50051")
	s.Serve(lis)
}
