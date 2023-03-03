package database

import (
	"context"
	"time"

	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/ent"

	// Import runtime to register the schema hooks
	_ "github.com/naofel1/api-golang-template/internal/ent/runtime"

	// Import mySQL database connection method
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

// ConnectDatabase make connection to the database
func ConnectDatabase(ctx context.Context, logger *otelzap.Logger, conf *configs.Mariadb) *ent.Client {
	// Register the otelsql wrapper for the provided mariadb driver.
	regstr, err := otelsql.Register(dialect.MySQL, otelsql.WithAttributes(
		semconv.DBSystemMySQL,
	))
	if err != nil {
		logger.Ctx(ctx).Fatal("failed to register mariadb",
			zap.String("DB Name: ", conf.DBName),
			zap.String("DB User: ", conf.User),
			zap.String("DB Host: ", conf.Host),
			zap.String("DB Port: ", conf.Port),
			zap.Error(err),
		)
	}

	params := map[string]string{
		"charset": conf.MariaParams.Charset,
		"loc":     conf.MariaParams.Location,
	}

	mysqlConfig := mysql.Config{
		DBName:               conf.DBName,
		User:                 conf.User,
		Passwd:               conf.Password,
		Addr:                 conf.Host + ":" + conf.Port,
		Net:                  conf.Net,
		ParseTime:            conf.ParseTime,
		AllowNativePasswords: true,
		InterpolateParams:    true,
		Params:               params,
	}

	// Connect to a mysql database using the mariadb wrapper.
	// create connection with sql.OpenDB
	db, err := otelsql.Open(regstr, mysqlConfig.FormatDSN())
	if err != nil {
		logger.Ctx(ctx).Fatal("failed to open DB",
			zap.Any("Config", conf),
			zap.Error(err),
		)
	}

	if err := otelsql.RegisterDBStatsMetrics(db, otelsql.WithAttributes(
		semconv.DBSystemMySQL,
	)); err != nil {
		logger.Ctx(ctx).Info("Failed to record stats", zap.Error(err))
	}

	const maxConnection = 10

	db.SetMaxIdleConns(maxConnection)
	db.SetMaxOpenConns(0)
	db.SetConnMaxLifetime(time.Hour)

	logger.Ctx(ctx).Info("Trying to open database connection...")

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)

	// Return a new client of the database
	client := ent.NewClient(ent.Driver(drv))

	logger.Ctx(ctx).Info("Successfully opened database connection")

	logger.Ctx(ctx).Info("Begin database migration...")

	// Run the auto migration tool.
	if err := client.Schema.Create(
		ctx,
	); err != nil {
		logger.Ctx(ctx).Fatal("Database Migration Failed",
			zap.String("DB Name: ", conf.DBName),
			zap.String("DB User: ", conf.User),
			zap.String("DB Host: ", conf.Host),
			zap.String("DB Port: ", conf.Port),
			zap.Error(err),
		)
	}

	logger.Ctx(ctx).Info("Database Migration Completed")

	return client
}
