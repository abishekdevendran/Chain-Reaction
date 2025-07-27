// In backend/cmd/server/main.go
package main

import (
	"context"
	"log"
	"net"
	"os"

	// Import our new db package and other dependencies
	db "github.com/abishekdevendran/Chain-Reaction/backend/db/sqlc"
	auth_service "github.com/abishekdevendran/Chain-Reaction/backend/internal/auth"
	pb "github.com/abishekdevendran/Chain-Reaction/backend/gen/go/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load .env file from the root directory
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not found in .env file")
	}

	// Create a database connection pool
	connPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer connPool.Close() // Close the pool when main() exits

	// Use our generated db code to create a "store"
	store := db.New(connPool)

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()

	// IMPORTANT: Inject the database store into our auth server
	authServer := &auth_service.Server{
		Store: store,
	}

	pb.RegisterAuthServiceServer(s, authServer)
	log.Printf("gRPC server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}