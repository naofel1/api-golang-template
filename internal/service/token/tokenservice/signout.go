package tokenservice

import (
	"context"

	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Signout reaches out to the repository layer to delete all valid tokens for a user
func (s *tokenService) Signout(ctx context.Context, uid uuid.UUID) error {
	if err := s.TokenRepository.DeleteUserRefreshTokens(ctx, uid.String()); err != nil {
		s.Logger.Ctx(ctx).Info("Failed to delete token",
			zap.Error(err),
		)

		return apistatus.NewInternal()
	}

	return nil
}
