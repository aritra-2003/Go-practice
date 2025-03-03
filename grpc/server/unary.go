package main

import (
	"context"

	pb "example.com/grpc/proto"
)

// Implementing the SayHello method for the gRPC server
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}