package twilio

import (
	"log"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mpon/terraform-provider-twilio/twilio-go"
)

func resourceTwilioApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceTwilioApplicationCreate,
		Read:   resourceTwilioApplicationRead,
		Update: resourceTwilioApplicationUpdate,
		Delete: resourceTwilioApplicationDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"voice_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voice_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voice_fallback_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voice_fallback_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_callback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_callback_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voice_caller_id_lookup": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sms_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_fallback_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_fallback_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_status_callback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_status_callback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTwilioApplicationCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*twilio.Client)
	params := url.Values{}
	params.Set("FriendlyName", d.Get("friendly_name").(string))
	if r, ok := d.GetOk("voice_url"); ok {
		params.Set("VoiceUrl", r.(string))
	}
	if r, ok := d.GetOk("voice_method"); ok {
		params.Set("VoiceMethod", r.(string))
	}
	if r, ok := d.GetOk("voice_fallback_url"); ok {
		params.Set("VoiceFallbackUrl", r.(string))
	}
	if r, ok := d.GetOk("voice_fallback_method"); ok {
		params.Set("VoiceFallbackMethod", r.(string))
	}
	if r, ok := d.GetOk("status_callback"); ok {
		params.Set("StatusCallback", r.(string))
	}
	if r, ok := d.GetOk("status_callback_method"); ok {
		params.Set("StatusCallbackMethod", r.(string))
	}
	if r, ok := d.GetOk("voice_caller_id_lookup"); ok {
		params.Set("VoiceCallerIdLookup", strconv.FormatBool(r.(bool)))
	}
	if r, ok := d.GetOk("sms_url"); ok {
		params.Set("SmsUrl", r.(string))
	}
	if r, ok := d.GetOk("sms_method"); ok {
		params.Set("SmsMethod", r.(string))
	}
	if r, ok := d.GetOk("sms_fallback_url"); ok {
		params.Set("SmsFallbackUrl", r.(string))
	}
	if r, ok := d.GetOk("sms_fallback_method"); ok {
		params.Set("SmsFallbackMethod", r.(string))
	}
	if r, ok := d.GetOk("sms_status_callback"); ok {
		params.Set("SmsStatusCallback", r.(string))
	}
	if r, ok := d.GetOk("message_status_callback"); ok {
		params.Set("MessageStatusCallback", r.(string))
	}
	output := twilio.Application{}
	if err := client.CreateApplication(params, &output); err != nil {
		log.Println(err.Error())
		return err
	}
	d.SetId(output.Sid)
	return resourceTwilioApplicationRead(d, m)
}

func resourceTwilioApplicationRead(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.Application{}
	if err := client.GetApplication(sid, &output); err != nil {
		log.Println(err.Error())
		return err
	}
	d.Set("account_sid", output.AccountSid)
	d.Set("api_version", output.APIVersion)
	d.Set("friendly_name", output.FriendlyName)
	d.Set("date_created", output.DateCreated)
	d.Set("date_updated", output.DateUpdated)
	d.Set("message_status_callback", output.MessageStatusCallback)
	d.Set("sms_fallback_method", output.SMSFallbackMethod)
	d.Set("sms_fallback_url", output.SMSFallbackURL)
	d.Set("sms_method", output.SMSMethod)
	d.Set("sms_status_callback", output.SMSStatusCallback)
	d.Set("status_callback", output.StatusCallback)
	d.Set("status_callback_method", output.StatusCallbackMethod)
	d.Set("uri", output.URI)
	d.Set("voice_caller_id_lookup", output.VoiceCallerIDLookup)
	d.Set("voice_fallback_method", output.VoiceFallbackMethod)
	d.Set("voice_fallback_url", output.VoiceFallbackURL)
	d.Set("voice_method", output.VoiceMethod)
	d.Set("voice_url", output.VoiceURL)
	return nil
}

func resourceTwilioApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.Application{}
	if err := client.GetApplication(sid, &output); err != nil {
		log.Println(err.Error())
		return err
	}

	if output.Sid == "" {
		log.Println("application does not exist")
		d.SetId("")
		return nil
	}

	var params = url.Values{}
	if r, ok := d.GetOk("friendly_name"); ok {
		params.Set("FriendlyName", r.(string))
	}
	if r, ok := d.GetOk("voice_url"); ok {
		params.Set("VoiceUrl", r.(string))
	}
	if r, ok := d.GetOk("voice_method"); ok {
		params.Set("VoiceMethod", r.(string))
	}
	if r, ok := d.GetOk("voice_fallback_url"); ok {
		params.Set("VoiceFallbackUrl", r.(string))
	}
	if r, ok := d.GetOk("voice_fallback_method"); ok {
		params.Set("VoiceFallbackMethod", r.(string))
	}
	if r, ok := d.GetOk("status_callback"); ok {
		params.Set("StatusCallback", r.(string))
	}
	if r, ok := d.GetOk("status_callback_method"); ok {
		params.Set("StatusCallbackMethod", r.(string))
	}
	if r, ok := d.GetOk("voice_caller_id_lookup"); ok {
		params.Set("VoiceCallerIdLookup", strconv.FormatBool(r.(bool)))
	}
	if r, ok := d.GetOk("sms_url"); ok {
		params.Set("SmsUrl", r.(string))
	}
	if r, ok := d.GetOk("sms_method"); ok {
		params.Set("SmsMethod", r.(string))
	}
	if r, ok := d.GetOk("sms_fallback_url"); ok {
		params.Set("SmsFallbackUrl", r.(string))
	}
	if r, ok := d.GetOk("sms_fallback_method"); ok {
		params.Set("SmsFallbackMethod", r.(string))
	}
	if r, ok := d.GetOk("sms_status_callback"); ok {
		params.Set("SmsStatusCallback", r.(string))
	}
	if r, ok := d.GetOk("message_status_callback"); ok {
		params.Set("MessageStatusCallback", r.(string))
	}

	if len(params) > 0 {
		updated := twilio.Application{}
		if err := client.UpdateApplication(sid, params, updated); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return resourceTwilioApplicationRead(d, m)
}

func resourceTwilioApplicationDelete(d *schema.ResourceData, m interface{}) error {
	sid := d.Id()
	client := m.(*twilio.Client)
	output := twilio.Application{}
	if err := client.GetApplication(sid, &output); err != nil {
		log.Println(err.Error())
		return err
	}

	if output.Sid == "" {
		log.Println("application does not exist")
		d.SetId("")
		return nil
	}

	if err := client.DeleteApplication(sid); err != nil {
		log.Println(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
