package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func parseArgs() (string, string, string) {
	profile := flag.String("profile", "", "aws profile to use")
	region := flag.String("region", "", "aws region override")
	allLambda := flag.Bool("all", false, "scan all the lambdas")
	singleLambda := flag.String("lambda", "", "function name to analyze")

	flag.Parse()

	if *singleLambda != "" && *allLambda {
		err := fmt.Errorf("You cannot set --lambda and --all at the same time")
		log.Fatal(err)
	}

	lambda := ""
	if *allLambda {
		lambda = "all"
	} else {
		lambda = *singleLambda
	}

	return *profile, *region, lambda
}

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
		result, err := svc.ListFunctions(nil)
		if err != nil {
			log.Fatal("Cannot list functions")
		}

		for _, f := range result.Functions {
			functions = append(functions, aws.StringValue(f.FunctionName))
		}
	} else {
		functions = append(functions, function)
	}

	return functions
}

func getLogs(sess *session.Session, function string) []string {
	svc := cloudwatchlogs.New(sess)

	logGroupName := fmt.Sprint("/aws/lambda/", function)
	// var limit int64 = 10

	result, err := svc.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &logGroupName,
		// Limit:        &limit,
	})

	if err != nil {
		log.Fatal(err)
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

func main() {
	profile, region, function := parseArgs()

	sess, err := initAWS(profile, region)

	if err != nil {
		log.Fatal(err)
	}

	functions := listLambdas(sess, function)

	logs := []string{}

	for _, f := range functions {
		logs = append(logs, getLogs(sess, f)...)
	}

	fmt.Println(logs)
}
