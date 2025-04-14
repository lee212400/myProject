package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/lee212400/myProject/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// gRPCサーバー接続
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// gRPC-Gateway router設定
	gwmux := runtime.NewServeMux(
		// httpHeader->metadata
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
	)
	err = pb.RegisterSampleServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway handler: %v", err)
	}

	// HTTPサーバー開始 (gRPC-Gateway)
	http.Handle("/", gwmux)
	fmt.Println("gRPC-Gateway server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func customHeaderMatcher(key string) (string, bool) {
	switch key {
	case "Authorization":
		return "authorization", true
	case "X-Trace-Id":
		return "x-trace-id", true
	case "X-Request-Id":
		return "x-request-id", true
	case "Lang":
		return "lang", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
