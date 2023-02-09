dev:
	export $(grep -v '^#' .env)
	app_env=dev air

run:
	app_env=dev go run .

build: clean
	GOOS=linux GOARCH=amd64 go build -v -a -o build/dev/bin/app .

init:
	terraform -chdir=infra/dev init

plan:
	terraform -chdir=infra/dev plan -var-file=variables.tfvars -out=plan_outfile

apply:
	terraform -chdir=infra/dev apply --auto-approve "plan_outfile"
	curl -X POST "`terraform -chdir=infra/dev output -raw api_url`/init-bot"
	echo "\n"

destroy:
	terraform -chdir=infra/dev destroy -var-file=variables.tfvars

clean:
	rm -rf ./build/dev
