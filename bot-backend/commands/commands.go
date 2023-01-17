package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandFunc func(update *tgbotapi.Update, bot *tgbotapi.BotAPI)

type CommandInfo struct {
	Description string
	Func        CommandFunc
}

var mapping map[string]*CommandInfo

func sendMessage(msg tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Webhook unable to send message")
	}
}

func ParseCommand(update *tgbotapi.Update) string {
	if update.Message != nil && update.Message.IsCommand() {
		return update.Message.Command()
	}
	return ""
}

func GetCommands() map[string]*CommandInfo {
	if mapping == nil {
		log.Println("Building command mapping...")
		mapping = make(map[string]*CommandInfo)
		mapping["help"] = &CommandInfo{
			Description: "Get list of available commands",
			Func:        Help,
		}
		mapping["dropbox"] = &CommandInfo{
			Description: "Get list of available Dropbox commands",
			Func:        GetDropboxCommands,
		}
		mapping["makefilerequest"] = &CommandInfo{
			Description: "Make a new Dropbox file request",
			Func:        MakeDropboxFileRequest,
		}
	}
	return mapping
}

func GetCommandFunc(command string) CommandFunc {
	commands := GetCommands()
	if cmd := commands[command]; cmd != nil {
		return cmd.Func
	}
	log.Println("Returning default command...")
	return NotFoundCommand
}

func Help(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Hi, trying using /dropbox"
		sendMessage(msg, bot)
	}
}

func GetDropboxCommands(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Hi, the following commands are available:\n/makefilerequest : Make a new Dropbox file request"
		sendMessage(msg, bot)
	}
}

func NotFoundCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil && update.Message.Chat != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand the command, try /help")
		sendMessage(msg, bot)
	}
}
