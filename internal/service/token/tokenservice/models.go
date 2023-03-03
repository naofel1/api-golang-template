package tokenservice

import (
	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RefreshToken model info
//
//	@Description	RefreshToken stores token properties that
//	@Description	are accessed in multiple application layers
type RefreshToken struct {
	SignedToken string    `json:"refreshToken"`
	ID          uuid.UUID `json:"-"`
	UID         uuid.UUID `json:"-"`
}

// IDToken model info
//
//	@Description	IDToken stores token properties that
//	@Description	are accessed in multiple application layers
type IDToken struct {
	SignedToken string `json:"idToken"`
}

// PairToken model info
//
//	@Description	PairToken is used for returning pairs
//	@Description	of id and refresh tokens
type PairToken struct {
	*IDToken
	*RefreshToken
}

// ToFront is the structure returned to the frontend
func (r *PairToken) ToFront() *PairToken {
	return &PairToken{
		IDToken:      r.IDToken,
		RefreshToken: r.RefreshToken,
	}
}

// SetCookies will set the cookie for the id and refresh tokens
func (r *PairToken) SetCookies(c *gin.Context, duration *configs.Jwt, host *configs.Host) {
	c.SetCookie("idToken", r.IDToken.SignedToken, int(duration.TokenDuration.Seconds()), "/", host.Address, false, true)
	c.SetCookie("refreshToken", r.RefreshToken.SignedToken, int(duration.RefreshDuration.Seconds()), "/", host.Address, false, true)
}

// ClearCookies will clear the cookie for the id and refresh tokens
func ClearCookies(c *gin.Context, host *configs.Host) {
	c.SetCookie("idToken", "", 0, "/", host.Address, false, true)
	c.SetCookie("refreshToken", "", 0, "/", host.Address, false, true)
}
