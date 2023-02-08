package commands

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackQueryInfo struct {
	QueryType string
	Params    []string
}

func ParseCallbackQuery(update *tgbotapi.Update, bot *tgbotapi.BotAPI) *CallbackQueryInfo {
	if update.CallbackQuery != nil {
		parts := strings.Split(update.CallbackQuery.Data, ",")
		if len(parts) > 0 {
			info := &CallbackQueryInfo{}
			info.QueryType = parts[0]

			if len(parts) > 1 {
				info.Params = parts[1:]
			} else {
				info.Params = []string{}
			}

			return info
		}
	}

	return nil
}

func HandleReplyToCommand(queryInfo *CallbackQueryInfo, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if queryInfo == nil {
		log.Println("Invalid query info")
		return
	}

	switch queryInfo.QueryType {
	case "replytocommand":
		if len(queryInfo.Params) < 2 {
			log.Println("Insufficient params for replytocommand")
			return
		}
		switch queryInfo.Params[0] {
		// case
		}
	}
}
