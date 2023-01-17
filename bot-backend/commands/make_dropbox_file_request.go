package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MakeDropboxFileRequest(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		fileRequestName := update.Message.CommandArguments()
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf(
				"You wanted to create Dropbox file request with name: %s\nIs that correct?",
				fileRequestName,
			),
		)
		msg.ReplyMarkup = NewYesNoKeyboard(
			fmt.Sprintf("replytocommand,makefilerequest,yes,%s", fileRequestName),
			fmt.Sprintf("replytocommand,makefilerequest,no,%s", fileRequestName),
		)
		sendMessage(msg, bot)
	}
}
