package twilio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TwilioClient is the object that handles talking to the Twilio API. This maintains
// state information for a particular application connection.
type TwilioClient struct {
	accountSid, authToken, baseURL string

	//The Http Client that is used to make requests
	HTTPClient   *http.Client
	RetryTimeout time.Duration
}

// NewTwilioClient returns a new Twilio HTTP Client which can be used to access the API
func NewTwilioClient(accountSid, authToken string) *TwilioClient {
	baseURL := "https://chat.twilio.com/v2/"

	return &TwilioClient{
		accountSid:   accountSid,
		authToken:    authToken,
		baseURL:      baseURL,
		HTTPClient:   http.DefaultClient,
		RetryTimeout: time.Duration(60 * time.Second),
	}
}

func (client *TwilioClient) post(url string, val url.Values, out interface{}) error {
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
