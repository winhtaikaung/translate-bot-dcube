build:
	GOOS=linux GOARCH=amd64 go build -v -a -o build/bin/app .

init:
	terraform -chdir=infra init

plan:
	terraform -chdir=infra plan -out=plan_outfile

apply:
	terraform -chdir=infra apply --auto-approve "plan_outfile"

destroy:
	terraform -chdir=infra destroy --auto-approve