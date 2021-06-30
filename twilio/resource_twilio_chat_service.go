package twilio

import (
	"context"
	"log"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"default_channel_creator_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_channel_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_service_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_webhook_retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"post_webhook_retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"webhook_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"limits": {
				Type:     schema.TypeSet,
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
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceTwilioChatServiceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceTwilioChatServiceStateUpgradeV0,
				Version: 0,
			},
		},
	}
}

func resourceTwilioChatServiceCreate(d *schema.ResourceData, m interface{}) error {
	friendlyName := d.Get("friendly_name").(string)
	client := m.(*twilio.Client)
	params := url.Values{}
	params.Set("FriendlyName", friendlyName)
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
	d.Set("default_channel_creator_role_sid", output.DefaultChannelCreatorRoleSid)
	d.Set("default_channel_role_sid", output.DefaultChannelRoleSid)
	d.Set("default_service_role_sid", output.DefaultServiceRoleSid)
	d.Set("pre_webhook_url", output.PreWebhookURL)
	d.Set("post_webhook_url", output.PostWebhookURL)
	d.Set("pre_webhook_retry_count", output.PreWebhookRetryCount)
	d.Set("post_webhook_retry_count", output.PostWebhookRetryCount)
	d.Set("webhook_method", output.WebhookMethod)
	d.Set("limits", output.Limits.ToList())

	// The webhook filter has different API response order and resource data order.
	// So we need to match the order of the resource data.
	if r, ok := d.GetOk("webhook_filters"); ok {
		data := r.([]interface{})
		for _, v := range output.WebhookFilters {
			if isNewWebhookFilter(data, v) {
				data = append(data, v)
			}
		}
		d.Set("webhook_filters", data)
	} else {
		d.Set("webhook_filters", output.WebhookFilters)
	}

	return nil
}

func isNewWebhookFilter(filters []interface{}, filter string) bool {
	for _, v := range filters {
		if v.(string) == filter {
			return false
		}
	}
	return true
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
	if r, ok := d.GetOk("friendly_name"); ok {
		params.Set("FriendlyName", r.(string))
	}
	if r, ok := d.GetOk("default_channel_creator_role_sid"); ok {
		params.Set("DefaultChannelCreatorRoleSid", r.(string))
	}
	if r, ok := d.GetOk("default_channel_role_sid"); ok {
		params.Set("DefaultChannelRoleSid", r.(string))
	}
	if r, ok := d.GetOk("default_service_role_sid"); ok {
		params.Set("DefaultServiceRoleSid", r.(string))
	}
	if r, ok := d.GetOk("pre_webhook_url"); ok {
		params.Set("PreWebhookUrl", r.(string))
	}
	if r, ok := d.GetOk("post_webhook_url"); ok {
		params.Set("PostWebhookUrl", r.(string))
	}
	if r, ok := d.GetOk("pre_webhook_retry_count"); ok {
		params.Set("PreWebhookRetryCount", r.(string))
	}
	if r, ok := d.GetOk("post_webhook_retry_count"); ok {
		params.Set("PostWebhookRetryCount", r.(string))
	}
	if r, ok := d.GetOk("webhook_method"); ok {
		params.Set("WebhookMethod", r.(string))
	}
	if r, ok := d.GetOk("webhook_filters"); ok {
		for _, v := range r.([]interface{}) {
			params.Add("WebhookFilters", v.(string))
		}
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

func resourceTwilioChatServiceResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_channel_creator_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_channel_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_service_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_webhook_retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"post_webhook_retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"webhook_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceTwilioChatServiceStateUpgradeV0(r_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	if rawStateLimits, ok := rawState["limits"].(map[string]interface{}); ok {
		channelMembers := rawStateLimits["channel_members"]
		userChannels := rawStateLimits["user_channels"]
		limits := []interface{}{
			map[string]interface{}{
				"channel_members": channelMembers,
				"user_channels":   userChannels,
			},
		}
		rawState["limits"] = limits
	}

	return rawState, nil
}
