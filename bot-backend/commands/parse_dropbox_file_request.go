package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

func HandleDropboxFileRequest(update *tgbotapi.Update, bot *tgbotapi.BotAPI, params ...string) {
	if len(params) > 1 && params[0] == "yes" {
		fileRequestName := params[1]
		log.Printf("file request name: %s\n", fileRequestName)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Roger that! Creating Dropbox file request %s, give me a moment", fileRequestName),
		)
		SendMessage(msg, bot)

		// Get Dropbox client
		dbx := utils.NewDropboxClient(utils.GetDropboxAccessToken())
		if allFileRequests, err := dbx.GetFileRequests(); err == nil {
			log.Println(allFileRequests)
		}
	}
}
