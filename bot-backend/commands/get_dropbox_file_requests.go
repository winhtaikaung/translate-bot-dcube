package commands

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func GetDropboxFileRequests(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Got it, fetching all Dropbox file requests...",
		)
		SendMessage(msg, bot)

		// Get Dropbox client
		dbx := utils.NewDropboxClient(
			utils.GetDropboxAppKey(),
			utils.GetDropboxAppSecret(),
			utils.GetDropboxRefreshToken(),
		)

		if allFileRequests, err := dbx.GetFileRequests(); err == nil {
			reply := "Here's the list of available file requests:\n"
			list := []string{}
			count := 1

			for _, fileRequest := range *allFileRequests {
				if fileRequest.IsOpen {
					getInfoCommand := fmt.Sprintf("/getfilerequest %s", fileRequest.ID)
					list = append(list,
						fmt.Sprintf(
							"%d. [%s][%s]: %s",
							count,
							fileRequest.Title,
							getInfoCommand,
							fileRequest.URL,
						),
					)
					count++
				}
			}

			reply += strings.Join(list, "\n")
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		} else {
			msg = tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Oops, something went wrong. I am unable to get file requests now.",
			)
		}

		SendMessage(msg, bot)
	}
}
