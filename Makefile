export GO111MODULE=on

build:
	go build -o terraform-provider-twilio

plan: build
	terraform init
	terraform plan

apply: plan
	terraform apply

tidy:
	go mod tidy

update:
	go get -u
