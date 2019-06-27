provider "twilio" {
    # account_sid = "" // You can set env variables TWILIO_ACCOUNT_SID
    # auth_token = ""  // You can set env variables TWILIO_AUTH_TOKEN
}


resource "twilio_chat_service" "test" {
    friendly_name = "test15"
}