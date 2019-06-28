provider "twilio" {
    # account_sid = "" // You can set env variables TWILIO_ACCOUNT_SID
    # auth_token = ""  // You can set env variables TWILIO_AUTH_TOKEN
}


resource "twilio_chat_service" "test" {
    friendly_name = "test17"
    limits {
        channel_members = 1
        user_channels = 1
    }
}

resource "twilio_chat_service" "test2" {
    friendly_name = "test22"
    limits {
        channel_members = 4
        user_channels = 5
    }
}