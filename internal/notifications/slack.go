package notifications

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (n *Notifications) SendNotificationToSlack(message string) error {
	messageJSON, err := json.Marshal(map[string]string{
		"text": message,
	})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", n.slack.WebhookURL, bytes.NewBuffer([]byte(messageJSON)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
