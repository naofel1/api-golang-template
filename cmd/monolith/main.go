// Monolith contain all microservice that will be launched as webserver
// All initialization of the required service, repository and handler are here.
package main

import (
	"context"

	"github.com/naofel1/api-golang-template/cmd/monolith/initsrv"
	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/naofel1/api-golang-template/internal/handler/rest"

	"github.com/naofel1/api-golang-template/internal/tracing"
	"github.com/naofel1/api-golang-template/internal/utils"

	"go.uber.org/zap"
)

// @title						API
// @version					1.0
// @description				This document's purpose is to document the API
// @description				used by this backend for further interaction with the front-end
//
// @termsOfService				https://domain.TLD
// @contact.name				Your name
// @contact.url				https://domain.TLD
// @contact.email				info@domain.TLD
// @schemes					https http
//
// @securitydefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				JWT Token
func main() {
	// Take the context from background
	ctx := context.Background()

	// Initialize Config
	conf := configs.Init()

	// Initialize Zap Log (With wrapper OpenTelemetry)
	logger := utils.InitLogger(ctx, conf.AppInfo)

	// Load Configurations from config.json
	conf.GetConfig(ctx, logger)

	// Set OpenTelemetry Information
	tracing.SetOpenTelemetryInfo(ctx, logger, tracing.JaegerTraceProvider(ctx, logger, conf))

	// Load all used utility
	utls := initsrv.InitUtils(ctx, logger, conf)

	// Load all used client
	cl := initsrv.InitClient(ctx, logger, conf)

	// Close all client connection when the server is done
	defer cl.CloseClient(ctx, logger)

	router := utils.InitGin(ctx, logger, conf)

	// Create new repository used in this microservice
	repo := initsrv.InitRepository(cl)

	// Create new service used in this microservice
	srvice := initsrv.InitService(&initsrv.ServiceConfig{
		Logger: logger,
		Conf:   conf,
		Repo:   repo,
		Utils:  utls,
	})

	// Initialize all handler used in this microservice
	initsrv.InitHandler(ctx, &initsrv.HandlerConfig{
		Conf:    conf,
		Logger:  logger,
		Router:  router,
		Service: srvice,
	})

	srv := rest.NewServer(logger, router, conf)

	if err := srv.Run(ctx); err != nil {
		logger.Error("Cannot start server",
			zap.Error(err),
		)
	}
}
