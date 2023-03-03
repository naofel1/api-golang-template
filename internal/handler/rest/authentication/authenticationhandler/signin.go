package authenticationhandler

import (
	"net/http"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Binding from JSON
type studentLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r studentLogin) toDomain() *ent.Student {
	return &ent.Student{
		Pseudo:       r.Login,
		PasswordHash: []byte(r.Password),
	}
}

// Signin for Student Handler
//
//	@Summary		Sign in a student
//	@Description	This handler purpose is to log in a student by the specified login
//	@Description	and password, it then returns a specific Token Pair associated with the student.
//	@Tags			Student Authentication
//	@produce		json
//	@Accept			json
//	@Param			student	body		studentLogin	true	"Student login"
//	@Success		200		{object}	tokenservice.PairToken
//	@Failure		400		{object}	apistatus.ErrorInvalidArgs
//	@Failure		401		{object}	apistatus.ErrorAPI
//	@Failure		500		{object}	apistatus.ErrorAPI
//	@Router			/student/login [post]
func (h *authenticationHandler) handleSigninStudent(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	// Start to trace handler
	ctx, span := h.Tracer.Start(ctx, "Student Signin Handler")
	defer span.End()

	var sLogin studentLogin
	// Bind incoming json to struct and check for validation errors
	if ok := utils.BindData(c, h.Logger, &sLogin); !ok {
		return
	}

	stud := sLogin.toDomain()

	// Signin student
	if err := h.StudentService.Signin(ctx, stud); err != nil {
		h.Logger.Ctx(ctx).Info("Failed to sign in user",
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	tokens, err := h.TokenService.NewPairFromStudent(ctx, stud, "")
	if err != nil {
		h.Logger.Ctx(ctx).Info("Failed to create tokens",
			zap.String("Student pseudo", stud.Pseudo),
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	// Set new cookie
	tokens.SetCookies(c, h.JwtConfig, h.HostConfig)

	// Return the created entry in the table
	c.JSON(http.StatusOK, tokens.ToFront())
}

// Binding from JSON
type adminLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r adminLogin) toDomain() *ent.Admin {
	return &ent.Admin{
		Pseudo:       r.Username,
		PasswordHash: []byte(r.Password),
	}
}

// Signin for Admin Handler
//
//	@Summary		Sign in an admin
//	@Description	This handler purpose is to log in an admin by the specified email
//	@Description	and password, it then returns a specific Token Pair associated with the admin.
//	@Tags			Admin Authentication
//	@produce		json
//	@Accept			json
//	@Param			admin	body		adminLogin	true	"Admin login"
//	@Success		200		{object}	tokenservice.PairToken
//	@Failure		400		{object}	apistatus.ErrorInvalidArgs
//	@Failure		401		{object}	apistatus.ErrorAPI
//	@Failure		500		{object}	apistatus.ErrorAPI
//	@Router			/admin/login [post]
func (h *authenticationHandler) handleSigninAdmin(c *gin.Context) {
	// Request gin context
	ctx := c.Request.Context()

	var aLogin adminLogin
	// Bind incoming json to struct and check for validation errors
	if ok := utils.BindData(c, h.Logger, &aLogin); !ok {
		return
	}

	u := aLogin.toDomain()

	if err := h.AdminService.Signin(ctx, u); err != nil {
		h.Logger.Ctx(ctx).Info("Failed to sign in user",
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	tokens, err := h.TokenService.NewPairFromAdmin(ctx, u, "")
	if err != nil {
		h.Logger.Ctx(ctx).Info("Failed to create tokens",
			zap.String("Admin pseudo", u.Pseudo),
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}
	// Set new cookie
	tokens.SetCookies(c, h.JwtConfig, h.HostConfig)

	// Return the created entry in the table
	c.JSON(http.StatusOK, tokens.ToFront())
}
