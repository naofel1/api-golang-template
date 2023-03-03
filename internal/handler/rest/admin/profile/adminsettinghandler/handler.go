package adminsettinghandler

import (
	"github.com/naofel1/api-golang-template/internal/handler/rest/middlewares"
	"github.com/naofel1/api-golang-template/internal/service/admin/adminservice"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// adminSettingHandler struct holds required services for handler to function
type adminSettingHandler struct {
	AdminService adminservice.Interface
	TokenService tokenservice.Interface
	Tracer       trace.Tracer
	Logger       *otelzap.Logger
}

// Config will hold services that will be injected
// into this handler layer on handler initialization
type Config struct {
	R            *gin.Engine
	AdminService adminservice.Interface
	TokenService tokenservice.Interface
	Tracer       trace.Tracer
	Logger       *otelzap.Logger
	BaseURL      string
}

// New will initialize all accessible
// route for admin microservice
func New(c *Config) {
	// Create a handler who have injected services
	h := &adminSettingHandler{
		AdminService: c.AdminService,
		TokenService: c.TokenService,
		Tracer:       c.Tracer,
		Logger:       c.Logger,
	}

	// Create an admin group
	g := c.R.Group(c.BaseURL)

	// Define /api prefix when accessing to the route
	adminAuth := g.Group(c.BaseURL + "/admin")
	if gin.Mode() != gin.TestMode {
		adminAuth.Use(middlewares.AuthAdmin(h.Logger, otel.Tracer("Middleware"), h.TokenService))
	}

	// Register all accessible route
	registerAdminRoute(adminAuth, h)
}

func registerAdminRoute(s *gin.RouterGroup, h *adminSettingHandler) {
	s.PATCH("/me", h.handleUpdateAdminProfile)
}
