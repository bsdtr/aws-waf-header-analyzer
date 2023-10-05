package config

import (
	"github.com/spf13/viper"
)

type NotificationsConfig struct {
	Slack    SlackConfig    `mapstructure:"slack"`
	Telegram TelegramConfig `mapstructure:"telegram"`
}

type SlackConfig struct {
	WebhookURL string `mapstructure:"webhook-url"`
}

type TelegramConfig struct {
	BotToken string `mapstructure:"bot-token"`
	ChatID   string `mapstructure:"chat-id"`
}

type AWSConfig struct {
	WAFLogGroupName       string `mapstructure:"waf-log-group-name"`
	Region                string `mapstructure:"region"`
	RetriveLogsMinutesAgo int64  `mapstructure:"retrive-logs-minutes-ago"`
}

type Config struct {
	Notifications NotificationsConfig `mapstructure:"notifications"`
	AWS           AWSConfig           `mapstructure:"aws"`
}

func NewConfig() *Config {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
