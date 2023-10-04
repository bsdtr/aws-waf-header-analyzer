package logparser

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type LogGroupQueryConfig struct {
	AWSRegion    string
	LogGroupName string
	MinutesAgo   int64
}

func NewLogGroupQuery(awsRegion string, logGroupName string, MinutesAgo int64) *LogGroupQueryConfig {
	return &LogGroupQueryConfig{
		AWSRegion:    awsRegion,
		LogGroupName: logGroupName,
		MinutesAgo:   MinutesAgo,
	}
}

func (l *LogGroupQueryConfig) LogGroupQueryResults() (*cloudwatchlogs.GetQueryResultsOutput, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(l.AWSRegion)},
	}))

	svc := cloudwatchlogs.New(sess)

	currentTime := time.Now()
	endTime := currentTime.UnixNano() / int64(time.Millisecond)
	minutesAgo := currentTime.Add(-(time.Duration(l.MinutesAgo)) * time.Minute)
	startTime := minutesAgo.UnixNano() / int64(time.Millisecond)

	queryInput := &cloudwatchlogs.StartQueryInput{
		LogGroupName: aws.String(l.LogGroupName),
		StartTime:    aws.Int64(startTime),
		EndTime:      aws.Int64(endTime),
		QueryString:  aws.String("fields @message"),
	}

	queryOutput, err := svc.StartQuery(queryInput)
	if err != nil {
		fmt.Println("Error on start query:", err)
		return nil, err
	}

	queryID := queryOutput.QueryId

	for {
		queryStatusInput := &cloudwatchlogs.GetQueryResultsInput{
			QueryId: queryID,
		}
		queryResults, err := svc.GetQueryResults(queryStatusInput)
		if err != nil {
			fmt.Println("Error on get query results:", err)
			return nil, err
		}

		if *queryResults.Status == "Complete" {
			break
		}

		time.Sleep(1 * time.Second)
	}

	queryResultsInput := &cloudwatchlogs.GetQueryResultsInput{
		QueryId: queryID,
	}

	queryResults, err := svc.GetQueryResults(queryResultsInput)
	if err != nil {
		fmt.Println("Error on get query results:", err)
		return nil, err
	}

	return queryResults, err
}
