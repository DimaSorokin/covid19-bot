package telegram

import (
	"github.com/DimaSorokin/covid-bot/pkg/logging"
	"github.com/DimaSorokin/go-medzakupivli-sdk/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	client *client.Client
	logger *logging.Logger
}

func NewBot(bot *tgbotapi.BotAPI, client *client.Client, logger *logging.Logger) *Bot {
	return &Bot{
		bot:    bot,
		client: client,
		logger: logger,
	}
}

func (b *Bot) Start() error {
	b.logger.Info("start bot")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ignore any non-Message Updates
			//Handle commands
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					b.handleError(update.Message.Chat.ID, err)
				}
				continue
			}
		} else if update.CallbackQuery != nil {
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := b.bot.Send(msg); err != nil {
				b.logger.Fatal(err)
			}

			if err := b.handleCallback(update.CallbackQuery); err != nil {
				b.logger.Fatal(err)
			}

		}
	}

	return nil
}
