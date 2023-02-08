package utils

import "os"

func GetAppEnv() string {
	return os.Getenv("app_env")
}

func IsProductionEnv() bool {
	return GetAppEnv() == "prod"
}

func GetTelegramBotToken() string {
	return os.Getenv("bot_token")
}

func GetLambdaInvokeUrl() string {
	return os.Getenv("lambda_invoke_url")
}

func GetOpenAIAPIKey() string {
	return os.Getenv(("openai_api_key"))
}
