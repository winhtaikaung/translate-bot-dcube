# bot-backend

Code is based on snippets provided [here](https://medium.com/swlh/serverless-super-simple-api-development-with-go-terraform-and-aws-lambda-cc2dd6c531cb) with some amendments:

## Makefile

1. `terraform init|plan|apply|destroy infra` no longer works. To execute terraform commands using configuration in another folder, use `terraform -chdir=infra init|plan|apply|destroy`.
2. `go build` command used in `build` target is simplified to `GOOS=linux GOARCH=amd64 go build -v -a -o build/bin/app .`. Build produced tested to be working fine in AWS Lambda.
3. Modified terraform commands to take in vars from .tfvars file.

## References

### https://blog.canopas.com/golang-serverless-microservices-with-gin-f3c2a4943a6d

Used logic for setting up routes and detecting AWS Lambda environment via env var, and allowing endpoints to be run locally. However usage of `apex/gateway` library is outdated, used next link instead.

### https://blog.0x427567.com/how-to-create-a-serverless-api-with-golang-gin-framework-aws-lambda-and-api-gateway-8f16458a0189

This page provided much updated code snippets and uses `awslabs/aws-lambda-go-api-proxy` which is the official AWS Lambda proxy library for Golang.
