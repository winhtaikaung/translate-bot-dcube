package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/routes"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

var ginLambda *ginadapter.GinLambda

func isRunningInLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/env", routes.Env)
	r.GET("/about-bot", routes.AboutBot)
	r.POST("/init-bot", routes.InitBot)
	r.POST("/bot"+utils.GetTelegramBotToken(), routes.WebHook)
	return r
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %v\n", request.RequestContext)
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	log.Printf("Start lambda")

	if isRunningInLambda() {
		ginLambda = ginadapter.New(setupRouter())
		lambda.Start(handler)
	} else {
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8080", setupRouter()))
	}
}
