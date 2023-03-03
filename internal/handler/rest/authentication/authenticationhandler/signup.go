package authenticationhandler

import (
	"net/http"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/primitive"
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Binding from JSON
type registerStudent struct {
	FirstName string `json:"firstname" binding:"required,gte=1,lte=32"`
	LastName  string `json:"lastname" binding:"required,gte=1,lte=32"`
	Pseudo    string `json:"pseudo" binding:"required"`
	Gender    string `json:"gender" binding:"required,oneof=men women neutral"`
	Password  string `json:"password" binding:"required,gte=6,lte=32"`
}

func (r *registerStudent) toDomain() *ent.Student {
	return &ent.Student{
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		Pseudo:       r.Pseudo,
		Gender:       primitive.Gender(r.Gender),
		PasswordHash: []byte(r.Password),
	}
}

// Signup Student Handler
//
//	@Summary	Signup student entity
//	@Tags		Student Authentication
//	@produce	json
//	@Accept		json
//	@Param		student	body		registerStudent	true	"Student registration"
//	@Success	201		{object}	tokenservice.PairToken
//	@Failure	400		{object}	apistatus.ErrorInvalidArgs
//	@Failure	409		{object}	apistatus.ErrorAPI
//	@Failure	500		{object}	apistatus.ErrorAPI
//	@Router		/student/register [post]
func (h *authenticationHandler) handleSignupStudent(c *gin.Context) {
	ctx := c.Request.Context()

	var newStudent registerStudent

	// Bind incoming newStudent to struct and check for validation errors
	if ok := utils.BindData(c, h.Logger, &newStudent); !ok {
		return
	}

	// Parse student information and return Student in ent Format
	stud := newStudent.toDomain()

	if err := h.StudentService.Signup(ctx, stud); err != nil {
		h.Logger.Ctx(ctx).Info("Failed to sign up user",
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	tokens, err := h.TokenService.NewPairFromStudent(ctx, stud, "")
	if err != nil {
		h.Logger.Ctx(ctx).Info("Failed to create tokens",
			zap.String("User Pseudo", stud.Pseudo),
			zap.Error(err),
		)
		c.JSON(apistatus.Status(err), apistatus.NewErrorAPI(err))

		return
	}

	tokens.SetCookies(c, h.JwtConfig, h.HostConfig)

	// Return the created entry in the table
	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}

// Binding from JSON
type registerAdmin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6,lte=32"`
}

func (r registerAdmin) toDomain() *ent.Admin {
	return &ent.Admin{
		Pseudo:       r.Username,
		PasswordHash: []byte(r.Password),
	}
}

// Signup Admin Handler
//
//	@Summary	Signup admin entity
//	@Tags		Admin Authentication
//	@produce	json
//	@Accept		json
//	@Param		admin	body		registerAdmin	true	"Admin registration"
//	@Success	201		{object}	tokenservice.PairToken
//	@Failure	400		{object}	apistatus.ErrorInvalidArgs
//	@Failure	409		{object}	apistatus.ErrorAPI
//	@Failure	500		{object}	apistatus.ErrorAPI
//	@Router		/admin/register [post]
func (h *authenticationHandler) handleSignupAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	var newAdmin registerAdmin

	// Bind incoming newAdmin to struct and check for validation errors
	if ok := utils.BindData(c, h.Logger, &newAdmin); !ok {
		return
	}

	// Parse admin information and return Admin in ent Format
	u := newAdmin.toDomain()

	if err := h.AdminService.Signup(ctx, u); err != nil {
		h.Logger.Ctx(ctx).Info("Failed to sign up user",
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

	tokens.SetCookies(c, h.JwtConfig, h.HostConfig)

	// Return the created entry in the table
	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
