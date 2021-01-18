package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strings"
)

type FundStock struct {
	FundStockKey string
	Company    	 string
	Symbol       string
	FundName 	 string
}

const FundQuotesSeparator = "-"

func processFundsGlobalQuotes(config *AWSSessionConfig){

	funds := findFundsFromDynamodb(config.Dynamodb)

	fundsMap := mapQuotesByFund(funds)

	fundsQuotesPerformanceMap := loadFundsQuotesPerformance(fundsMap)

	email := generateStockPerformanceEmail(fundsQuotesPerformanceMap)

	sendEmail(config, email)
}

func mapQuotesByFund(funds *dynamodb.ScanOutput)map[string][]FundStock{
	fundsMap := make(map[string][]FundStock)

	for _, fund := range funds.Items {
		fundStock := FundStock{}

		err := dynamodbattribute.UnmarshalMap(fund, &fundStock)

		separatorPosition := strings.Index(fundStock.FundStockKey, FundQuotesSeparator)

		fundCode := fundStock.FundStockKey[:separatorPosition]

		fundsMap[fundCode] = append(fundsMap[fundCode], fundStock)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	return fundsMap
}