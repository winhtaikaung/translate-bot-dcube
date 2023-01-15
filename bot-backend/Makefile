build: clean
	GOOS=linux GOARCH=amd64 go build -v -a -o build/bin/app .

init:
	terraform -chdir=infra init

plan:
	terraform -chdir=infra plan -var-file=dev.tfvars -out=plan_dev_outfile

apply:
	terraform -chdir=infra apply --auto-approve "plan_dev_outfile"

destroy:
	terraform -chdir=infra apply -destroy "plan_dev_outfile"

clean:
	rm -rf ./build
