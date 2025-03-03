package main

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc/proto"
)

// Function to call SayHello RPC
func callSayHello(client pb.GreetServcieClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // Corrected

	res, err := client.SayHello(ctx, &pb.NoParam{}) // Corrected
	if err != nil {
		log.Fatalf("could not greet: %v", err) // Corrected
	}

	log.Printf("%s", res.Message) // Corrected
}