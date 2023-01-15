package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupTelegramBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(GetTelegramBotToken())
	if bot != nil {
		bot.Debug = !IsProductionEnv()
	}
	return bot, err
}
