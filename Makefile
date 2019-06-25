export GO111MODULE=on

build:
	go build -o terraform-provider-twilio

plan: build
	terraform init
	terraform plan

apply:
	terraform apply

update:
	go get -u
