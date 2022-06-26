package telegram

import (
	"fmt"
	"github.com/DimaSorokin/covid-bot/internal/region"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) error {
	regionCurrent := region.CheckRegion(callback.Data)
	b.logger.Info(regionCurrent)
	if "unknown" == regionCurrent {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Такого міста не існує, спробуйте ще")
		_, err := b.bot.Send(msg)
		return err
	}

	hospitals, err := b.client.GetHostitals()
	if err != nil {
		log.Fatal(err)
	}
	var prepHospitalName string
	for _, hospital := range *hospitals {
		if regionCurrent == hospital.Region {
			prepHospitalName += fmt.Sprintf(" %s\n\n\n", hospital.Name)
		}
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, prepHospitalName)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	regions := region.GetRegion()
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for key, _ := range regions {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(key, key)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Введіть назву міста, яке вас цікавить")
	msg.ReplyMarkup = keyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Невідома команда")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	regionCurrent := region.CheckRegion(message.Text)
	if "unknown" == regionCurrent {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Такого міста не існує, спробуйте ще")
		_, err := b.bot.Send(msg)
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	_, err := b.bot.Send(msg)
	return err
}
