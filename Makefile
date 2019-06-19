GO111MODULE=on

build:
	go build -o terraform-provider-httpbin

update:
	go get -u