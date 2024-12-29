package bot

import (
	"fmt"
	"os"
	"os/signal"
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
func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
	}

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
}

// messageHandler is a callback function that processes messages from the channel
// and returns an appropriate response.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	var currencySymbol string
	var quoteCurrencys []string
	err := error(nil)
	result := economia.QuoteCurrency{}

	// Compares the received message and responds if the condition is true.
	// economia.GetQuote return quote currency
	switch m.Content {
	case "!dollar ptax, !usd ptax":
		result, err = economia.GetQuote("USD-BRLPTAX")
	case "!dollar, !usd":
		result, err = economia.GetQuote("USD-BRL")
	case "!euro, !eur":
		result, err = economia.GetQuote("EUR-BRL")
	case "!btc", "!bitcoin":
		result, err = economia.GetQuote("BTC-BRL")
	case "!btc dollar", "!btc usd", "!bitcoin dollar", "!bitcoin usd":
		result, err = economia.GetQuote("BTC-USD")
	case "!eth", "!etherium":
		result, err = economia.GetQuote("ETH-BRL")
	case "!help":
		s.ChannelMessageSend(m.ChannelID, fmt.Sprint(
			"```makefile\n",
			"!dollar ptax or !usd ptax -> Show Current PTAX Dollar in BRL\n",
			"!dollar or !usd -> Show Current Dollar in BRL\n",
			"!euro or !eur -> Show Current Euro in BRL\n",
			"!btc or !bitcoin -> Show Current BTC in BRL\n",
			"!btc dollar or !bitcoin dollar or !btc usd or !bitcoin usd -> Show Current BTC in USD-Dollar \n",
			"!eth or !etherium -> Show Current ETH in BRL\n",
			"!q <Abbreviation1> <Abbreviation2> -> Show Current <Abbreviation1> in <Abbreviation2> \n",
			"Example: \n!q usd brl -> Show Current Dollar in BR\n",
			"```\n",
		))
	// Do nothing if none of the conditions are satisfied.
	default:
		if !strings.HasPrefix(m.Content, "!q") {
			return
		}
		currencySymbol, quoteCurrencys, err = NomalizationQuotes(m.Content)
		if err != nil {
			fmt.Print(err)
			return
		}
		result, err = economia.GetQuote(currencySymbol)
	}

	if err != nil {
		fmt.Print(err)
		s.ChannelMessageSend(m.ChannelID, "failed to search quote currency")
	}

	// send quote current message in the channel if result.Bid has a value
	if result.Bid > 0 {
		// choose the correct symbol to send in discord message
		quoteSymbol := "$"
		switch quoteCurrencys[1] {
		case "brl":
			quoteSymbol = "R$"
		case "eur":
			quoteSymbol = "€"
		}

		// send discord message
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %.2f", quoteSymbol, result.Bid))
	}

}
