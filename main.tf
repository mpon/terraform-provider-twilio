provider "twilio" {
    # account_sid = ""
    # auth_token = ""
}


resource "twilio_chat_service" "test" {
    friendly_name = "test10"
}