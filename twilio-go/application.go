package twilio

import (
	"net/url"
	"path"
)

// Application represents Twilio Application
type Application struct {
	AccountSid            string `json:"account_sid"`
	APIVersion            string `json:"api_version"`
	DateCreated           string `json:"date_created"`
	DateUpdated           string `json:"date_updated"`
	FriendlyName          string `json:"friendly_name"`
	MessageStatusCallback string `json:"message_status_callback"`
	Sid                   string `json:"sid"`
	SMSFallbackMethod     string `json:"sms_fallback_method"`
	SMSFallbackURL        string `json:"sms_fallback_url"`
	SMSMethod             string `json:"sms_method"`
	SMSURL                string `json:"sms_url"`
	SMSStatusCallback     string `json:"sms_status_callback"`
	StatusCallback        string `json:"status_callback"`
	StatusCallbackMethod  string `json:"status_callback_method"`
	URI                   string `json:"uri"`
	VoiceCallerIDLookup   bool   `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod   string `json:"voice_fallback_method"`
	VoiceFallbackURL      string `json:"voice_fallback_url"`
	VoiceMethod           string `json:"voice_method"`
	VoiceURL              string `json:"voice_url"`
}

// GetApplication retrieve a TwiML Application
func (client *Client) GetApplication(sid string, out interface{}) error {
	u, err := url.Parse(client.applicationBaeURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "Accounts", client.accountSid, "Applications", sid+".json")
	return client.getRequest(u.String(), out)
}

// CreateApplication creates a TwiML Application
func (client *Client) CreateApplication(params url.Values, out interface{}) error {
	u, err := url.Parse(client.applicationBaeURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "Accounts", client.accountSid, "Applications.json")
	return client.postRequest(u.String(), params, out)
}

// UpdateApplication updates a TwiML Application
func (client *Client) UpdateApplication(sid string, params url.Values, out interface{}) error {
	u, err := url.Parse(client.applicationBaeURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "Accounts", client.accountSid, "Applications", sid+".json")
	return client.postRequest(u.String(), params, out)
}

// DeleteApplication delete a TwiML Application
func (client *Client) DeleteApplication(sid string) error {
	u, err := url.Parse(client.applicationBaeURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "Accounts", client.accountSid, "Applications", sid+".json")
	return client.deleteRequest(u.String())
}
