package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Language(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	languages := [5]string{"Burmese", "Chinese/Mandarin", "Malay", "Tamil", "English"}

	if update.Message != nil && update.Message.Chat != nil {
		reply := "Here's the following languages I can understand:\n"

		for i, language := range languages {
			reply += fmt.Sprintf("%d) %s\n", i+1, language)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		SendMessage(msg, bot)
	}
}
