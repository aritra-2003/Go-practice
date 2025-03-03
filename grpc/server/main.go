package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "example.com/grpc/proto" // Replace with actual path
)

const port = ":8080"

// Implementing the gRPC service
type helloServer struct {
	pb.GreetServcieServer // Correct way to embed the service
}

func main() {
	// Start TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	// Create a single gRPC server
	grpcServer := grpc.NewServer()

	// Register the gRPC service
	pb. RegisterGreetServcieServer(grpcServer, &helloServer{})

	log.Printf("Server started at %v", lis.Addr())

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}