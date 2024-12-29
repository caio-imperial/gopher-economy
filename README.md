<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="docs/images/logo.png" alt="Logo" height="150">
  </a>

  <h3 align="center">Gopher Economy</h3>
</div>

GopherEconomy is a Discord bot built with [Go (Golang)](https://go.dev/) to provide currenty  quote, economic data, and market insights. Designed for simplicity and efficiency, Gophernomy is perfect for anyone who wants quick access to market data.


## Features 1.0.0

- [ ] 💱 Delete all commands except !help and !q.

- [ ] 🧪 Implement Tests: Create and add tests to ensure functionality and code quality.

## Fix 0.3.1

- [x] Fix: run application without .env.

## Features 0.3.0

- [x] 🚀 Implement Deploy: Set up and automate the deployment process to ensure continuous and seamless delivery of new versions.

## Features 0.2.0

- [x] 💱 Universal Conversion: Check the current price of any QuoteCurrency.


## Features 0.1.0

- [x] 💱 BTC/USD Conversion: Check the current price of Bitcoin in US Dollars.

## Features 0.0.1

- [x] 💱 USD/BRL Conversion: Check the current price of US Dollars in Brazilian Real.

- [x] 💱 BTC/USD Conversion: Check the current price of Bitcoin in Brazilian Real.

- [x] 💱 EUR/BRL Conversion: Check the current price of Euro in Brazilian Real.

- [x] 💱 ETH/BRL Conversion: Check the current price of Ethereum in Brazilian Real.

- [x] 📜 HELP command: List all the commands.

# Getting Started

### Prerequisites:

- Go (Golang) installed on your system.

- A Discord bot token.

### Installation:

- Clone the repository:

```bash
git clone https://github.com/yourusername/GopherEconomy.git
cd GopherEconomy
```

- Install dependencies:

```bash
go mod tidy
```

- create .env

```bash
cp .env.template .env
```

- Set up your environment variables in file `./.env`:

```makefile
DISCORD_TOKEN=your-discord-bot-token
BASE_URL="url"
```

- Run the bot:

```bash
go run main.go
```

# Discord Commands

```makefile
!dollar ptax or !usd ptax -> Show Current PTAX Dollar in BRL (deprecated)
!dollar or !usd -> Show Current Dollar in BRL (deprecated)
!euro or !eur -> Show Current Euro in BRL (deprecated)
!btc or !bitcoin -> Show Current BTC in BRL (deprecated)
!btc dollar or !bitcoin dollar or !btc usd or !bitcoin usd -> Show Current BTC in USD-Dollar (deprecated)
!q <Abbreviation1> <Abbreviation2> -> Show Current <Abbreviation1> in <Abbreviation2>
Example:
!q usd brl -> Show Current Dollar in BR
!eth or !etherium -> Show Current ETH in BRL
```

# Contributing

Contributions are welcome!

# Support

If you encounter any issues or have suggestions, feel free to open an issue.

With just a few commands, GopherEconomy keeps your server updated with BRL/USD and BTC/USD rates. Happy coding! 😊