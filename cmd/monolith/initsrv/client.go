package initsrv

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/naofel1/api-golang-template/internal/client"
	"github.com/naofel1/api-golang-template/internal/client/database"
	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// Client will hold client that will be injected
// into the Service layer on service initialization
type Client struct {
	Mailer  *mailgun.MailgunImpl
	Discord *discordgo.Session
	Redis   *redis.Client
	MariaDB *ent.Client
}

// InitClient will initialize all client
func InitClient(ctx context.Context, logger *otelzap.Logger, conf *configs.Config) *Client {
	return &Client{
		MariaDB: database.ConnectDatabase(ctx, logger, conf.Mariadb),
		Redis:   database.ConnectRedis(ctx, logger, conf.Redis),
		Discord: client.InitDiscord(ctx, logger, conf.Discord),
		Mailer:  client.InitMailGun(ctx, logger, conf.Mailgun),
	}
}
