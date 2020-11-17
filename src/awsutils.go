package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func initAWS(profile string, region string) (*session.Session, error) {
	awsConfig := &aws.Config{}

	if profile != "" {
		awsConfig.Credentials = credentials.NewSharedCredentials("", profile)
	}

	if region != "" {
		awsConfig.Region = aws.String(region)
	}

	sess, err := session.NewSession(awsConfig)

	return sess, err
}

func listLambdas(sess *session.Session, function string) []string {
	svc := lambda.New(sess)

	functions := []string{}

	if function == "all" {
		marker := "start"
		for marker != "" {
			listFunctionsConfig := &lambda.ListFunctionsInput{}
			if marker != "start" {
				listFunctionsConfig.Marker = &marker
			}

			result, err := svc.ListFunctions(listFunctionsConfig)
			if err != nil {
				log.Fatal("Cannot list functions")
			}
			marker = aws.StringValue(result.NextMarker)

			for _, f := range result.Functions {
				functions = append(functions, aws.StringValue(f.FunctionName))
			}
		}
	} else {
		functions = append(functions, function)
	}

	return functions
}

func getLogs(sess *session.Session, function string, limit int64) []string {
	svc := cloudwatchlogs.New(sess)

	logGroupName := fmt.Sprint("/aws/lambda/", function)

	result, err := svc.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &logGroupName,
		Limit:        &limit,
	})

	if err != nil {
		switch err := err.(type) {
		case *cloudwatchlogs.ResourceNotFoundException:
			fmt.Printf("\nCannot find the log group: %s\n", logGroupName)
		default:
			log.Fatal(err)
		}
	}

	logStreamNames := []string{}

	for _, l := range result.LogStreams {
		logStreamNames = append(logStreamNames, aws.StringValue(l.LogStreamName))
	}

	events := []string{}

	startFromHead := true

	for _, s := range logStreamNames {
		result, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
			LogGroupName:  &logGroupName,
			LogStreamName: &s,
			StartFromHead: &startFromHead,
		})

		if err != nil {
			log.Fatal(err)
		}

		for _, e := range result.Events {
			msg := aws.StringValue(e.Message)
			if strings.HasPrefix(msg, "REPORT") {
				events = append(events, msg)
			}
		}
	}

	return events
}
