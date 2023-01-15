resource "aws_ssm_parameter" "prod_params" {
  name        = "/psg_navi_bot/prod/telegram_api_token"
  description = "API token of PSGNaviBot Telegram bot"
  type        = "SecureString"
  value       = var.bot_token

  tags = {
    environment = "prod"
  }
}