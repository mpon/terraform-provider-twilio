package twilio

import (
	"net/url"
	"testing"
)

func TestCreatePostRequest(t *testing.T) {
	c := Client{
		accountSid: "accountSid",
		authToken:  "authToken",
	}
	params := url.Values{
		"key1": {"value1"},
	}
	req, err := c.createRequest("https://host/path", "POST", params)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if username, password, ok := req.BasicAuth(); ok {
		if username != "accountSid" {
			t.Fatal("username is not accountSid")
		}
		if password != "authToken" {
			t.Fatal("password is not authToken")
		}
	} else {
		t.Fatal("basic auth is not got")
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fatal("Content Type is not set")
	}

	if req.FormValue("key1") != "value1" {
		t.Fatalf("form value is not set")
	}
}

func TestCreateGetRequest(t *testing.T) {
	c := Client{
		accountSid: "accountSid",
		authToken:  "authToken",
	}
	req, err := c.createRequest("https://host/path", "GET", nil)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if username, password, ok := req.BasicAuth(); ok {
		if username != "accountSid" {
			t.Fatal("username is not accountSid")
		}
		if password != "authToken" {
			t.Fatal("password is not authToken")
		}
	} else {
		t.Fatal("basic auth is not got")
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fatal("Content Type is not set")
	}
}

func TestDeleteGetRequest(t *testing.T) {
	c := Client{
		accountSid: "accountSid",
		authToken:  "authToken",
	}
	req, err := c.createRequest("https://host/path", "DELETE", nil)

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if username, password, ok := req.BasicAuth(); ok {
		if username != "accountSid" {
			t.Fatal("username is not accountSid")
		}
		if password != "authToken" {
			t.Fatal("password is not authToken")
		}
	} else {
		t.Fatal("basic auth is not got")
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fatal("Content Type is not set")
	}
}
