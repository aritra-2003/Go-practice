package main

import (
	
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/grpc/proto"  // Replace with actual path
)

const port = ":8080"

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // Corrected

	// Create a client instance
	client := pb.NewGreetServcieClient(conn)

	// Define a NamesList
	// names := &pb.NamesList{
	// 	Names: []string{"Akhil", "Alice", "Bob"},
	// }

	// Call the SayHello function
	callSayHello(client)
}

// Example function for calling SayHello
