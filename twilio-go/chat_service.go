package twilio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ChatService represents Twilio Chat Service
type ChatService struct {
	Sid string `json:"sid"`
}

// CreateChatService creates chat service
func (client *Client) CreateChatService(val url.Values, out interface{}) error {
	url := client.chatBaseURL + "Services"
	req, err := http.NewRequest("POST", url, strings.NewReader(val.Encode()))
	req.SetBasicAuth(client.accountSid, client.authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		log.Println("response error", err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("body read error", err.Error())
		return err
	}

	return json.Unmarshal(body, &out)
}
