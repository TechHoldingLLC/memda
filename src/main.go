package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type lambdaLogs struct {
	name string
	logs []string
}

func printHeader() {
	fmt.Println("                           _       ")
	fmt.Println("  /\\/\\   ___ _ __  ___  __| | __ _ ")
	fmt.Println(" /    \\ / _ \\ '_ \\` _ \\/ _\\`|/ _\\`|")
	fmt.Println("/ /\\/\\ \\  __/ | | | | | (_| | (_| |")
	fmt.Println("\\/    \\/\\___|_| |_| |_|\\__,_|\\__,_|")
	fmt.Println("                                   ")
}

func parseArgs() (string, string, string, int64) {
	profile := flag.String("profile", "", "aws profile to use")
	region := flag.String("region", "", "aws region override")
	allLambda := flag.Bool("all", false, "scan all the lambdas")
	singleLambda := flag.String("lambda", "", "function name to analyze")
	limit := flag.Int("limit", 10, "number of log streams to retrieve")

	flag.Parse()

	if *region == "" {
		err := fmt.Errorf("You need to specify a region")
		fmt.Println(err)
		os.Exit(0)
	}

	if *singleLambda == "" && !*allLambda {
		err := fmt.Errorf("You need to set --lambda or --all")
		fmt.Println(err)
		os.Exit(0)
	}

	if *singleLambda != "" && *allLambda {
		err := fmt.Errorf("You cannot set --lambda and --all at the same time")
		fmt.Println(err)
		os.Exit(0)
	}

	lambda := ""
	if *allLambda {
		lambda = "all"
	} else {
		lambda = *singleLambda
	}

	return *profile, *region, lambda, int64(*limit)
}

func main() {
	printHeader()

	profile, region, function, limit := parseArgs()

	sess, err := initAWS(profile, region)

	if err != nil {
		log.Fatal(err)
	}

	functions := listLambdas(sess, function)
	totalFunctions := len(functions)
	fmt.Printf("Retrieved %d functions\n", totalFunctions)

	logs := []lambdaLogs{}

	for i, f := range functions {
		fmt.Printf("\rGetting logs: %d / %d", i+1, totalFunctions)
		logs = append(logs, lambdaLogs{f, getLogs(sess, f, limit)})
	}
	fmt.Println("\nParsing logs...")

	report(logs)
}
