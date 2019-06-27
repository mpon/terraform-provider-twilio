package twilio

import (
	"fmt"
	"net/url"
	"strings"
)

// ChatService represents Twilio Chat Service
type ChatService struct {
	Sid                          string             `json:"sid"`
	AccountSid                   string             `json:"account_sid"`
	ConsumptionReportInterval    int           		`json:"consumption_report_interval"`
	DateCreated                  string             `json:"date_created"`
	DateUpdated                  string             `json:"date_updated"`
	DefaultChannelCreatorRoleSid string             `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string             `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string             `json:"default_service_role_sid"`
	FriendlyName                 string             `json:"friendly_name"`
	Limits                       ChatServiceLimit   `json:"limits"`
	Links                        map[string]string  `json:"links"`
	PostWebhookUrl               string             `json:"post_webhook_url"`
	PreWebhookUrl                string             `json:"pre_webhook_url"`
	PostWebhookRetryCount        int                `json:"post_webhook_retry_count"`
	PreWebhookRetryCount         int                `json:"pre_webhook_retry_count"`
	ReachabilityEnabled          bool               `json:"reachability_enabled"`
	ReadStatusEnabled            bool               `json:"read_status_enabled"`
	TypingIndicatorTimeout       int                `json:"typing_indicator_timeout"`
	Url                          string             `json:"url"`
	WebhookFilters               []ChatWebhookEvent `json:"webhook_filters"`
	WebhookMethod                string             `json:"webhook_method"`
	Media                        ChatServiceMedia   `json:"media"`
}

type ChatServiceLimit struct {
	ChannelMembers int `json:"channel_members"`
	UserChannels   int `json:"user_channels"`
}

type ChatServiceMedia struct {
	SizeLimitMb          int `json:"size_limit_mb"`
	CompatibilityMessage string `json:"compatibility_message"`
}

type ChatWebhookEvent string

const OnMessageSend = ChatWebhookEvent("onMessageSend")
const OnMessageRemove = ChatWebhookEvent("onMessageRemove")
const OnMessageUpdate = ChatWebhookEvent("onMessageUpdate")
const OnMediaMessageSend = ChatWebhookEvent("onMediaMessageSend")
const OnChannelAdd = ChatWebhookEvent("onChannelAdd")
const OnChannelUpdate = ChatWebhookEvent("onChannelUpdate")
const OnChannelDestroy = ChatWebhookEvent("onChannelDestroy")
const OnMemberAdd = ChatWebhookEvent("onMemberAdd")
const OnMemberRemove = ChatWebhookEvent("onMemberRemove")
const OnUserAdded = ChatWebhookEvent("onUserAdded")
const OnUserUpdate = ChatWebhookEvent("onUserUpdate")

func (s ChatWebhookEvent) ToString() string {
	switch s {
	case OnMessageSend:
		return "onMessageSend"
	case OnMessageRemove:
		return "onMessageRemove"
	case OnMessageUpdate:
		return "onMessageUpdate"
	case OnMediaMessageSend:
		return "onMediaMessageSend"
	case OnChannelAdd:
		return "onChannelAdd"
	case OnChannelUpdate:
		return "onChannelUpdate"
	case OnChannelDestroy:
		return "onChannelDestroy"
	case OnMemberAdd:
		return "onMemberAdd"
	case OnMemberRemove:
		return "onMemberRemove"
	case OnUserAdded:
		return "onUserAdded"
	case OnUserUpdate:
		return "onUserUpdate"
	default:
		return strings.Title(string(s))
	}
}

// GetChatService retrieve a chat service
func (client *Client) GetChatService(sid string, out interface{}) error {
	endPoint := fmt.Sprintf("%s/Services/%s", client.chatBaseURL, sid)
	return client.getRequest(endPoint, out)
}

// CreateChatService creates chat service
func (client *Client) CreateChatService(params url.Values, out interface{}) error {
	endPoint := fmt.Sprintf("%s/Services", client.chatBaseURL)
	return client.postRequest(endPoint, params, out)
}

func (client *Client) UpdateChatService(sid string, params url.Values, out interface{}) error {
	endPoint := fmt.Sprintf("%s/Services/%s", client.chatBaseURL, sid)
	return client.postRequest(endPoint, params, out)
}

func (client *Client) DeleteChatService(sid string) error {
	endPoint := fmt.Sprintf("%s/Services/%s", client.chatBaseURL, sid)
	return client.deleteRequest(endPoint)
}