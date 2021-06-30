export GO111MODULE=on
version = 99.0.0

fmt:
	go fmt ./...

build: fmt
	go build -o terraform-provider-twilio

install_macos: build
	mkdir -p ~/.terraform.d/plugins/local/mpon/twilio/$(version)/darwin_amd64
	cp terraform-provider-twilio ~/.terraform.d/plugins/local/mpon/twilio/$(version)/darwin_amd64/terraform-provider-twilio_v$(version)

test: fmt
	go test -v ./...

lint: fmt
	golangci-lint run

plan: install_macos
	rm .terraform.lock.hcl
	terraform init -upgrade
	terraform fmt
	terraform plan

apply: plan
	terraform apply

.PHONY: fmt build test plan apply
