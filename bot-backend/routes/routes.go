package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type BotInfoResult struct {
	ID       int64  `json:"id"`
	Name     string `json:"first_name"`
	UserName string `json:"username"`
}

type BotInfo struct {
	Result BotInfoResult `json:"result"`
}

func Env(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"task_root": os.Getenv("LAMBDA_TASK_ROOT"), "app_env": os.Getenv("app_env")})
}

func AboutBot(c *gin.Context) {
	if res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", os.Getenv("bot_token"))); err != nil {
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
