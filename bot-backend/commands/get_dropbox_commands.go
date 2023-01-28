package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetDropboxCommands(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		reply := "Hi, the following Dropbox commands are available:"

		// Get all Dropbox related commands
		commands := GetCommands(CATEGORY_DROPBOX)
		for command, info := range commands {
			reply += GetCommandOneLinerDesc(command, info, true)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		SendMessage(msg, bot)
	}
}
