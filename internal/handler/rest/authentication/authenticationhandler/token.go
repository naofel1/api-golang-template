package authenticationhandler

import (
	"net/http"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type tokensResponse struct {
	Tokens   *tokenservice.PairToken `json:"tokens"`
	Duration int64                   `json:"duration"`
}

type tokensReq struct {
	RefreshToken string `json:"refreshToken"`
}

// Tokens Student handler
//
//	@Summary	Return new token and refresh token for student
//	@Tags		Student Authentication
//	@produce	json
//	@Accept		json
//	@Param		refreshToken	body	tokensReq	true	"Your refresh token"
//	@Security	ApiKeyAuth
//	@Success	200	{object}	tokensResponse
//	@Failure	401	{object}	apistatus.ErrorAPI
//	@Failure	500	{object}	apistatus.ErrorAPI
//	@Router		/student/tokens [post]
func (h *authenticationHandler) handleTokensStudent(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	refresh, err := checkRefreshToken(c, h.Logger)
	if err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// verify refresh JWT
	refreshToken, err := h.TokenService.ValidateRefreshToken(ctx, refresh)
	if err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	student := &ent.Student{
		ID: refreshToken.UID,
	}

	// get up-to-date student
	if err = h.StudentService.GetStudent(ctx, student); err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// create fresh pair of tokens
	tokens, err := h.TokenService.NewPairFromStudent(ctx, student, refreshToken.ID.String())
	if err != nil {
		h.Logger.Ctx(ctx).Error("Failed to create tokens for student",
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// Set new cookie
	tokens.SetCookies(c, h.JwtConfig, h.HostConfig)

	c.JSON(http.StatusOK, &tokensResponse{
		Tokens:   tokens.ToFront(),
		Duration: h.JwtConfig.TokenDuration.Milliseconds(),
	})
}

// Tokens admin handler
//
//	@Summary	Return new token and refresh token for admin
//	@Tags		Admin Authentication
//	@produce	json
//	@Accept		json
//	@Param		refreshToken	body	tokensReq	true	"Your refresh token"
//	@Security	ApiKeyAuth
//	@Success	200	{object}	tokensResponse
//	@Failure	401	{object}	apistatus.ErrorAPI
//	@Failure	500	{object}	apistatus.ErrorAPI
//	@Router		/admin/tokens [post]
func (h *authenticationHandler) handleTokensAdmin(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	refresh, err := checkRefreshToken(c, h.Logger)
	if err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// verify refresh JWT
	refreshToken, err := h.TokenService.ValidateRefreshToken(ctx, refresh)
	if err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	admin := &ent.Admin{
		ID: refreshToken.UID,
	}

	// get up-to-date admin
	if err = h.AdminService.GetAdmin(ctx, admin); err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// create fresh pair of tokens
	tokens, err := h.TokenService.NewPairFromAdmin(ctx, admin, refreshToken.ID.String())
	if err != nil {
		h.Logger.Ctx(ctx).Info("Failed to create tokens for admin",
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// Set new cookie
	tokens.SetCookies(c, h.JwtConfig, h.HostConfig)

	c.JSON(http.StatusOK, &tokensResponse{
		Tokens:   tokens.ToFront(),
		Duration: h.JwtConfig.TokenDuration.Milliseconds(),
	})
}

func checkRefreshToken(c *gin.Context, logger *otelzap.Logger) (string, error) {
	// Request gin context
	ctx := c.Request.Context()

	var req *tokensReq

	refresh, err := c.Cookie("refreshToken")
	if refresh == "" {
		logger.Ctx(ctx).Info("No refreshToken cookie set",
			zap.Error(err),
		)

		if ok := utils.BindData(c, logger, &req); !ok {
			return "", err
		}

		if req.RefreshToken == "" {
			return "", apistatus.NewBadRequest("Invalid refresh token")
		}

		refresh = req.RefreshToken
	}

	return refresh, nil
}
