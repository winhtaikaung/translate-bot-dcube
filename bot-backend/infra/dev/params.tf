resource "aws_ssm_parameter" "dev_params" {
  name        = "/psg_navi_bot/dev/telegram_api_token"
  description = "API token of PSGNaviBot Telegram bot"
  type        = "SecureString"
  value       = var.bot_token

  tags = {
    environment = "dev"
  }
}

resource "aws_ssm_parameter" "dev_openai_api_key" {
  name        = "/psg_navi_bot/dev/openai_api_key"
  description = "API key for OpenAI"
  type        = "SecureString"
  value       = var.openai_api_key

  tags = {
    environment = "dev"
  }
}