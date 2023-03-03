package tokenrepo

import (
	"context"
	"fmt"
	"time"
)

// SetRefreshToken stores a refresh token with an expiry time
func (r *Repository) SetRefreshToken(ctx context.Context, userID string,
	tokenID string, expiresIn time.Duration,
) error {
	// We'll store userID with token id, so we can scan (non-blocking)
	// over the user's tokens and delete them in case of token leakage
	key := fmt.Sprintf("%s:refresh_token:%s", userID, tokenID)
	if err := r.Redis.Set(ctx, key, 0, expiresIn).Err(); err != nil {
		return err
	}

	return nil
}
