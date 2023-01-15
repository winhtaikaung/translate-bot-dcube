# bot-backend

Code is based on snippets provided [here](https://medium.com/swlh/serverless-super-simple-api-development-with-go-terraform-and-aws-lambda-cc2dd6c531cb) with some amendments:

## Makefile

1. `terraform init|plan|apply|destroy infra` no longer works. To execute terraform commands using configuration in another folder, use `terraform -chdir=infra init|plan|apply|destroy`.
2. `go build` command used in `build` target is simplified to `GOOS=linux GOARCH=amd64 go build -v -a -o build/bin/app .`. Build produced tested to be working fine in AWS Lambda.
