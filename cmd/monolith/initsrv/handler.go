package initsrv

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/handler/rest/admin/profile/adminsettinghandler"
	"github.com/naofel1/api-golang-template/internal/handler/rest/authentication/authenticationhandler"
	"github.com/naofel1/api-golang-template/internal/primitive"
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
)

// HandlerConfig will hold parameter that will be injected into handler
type HandlerConfig struct {
	Logger  *otelzap.Logger
	Conf    *configs.Config
	Router  *gin.Engine
	Service *Service
}

// InitHandler will initialize all handler that will be used in the application
func InitHandler(ctx context.Context, c *HandlerConfig) {
	// Define the authentication handler that will be started as web server
	authenticationhandler.New(&authenticationhandler.Config{
		Tracer:     otel.Tracer("Authentication Handler"),
		R:          c.Router,
		Logger:     c.Logger,
		JwtConfig:  c.Conf.Jwt,
		HostConfig: c.Conf.Host,
		BaseURL:    c.Conf.Host.BaseURL,

		AdminService:   c.Service.Admin,
		TokenService:   c.Service.Token,
		MailerService:  c.Service.Mailer,
		StudentService: c.Service.Student,
	})

	// Define the authentication handler that will be started as web server
	adminsettinghandler.New(&adminsettinghandler.Config{
		Tracer:  otel.Tracer("Admin Setting Handler"),
		R:       c.Router,
		Logger:  c.Logger,
		BaseURL: c.Conf.Host.BaseURL,

		AdminService: c.Service.Admin,
		TokenService: c.Service.Token,
	})

	if c.Conf.AppInfo.Mode == primitive.AppModeDev.String() {
		utils.InitSwagger(ctx, c.Logger, c.Router, c.Conf)
	}
}
