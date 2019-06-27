export GO111MODULE=on

fmt:
	go fmt ./...

build: fmt
	go build -o terraform-provider-twilio

test: fmt
	go test -v ./...

plan: build
	terraform init
	terraform plan

apply: plan
	terraform apply

.PHONY: fmt build test plan apply
