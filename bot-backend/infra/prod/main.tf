provider "aws" {
  region = "ap-southeast-1"
}

variable "app_name" {
  description = "Application name"
  default     = "psg-navi-bot-backend"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "prod"
}

variable "bot_token" {
  description = "API token of Telegram bot"
}

variable "lambda_invoke_url" {
  description = "Url to invoke Lambda function"
}

locals {
  app_id = "${lower(var.app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../../build/prod/bin/app"
  output_path = "../../build/prod/bin/app.zip"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}

output "api_url" {
  value = aws_api_gateway_deployment.api_deployment.invoke_url
}