package counter

import (
	"aws-waf-header-analyzer/internal/logparser"
	"aws-waf-header-analyzer/internal/rules"
)

type threshold struct {
	Name             string
	Value            string
	NumberOfRequests int
	IP               string
}

func CounterExceededThresholdHeader(headerAndValues []logparser.QueryResultsOutput) []threshold {
	headerRules := rules.HeaderRules()
	headerCount := make(map[logparser.QueryResultsOutput]int)
	headersExceededThreshold := []threshold{}

	for _, header := range headerAndValues {
		headerCount[header]++
	}

	for header, count := range headerCount {
		if count > headerRules[header.Name] {
			headersExceededThreshold = append(headersExceededThreshold, threshold{
				Name:             header.Name,
				Value:            header.Value,
				NumberOfRequests: count,
				IP:               header.IP,
			})
		}
	}

	return headersExceededThreshold
}
