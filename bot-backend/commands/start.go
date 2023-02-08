package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Hi, nice to meet you! I am DCube Translate bot, and I am here to help although you don'y know English."
		SendMessage(msg, bot)
	}
}
