package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Hi, nice to meet you! I am PSGNaviBot, and I am here to help with Dropbox requests and answer some NVPS PSG questions."
		SendMessage(msg, bot)
	}
}
