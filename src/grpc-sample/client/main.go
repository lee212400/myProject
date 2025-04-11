package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSampleServiceClient(conn)

	md := metadata.New(map[string]string{
		"authorization": "Bearer abc123",
		"lang":          "ja",
		"x-request-id":  uuid.New().String(),
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ctx4, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := client.GetUser(ctx4, &pb.GetUserRequest{Id: "testId"})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("GetUser Response:", res)
	}
}
