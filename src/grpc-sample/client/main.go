package main

import (
	"context"
	"io"
	"log"

	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSampleServiceClient(conn)

	stream, _ := client.GetData(context.Background(), &pb.StreamRequest{})

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream recv faild: %v", err)
		}
		log.Printf("name:%s, email:%s", msg.Name, msg.Email)
	}
}
