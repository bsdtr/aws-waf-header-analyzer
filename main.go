package main

import (
	"fmt"
	"log"
	"traffic-analyzer/internal/config"
	"traffic-analyzer/internal/counter"
	"traffic-analyzer/internal/logparser"
	"traffic-analyzer/internal/notifications"
)

func main() {
	cfg := config.NewConfig()
	logQuery := logparser.NewLogGroupQuery(cfg.AWS.Region, cfg.AWS.WAFLogGroupName, cfg.AWS.RetriveLogsMinutesAgo)
	slackNotifications := notifications.NewSlackNotifications(cfg.Notifications.Slack.WebhookURL)
	telegramNotifications := notifications.NewTelegramNotifications(cfg.Notifications.Telegram.BotToken, cfg.Notifications.Telegram.ChatID)

	var notificationMessage string

	queryResults, err := logQuery.LogGroupQueryResults()
	if err != nil {
		panic(err)
	}

	headerAndValues := logQuery.ParserLogQueryResults(queryResults)
	exceededThresholdHeader := counter.CounterExceededThresholdHeader(headerAndValues)

	for _, header := range exceededThresholdHeader {
		notificationMessage += fmt.Sprintf("The '%s:%s' with IP address '%s' has exceeded the request limit with %d requests in %d minutes.\n", header.Name, header.Value, header.IP, header.NumberOfRequests, cfg.AWS.RetriveLogsMinutesAgo)
	}

	log.Println(notificationMessage)

	err = slackNotifications.SendNotificationToSlack(notificationMessage)

	if err != nil {
		log.Printf("Error on send notification to slack: %v\n", err)
	}

	err = telegramNotifications.SendNotificationToTelegram(notificationMessage)

	if err != nil {
		log.Printf("Error on send notification to telegram: %v\n", err)
	}
}
