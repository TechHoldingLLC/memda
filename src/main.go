package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func printHeader() {
	fmt.Println("                           _       ")
	fmt.Println("  /\\/\\   ___ _ __  ___  __| | __ _ ")
	fmt.Println(" /    \\ / _ \\ '_ \\` _ \\/ _\\`|/ _\\`|")
	fmt.Println("/ /\\/\\ \\  __/ | | | | | (_| | (_| |")
	fmt.Println("\\/    \\/\\___|_| |_| |_|\\__,_|\\__,_|")
	fmt.Println("                                   ")
}

func parseArgs() (string, string, string) {
	profile := flag.String("profile", "", "aws profile to use")
	region := flag.String("region", "", "aws region override")
	allLambda := flag.Bool("all", false, "scan all the lambdas")
	singleLambda := flag.String("lambda", "", "function name to analyze")

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

	return *profile, *region, lambda
}

func main() {
	printHeader()

	profile, region, function := parseArgs()

	sess, err := initAWS(profile, region)

	if err != nil {
		log.Fatal(err)
	}

	functions := listLambdas(sess, function)

	for _, f := range functions {
		logs := getLogs(sess, f)
		report(f, logs)
	}

}