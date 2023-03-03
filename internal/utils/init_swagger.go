package utils

import (
	"context"

	"github.com/gin-gonic/gin"
	docs "github.com/naofel1/api-golang-template/api/rest"
	"github.com/naofel1/api-golang-template/internal/configs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

// InitSwagger init swagger in the router
func InitSwagger(ctx context.Context, logger *otelzap.Logger, router *gin.Engine, conf *configs.Config) {
	docs.SwaggerInfo.Host = conf.Host.Address
	docs.SwaggerInfo.BasePath = conf.Host.BaseURL

	// Set url for swagger documentation
	router.GET(conf.Host.BaseURL+"/swagger/*any", otelgin.Middleware("Documentation"), ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Ctx(ctx).Info("Swagger documentation is available",
		zap.String("URL", conf.Host.BaseURL+"/swagger/index.html"))
}
