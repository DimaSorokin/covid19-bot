package main

import (
	"github.com/DimaSorokin/covid-bot/pkg/logging"
	"github.com/DimaSorokin/covid-bot/pkg/telegram"
	"github.com/DimaSorokin/go-medzakupivli-sdk/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	logging.Init("info")
	logger := logging.GetLogger()
	logger.Println("Running Application")
	botApi, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		logger.Fatal(err)
	}
	botApi.Debug = true

	mozClient := client.New()
	if err != nil {
		logger.Fatal(err)
	}
	bot := telegram.NewBot(botApi, mozClient, logger)
	if err := bot.Start(); err != nil {
		logger.Fatal(err)
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
