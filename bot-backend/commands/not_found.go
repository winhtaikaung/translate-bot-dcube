package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NotFoundCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand the command, try /help")
		SendMessage(msg, bot)
	}
}
