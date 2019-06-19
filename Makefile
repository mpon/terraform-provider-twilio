GO111MODULE=on

build:
	go build -o terraform-provider-wordpress

update:
	go get -u