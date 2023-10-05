package logparser

import (
	"encoding/json"
	"fmt"
)

type HTTPHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HTTPRequest struct {
	ClientIP    string       `json:"clientIp"`
	Country     string       `json:"country"`
	Headers     []HTTPHeader `json:"headers"`
	URI         string       `json:"uri"`
	Args        string       `json:"args"`
	HTTPVersion string       `json:"httpVersion"`
	HTTPMethod  string       `json:"httpMethod"`
	RequestID   string       `json:"requestId"`
}

type RuleGroup struct {
	RuleGroupId                 string        `json:"ruleGroupId"`
	TerminatingRule             interface{}   `json:"terminatingRule"`
	NonTerminatingMatchingRules []interface{} `json:"nonTerminatingMatchingRules"`
	ExcludedRules               interface{}   `json:"excludedRules"`
	CustomerConfig              interface{}   `json:"customerConfig"`
}

type LogEntry struct {
	Timestamp                   int64         `json:"timestamp"`
	FormatVersion               int           `json:"formatVersion"`
	WebACLId                    string        `json:"webaclId"`
	TerminatingRuleId           string        `json:"terminatingRuleId"`
	TerminatingRuleType         string        `json:"terminatingRuleType"`
	Action                      string        `json:"action"`
	TerminatingRuleMatchDetails []interface{} `json:"terminatingRuleMatchDetails"`
	HTTPSourceName              string        `json:"httpSourceName"`
	HTTPSourceId                string        `json:"httpSourceId"`
	RuleGroupList               []RuleGroup   `json:"ruleGroupList"`
	RateBasedRuleList           []interface{} `json:"rateBasedRuleList"`
	NonTerminatingMatchingRules []interface{} `json:"nonTerminatingMatchingRules"`
	RequestHeadersInserted      interface{}   `json:"requestHeadersInserted"`
	ResponseCodeSent            interface{}   `json:"responseCodeSent"`
	HTTPRequest                 HTTPRequest   `json:"httpRequest"`
	Ja3Fingerprint              string        `json:"ja3Fingerprint"`
}

func LogUnmarshal(logMessage string) (*LogEntry, error) {
	var logEntry LogEntry

	err := json.Unmarshal([]byte(logMessage), &logEntry)
	if err != nil {
		fmt.Println("Error on unmarshal JSON:", err)
		return nil, err
	}

	return &logEntry, nil
}
