package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type TelegramUser struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

type BotUpdateMessage struct {
	MessageId       int          `json:"message_id"`
	MessageThreadId int          `json:"message_thread_id"`
	From            TelegramUser `json:"from"`
	Text            string       `json:"text"`
}

type BotUpdate struct {
	UpdateId int              `json:"update_id"`
	Message  BotUpdateMessage `json:"message"`
}

type BotInfoResult struct {
	ID       int64  `json:"id"`
	Name     string `json:"first_name"`
	UserName string `json:"username"`
}

type BotInfo struct {
	Result BotInfoResult `json:"result"`
}

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
		var botInfo BotInfo
		if err := json.NewDecoder(res.Body).Decode(&botInfo); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode bot info"})
		} else {
			c.JSON(http.StatusOK, botInfo)
		}
	}
}

func SetWebHook(c *gin.Context) {
	if _, err := utils.NewTelegramBot(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to setup Telegram bot API"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func WebHook(c *gin.Context) {
	// var reqBody BotUpdate
	// defer c.Request.Body.Close()
	// err := json.NewDecoder(c.Request.Body).Decode(&reqBody)

	// get bot
	if bot, err := utils.NewTelegramBot(); err != nil {
		log.Println("Webhook unable to parse update")
	} else {
		update, err2 := bot.HandleUpdate(c.Request)
		log.Println(update.Message.Text, update.Message.From.UserName, err2)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
