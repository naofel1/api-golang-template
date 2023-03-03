package tokenservice

import (
	"context"
	"fmt"
	"time"

	"github.com/naofel1/api-golang-template/internal/primitive"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// refreshTokenData holds the actual signed jwt string along with the ID
// We return the id, so it can be used without parsing again the JWT from signed string
type refreshTokenData struct {
	SignedToken string
	Role        primitive.Roles
	ExpiresIn   time.Duration
	ID          uuid.UUID
}

// refreshTokenCustomClaims holds the payload of a refresh token
// This can be used to extract user id for subsequent
// application operations (IE, fetch user in Redis)
type refreshTokenCustomClaims struct {
	jwt.RegisteredClaims
	Role primitive.Roles `json:"role"`
	UID  uuid.UUID       `json:"uid"`
}

// ValidateRefreshToken checks to make sure the JWT provided by a string is valid
// and returns a RefreshToken if valid
func (s *tokenService) ValidateRefreshToken(ctx context.Context, tokenString string) (*RefreshToken, error) {
	// validate actual JWT with string a secret
	claims, err := validateRefreshToken(tokenString, s.RefreshSecret)
	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		s.Logger.Ctx(ctx).Info("Unable to validate or parse refreshToken",
			zap.String("tokenString", tokenString),
			zap.Error(err),
		)

		return nil, apistatus.NewAuthorization("Unable to verify user from refresh token")
	}

	// Standard claims store ID as a string. I want "models" to be clear our string
	// is a UUID. So we parse claims.ID as UUID
	tokenUUID, err := uuid.Parse(claims.ID)
	if err != nil {
		s.Logger.Ctx(ctx).Info("Claims ID could not be parsed",
			zap.String("UUID", claims.ID),
			zap.Error(err),
		)

		return nil, apistatus.NewAuthorization("Unable to verify user from refresh token")
	}

	return &RefreshToken{
		SignedToken: tokenString,
		ID:          tokenUUID,
		UID:         claims.UID,
	}, nil
}

// validateRefreshToken uses the secret key to validate a refresh token
func validateRefreshToken(tokenString, key string) (*refreshTokenCustomClaims, error) {
	claims := &refreshTokenCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	// For now, we'll just return the error and handle logging in service level
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("refresh token is invalid")
	}

	claims, ok := token.Claims.(*refreshTokenCustomClaims)
	if !ok {
		return nil, fmt.Errorf("refresh token valid but couldn't parse claims")
	}

	return claims, nil
}
