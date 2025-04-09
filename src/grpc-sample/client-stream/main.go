package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/lee212400/myProject/rpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := c.ChatStream(ctx)
	if err != nil {
		log.Fatalf("could not initiate stream: %v", err)
	}

	messages := []string{
		"Hello client",
		"How are you",
		"Goodbye",
	}

	for _, message := range messages {
		if err := stream.Send(&pb.ChatRequest{
			User:    "User1",
			Message: message,
		}); err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}

	fmt.Printf("Received from server: %s - %s\n", resp.ResUser, resp.ResMessage)
}
