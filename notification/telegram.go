package notification

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

// SendMessageTelegram send a message to the telegram bot
func SendMessageTelegram(message string) {
	queryParams := "chat_id=" + os.Getenv("CHAT_ID") + "&text=" + url.QueryEscape(message) + "&parse_mode=HTML&disable_web_page_preview=true"

	resp, err := http.Get(os.Getenv("TG_URL") + os.Getenv("TG_TOKEN") + "/sendMessage?" + queryParams)
	if err != nil {
		log.Fatalln("Error sending message to Telegram")
	}

	defer resp.Body.Close()
}
