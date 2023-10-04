package counter

import (
	"traffic-analyzer/internal/logparser"
	"traffic-analyzer/internal/rules"
)

type threshold struct {
	Name             string `json:"name"`
	NumberOfRequests int    `json:"numberOfRequests"`
}

func CounterHeader(headerAndValues []logparser.HTTPHeader) []threshold {
	headerRules := rules.HeaderRules()
	headerCount := make(map[logparser.HTTPHeader]int)
	headersExceededThreshold := []threshold{}

	for _, header := range headerAndValues {
		headerCount[header]++
	}

	for header, count := range headerCount {
		if count > headerRules[header.Name] {
			headersExceededThreshold = append(headersExceededThreshold, threshold{
				Name:             header.Name,
				NumberOfRequests: count,
			})
		}
	}

	return headersExceededThreshold
}
