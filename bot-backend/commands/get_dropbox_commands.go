package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetDropboxCommands(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		reply := "Hi, the following Dropbox commands are available:"

		// Get all Dropbox related commands
		commands := GetCommands(CATEGORY_DROPBOX)
		for command, info := range commands {
			reply += GetCommandOneLinerDesc(command, info, true)
		}

		msg.Text = `Hi, the following commands are available:
/makefilerequest : Make a new Dropbox file request
/listfilerequest : List all file requests
/getfilerequest  : Get info on a file request`
		SendMessage(msg, bot)
	}
}
