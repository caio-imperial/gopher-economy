package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/caiosilvestre/gopher-economy/integrations/awesomeapi/economia"
	"github.com/caiosilvestre/gopher-economy/logger"
)

var (
	//discord_token is a token using to connect with discord bot
	discord_token string
)

// Init initializes all environment variables, configures the Discord bot,
// and registers callback functions on the handler.
func Init() (err error) {
	appLogger := logger.GetLogger()

	discord_token = os.Getenv("DISCORD_TOKEN")
	if discord_token == "" {
		appLogger.Error("DISCORD_TOKEN not set in environment")
		return
	}

	//start package economia
	err = economia.Init()
	if err != nil {
		appLogger.Error("error start economia", "err", err)
		return
	}

	//init discord bot session
	dg, err := discordgo.New("Bot " + discord_token)
	if err != nil {
		appLogger.Error("error creating Discord session", "err", err)
		return
	}

	//register messageHandler in discord bot
	dg.AddHandler(messageHandler)

	// Set the bot's intents to include all non-privileged events.
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	//Open conection with discord
	err = dg.Open()
	if err != nil {
		appLogger.Error("error opening connection", "err", err)
		return
	}

	//Close conection with discord
	defer dg.Close()
	appLogger.Info("Bot is Online", "status", "active")

	// Close the connection with Discord gracefully when Ctrl+C is pressed in the terminal.
	// A channel is created to listen for termination signals (SIGINT, SIGTERM, or os.Interrupt).
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return
}

// messageHandler is a callback function that processes messages from the channel
// and returns an appropriate response.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Init message logger with credentials
	messageLogger := InitMessageLogger(m)

	// Compares the received message and responds if the condition is true.
	switch {

	// Convert currency amount
	case strings.HasPrefix(m.Content, "!q"):
		messageLogger.Info("Valid command detected", "command_received", "!q")
		s.ChannelMessageSend(m.ChannelID, ConvertMessage(messageLogger, m.Content))
		messageLogger.Info("Command executed successfully", "command_execute", "!q")

	// List all commands
	case strings.HasPrefix(m.Content, "!help"):
		messageLogger.Info("Valid command detected", "command_received", "!help")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprint(
			"```makefile\n",
			"!help -> List commands\n",
			"!q <Abbreviation1> <Abbreviation2> -> Show Current <Abbreviation1> in <Abbreviation2> \n",
			"Example: \n!q usd brl -> Show Current Dollar in BR\n",
			"!q <Abbreviation1> <Abbreviation2> <Amount value> -> Show Current <Abbreviation1> in <Abbreviation2><Amount value>\n",
			"Example: \n!q usd brl 2.50 -> Show Current Amount in BR, R$ 15,00\n",
			"```\n",
		))
		messageLogger.Info("Command executed successfully", "command_execute", "!help")

	// Do nothing if none of the conditions are satisfied.
	default:
		return
	}

}
