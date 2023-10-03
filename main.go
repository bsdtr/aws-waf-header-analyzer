package main

import (
	"fmt"
	"traffic-analyzer/internal/logparser"
	"traffic-analyzer/internal/rules"
)

var logGroupName = ""

func main() {

	queryResults, err := logparser.LogGroupQueryResults(logGroupName)
	if err != nil {
		panic(err)
	}

	headerRules := rules.HeaderRules()

	headerCount := make(map[logparser.HTTPHeader]int)
	headerAndValues := []logparser.HTTPHeader{}

	for _, result := range queryResults.Results {
		for _, field := range result {
			fieldName := *field.Field
			fieldValue := *field.Value

			if fieldName == "@timestamp" || fieldName == "@message" {
				logMessage, err := logparser.LogParser(fieldValue)
				if err != nil {
					continue
				}

				for _, header := range logMessage.HTTPRequest.Headers {
					if _, ok := headerRules[header.Name]; ok {
						fmt.Println(header.Name, header.Value)

						headerAndValues = append(headerAndValues, logparser.HTTPHeader{
							Name:  header.Name,
							Value: header.Value,
						})
					}
				}
			}
		}
	}

	for _, header := range headerAndValues {
		headerCount[header]++
	}

	for header, count := range headerCount {
		fmt.Printf("(%s, %s): %d vezes\n", header.Name, header.Value, count)
	}
}
