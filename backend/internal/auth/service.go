// In backend/internal/auth/service.go

package auth

import (
	"context"
	"log"

	db "github.com/abishekdevendran/Chain-Reaction/backend/db/sqlc"
	pb "github.com/abishekdevendran/Chain-Reaction/backend/gen/go/user"
	"github.com/abishekdevendran/Chain-Reaction/backend/internal/token"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server now holds a reference to our database Querier interface.
type Server struct {
	pb.UnimplementedAuthServiceServer
	Store db.Querier
}

// Register implements the Register RPC method.
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	log.Printf("Received Register request for user: %s", req.GetUsername())

	// 1. Validate the request
	if len(req.GetUsername()) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "Username must be at least 3 characters long")
	}
	if len(req.GetPassword()) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "Password must be at least 8 characters long")
	}

	// 2. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %v", err)
	}

	// 3. Create the user in the database
	params := db.CreateUserParams{
		Username:     req.GetUsername(),
		PasswordHash: string(hashedPassword),
	}

	createdUser, err := s.Store.CreateUser(ctx, params)
	if err != nil {
		// Specifically check for a unique constraint violation on the username.
		// The exact error message can depend on the database driver.
		// if db.IsUniqueViolationError(err) { // You would need to define this helper
		// 	return nil, status.Errorf(codes.AlreadyExists, "Username '%s' is already taken", req.GetUsername())
		// }
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	// 4. Generate JWT and Refresh Token
	accessToken, err := token.CreateAccessToken(createdUser.ID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %v", err)
	}
	refreshToken, err := token.CreateRefreshToken(createdUser.ID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token: %v", err)
	}

	// 5. Return the full AuthResponse
	return &pb.AuthResponse{
		User: &pb.User{
			Id:        createdUser.ID.String(),
			Username:  createdUser.Username,
			CreatedAt: createdUser.CreatedAt.Time.Unix(),
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Login implements the Login RPC method.
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	log.Printf("Received Login request for user: %s", req.GetUsername())

	// 1. Find the user in the database by username
	user, err := s.Store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		// Map the specific DB "not found" error to a gRPC "NotFound" status
		if err.Error() == "no rows in result set" {
			return nil, status.Errorf(codes.NotFound, "invalid credentials")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	// 2. Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.GetPassword()))
	if err != nil {
		// If they don't match, bcrypt returns an error. Return a generic message.
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// 3. Generate JWT and Refresh Token
	accessToken, err := token.CreateAccessToken(user.ID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %v", err)
	}
	refreshToken, err := token.CreateRefreshToken(user.ID.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token: %v", err)
	}

	// 4. Return the full AuthResponse
	return &pb.AuthResponse{
		User: &pb.User{
			Id:        user.ID.String(),
			Username:  user.Username,
			CreatedAt: user.CreatedAt.Time.Unix(),
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}