# Dev infra

## Getting Started

1. Create `variables.tfvars` file in this folder, and fill in with following content, replace <...> where necessary.

```
app_env = "dev"
dropbox_app_key = "<YOUR_DROPBOX_API_APP_KEY>"
dropbox_app_secret = "<YOUR_DROPBOX_API_APP_SECRET>"
dropbox_refresh_token = "<YOUR_DROPBOX_API_REFRESH_TOKEN>"
openai_api_key = "<YOUR_OPENAI_API_KEY>"
lambda_invoke_url = "<INVOKE_URL_OF_LAMBDA_FUNCTION>"
```
