<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="docs/images/logo.png" alt="Logo" height="150">
  </a>

  <h3 align="center">Gopher Economy</h3>
</div>

GopherEconomy is a Discord bot built with [Go (Golang)](https://go.dev/) to provide currenty  quote, economic data, and market insights. Designed for simplicity and efficiency, Gophernomy is perfect for anyone who wants quick access to market data.

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
!help -> List commands
!q <Abbreviation1> <Abbreviation2> -> Show Current <Abbreviation1> in <Abbreviation2>
Example:
!q usd brl -> Show Current Dollar in BR
!q <Abbreviation1> <Abbreviation2> <Amount value> -> Show Current <Abbreviation1> in <Abbreviation2><Amount value>
Example:
!q usd brl 2.50 -> Show Current Amount in BR, R$ 15,00
```
# Using Git in Jupyter Notebooks or Google Colab

If you're working in a Jupyter Notebook or Google Colab environment, you can use !git commands to manage your repository directly from the notebook.

###Common Git Commands:

- Clone the repository:

```bash
!git clone https://github.com/yourusername/GopherEconomy.git
```
- Check the status of the repository:

```bash
!git status
```

- Pull the latest changes from the remote repository:

```bash
!git pull origin main
```

- Add changes to staging:

```bash
!git add .
```

- Commit changes:

```bash
!git commit -m "Your commit message"
```

- Push changes to GitHub:

```bash
!git push origin main
```

# Contributing

Contributions are welcome!

# Support

If you encounter any issues or have suggestions, feel free to open an issue.

With just a few commands, GopherEconomy keeps your server updated with BRL/USD and BTC/USD rates. Happy coding! 😊
