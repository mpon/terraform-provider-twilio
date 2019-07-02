export GO111MODULE=on

fmt:
	go fmt ./...

build: fmt
	go build -o terraform-provider-twilio

test: fmt
	go test -v `go list ./... | tail -n +2`

plan: build
	terraform init
	terraform fmt
	terraform plan

apply: plan
	terraform apply

.PHONY: fmt build test plan apply
