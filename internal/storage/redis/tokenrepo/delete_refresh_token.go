package tokenrepo

import (
	"context"
	"fmt"

	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
)

// DeleteRefreshToken used to delete old refresh tokens
// Services my access this to revolve tokens
func (r *Repository) DeleteRefreshToken(ctx context.Context, userID, tokenID string) error {
	key := fmt.Sprintf("%s:refresh_token:%s", userID, tokenID)

	result := r.Redis.Del(ctx, key)

	if err := result.Err(); err != nil {
		return apistatus.NewInternal()
	}

	// Val returns count of deleted keys.
	// If no key was deleted, the refresh token is invalid
	if result.Val() < 1 {
		return tokenservice.NewInvalidToken(tokenID)
	}

	return nil
}
