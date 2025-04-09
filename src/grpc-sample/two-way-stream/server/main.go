package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) ChatToWay(stream pb.SampleService_ChatToWayServer) error {
	// context活用して、cancel,timout制御可能
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			log.Println("Client canceled or timed out:", ctx.Err())
			return ctx.Err()

		default:
			req, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("Error receiving message: %v", err)
				return err
			}

			log.Printf("Received message: %s from %s", req.Message, req.User)

			res := &pb.Chat{
				User:    "Server",
				Message: "Echo: " + req.Message,
			}
			if err := stream.Send(res); err != nil {
				log.Printf("Error sending message: %v", err)
				return err
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	g := grpc.NewServer()
	pb.RegisterSampleServiceServer(g, &server{})

	fmt.Println("Server is running on port 50051")
	if err := g.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
