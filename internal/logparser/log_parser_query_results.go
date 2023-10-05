package logparser

import (
	"traffic-analyzer/internal/rules"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type QueryResultsOutput struct {
	Name  string
	Value string
	IP    string
}

func (l *LogGroupQueryConfig) ParserLogQueryResults(queryResults *cloudwatchlogs.GetQueryResultsOutput) []QueryResultsOutput {

	headerRulesConfig := rules.HeaderRules()

	headerAndValues := []QueryResultsOutput{}

	for _, result := range queryResults.Results {
		for _, field := range result {
			fieldName := *field.Field
			fieldValue := *field.Value

			if fieldName == "@message" {
				logMessage, err := LogUnmarshal(fieldValue)
				if err != nil {
					continue
				}

				for _, header := range logMessage.HTTPRequest.Headers {
					if _, ok := headerRulesConfig[header.Name]; ok {
						headerAndValues = append(headerAndValues, QueryResultsOutput{
							Name:  header.Name,
							Value: header.Value,
							IP:    logMessage.HTTPRequest.ClientIP,
						})
					}
				}
			}
		}
	}

	return headerAndValues
}
