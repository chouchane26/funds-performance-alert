package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func findFundsFromDynamodb(dynamodbConfig Dynamodb) *dynamodb.ScanOutput {
	input := &dynamodb.ScanInput{
		TableName: aws.String(dynamodbConfig.TableName),
	}

	result, err := dynamodbConfig.DynamodbSession.Scan(input)

	if err != nil {
		fmt.Errorf("Error related to DynamodbScan %v\n", err)
		os.Exit(1)
	}

	return result
}