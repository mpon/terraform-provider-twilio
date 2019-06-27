package twilio

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	accountSid := "accountSid"
	authToken := "authToken"
	client := NewClient(accountSid, authToken)
	if client.accountSid != accountSid {
		t.Fatal("accountSid is not saved")
	}
	if client.authToken != authToken {
		t.Fatal("authToken is not saved")
	}
}