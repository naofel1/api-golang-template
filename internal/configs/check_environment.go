package configs

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// newHostFromEnv return a config.Host struct filled with env variable
func newHostFromEnv(ctx context.Context, logger *otelzap.Logger) *Host {
	hMode, exists := os.LookupEnv("HOST_MODE")
	if !exists {
		logger.Ctx(ctx).Fatal("No HOST_MODE env variable")
	}

	hAddress, exists := os.LookupEnv("HOST_ADDRESS")
	if !exists {
		logger.Ctx(ctx).Fatal("No HOST_ADDRESS env variable")
	}

	hPort, exists := os.LookupEnv("HOST_PORT")
	if !exists {
		logger.Ctx(ctx).Fatal("No HOST_PORT env variable")
	}

	hPortInt, err := strconv.Atoi(hPort)
	if err != nil {
		logger.Ctx(ctx).Fatal("No HOST_PORT env variable must be a number",
			zap.Error(err),
		)
	}

	hBaseURL, exists := os.LookupEnv("HOST_BASE_URL")
	if !exists {
		logger.Ctx(ctx).Fatal("No HOST_BASE_URL env variable")
	}

	return &Host{
		Mode:    hMode,
		Address: hAddress,
		Port:    hPortInt,
		BaseURL: hBaseURL,
	}
}

// newJWTFromEnv return a config.Jwt struct filled with env variable
func newJWTFromEnv(ctx context.Context, logger *otelzap.Logger) *Jwt {
	jTokenDuration, exists := os.LookupEnv("JWT_TOKEN_DURATION")
	if !exists {
		logger.Ctx(ctx).Fatal("No JWT_TOKEN_DURATION env variable")
	}

	jTokenDurationInt, err := strconv.Atoi(jTokenDuration)
	if err != nil {
		logger.Ctx(ctx).Fatal("No JWT_TOKEN_DURATION env variable must be a number")
	}

	jRefreshDuration, exists := os.LookupEnv("JWT_REFRESH_DURATION")
	if !exists {
		logger.Ctx(ctx).Fatal("No JWT_REFRESH_DURATION env variable")
	}

	jRefreshDurationInt, err := strconv.Atoi(jRefreshDuration)
	if err != nil {
		logger.Ctx(ctx).Fatal("No JWT_REFRESH_DURATION env variable must be a number")
	}

	return &Jwt{
		TokenDuration:   time.Duration(jTokenDurationInt),
		RefreshDuration: time.Duration(jRefreshDurationInt),
	}
}

// newServerFromEnv return a config.Server struct filled with env variable
func newServerFromEnv(ctx context.Context, logger *otelzap.Logger) *Server {
	cTimeout, exists := os.LookupEnv("SERVER_CLIENT_TIMEOUT")
	if !exists {
		logger.Ctx(ctx).Fatal("No SERVER_CLIENT_TIMEOUT env variable")
	}

	cTime, err := time.ParseDuration(cTimeout)
	if err != nil {
		logger.Ctx(ctx).Fatal("Incorrect format SERVER_CLIENT_TIMEOUT",
			zap.Error(err),
		)
	}

	rTimeout, exists := os.LookupEnv("SERVER_READ_TIMEOUT")
	if !exists {
		logger.Ctx(ctx).Fatal("No SERVER_READ_TIMEOUT env variable")
	}

	rTime, err := time.ParseDuration(rTimeout)
	if err != nil {
		logger.Ctx(ctx).Fatal("Incorrect format SERVER_READ_TIMEOUT",
			zap.Error(err),
		)
	}

	wTimeout, exists := os.LookupEnv("SERVER_WRITE_TIMEOUT")
	if !exists {
		logger.Ctx(ctx).Fatal("No SERVER_WRITE_TIMEOUT env variable")
	}

	wTime, err := time.ParseDuration(wTimeout)
	if err != nil {
		logger.Ctx(ctx).Fatal("Incorrect format SERVER_WRITE_TIMEOUT",
			zap.Error(err),
		)
	}

	cGracePeriod, exists := os.LookupEnv("SERVER_SHUTDOWN_GRACE_PERIOD")
	if !exists {
		logger.Ctx(ctx).Fatal("No SERVER_SHUTDOWN_GRACE_PERIOD env variable")
	}

	cGrace, err := time.ParseDuration(cGracePeriod)
	if err != nil {
		logger.Ctx(ctx).Fatal("Incorrect format SERVER_SHUTDOWN_GRACE_PERIOD",
			zap.Error(err),
		)
	}

	return &Server{
		ClientTimeout:       cTime,
		ReadTimeout:         rTime,
		WriteTimeout:        wTime,
		ShutdownGracePeriod: cGrace,
	}
}

