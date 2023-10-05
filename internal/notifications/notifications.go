package notifications

type Slack struct {
	WebhookURL string
}

type Telegram struct {
	BotToken string
	ChatID   string
}

type Notifications struct {
	slack    Slack
	telegram Telegram
}

func NewSlackNotifications(webhookURL string) *Notifications {
	slack := Slack{
		WebhookURL: webhookURL,
	}

	return &Notifications{
		slack: slack,
	}
}

func NewTelegramNotifications(botToken, chatID string) *Notifications {
	telegram := Telegram{
		BotToken: botToken,
		ChatID:   chatID,
	}

	return &Notifications{
		telegram: telegram,
	}
}
