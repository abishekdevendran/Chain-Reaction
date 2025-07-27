package auth

import (
	"context"
	"log"

	pb "github.com/abishekdevendran/Chain-Reaction/backend/gen/go/user"
)

// Server is the struct that will implement our AuthService.
// It needs to embed the UnimplementedAuthServiceServer type, which is generated
// by protoc. This ensures forward compatibility - if we add new RPCs to our
// .proto file later, our server will still compile.
type Server struct {
	pb.UnimplementedAuthServiceServer
}

// Register implements the Register RPC method.
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	log.Printf("Received Register request for user: %s", req.GetUsername())

	// TODO:
	// 1. Validate the request (username length, password strength).
	// 2. Hash the password.
	// 3. Create the user in the database.
	// 4. Generate JWT and Refresh Token.
	// 5. Return the AuthResponse.

	// For now, we return an empty response.
	return &pb.AuthResponse{}, nil
}

// Login implements the Login RPC method.
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	log.Printf("Received Login request for user: %s", req.GetUsername())

	// TODO:
	// 1. Find the user in the database by username.
	// 2. Compare the provided password with the stored hash.
	// 3. Generate JWT and Refresh Token.
	// 4. Return the AuthResponse.

	return &pb.AuthResponse{}, nil
}

// RefreshAccessToken implements the RefreshAccessToken RPC method.
func (s *Server) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	log.Printf("Received RefreshAccessToken request")

	// TODO:
	// 1. Validate the refresh token.
	// 2. Generate a new access token.
	// 3. Return the new token.

	return &pb.RefreshAccessTokenResponse{}, nil
}