// newCertsFromEnv return a config.Certs struct filled with env variable
func newCertsFromEnv(ctx context.Context, logger *otelzap.Logger) *Certs {
	cPubStudent, exists := os.LookupEnv("CERT_PUBLIC_STUDENT")
	if !exists {
		logger.Ctx(ctx).Fatal("No CERT_PUBLIC_STUDENT env variable")
	}

	cPrivStudent, exists := os.LookupEnv("CERT_PRIVATE_STUDENT")
	if !exists {
		logger.Ctx(ctx).Fatal("No CERT_PRIVATE_STUDENT env variable")
	}

	cPubAdmin, exists := os.LookupEnv("CERT_PUBLIC_ADMIN")
	if !exists {
		logger.Ctx(ctx).Fatal("No CERT_PUBLIC_ADMIN env variable")
	}

	cPrivAdmin, exists := os.LookupEnv("CERT_PRIVATE_ADMIN")
	if !exists {
		logger.Ctx(ctx).Fatal("No CERT_PRIVATE_ADMIN env variable")
	}

	return &Certs{
		PubStudent:  cPubStudent,
		PrivStudent: cPrivStudent,
		PubAdmin:    cPubAdmin,
		PrivAdmin:   cPrivAdmin,
	}
}

// newMariaDBFromEnv return a config.Mariadb struct filled with env variable
func newMariaDBFromEnv(ctx context.Context, logger *otelzap.Logger) *Mariadb {
	dbName, exists := os.LookupEnv("MARIADB_DATABASE_NAME")
	if !exists {
		logger.Ctx(ctx).Fatal("No MARIADB_DATABASE_NAME env variable")
	}

	dbUser, exists := os.LookupEnv("MARIADB_USER")
	if !exists {
		logger.Ctx(ctx).Fatal("No MARIADB_USER env variable")
	}

	dbHost, exists := os.LookupEnv("MARIADB_HOST")
	if !exists {
		logger.Ctx(ctx).Fatal("No MARIADB_HOST env variable")
	}

	dbPort, exists := os.LookupEnv("MARIADB_PORT")
	if !exists {
		logger.Ctx(ctx).Fatal("No MARIADB_PORT env variable")
	}

	return &Mariadb{
		DBName: dbName,
		User:   dbUser,
		Host:   dbHost,
		Port:   dbPort,
	}
}

// newRedisFromEnv return a config.redis struct filled with env variable
func newRedisFromEnv(ctx context.Context, logger *otelzap.Logger) *Redis {
	rConnectionString, exists := os.LookupEnv("REDIS_CONNECTION_STRING")
	if !exists {
		logger.Ctx(ctx).Fatal("No REDIS_CONNECTION_STRING env variable")
	}

	rSelectedDB, exists := os.LookupEnv("REDIS_DB")
	if !exists {
		logger.Ctx(ctx).Fatal("No REDIS_DB env variable")
	}

	rSelectedDBInt, err := strconv.Atoi(rSelectedDB)
	if err != nil {
		logger.Ctx(ctx).Fatal("No REDIS_DB env variable must be a number")
	}

	return &Redis{
		ConnectionString: rConnectionString,
		SelectedDB:       rSelectedDBInt,
	}
}

