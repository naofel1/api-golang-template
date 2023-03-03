package tokenservice

import (
	"fmt"

	"github.com/naofel1/api-golang-template/internal/primitive"

	"github.com/golang-jwt/jwt/v4"
)

func (s *tokenService) GetRoleFromIDToken(tokenString string) (primitive.Roles, error) {
	var ok bool

	claims := &idTokenCustomClaims{}

	token, _ := jwt.ParseWithClaims(tokenString, claims, nil)

	claims, ok = token.Claims.(*idTokenCustomClaims)
	if !ok {
		return "", fmt.Errorf("ID token valid but couldn't parse claims")
	}

	return claims.Role, nil
}
