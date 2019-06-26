package main

import (
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is the root of terraform provider plugin
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_ACCOUNT_SID", nil),
			},
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_chat_service": resourceTwilioChatService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := NewTwilioClient(
		d.Get("account_sid").(string),
		d.Get("auth_token").(string),
	)

	c := cleanhttp.DefaultClient()
	c.Transport = logging.NewTransport("Twilio", c.Transport)
	client.HTTPClient = c

	return client, nil
}
