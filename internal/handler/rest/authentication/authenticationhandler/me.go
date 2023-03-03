package authenticationhandler

import (
	"net/http"

	"github.com/naofel1/api-golang-template/internal/handler/rest/gincontext"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Me handler return student's details
//
//	@Summary	Me student entity
//	@Tags		Student Authentication
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	ent.StudentFront
//	@Failure	401	{object}	apistatus.ErrorAPI
//	@Failure	500	{object}	apistatus.ErrorAPI
//	@Router		/student/me [get]
func (h *authenticationHandler) handleMeStudent(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	// Start to trace handler
	ctx, span := h.Tracer.Start(ctx, "handleMeStudent Handler")
	defer span.End()

	stud, err := gincontext.GetStudentFromContext(c)
	if err != nil {
		h.Logger.Ctx(ctx).Info("Error getting student from context", zap.Error(err))

		err := apistatus.NewInternal()
		c.JSON(err.Status(), apistatus.NewErrorAPI(err))

		return
	}

	if err := h.StudentService.GetStudent(ctx, stud); err != nil {
		h.Logger.Ctx(ctx).Info("Unable to find student",
			zap.String("UUID", stud.ID.String()),
			zap.Error(err),
		)

		e := apistatus.NewNotFound("student", stud.ID.String())
		c.JSON(e.Status(), apistatus.NewErrorAPI(e))

		return
	}

	c.JSON(http.StatusOK, stud.ToFront())
}

// Me handler return admin's details
//
//	@Summary	Me admin entity
//	@Tags		Admin Authentication
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	ent.AdminFront
//	@Failure	401	{object}	apistatus.ErrorAPI
//	@Failure	500	{object}	apistatus.ErrorAPI
//	@Router		/admin/me [get]
func (h *authenticationHandler) handleMeAdmin(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	// Start to trace handler
	ctx, span := h.Tracer.Start(ctx, "handleMeAdmin Handler")
	defer span.End()

	adm, err := gincontext.GetAdminFromContext(c)
	if err != nil {
		h.Logger.Ctx(ctx).Info("Error getting admin from context", zap.Error(err))

		err := apistatus.NewInternal()
		c.JSON(err.Status(), apistatus.NewErrorAPI(err))

		return
	}

	if err := h.AdminService.GetAdmin(ctx, adm); err != nil {
		h.Logger.Ctx(ctx).Info("Unable to find admin",
			zap.String("UUID", adm.ID.String()),
			zap.Error(err),
		)

		e := apistatus.NewNotFound("admin", adm.ID.String())
		c.JSON(e.Status(), apistatus.NewErrorAPI(e))

		return
	}

	c.JSON(http.StatusOK, adm.ToFront())
}
