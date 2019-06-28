package twilio

import (
	"log"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mpon/terraform-provider-twilio/twilio-go"
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"limits": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_members": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"user_channels": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
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
	if err := client.CreateChatService(params, &output); err != nil {
		log.Println(err.Error())
		return err
	}
	d.SetId(output.Sid)
	if err := resourceTwilioChatServiceUpdate(d, m); err != nil {
		log.Println(err.Error())
		return err
	}
	return resourceTwilioChatServiceRead(d, m)
}

func resourceTwilioChatServiceRead(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.ChatService{}
	if err := client.GetChatService(sid, &output); err != nil {
		log.Println(err.Error())
		return err
	}
	d.Set("account_sid", output.AccountSid)
	d.Set("friendly_name", output.FriendlyName)
	d.Set("limits", output.Limits.ToMap())
	return nil
}

func resourceTwilioChatServiceUpdate(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.ChatService{}
	if err := client.GetChatService(sid, &output); err != nil {
		log.Println(err.Error())
		return err
	}

	if output.Sid == "" {
		log.Println("chat service does not exist")
		d.SetId("")
		return nil
	}

	var params = url.Values{}
	if d.HasChange("friendly_name") {
		params.Add("FriendlyName", d.Get("friendly_name").(string))
	}
	if r, ok := d.GetOk("limits.channel_members"); ok {
		params.Add("Limits.ChannelMembers", r.(string))
	}
	if r, ok := d.GetOk("limits.user_channels"); ok {
		params.Add("Limits.UserChannels", r.(string))
	}

	if len(params) > 0 {
		updated := twilio.ChatService{}
		if err := client.UpdateChatService(sid, params, updated); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return resourceTwilioChatServiceRead(d, m)
}

func resourceTwilioChatServiceDelete(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.ChatService{}
	if err := client.GetChatService(sid, &output); err != nil {
		log.Println(err.Error())
		return err
	}

	if output.Sid == "" {
		log.Println("chat service does not exist")
		d.SetId("")
		return nil
	}

	if err := client.DeleteChatService(sid); err != nil {
		log.Println(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
