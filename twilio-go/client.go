package twilio

import (
	"net/http"
	"time"
)

// Client is the object that handles talking to the Twilio API. This maintains
// state information for a particular application connection.
type Client struct {
	accountSid, authToken, chatBaseURL, applicationBaeURL string

	//The HTTP Client that is used to make requests
	HTTPClient   *http.Client
	RetryTimeout time.Duration
}

// NewClient returns a new Twilio HTTP Client which can be used to access the API
func NewClient(accountSid, authToken string) *Client {
	chatBaseURL := "https://chat.twilio.com/v2"
	applicationBaseURL := "https://api.twilio.com/2010-04-01"

	return &Client{
		accountSid:        accountSid,
		authToken:         authToken,
		chatBaseURL:       chatBaseURL,
		applicationBaeURL: applicationBaseURL,
		HTTPClient:        http.DefaultClient,
		RetryTimeout:      time.Duration(60 * time.Second),
	}
}
