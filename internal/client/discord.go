package client

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/bwmarrin/discordgo"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// InitDiscord will return a discord client initialized
func InitDiscord(ctx context.Context, logger *otelzap.Logger, conf *configs.Discord) *discordgo.Session {
	discord, err := discordgo.New("Bot " + conf.AuthToken)
	if err != nil {
		logger.Ctx(ctx).Fatal("Discord client failed to init",
			zap.Error(err),
		)
	}

	// Open a websocket connection to Discord and begin listening.
	if err := discord.Open(); err != nil {
		logger.Ctx(ctx).Info("Error while oppening the discord WS connection", zap.Error(err))
	}

	logger.Ctx(ctx).Info("Discord client initialized")

	return discord
}
