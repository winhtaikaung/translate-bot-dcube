resource "aws_ssm_parameter" "dev_params" {
  name        = "/psg_navi_bot/dev/telegram_api_token"
  description = "API token of PSGNaviBot Telegram bot"
  type        = "SecureString"
  value       = var.bot_token

  tags = {
    environment = "dev"
  }
}