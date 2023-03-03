package configs

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// GetConfig is used to choose the right manner to choose configuration (kube or local)
func (conf *Config) GetConfig(ctx context.Context, logger *otelzap.Logger) *Config {
	configPath := os.Getenv("CONF_PATH")

	// If the env is not set, it will take the configuration from env variable
	if configPath == "" {
		logger.Ctx(ctx).Info("As there are not CONF_PATH set, switching to env variable")

		newConfigFromEnv(ctx, logger, conf)
		newSecretFromEnv(ctx, logger, conf)

		return conf
	}

	if err := newConfigFromFile(ctx, logger, conf, configPath); err != nil {
		logger.Ctx(ctx).Fatal("Cannot read configuration",
			zap.Error(err),
		)
	}

	newSecretFromEnv(ctx, logger, conf)

	return conf
}

// NewSecretFromEnv is used to get secret used by api from environment
func newSecretFromEnv(ctx context.Context, logger *otelzap.Logger, conf *Config) {
	// Initialize config interface
	jRefreshSecret, exists := os.LookupEnv("JWT_REFRESH_SECRET")
	if !exists {
		logger.Ctx(ctx).Fatal("No JWT_REFRESH_SECRET env variable")
	}

	rPassword, exists := os.LookupEnv("REDIS_PASSWORD")
	if !exists {
		logger.Ctx(ctx).Fatal("No REDIS_PASSWORD env variable")
	}

	mPassword, exists := os.LookupEnv("MARIADB_PASSWORD")
	if !exists {
		logger.Ctx(ctx).Fatal("No MARIADB_PASSWORD env variable")
	}

	sAccessKey, exists := os.LookupEnv("SCW_ACCESS_KEY")
	if !exists {
		logger.Ctx(ctx).Fatal("No SCW_ACCESS_KEY env variable")
	}

	sSecretKey, exists := os.LookupEnv("SCW_SECRET_KEY")
	if !exists {
		logger.Ctx(ctx).Fatal("No SCW_SECRET_KEY env variable")
	}

	mgAPIKeyPrivate, exists := os.LookupEnv("MG_APIKEY_PRIVATE")
	if !exists {
		logger.Ctx(ctx).Fatal("No MG_APIKEY_PRIVATE env variable")
	}

	centrifugoAPIKey, exists := os.LookupEnv("CENTRIFUGO_API_KEY")
	if !exists {
		logger.Ctx(ctx).Fatal("No CENTRIFUGO_API_KEY env variable")
	}

	discordToken, exists := os.LookupEnv("DISCORD_TOKEN")
	if !exists {
		logger.Ctx(ctx).Fatal("No DISCORD_TOKEN env variable")
	}

	conf.Jwt.RefreshSecret = jRefreshSecret
	conf.Redis.Password = rPassword
	conf.Mariadb.Password = mPassword
	conf.AWS.SCWAccessKey = sAccessKey
	conf.AWS.SCWSecretKey = sSecretKey
	conf.Mailgun.ClientSecret = mgAPIKeyPrivate
	conf.Centrifugo.APIKey = centrifugoAPIKey
	conf.Discord.AuthToken = discordToken
}

// newConfigFromEnv is used for initialize the configuration set in environment variable
func newConfigFromEnv(ctx context.Context, logger *otelzap.Logger, conf *Config) {
	// Initialize config from env
	conf.Host = newHostFromEnv(ctx, logger)
	conf.Server = newServerFromEnv(ctx, logger)
	conf.Jwt = newJWTFromEnv(ctx, logger)
	conf.Certs = newCertsFromEnv(ctx, logger)
	conf.Mariadb = newMariaDBFromEnv(ctx, logger)
	conf.Redis = newRedisFromEnv(ctx, logger)
	conf.Centrifugo = newCentrifugoFromEnv(ctx, logger)
	conf.AppInfo = newAppInfoFromEnv(ctx, logger)
	conf.Jaeger = newJaegerFromEnv(ctx, logger)
	conf.AWS = newAWSFromEnv(ctx, logger)
	conf.Mailgun = newMailGunFromEnv(ctx, logger)
}

// NewConfigFromFile is used for parse configuration in file
func newConfigFromFile(ctx context.Context, logger *otelzap.Logger, conf *Config, configPath string) error {
	// Read file content
	file, err := os.ReadFile(configPath)
	if err != nil {
		logger.Ctx(ctx).Info("error reading config file",
			zap.String("Config path", configPath),
			zap.Error(err),
		)

		return err
	}

	// Choose the right file type
	switch filepath.Ext(configPath) {
	case ".json":
		if err := newConfigFromJSON(ctx, logger, conf, string(file)); err != nil {
			return err
		}
	case ".yaml":
		if err := newConfigFromYAML(ctx, logger, conf, string(file)); err != nil {
			return err
		}
	default:
		newConfigFromEnv(ctx, logger, conf)
	}

	return nil
}

// NewConfigFromJSON is used for initialize the configuration set in config.json
func newConfigFromJSON(ctx context.Context, logger *otelzap.Logger, conf *Config, jsonFile string) error {
	confContent := []byte(os.ExpandEnv(jsonFile))
	if err := json.Unmarshal(confContent, &conf); err != nil {
		logger.Ctx(ctx).Info("error reading json config file",
			zap.Error(err),
		)

		return err
	}

	return nil
}

// NewConfigFromYAML is used for initialize the configuration set in config.yaml
func newConfigFromYAML(ctx context.Context, logger *otelzap.Logger, conf *Config, yamlFile string) error {
	confContent := []byte(os.ExpandEnv(yamlFile))
	if err := yaml.Unmarshal(confContent, &conf); err != nil {
		logger.Ctx(ctx).Info("error reading YAML config file",
			zap.Error(err),
		)

		return err
	}

	return nil
}
