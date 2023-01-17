package utils

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	existingWebHook, err3 := bot.GetWebhookInfo()

	log.Println(existingWebHook.URL, err3)
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
