package main

import (
	"log"

	"github.com/caiosilvestre/gopher-economy/bot"
	"github.com/caiosilvestre/gopher-economy/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Erro ao configurar o logger: %v", err)
	}

	appLogger := logger.GetLogger()

	appLogger.Info("Application initialized successfully!")

	defer appLogger.Sync()

	bot.Init()
}
