GO111MODULE=on

build:
	go build -o terraform-provider-twilio

update:
	go get -u