package tokenservice

import (
	"time"

	"github.com/naofel1/api-golang-template/internal/primitive"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// generateRefreshToken creates a refresh token
// The refresh token stores only the user's ID, a string
func generateRefreshToken(uid uuid.UUID, role primitive.Roles, key string, exp time.Duration) (*refreshTokenData, error) {
	currentTime := time.Now()
	tokenExp := currentTime.Add(exp)

	tokenID, err := uuid.NewRandom() // v4 uuid in the google uuid lib
	if err != nil {
		return nil, err
	}

	claims := refreshTokenCustomClaims{
		UID:  uid,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(tokenExp),
			ID:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	return &refreshTokenData{
		SignedToken: signedToken,
		ID:          tokenID,
		Role:        role,
		ExpiresIn:   tokenExp.Sub(currentTime),
	}, nil
}
