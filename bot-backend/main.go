package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/translate-bot-dcube/bot-backend/routes"
	"github.com/translate-bot-dcube/bot-backend/utils"
	// "net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// BOT ROUTES
	r.GET("/env", routes.Env)
	r.GET("/about-bot", routes.AboutBot)
	r.POST("/init-bot", routes.InitBot)
	r.POST("/bot"+utils.GetTelegramBotToken(), routes.WebHook)

	// MOCK ICA SERVICES ROUTES
	r.POST("/ica/status/:service", routes.QueryICAStatus)
	r.POST("/ica/sg-arrival-card", routes.GenerateSGArrivalCard)

	// MOCK MOM SERVICES ROUTES
	r.POST("/mom/status/:service", routes.QueryMOMStatus)

	return r
}

// func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	// stdout and stderr are sent to AWS CloudWatch Logs
// 	log.Printf("Processing Lambda request %v\n", request.RequestContext)
// 	return ginLambda.ProxyWithContext(ctx, request)
// }

func main() {
	log.Printf("Start Bot")

	log.Fatal(http.ListenAndServe(":8080", setupRouter()))

}
