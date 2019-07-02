package twilio

import (
	"fmt"
	"net/url"
	"strconv"
)

// ChatService represents Twilio Chat Service
type ChatService struct {
	Sid                          string            `json:"sid"`
	AccountSid                   string            `json:"account_sid"`
	ConsumptionReportInterval    int               `json:"consumption_report_interval"`
	DateCreated                  string            `json:"date_created"`
	DateUpdated                  string            `json:"date_updated"`
	DefaultChannelCreatorRoleSid string            `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string            `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string            `json:"default_service_role_sid"`
	FriendlyName                 string            `json:"friendly_name"`
	Limits                       ChatServiceLimit  `json:"limits"`
	Links                        map[string]string `json:"links"`
	PreWebhookURL                string            `json:"pre_webhook_url"`
	PostWebhookURL               string            `json:"post_webhook_url"`
	PreWebhookRetryCount         int               `json:"pre_webhook_retry_count"`
	PostWebhookRetryCount        int               `json:"post_webhook_retry_count"`
	ReachabilityEnabled          bool              `json:"reachability_enabled"`
	ReadStatusEnabled            bool              `json:"read_status_enabled"`
	TypingIndicatorTimeout       int               `json:"typing_indicator_timeout"`
	URL                          string            `json:"url"`
	WebhookFilters               []string          `json:"webhook_filters"`
	WebhookMethod                string            `json:"webhook_method"`
	Media                        ChatServiceMedia  `json:"media"`
}

// ChatServiceLimit represents Twilio Chat Service Limit property
type ChatServiceLimit struct {
	ChannelMembers int `json:"channel_members"`
	UserChannels   int `json:"user_channels"`
}

// ToMap returns limits property to map
func (limits ChatServiceLimit) ToMap() map[string]string {
	return map[string]string{
		"channel_members": strconv.Itoa(limits.ChannelMembers),
		"user_channels":   strconv.Itoa(limits.UserChannels),
	}
}

// ChatServiceMedia represents Twilio Chat Service Media property
type ChatServiceMedia struct {
	SizeLimitMb          int    `json:"size_limit_mb"`
	CompatibilityMessage string `json:"compatibility_message"`
}

// GetChatService retrieve a chat service
func (client *Client) GetChatService(sid string, out interface{}) error {
	endPoint := fmt.Sprintf("%s/Services/%s", client.chatBaseURL, sid)
	return client.getRequest(endPoint, out)
}

// CreateChatService creates a chat service
func (client *Client) CreateChatService(params url.Values, out interface{}) error {
	endPoint := fmt.Sprintf("%s/Services", client.chatBaseURL)
	return client.postRequest(endPoint, params, out)
}

// UpdateChatService update a chat service
func (client *Client) UpdateChatService(sid string, params url.Values, out interface{}) error {
	endPoint := fmt.Sprintf("%s/Services/%s", client.chatBaseURL, sid)
	return client.postRequest(endPoint, params, out)
}

// DeleteChatService delete a chat service
func (client *Client) DeleteChatService(sid string) error {
	endPoint := fmt.Sprintf("%s/Services/%s", client.chatBaseURL, sid)
	return client.deleteRequest(endPoint)
}
