package logparser

import (
	"fmt"
	"testing"
)

func TestLogParser(t *testing.T) {

	logMessage := `{
		        "timestamp": 1696294488732,
		        "formatVersion": 1,
		        "webaclId": "arn:aws:wafv2:us-east-1:123456789:global/webacl/aws-wafv2/fcd56846-1476-41b4-8e2d-f44d81e78c26",
		        "terminatingRuleId": "Default_Action",
		        "terminatingRuleType": "REGULAR",
		        "action": "ALLOW",
		        "terminatingRuleMatchDetails": [],
		        "httpSourceName": "CF",
		        "httpSourceId": "E1B4MZC46KTIDU",
		        "ruleGroupList": [
		            {
		                "ruleGroupId": "AWS#AWSManagedRulesAmazonIpReputationList",
		                "terminatingRule": null,
		                "nonTerminatingMatchingRules": [],
		                "excludedRules": null,
		                "customerConfig": null
		            }
		        ],
		        "rateBasedRuleList": [],
		        "nonTerminatingMatchingRules": [],
		        "requestHeadersInserted": null,
		        "responseCodeSent": null,
		        "httpRequest": {
		            "clientIp": "187.109.199.89",
		            "country": "BR",
		            "headers": [
		                {
		                    "name": "host",
		                    "value": "api.example.com"
		                },
		                {
		                    "name": "user-agent",
		                    "value": "curl/8.2.1"
		                },
		                {
		                    "name": "accept",
		                    "value": "**"
		                },
		                {
		                    "name": "x-custom-header-1",
		                    "value": "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		                },
		                {
		                    "name": "x-custom-header-2",
		                    "value": "zzzzzzzzzzzzzzzzzzzzzzzzzzzz"
		                }
		            ],
		            "uri": "/v1/health",
		            "args": "",
		            "httpVersion": "HTTP/2.0",
		            "httpMethod": "GET",
		            "requestId": "v1pZauSWwgou4bt7afNwHL-Eart-KrBp492Yiz51dmJnBDvuB9hAbA=="
		        },
		        "ja3Fingerprint": "0149f47eabf9a20d0893e2a44e5a6323"
		    }`

	fmt.Println(logMessage)
}
