package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/translate-bot-dcube/bot-backend/bot"
	"github.com/translate-bot-dcube/bot-backend/commands"
	"github.com/translate-bot-dcube/bot-backend/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Env(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"task_root":         os.Getenv("LAMBDA_TASK_ROOT"),
		"app_env":           os.Getenv("app_env"),
		"lambda_invoke_url": utils.GetLambdaInvokeUrl(),
	})
}

func AboutBot(c *gin.Context) {
	botToken := utils.GetTelegramBotToken()
	if res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", botToken)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch bot info"})
	} else {
		defer res.Body.Close()
		var botInfo bot.BotInfo
		if err := json.NewDecoder(res.Body).Decode(&botInfo); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode bot info"})
		} else {
			c.JSON(http.StatusOK, botInfo)
		}
	}
}

func InitBot(c *gin.Context) {
	if _, err := bot.NewTelegramBot(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to setup Telegram bot API"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func WebHook(c *gin.Context) {
	// get bot
	if bot, err := bot.NewTelegramBot(); err != nil {
		log.Println("Webhook unable to init bot")
	} else {
		if update, err2 := bot.HandleUpdate(c.Request); err2 != nil {
			log.Println("Webhook unable to parse update")
		} else {
			if update.Message != nil {
				cmdStr := commands.ParseCommand(update)
				if cmdStr != "" {
					log.Println("Received command: ", cmdStr)
					cmd := commands.GetCommandFunc(cmdStr)
					cmd(update, bot)
				} else {
					log.Println("Message received: ", update.Message.Text)
					// translated, translatedStr, language := commands.TranslateToEnglishByOpenAI(update.Message.Text)
					// if translated {
					// fmt.Print(translated, translatedStr, language)
					// here should use the text to do government searching
					// then translate back to the original language and return
					translatedStr := ""
					searchedText := update.Message.Text
					translatedStr = commands.TranslateBackByOpenAI(searchedText)
					// }
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, translatedStr))
					if err != nil {
						log.Println("Webhook unable to send message")
					}
				}
			} else if update.CallbackQuery != nil {
				/*
					// Flashes update.CallbackQuery.Data in Telegram window like a toast
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
					if _, err := bot.Request(callback); err != nil {
						log.Println(err)
					}
				*/

				/*
					// Sends query data back to user
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
				*/

				queryInfo := commands.ParseCallbackQuery(update, bot)
				log.Println("Received query: ", queryInfo.QueryType)
				commands.HandleReplyToCommand(queryInfo, update, bot)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
