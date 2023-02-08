package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CATEGORY_ALL      = "all"
	CATEGORY_TOP      = "top"
	CATEGORY_LANGUAGE = "language"
)

type CommandFunc func(update *tgbotapi.Update, bot *tgbotapi.BotAPI)

type CommandInfo struct {
	Name        string
	Description string
	Category    string
	Func        CommandFunc
}

var mapping map[string]*CommandInfo

func SendMessage(msg tgbotapi.MessageConfig, bot *tgbotapi.BotAPI) {
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

func GetCommandOneLinerDesc(command string, info *CommandInfo, addLineBreak bool) string {
	lineBreak := ""
	if addLineBreak {
		lineBreak = "\n"
	}
	return fmt.Sprintf("/%s (%s) : %s%s", command, info.Name, info.Description, lineBreak)
}

func GetCommands(category string) map[string]*CommandInfo {
	if mapping == nil {
		log.Println("Building command mapping...")
		mapping = make(map[string]*CommandInfo)
		mapping["help"] = &CommandInfo{
			Name:        "Help command",
			Description: "Get list of available commands",
			Category:    CATEGORY_TOP,
			Func:        Help,
		}
		mapping["languages"] = &CommandInfo{
			Name:        "Languages support command",
			Description: "Get list of supported languages",
			Category:    CATEGORY_LANGUAGE,
			Func:        Language,
		}
	}

	if category == CATEGORY_ALL {
		return mapping
	}

	filteredMapping := make(map[string]*CommandInfo)
	for command, info := range mapping {
		if info.Category == category {
			filteredMapping[command] = info
		}
	}

	return filteredMapping
}

func GetCommandFunc(command string) CommandFunc {
	commands := GetCommands(CATEGORY_ALL)

	if command == "start" {
		log.Println("Returning Start command...")
		return Start
	}

	if cmd := commands[command]; cmd != nil {
		return cmd.Func
	}

	log.Println("Returning default command...")
	return NotFoundCommand
}
