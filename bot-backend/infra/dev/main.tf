provider "aws" {
  region = "ap-southeast-1"
}

variable "app_name" {
  description = "Application name"
  default     = "dcube-translate-bot-backend"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "dev"
}

variable "bot_token" {
  description = "API token of Telegram bot"
}

variable "openai_api_key" {
  description = "API key for OpenAI"
}

variable "lambda_invoke_url" {
  description = "Url to invoke Lambda function"
}

locals {
  app_id = "${lower(var.app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_dir = "../../build/dev/bin"
  output_path = "../../build/dev/app.zip"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}

output "api_url" {
  value = aws_api_gateway_deployment.api_deployment.invoke_url
}