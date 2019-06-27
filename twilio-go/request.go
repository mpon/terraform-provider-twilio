package twilio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (client *Client) postRequest(endPoint string, params url.Values, out interface{}) error {
	req, err := client.createRequest(endPoint, params)
	if err != nil {
		log.Println("create request error", err.Error())
		return err
	}

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

func (client *Client) createRequest(endPoint string, params url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(params.Encode()))
	req.SetBasicAuth(client.accountSid, client.authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}