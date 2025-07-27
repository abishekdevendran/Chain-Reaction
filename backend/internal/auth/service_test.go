package auth

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	db "github.com/abishekdevendran/Chain-Reaction/backend/db/sqlc"
	pb "github.com/abishekdevendran/Chain-Reaction/backend/gen/go/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// testAuthClient is a helper function to set up and tear down a test server and client
func testAuthClient(t *testing.T) pb.AuthServiceClient {
	// 1. Load environment variables
	err := godotenv.Load("../../../.env") // Note the path from the test file
	require.NoError(t, err)

	dbURL := os.Getenv("DATABASE_URL")
	require.NotEmpty(t, dbURL)

	// 2. Set up database connection
	connPool, err := pgxpool.New(context.Background(), dbURL)
	require.NoError(t, err)
	store := db.New(connPool)

	// 3. Set up a gRPC server on a random available port
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)

	server := &Server{Store: store}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Test server failed: %v", err)
		}
	}()

	// 4. Create a client connection to the test server
	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	// 5. Add a cleanup function to close connections when the test is done
	t.Cleanup(func() {
		grpcServer.Stop()
		conn.Close()
		connPool.Close()
	})

	return pb.NewAuthServiceClient(conn)
}

// TestRegister is our actual test case
func TestRegister(t *testing.T) {
	// Get a client connected to a test instance of our server
	client := testAuthClient(t)

	// Define the test case
	username := "testuser-from-gocode"
	password := "a-very-good-password"

	req := &pb.RegisterRequest{
		Username: username,
		Password: password,
	}

	// Make the RPC call
	res, err := client.Register(context.Background(), req)

	// Assert the results
	require.NoError(t, err) // We expect no error
	require.NotNil(t, res)  // The response should not be nil

	require.Equal(t, username, res.GetUser().GetUsername())
	require.NotEmpty(t, res.GetUser().GetId())
	require.NotEmpty(t, res.GetAccessToken())
	require.NotEmpty(t, res.GetRefreshToken())
}