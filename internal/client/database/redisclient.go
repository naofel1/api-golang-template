package database

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"

	// Import mySQL database connection method
	_ "github.com/go-sql-driver/mysql"
)

// ConnectRedis make connection to the database
func ConnectRedis(ctx context.Context, logger *otelzap.Logger, conf *configs.Redis) *redis.Client {
	// 👇 Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.ConnectionString,
		Password: conf.Password,
		DB:       conf.SelectedDB,
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	// Try to ping Redis Server
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		logger.Ctx(ctx).Fatal("Redis Ping error",
			zap.Error(err),
		)
	}

	logger.Ctx(ctx).Info("Redis client connected successfully...")

	return rdb
}
