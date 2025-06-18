package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server starting on port 50051...")
	s.Serve(lis)

}
