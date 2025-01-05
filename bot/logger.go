package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/caiosilvestre/gopher-economy/logger"
)

func InitMessageLogger(message *discordgo.MessageCreate) *logger.AppLogger {
	messageLogger := logger.GetLogger()
	messageLogger.Sugar = messageLogger.Sugar.With(
		"author_id", message.Author.ID,
		"server_id", message.GuildID,
		"channel_id", message.ChannelID,
		"message_id", message.ID,
		"message_text", message.Content,
	)
	messageLogger.Info("teste config")
	return messageLogger
}
