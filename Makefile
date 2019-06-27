export GO111MODULE=on

build:
	go build -o terraform-provider-twilio

test:
	go test -v ./...

plan: build
	terraform init
	terraform plan

apply: plan
	terraform apply

.PHONY: build test plan apply
