package adminsettinghandler

import (
	"net/http"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/handler/rest/gincontext"
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Binding from JSON
type updateAdmin struct {
	Password string `form:"password" json:"password" binding:"omitempty,gte=6,lte=32"`
}

func (t updateAdmin) toDomain(id uuid.UUID) *ent.Admin {
	return &ent.Admin{
		ID:           id,
		PasswordHash: []byte(t.Password),
	}
}

// UpdateAdminProfile handler
//
//	@Summary	Update the Admin Profile
//	@Tags		Admin Setting
//	@Security	ApiKeyAuth
//	@Produce	json
//	@Param		adminInfo	body		updateAdmin	true	"Admin new info"
//	@Success	200			{object}	ent.AdminFront
//	@Failure	400			{object}	apistatus.ErrorInvalidArgs
//	@Failure	401			{object}	apistatus.ErrorAPI
//	@Failure	500			{object}	apistatus.ErrorAPI
//	@Router		/admin/me [patch]
func (h *adminSettingHandler) handleUpdateAdminProfile(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	adm, err := gincontext.GetAdminFromContext(c)
	if err != nil {
		h.Logger.Ctx(ctx).Info("Error getting admin from context", zap.Error(err))

		err := apistatus.NewInternal()
		c.JSON(err.Status(), apistatus.NewErrorAPI(err))

		return
	}

	var tReqUpdate updateAdmin
	if ok := utils.BindData(c, h.Logger, &tReqUpdate); !ok {
		return
	}

	admToUpdate := tReqUpdate.toDomain(adm.ID)

	if err := h.AdminService.ModifyProfile(ctx, admToUpdate); err != nil {
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	c.JSON(http.StatusOK, admToUpdate.ToFront())
}
