package main

import (
	"context"
	"fmt"
	"log"

	"time"

	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendMessages(stream pb.SampleService_ChatToWayClient) {
	for i := 0; i < 5; i++ {
		msg := &pb.Chat{
			User:    fmt.Sprintf("Client%d", i),
			Message: fmt.Sprintf("Message %d", i),
		}
		if err := stream.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func receiveMessages(stream pb.SampleService_ChatToWayClient) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		fmt.Printf("Received from server: %s\n", resp.Message)
	}
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSampleServiceClient(conn)

	stream, err := client.ChatToWay(context.Background())
	if err != nil {
		log.Fatalf("Failed to start stream: %v", err)
	}

	go sendMessages(stream) // 非同期でメッセージ送信
	receiveMessages(stream) // servcerから受信
}
