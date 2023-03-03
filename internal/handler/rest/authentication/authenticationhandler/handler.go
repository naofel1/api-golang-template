package authenticationhandler

import (
	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/handler/rest/middlewares"
	"github.com/naofel1/api-golang-template/internal/service/admin/adminservice"
	"github.com/naofel1/api-golang-template/internal/service/mailerservice"
	"github.com/naofel1/api-golang-template/internal/service/student/studentservice"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// authenticationHandler struct holds required services for handler to function
type authenticationHandler struct {
	AdminService   adminservice.Interface
	StudentService studentservice.Interface
	TokenService   tokenservice.Interface
	MailerService  mailerservice.Interface
	Tracer         trace.Tracer
	Logger         *otelzap.Logger
	HostConfig     *configs.Host
	JwtConfig      *configs.Jwt
}

// Config will hold services that will be injected
// into this handler layer on handler initialization
type Config struct {
	R          *gin.Engine
	Logger     *otelzap.Logger
	HostConfig *configs.Host
	JwtConfig  *configs.Jwt

	AdminService   adminservice.Interface
	StudentService studentservice.Interface
	TokenService   tokenservice.Interface
	MailerService  mailerservice.Interface
	Tracer         trace.Tracer
	BaseURL        string
}

// New will initialize all accessible
// route for authentication microservice
func New(c *Config) {
	// Create a handler who have injected services
	h := &authenticationHandler{
		AdminService:   c.AdminService,
		StudentService: c.StudentService,
		TokenService:   c.TokenService,
		MailerService:  c.MailerService,
		JwtConfig:      c.JwtConfig,
		HostConfig:     c.HostConfig,
		Tracer:         c.Tracer,
		Logger:         c.Logger,
	}

	// Create an authentication group
	g := c.R.Group(c.BaseURL)

	// Define the service to collect metric
	g.Use(otelgin.Middleware("Authentication"))

	// Define /api prefix when accessing to the route
	studentAuth := g.Group("/student")
	adminAuth := g.Group("/admin")

	// No token check when gin is on test mode
	if gin.Mode() != gin.TestMode {
		studentAuth.Use(
			middlewares.AuthStudent(h.Logger, otel.Tracer("Middleware"), h.TokenService),
		)
		adminAuth.Use(
			middlewares.AuthAdmin(h.Logger, otel.Tracer("Middleware"), h.TokenService),
		)
	}

	// Register all accessible route
	registerAuthenticationRoute(g, h)

	// Register all accessible route with provided token
	registerStudentRouteAuth(studentAuth, h)

	// Register all accessible route with provided token
	registerAdminRouteAuth(adminAuth, h)
}

func registerAuthenticationRoute(s *gin.RouterGroup, h *authenticationHandler) {
	s.POST("/student/register", h.handleSignupStudent)
	s.POST("/admin/register", h.handleSignupAdmin)

	s.POST("/student/login", h.handleSigninStudent)
	s.POST("/admin/login", h.handleSigninAdmin)

	s.POST("/student/tokens", h.handleTokensStudent)
	s.POST("/admin/tokens", h.handleTokensAdmin)
}

func registerStudentRouteAuth(s *gin.RouterGroup, h *authenticationHandler) {
	s.POST("/logout", h.handleSignoutStudent)
	s.GET("/me", h.handleMeStudent)
}

func registerAdminRouteAuth(s *gin.RouterGroup, h *authenticationHandler) {
	s.POST("/logout", h.handleSignoutAdmin)
	s.GET("/me", h.handleMeAdmin)
}
