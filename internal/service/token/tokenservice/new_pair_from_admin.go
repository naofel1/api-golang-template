package tokenservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/primitive"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// NewPairFromAdmin creates fresh id and refresh tokens for the current admin
// If a previous token is included, the previous token is removed from the tokens repository
func (s *tokenService) NewPairFromAdmin(ctx context.Context, u *ent.Admin, prevTokenID string) (*PairToken, error) {
	// Start to trace middleware
	ctx, span := s.Tracer.Start(ctx, "NewPairFromAdmin Service")
	defer span.End()

	if prevTokenID != "" {
		if err := s.TokenRepository.DeleteRefreshToken(ctx, u.ID.String(), prevTokenID); err != nil {
			s.Logger.Ctx(ctx).Info("Could not delete previous refreshToken",
				zap.String("uid", u.ID.String()),
				zap.String("tokenID", prevTokenID),
				zap.Error(err),
			)

			if IsInvalidToken(err) {
				return nil, apistatus.NewAuthorization("Invalid refresh token")
			}

			return nil, apistatus.NewInternal()
		}
	}

	// No need to use a repository for idToken as it is unrelated to any data source
	idToken, err := generateIDToken(u.ID, primitive.RoleAdmin, s.AdminCert.Priv, s.IDExpiration)
	if err != nil {
		s.Logger.Ctx(ctx).Info("Could not delete previous refreshToken",
			zap.String("uid", u.ID.String()),
			zap.Error(err),
		)

		return nil, apistatus.NewInternal()
	}

	refreshToken, err := generateRefreshToken(u.ID, primitive.RoleAdmin, s.RefreshSecret, s.RefreshExpiration)
	if err != nil {
		s.Logger.Ctx(ctx).Info("Error generating refreshToken",
			zap.String("uid", u.ID.String()),
			zap.Error(err),
		)

		return nil, apistatus.NewInternal()
	}

	// set freshly minted refresh token to valid list
	if err := s.TokenRepository.SetRefreshToken(ctx, u.ID.String(), refreshToken.ID.String(), refreshToken.ExpiresIn); err != nil {
		s.Logger.Ctx(ctx).Info("Error storing tokenID",
			zap.String("uid", u.ID.String()),
			zap.Error(err),
		)

		return nil, apistatus.NewInternal()
	}

	return &PairToken{
		IDToken: &IDToken{
			SignedToken: idToken,
		},
		RefreshToken: &RefreshToken{
			SignedToken: refreshToken.SignedToken,
			ID:          refreshToken.ID,
			UID:         u.ID,
		},
	}, nil
}
