package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func GetDropboxFileRequestInfo(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		fileRequestID := update.Message.CommandArguments()

		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Got it, fetching info on Dropbox file request...",
		)
		SendMessage(msg, bot)

		// Get Dropbox client
		dbx := utils.NewDropboxClient(
			utils.GetDropboxAppKey(),
			utils.GetDropboxAppSecret(),
			utils.GetDropboxRefreshToken(),
		)

		if fileRequest, err := dbx.GetFileRequestInfo(fileRequestID); err == nil {
			reply := fmt.Sprintf(
				`File request info:
Title: %s
URL: %s
Created On: %s`,
				fileRequest.Title,
				fileRequest.URL,
				fileRequest.Created,
			)

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		} else {
			msg = tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Oops, something went wrong. I am unable to get info on file request now.",
			)
		}

		SendMessage(msg, bot)
	}
}
