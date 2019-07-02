package twilio

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
	SMSURL                string `json:"sms_url"`
	StatusCallback        string `json:"status_callback"`
	StatusCallbackMethod  string `json:"status_callback_method"`
	URI                   string `json:"uri"`
	VoiceCallerIDLookup   bool   `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod   string `json:"voice_fallback_method"`
	VoiceFallbackURL      string `json:"voice_fallback_url"`
	VoiceMethod           string `json:"voice_method"`
	VoiceURL              string `json:"voice_url"`
}
