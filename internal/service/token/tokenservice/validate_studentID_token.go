package tokenservice

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// ValidateStudentIDToken validates the id token jwt string
// It returns the user extract from the IDTokenCustomClaims
func (s *tokenService) ValidateStudentIDToken(ctx context.Context, tokenString string) (*ent.Student, error) {
	claims, err := validateStudentIDToken(tokenString, s.StudCert.Pub) // uses public RSA key
	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		s.Logger.Ctx(ctx).Info("Unable to validate or parse idToken",
			zap.Error(err),
		)

		return nil, apistatus.NewAuthorization("Unable to verify user from idToken")
	}

	return &ent.Student{
		ID: claims.ID,
	}, nil
}

// validateIDToken returns the token's claims if the token is valid
func validateStudentIDToken(tokenString string, key *rsa.PublicKey) (*idTokenCustomClaims, error) {
	claims := &idTokenCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// For now, we'll just return the error and handle logging in service level
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("ID token is invalid")
	}

	claims, ok := token.Claims.(*idTokenCustomClaims)
	if !ok {
		return nil, fmt.Errorf("ID token valid but couldn't parse claims")
	}

	return claims, nil
}
