package twilio

import (
	"log"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	twilio "github.com/mpon/terraform-provider-twilio/twilio-go"
)

func resourceTwilioChatService() *schema.Resource {
	return &schema.Resource{
		Create: resourceTwilioChatServiceCreate,
		Read:   resourceTwilioChatServiceRead,
		Update: resourceTwilioChatServiceUpdate,
		Delete: resourceTwilioChatServiceDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTwilioChatServiceCreate(d *schema.ResourceData, m interface{}) error {
	friendlyName := d.Get("friendly_name").(string)
	client := m.(*twilio.Client)
	params := url.Values{
		"FriendlyName": {friendlyName},
	}
	output := twilio.ChatService{}
	err := client.CreateChatService(params, &output)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	d.SetId(output.Sid)
	return resourceTwilioChatServiceRead(d, m)
}

func resourceTwilioChatServiceRead(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.ChatService{}
	err := client.GetChatService(sid, &output)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	d.Set("account_sid", output.AccountSid)
	d.Set("friendly_name", output.FriendlyName)
	return nil
}

func resourceTwilioChatServiceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceTwilioChatServiceRead(d, m)
}

func resourceTwilioChatServiceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
