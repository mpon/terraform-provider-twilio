package twilio

import (
	"log"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}
}

// TwilioChatService represents Twilio Chat Service
type TwilioChatService struct {
	Sid string `json:"sid"`
}

func resourceTwilioChatServiceCreate(d *schema.ResourceData, m interface{}) error {
	friendlyName := d.Get("friendly_name").(string)
	client := m.(*TwilioClient)
	val := url.Values{
		"FriendlyName": {friendlyName},
	}
	output := TwilioChatService{}
	err := client.post("https://chat.twilio.com/v2/Services", val, &output)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	sid := output.Sid
	log.Println("sid:", sid)
	d.SetId(sid)
	return resourceTwilioChatServiceRead(d, m)
}

func resourceTwilioChatServiceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTwilioChatServiceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceTwilioChatServiceRead(d, m)
}

func resourceTwilioChatServiceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
