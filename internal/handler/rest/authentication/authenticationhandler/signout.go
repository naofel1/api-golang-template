package authenticationhandler

import (
	"net/http"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/handler/rest/gincontext"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Signout student handler
//
//	@Summary	Signout handler
//	@Tags		Student Authentication
//	@produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	apistatus.SuccessStatus
//	@Failure	401	{object}	apistatus.ErrorAPI
//	@Failure	500	{object}	apistatus.ErrorAPI
//	@Router		/student/logout [post]
func (h *authenticationHandler) handleSignoutStudent(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	var stud *ent.Student

	stud, err := gincontext.GetStudentFromContext(c)
	if err != nil {
		h.Logger.Ctx(ctx).Info("Error getting student from context", zap.Error(err))

		err := apistatus.NewInternal()
		c.JSON(err.Status(), apistatus.NewErrorAPI(err))

		return
	}

	if err := h.TokenService.Signout(ctx, stud.ID); err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}
	// Remove any previous set cookies
	tokenservice.ClearCookies(c, h.HostConfig)

	c.JSON(http.StatusOK, apistatus.NewSuccessStatus(
		"student signed out successfully!",
	))
}

// Signout admin handler
//
//	@Summary	Signout handler
//	@Tags		Admin Authentication
//	@produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	apistatus.SuccessStatus
//	@Failure	401	{object}	apistatus.ErrorAPI
//	@Failure	500	{object}	apistatus.ErrorAPI
//	@Router		/admin/logout [post]
func (h *authenticationHandler) handleSignoutAdmin(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	adm, err := gincontext.GetAdminFromContext(c)
	if err != nil {
		h.Logger.Ctx(ctx).Info("Error getting admin from context", zap.Error(err))

		err := apistatus.NewInternal()
		c.JSON(err.Status(), apistatus.NewErrorAPI(err))

		return
	}

	if err := h.TokenService.Signout(ctx, adm.ID); err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// Remove any previous set cookies
	tokenservice.ClearCookies(c, h.HostConfig)

	c.JSON(http.StatusOK, apistatus.NewSuccessStatus(
		"admin signed out successfully!",
	))
}
