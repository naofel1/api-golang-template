package tokenservice

import (
	"crypto/rsa"
	"time"

	"github.com/naofel1/api-golang-template/internal/primitive"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type idTokenCustomClaims struct {
	jwt.RegisteredClaims
	Role primitive.Roles `json:"role"`
	ID   uuid.UUID       `json:"uuid"`
}

func generateIDToken(id uuid.UUID, ro primitive.Roles, key *rsa.PrivateKey, exp time.Duration) (string, error) {
	// Fill token information
	currentTime := time.Now()
	expirationTime := currentTime.Add(exp)
	claims := &idTokenCustomClaims{
		ID:   id,
		Role: ro,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create a new token and claims information filled
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
