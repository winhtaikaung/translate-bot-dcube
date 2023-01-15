package utils

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func setupWebHook(bot *tgbotapi.BotAPI) {
	existingWebHook, err := bot.GetWebhookInfo()
	if existingWebHook.URL == "" || err != nil {
		// re-setup webhook
		// newWebHook, err := tgbotapi.NewWebhookWithCert(GetLambdaInvokeUrl() + "/bot" + GetTelegramBotToken())
		if dat, err := os.ReadFile("cert.pem"); err == nil {
			log.Println("got cert", string(dat))
		} else {
			log.Println("no cert")
		}
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

	return bot, err
}
