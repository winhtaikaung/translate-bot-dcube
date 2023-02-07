# psg-navi-bot
Telegram bot for Parent Support Group

## Project structure

- bot-backend: AWS Lambda functions to handle incoming requests from bot

## Please open in devcontainer 
- after running please run `ngrok http 8080` to expose local server 


## Authenticate with bot

Setup launch.json as follow if you use VScode

```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Run Bot Server",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {
                "bot_token": "your bot father token",
                "lambda_invoke_url": "https://5427-42-60-224-74.ngrok.io"
            }
        }
    ]
}
```




