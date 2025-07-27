// in backend/internal/token/jwt.go
package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateAccessToken generates a new JWT access token for a given user ID.
func CreateAccessToken(userID string) (string, error) {
	return createToken(userID, 15*time.Minute, os.Getenv("ACCESS_TOKEN_SECRET"))
}

// CreateRefreshToken generates a new JWT refresh token for a given user ID.
func CreateRefreshToken(userID string) (string, error) {
	// Refresh tokens typically have a much longer expiry.
	return createToken(userID, 7*24*time.Hour, os.Getenv("REFRESH_TOKEN_SECRET"))
}

// createToken is a helper function to generate a token with a specific duration and secret.
func createToken(userID string, expiryDuration time.Duration, secretKey string) (string, error) {
	if secretKey == "" {
		return "", fmt.Errorf("token secret not found in environment variables")
	}

	// Create the claims
	claims := jwt.MapClaims{
		"sub": userID,                          // 'sub' (subject) is a standard claim for the user ID
		"iat": time.Now().Unix(),               // 'iat' (issued at) is the time the token was created
		"exp": time.Now().Add(expiryDuration).Unix(), // 'exp' (expiration time)
	}

	// Create the token with the HS256 signing algorithm and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}