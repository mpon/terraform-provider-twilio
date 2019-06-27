package twilio

import (
	"net/url"
)

// ChatService represents Twilio Chat Service
type ChatService struct {
	Sid string `json:"sid"`
}

// CreateChatService creates chat service
func (client *Client) CreateChatService(params url.Values, out interface{}) error {
	endPoint := client.chatBaseURL + "Services"
	return client.postRequest(endPoint, params, out)
}
