package main

import (
	"log"
	"net"

	// Import our generated protobuf code and our service implementation
	auth_service "github.com/abishekdevendran/Chain-Reaction/backend/internal/auth"
	pb "github.com/abishekdevendran/Chain-Reaction/backend/gen/go/user"

	"google.golang.org/grpc"
)

func main() {
	// Define the port for our gRPC server to listen on.
	// 50051 is the conventional port for gRPC.
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Create a new gRPC server instance.
	s := grpc.NewServer()

	// Create an instance of our auth service implementation.
	authServer := &auth_service.Server{}

	// Register our service implementation with the gRPC server.
	// This is the crucial step that links the server to our logic.
	pb.RegisterAuthServiceServer(s, authServer)

	log.Printf("gRPC server listening on %s", port)

	// Start serving requests. This is a blocking call that will run until
	// the process is killed.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}