// newCentrifugoFromEnv return a config.Centrifugo struct filled with env variable
func newCentrifugoFromEnv(ctx context.Context, logger *otelzap.Logger) *Centrifugo {
	pConnectionString, exists := os.LookupEnv("CENTRIFUGO_CONNECTION_STRING")
	if !exists {
		logger.Ctx(ctx).Fatal("No CENTRIFUGO_CONNECTION_STRING env variable")
	}

	return &Centrifugo{
		ConnectionString: pConnectionString,
	}
}

// newMailGunFromEnv return a config.Mailgun struct filled with env variable
func newMailGunFromEnv(ctx context.Context, logger *otelzap.Logger) *Mailgun {
	mgDomain, exists := os.LookupEnv("MG_DOMAIN")
	if !exists {
		logger.Ctx(ctx).Fatal("No MG_DOMAIN env variable")
	}

	mgRPSender, exists := os.LookupEnv("MG_RESET_PASSWORD_SENDER")
	if !exists {
		logger.Ctx(ctx).Fatal("No MG_RESET_PASSWORD_SENDER env variable")
	}

	mgRPSubject, exists := os.LookupEnv("MG_RESET_PASSWORD_SUBJECT")
	if !exists {
		logger.Ctx(ctx).Fatal("No MG_RESET_PASSWORD_SUBJECT env variable")
	}

	return &Mailgun{
		ClientDomain: mgDomain,
		MailConfig: &MailgunMailer{
			ResetPassword: &MailgunResetPassword{
				Sender:  mgRPSender,
				Subject: mgRPSubject,
			},
		},
	}
}

// newAppInfoFromEnv return a config.AppInfo struct filled with env variable
func newAppInfoFromEnv(ctx context.Context, logger *otelzap.Logger) *AppInfo {
	appMode, exists := os.LookupEnv("APP_MODE")
	if !exists {
		logger.Ctx(ctx).Fatal("No APP_MODE env variable")
	}

	return &AppInfo{
		Mode: appMode,
	}
}

// newJaegerFromEnv return a config.Jaeger struct filled with env variable
func newJaegerFromEnv(ctx context.Context, logger *otelzap.Logger) *Jaeger {
	pConnectionString, exists := os.LookupEnv("JAEGER_CONNECTION_STRING")
	if !exists {
		logger.Ctx(ctx).Fatal("No JAEGER_CONNECTION_STRING env variable")
	}

	return &Jaeger{
		ConnectionString: pConnectionString,
	}
}

// newAWSFromEnv return a config.AWS struct filled with env variable
func newAWSFromEnv(ctx context.Context, logger *otelzap.Logger) *AWS {
	sAccessKey, exists := os.LookupEnv("SCW_ACCESS_KEY")
	if !exists {
		logger.Ctx(ctx).Fatal("No SCW_ACCESS_KEY env variable")
	}

	sSecretKey, exists := os.LookupEnv("SCW_SECRET_KEY")
	if !exists {
		logger.Ctx(ctx).Fatal("No SCW_SECRET_KEY env variable")
	}

	aRegion, exists := os.LookupEnv("AWS_REGION")
	if !exists {
		logger.Ctx(ctx).Fatal("No AWS_REGION env variable")
	}

	aEndpoint, exists := os.LookupEnv("AWS_ENDPOINT")
	if !exists {
		logger.Ctx(ctx).Fatal("No AWS_ENDPOINT env variable")
	}

	aName, exists := os.LookupEnv("AWS_BUCKET_NAME")
	if !exists {
		logger.Ctx(ctx).Fatal("No AWS_BUCKET_NAME env variable")
	}

	return &AWS{
		SCWAccessKey: sAccessKey,
		SCWSecretKey: sSecretKey,
		Region:       aRegion,
		Endpoint:     aEndpoint,
		BucketName:   aName,
	}
}

// newCorsFromEnv return a config.Cors struct filled with env variable
// func newCorsFromEnv(logger *otelzap.Logger) Cors {
// 	cAllowedOrigins, exists := os.LookupEnv("CORS_ALLOWED_ORIGINS")
// 	if !exists {
// 		logger.Ctx(ctx).Fatal("No CORS_ALLOWED_ORIGINS env variable")
// 	}

// 	return Cors{
// 		AllowedOrigins: ,
// 	}
// }
