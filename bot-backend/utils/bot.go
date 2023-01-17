package utils

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/commands"
)

var bot *tgbotapi.BotAPI

type BotInfoResult struct {
	ID       int64  `json:"id"`
	Name     string `json:"first_name"`
	UserName string `json:"username"`
}

type BotInfo struct {
	Result BotInfoResult `json:"result"`
}

func setupWebHook(bot *tgbotapi.BotAPI) {
	// re-setup webhook
	newWebHook, err := tgbotapi.NewWebhookWithCert(GetLambdaInvokeUrl()+"/bot"+GetTelegramBotToken(), nil)

	if err != nil {
		log.Println("unable to create new webhook", err)
	}

	_, err2 := bot.Request(newWebHook)
	if err != nil {
		log.Println("webhook request via bot failed", err2)
	}

	if !IsProductionEnv() {
		existingWebHook, err3 := bot.GetWebhookInfo()
		log.Println(existingWebHook.URL, err3)
	}
}

func setupCommands(bot *tgbotapi.BotAPI) {
	// get list of available commands
	commands := commands.GetCommands()

	botCommands := make([]tgbotapi.BotCommand, 0)

	for cmd, info := range commands {
		botCommands = append(botCommands, tgbotapi.BotCommand{
			Command:     "/" + cmd,
			Description: info.Description,
		})
	}

	cfg := tgbotapi.NewSetMyCommands(botCommands...)

	if _, err := bot.Request(cfg); err != nil {
		log.Println("Set bot commands failed", err)
	} else {
		log.Println("Bot commands registered", err)
	}
}

func NewTelegramBot() (*tgbotapi.BotAPI, error) {
	if bot != nil {
		return bot, nil
	}

	newBot, err := tgbotapi.NewBotAPI(GetTelegramBotToken())
	if err == nil {
		bot = newBot
		newBot.Debug = !IsProductionEnv()
	}

	setupWebHook(newBot)
	setupCommands(newBot)

	return bot, err
}
