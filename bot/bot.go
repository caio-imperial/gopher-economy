package bot

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/caiosilvestre/gopher-economy/integrations/awesomeapi/economia"
	"github.com/joho/godotenv"
)

var (
	//discord_token is a token using to connect with discord bot
	discord_token string
)

// Init initializes all environment variables, configures the Discord bot,
// and registers callback functions on the handler.
func Init() (err error) {
	godotenv.Load(".env")

	discord_token = os.Getenv("DISCORD_TOKEN")
	if discord_token == "" {
		fmt.Printf("DISCORD_TOKEN not set in environment\n")
		return
	}

	//start package economia
	err = economia.Init()
	if err != nil {
		fmt.Printf("error economia.Init(): %v\n", err)
		return
	}

	//init discord bot session
	dg, err := discordgo.New("Bot " + discord_token)
	if err != nil {
		fmt.Printf("error creating Discord session: %v\n", err)
		return
	}

	//register messageHandler in discord bot
	dg.AddHandler(messageHandler)

	// Set the bot's intents to include all non-privileged events.
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	//Open conection with discord
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	//Close conection with discord
	defer dg.Close()

	fmt.Println("Bot is Online")

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

	// Compares the received message and responds if the condition is true.
	switch {

	// Convert currency amount
	case strings.HasPrefix(m.Content, "!q"):
		// Clean up multiple spaces in the message
		message := cleanUpMultipleSpaces(m.Content)
		s.ChannelMessageSend(m.ChannelID, ConvertMessage(message))

	// Show the information about repository
	case strings.HasPrefix(m.Content, "!git"):
		s.ChannelMessageSend(m.ChannelID, fmt.Sprint(
			" Feel free to fork the repository, make your changes, and submit a pull request.",
			" We appreciate your contributions!\n",
			" [Git Repository URL](https://github.com/caio-imperial/gopher-economy)",
		))

		// List all commands
	case strings.HasPrefix(m.Content, "!help"):
		s.ChannelMessageSend(m.ChannelID, fmt.Sprint(
			"```makefile\n",
			"!help -> List commands\n",
			"!q <Abbreviation1> <Abbreviation2> -> Show Current <Abbreviation1> in <Abbreviation2> \n",
			"Example: \n!q usd brl -> Show Current Dollar in BR\n",
			"!q <Abbreviation1> <Abbreviation2> <Amount value> -> Show Current <Abbreviation1> in <Abbreviation2><Amount value>\n",
			"Example: \n!q usd brl 2.50 -> Show Current Amount in BR, R$ 15,00\n",
			"```\n",
		))

	// Do nothing if none of the conditions are satisfied.
	default:
		return
	}

}

// Function to clean up multiple spaces in a string.
func cleanUpMultipleSpaces(message string) string {
	spaceReCleanUp := regexp.MustCompile(`\s+`)
	return spaceReCleanUp.ReplaceAllString(strings.TrimSpace(message), " ")
}
