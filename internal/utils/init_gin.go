package utils

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/handler/rest"
	"github.com/naofel1/api-golang-template/internal/primitive"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// InitGin init gin router with middleware and cors configuration from config
func InitGin(ctx context.Context, logger *otelzap.Logger, conf *configs.Config) *gin.Engine {
	// Set gin in production mode
	if conf.AppInfo.Mode == primitive.AppModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}
	// Define router
	router := gin.New()

	router.Use(rest.GlobalServerMiddleware(logger)...)

	// Define CORS condition
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,

		AllowOrigins: conf.Cors.AllowedOrigins,
		AllowMethods: conf.Cors.AllowedMethods,
		AllowHeaders: conf.Cors.AllowedHeaders,
	}))

	logger.Ctx(ctx).Info("Gin router initialized")

	return router
}
