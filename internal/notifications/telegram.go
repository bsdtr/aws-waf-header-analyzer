package notifications

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (n *Notifications) SendNotificationToTelegram(message string) error {
	messageJSON, err := json.Marshal(map[string]string{
		"chat_id": n.telegram.ChatID,
		"text":    message,
	})

	if err != nil {
		return err
	}

	requestURL := "https://api.telegram.org/bot" + n.telegram.BotToken + "/sendMessage"

	req, _ := http.NewRequest("POST", requestURL, bytes.NewBuffer([]byte(messageJSON)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
