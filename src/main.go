package main

import "github.com/aws/aws-lambda-go/lambda"

func handleRequest () {
	awsConfig := getAWSConfig()

	processFundsGlobalQuotes(awsConfig)
}

func main() {
	lambda.Start(handleRequest)
}