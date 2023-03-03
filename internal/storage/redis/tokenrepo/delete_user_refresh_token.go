package tokenrepo

import (
	"context"
	"fmt"
)

// DeleteUserRefreshTokens looks for all tokens beginning with
// userID and scans to delete them in a non-blocking fashion
func (r *Repository) DeleteUserRefreshTokens(ctx context.Context, userID string) error {
	const nbRefreshToken = 5

	pattern := fmt.Sprintf("%s:refresh_token:*", userID)

	iter := r.Redis.Scan(ctx, 0, pattern, nbRefreshToken).Iterator()

	for iter.Next(ctx) {
		if err := r.Redis.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}

	// check last value
	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
